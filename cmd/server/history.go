package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx"
	pb "github.com/patterns/fhirbuffer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// The key/val pair for the tmp slice
type kvpair struct {
	k  string
	v  int
	ts time.Time
}

// Keeping a known-types map built from the results of
// SELECT Table_name FROM information_schema.tables WHERE Table_name LIKE '%_history';

var resourceHistory = map[string]int{
	"devicerequest":                  0,
	"servicerequest":                 0,
	"devicecomponent":                0,
	"devicemetric":                   0,
	"usersession":                    0,
	"careplan":                       0,
	"observation":                    0,
	"enrollmentrequest":              0,
	"group":                          0,
	"messagedefinition":              0,
	"appointment":                    0,
	"biologicallyderivedproduct":     0,
	"questionnaireresponse":          0,
	"episodeofcare":                  0,
	"substancepolymer":               0,
	"processresponse":                0,
	"supplydelivery":                 0,
	"adverseevent":                   0,
	"iteminstance":                   0,
	"endpoint":                       0,
	"substancereferenceinformation":  0,
	"compartmentdefinition":          0,
	"detectedissue":                  0,
	"medicationadministration":       0,
	"implementationguide":            0,
	"goal":                           0,
	"communication":                  0,
	"schedule":                       0,
	"documentreference":              0,
	"coverage":                       0,
	"auditevent":                     0,
	"messageheader":                  0,
	"contract":                       0,
	"sequence":                       0,
	"testreport":                     0,
	"codesystem":                     0,
	"plandefinition":                 0,
	"invoice":                        0,
	"claimresponse":                  0,
	"chargeitem":                     0,
	"bodystructure":                  0,
	"parameters":                     0,
	"clinicalimpression":             0,
	"familymemberhistory":            0,
	"medicinalproductauthorization":  0,
	"binary":                         0,
	"composition":                    0,
	"practitionerrole":               0,
	"healthcareservice":              0,
	"patient":                        0,
	"entrydefinition":                0,
	"medicationdispense":             0,
	"deviceusestatement":             0,
	"structuremap":                   0,
	"immunizationevaluation":         0,
	"library":                        0,
	"basic":                          0,
	"slot":                           0,
	"activitydefinition":             0,
	"specimen":                       0,
	"diagnosticreport":               0,
	"subscription":                   0,
	"requestgroup":                   0,
	"provenance":                     0,
	"medicinalproduct":               0,
	"organizationrole":               0,
	"practitioner":                   0,
	"medicinalproductpackaged":       0,
	"flag":                           0,
	"explanationofbenefit":           0,
	"linkage":                        0,
	"operationoutcome":               0,
	"medicinalproductpharmaceutical": 0,
	"immunization":                   0,
	"researchsubject":                0,
	"expansionprofile":               0,
	"eligibilityrequest":             0,
	"medicinalproductclinicals":      0,
	"occupationaldata":               0,
	"paymentnotice":                  0,
	"namingsystem":                   0,
	"medicationstatement":            0,
	"enrollmentresponse":             0,
	"nutritionorder":                 0,
	"questionnaire":                  0,
	"account":                        0,
	"eventdefinition":                0,
	"medicinalproductdevicespec":     0,
	"substancespecification":         0,
	"communicationrequest":           0,
	"specimendefinition":             0,
	"verificationresult":             0,
	"documentmanifest":               0,
	"eligibilityresponse":            0,
	"task":                           0,
	"valueset":                       0,
	"claim":                          0,
	"examplescenario":                0,
	"researchstudy":                  0,
	"medicationrequest":              0,
	"measure":                        0,
	"list":                           0,
	"encounter":                      0,
	"capabilitystatement":            0,
	"visionprescription":             0,
	"riskassessment":                 0,
	"immunizationrecommendation":     0,
	"processrequest":                 0,
	"relatedperson":                  0,
	"medication":                     0,
	"appointmentresponse":            0,
	"substance":                      0,
	"paymentreconciliation":          0,
	"testscript":                     0,
	"conceptmap":                     0,
	"person":                         0,
	"condition":                      0,
	"careteam":                       0,
	"structuredefinition":            0,
	"procedure":                      0,
	"consent":                        0,
	"observationdefinition":          0,
	"productplan":                    0,
	"location":                       0,
	"organization":                   0,
	"device":                         0,
	"supplyrequest":                  0,
	"allergyintolerance":             0,
	"operationdefinition":            0,
	"imagingstudy":                   0,
	"medicinalproductingredient":     0,
	"guidanceresponse":               0,
	"media":                          0,
	"measurereport":                  0,
	"graphdefinition":                0,
	"terminologycapabilities":        0,
	"metadataresource":               0,
	"concept":                        0,
}

func (s *fhirbuffer) History(req *pb.Search, stream pb.Fhirbuffer_HistoryServer) error {
	// The first and simplest history is the table list with a couple aggregate values.

	res := strings.ToLower(req.Type)
	if res != "change" {
		return status.Error(codes.Unimplemented, "Unsupported history type")
	}

	for k, v := range resourceHistory {
		stats, err := json.Marshal(map[string]string{"resourceType": k, "count": strconv.Itoa(v), "id": k})
		if err != nil {
			return status.Error(codes.Unknown, err.Error())
		}

		rec := &pb.Record{Resource: stats}

		if err := stream.Send(rec); err != nil {
			return status.Error(codes.Unknown, err.Error())
		}
	}
	return nil
}

func (s *fhirbuffer) loadHistory(ctx context.Context) error {
	conn, err := pgx.Connect(*databaseConfig)
	if err != nil {
		log.Printf("Database connection, %v", err)
		return err
	}
	defer conn.Close()

	// Use a tmp slice as accumulator
	tmp := make([]kvpair, 1)

	for key, _ := range resourceHistory {
		var histab = fmt.Sprintf("public.%s_history", key)

		qr := conn.QueryRow("SELECT COUNT(id), coalesce(MAX(ts), '0001-01-01') FROM " + histab)

		var total int
		var last time.Time
		err := qr.Scan(&total, &last)

		switch err {
		case nil:
			tmp = append(tmp, kvpair{k: key, v: total, ts: last})
			if key == "patient" {
				log.Println("patient hist == ", total)
				log.Println("patient last == ", last)
			}

		case pgx.ErrNoRows:
			continue

		default:
			log.Printf("Database error, %v", err)
			return err
		}
	}

	// Now mutate the map safely
	for _, pair := range tmp {
		if _, ok := resourceHistory[pair.k]; ok {
			resourceHistory[pair.k] = pair.v
		}
	}
	log.Println("mutated ", resourceHistory["patient"])

	return nil
}
