package mongodbapi

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/av-belyakov/shaper_stix_2.1/confighandler"
	wrap "github.com/av-belyakov/shaper_stix_2.1/databaseapi/mongodbapi/wrappers"
	"github.com/av-belyakov/shaper_stix_2.1/datamodels"
)

// NewClientMongoDB инициализация нового клиента СУБД MongoDB
func NewClientMongoDB(
	ctxDone context.Context,
	conf confighandler.AppConfigMongoDB,
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) (*MongoDBModule, error) {
	channels := &MongoDBModule{
		chanInputToModule:    make(chan ChanInput),
		chanOutputFromModule: make(chan ChanOutput),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	description := ConnectionDescriptorMongoDB{
		databaseName: conf.NameDB,
		ctx:          ctx,
		ctxCancel:    cancel,
		ctxRoute:     ctxDone,
	}

	conn, err := NewConnection(ctx, conf)
	if err != nil {
		return channels, err
	}

	description.connection = conn
	description.Routing(channels, logging, counting)

	return channels, nil
}

func NewConnection(ctx context.Context, conf confighandler.AppConfigMongoDB) (*mongo.Client, error) {
	confPath := fmt.Sprintf("mongodb://%s:%d/%s", conf.Host, conf.Port, conf.NameDB)

	connect, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		AuthSource:    conf.NameDB,
		Username:      conf.User,
		Password:      conf.Passwd,
	}).ApplyURI(confPath))
	if err != nil {
		return connect, err
	}

	if err = connect.Ping(ctx, readpref.Primary()); err != nil {
		return connect, err
	}

	log.Printf("Create connection with MongoDB (%s:%d)\n", conf.Host, conf.Port)

	return connect, nil
}

func (conn ConnectionDescriptorMongoDB) Routing(
	channels *MongoDBModule,
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings,
) {
	ws := wrap.Wrappers{
		NameDB: conn.databaseName,
		ConnDB: conn.connection,
	}

	go func() {
		defer func() {
			_ = conn.connection.Disconnect(context.TODO())
		}()

		for {
			select {
			case <-conn.ctxRoute.Done():
				return

			case data := <-channels.GetChanInput():
				switch data.ObjectType {

				//Domain Object STIX
				case "attack-pattern":
					go ws.AddAttackPatternDO(data.Data, logging, counting)

				case "campaign":
					go ws.AddCampaignDO(data.Data, logging, counting)

				case "course-of-action":
					go ws.AddCourseOfActionDO(data.Data, logging, counting)

				case "grouping":
					go ws.AddGroupingDO(data.Data, logging, counting)

				case "identity":
					go ws.AddIdentityDO(data.Data, logging, counting)

				case "indicator":
					go ws.AddIndicatorDO(data.Data, logging, counting)

				case "infrastructure":
					go ws.AddInfrastructureDO(data.Data, logging, counting)

				case "intrusion-set":
					go ws.AddIntrusionSetDO(data.Data, logging, counting)

				case "location":
					go ws.AddLocationDO(data.Data, logging, counting)

				case "malware":
					go ws.AddMalwareDO(data.Data, logging, counting)

				case "malware-analysis":
					go ws.AddMalwareAnalysisDO(data.Data, logging, counting)

				case "note":
					go ws.AddNoteDO(data.Data, logging, counting)

				case "observed-data":
					go ws.AddObservedDataDO(data.Data, logging, counting)

				case "opinion":
					go ws.AddOpinionDO(data.Data, logging, counting)

				case "report":
					go ws.AddReportDO(data.Data, logging, counting)

				case "threat-actor":
					go ws.AddThreatActorDO(data.Data, logging, counting)

				case "tool":
					go ws.AddToolDO(data.Data, logging, counting)

				case "vulnerability":
					go ws.AddVulnerabilityDO(data.Data, logging, counting)

				//Cyber Observable Object STIX
				case "artifact":
					go ws.AddArtifactCO(data.Data, logging, counting)

				case "autonomous-system":
					go ws.AddAutonomousSystemCO(data.Data, logging, counting)

				case "directory":
					go ws.AddDirectoryCO(data.Data, logging, counting)

				case "domain-name":
					go ws.AddDomainNameCO(data.Data, logging, counting)

				case "email-addr":
					go ws.AddEmailAddressCO(data.Data, logging, counting)

				case "email-message":
					go ws.AddEmailMessageCO(data.Data, logging, counting)

				case "file":
					go ws.AddFileCO(data.Data, logging, counting)

				case "ipv4-addr":
					go ws.AddIPv4AddressCO(data.Data, logging, counting)

				case "ipv6-addr":
					go ws.AddIPv6AddressCO(data.Data, logging, counting)

				case "mac-addr":
					go ws.AddMACAddressCO(data.Data, logging, counting)

				case "mutex":
					go ws.AddMutexCO(data.Data, logging, counting)

				case "network-traffic":
					go ws.AddNetworkTrafficCO(data.Data, logging, counting)

				case "process":
					go ws.AddProcessCO(data.Data, logging, counting)

				case "software":
					go ws.AddSoftwareCO(data.Data, logging, counting)

				case "url":
					go ws.AddURLCO(data.Data, logging, counting)

				case "user-account":
					go ws.AddUserAccountCO(data.Data, logging, counting)

				case "windows-registry-key":
					go ws.AddWindowsRegistryKeyCO(data.Data, logging, counting)

				case "x509-certificate":
					go ws.AddX509CertificateCO(data.Data, logging, counting)
				}
			}
		}

		//		for data := range channels.GetChanInput() {}
	}()
}
