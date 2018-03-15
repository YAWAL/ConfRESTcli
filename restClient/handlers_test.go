package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/YAWAL/ConfRESTcli/api"
	"github.com/YAWAL/ConfRESTcli/entities"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

type Mock struct {
	GrpcClient api.ConfigServiceClient
	configClient
}

type MockGrpcClient struct {
	grpc.ClientStream
	Results []*api.GetConfigResponce
}

func TestSelectType(t *testing.T) {

	mockConfigClient := &Mock{}

	mongoType := "mongodb"
	tsType := "tsconfig"
	tempType := "tempconfig"

	expMongoType := new(entities.Mongodb)
	expTsType := new(entities.Tsconfig)
	expTempType := new(entities.Tempconfig)

	result, err := mockConfigClient.selectType(mongoType)
	assert.Equal(t, expMongoType, result)
	result, err = mockConfigClient.selectType(tsType)
	assert.Equal(t, expTsType, result)
	result, err = mockConfigClient.selectType(tempType)
	assert.Equal(t, expTempType, result)

	expErr := errors.New("config does not exist")
	_, err = mockConfigClient.selectType("notExistingType")
	if assert.Error(t, err) {
		assert.Equal(t, expErr, err)
	}
}

//func (m MockGrpcClient) GetConfigByName(ctx context.Context, in *api.GetConfigByNameRequest, opts ...grpc.CallOption) (*api.GetConfigResponce, error) {
//	fmt.Println("from mock func")
//	return &api.GetConfigResponce{Config: []byte("1")}, errors.New("err ftom moc func")
//}
//func (m MockGrpcClient) GetConfigsByType(ctx context.Context, in *api.GetConfigsByTypeRequest, opts ...grpc.CallOption) (api.ConfigService_GetConfigsByTypeClient, error) {
//	return nil, nil
//}
//
//func (m MockGrpcClient) CreateConfig(ctx context.Context, in *api.Config, opts ...grpc.CallOption) (*api.Responce, error) {
//	return nil, nil
//}
//func (m MockGrpcClient) DeleteConfig(ctx context.Context, in *api.DeleteConfigRequest, opts ...grpc.CallOption) (*api.Responce, error) {
//	return nil, nil
//}
//func (m MockGrpcClient) UpdateConfig(ctx context.Context, in *api.Config, opts ...grpc.CallOption) (*api.Responce, error) {
//	return nil, nil
//}

func TestRetrieveConfig(t *testing.T) {
	mockConfigClient := Mock{}
	testMongoType := "mongodb"
	testTsType := "tsconfig"
	testTempType := "tempconfig"
	testName := "testName"
	notPresentedType := "errType"
	ctrl := gomock.NewController(t)
	mockGrpcClient := NewMockConfigServiceClient(ctrl)
	mongoConfig := entities.Mongodb{Domain: testName, Mongodb: true, Host: "testHost", Port: "testPort"}
	byteResMongo, err := json.Marshal(mongoConfig)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	tsConfig := entities.Tsconfig{Module: testName, Target: "testTarget", SourceMap: true, Excluding: 1}
	byteResTs, err := json.Marshal(tsConfig)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	tempConfig := entities.Tempconfig{RestApiRoot: testName, Host: "testHost", Port: "testPort", Remoting: "testRemoting", LegasyExplorer: true}
	byteResTemp, err := json.Marshal(tempConfig)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	mockGrpcClient.EXPECT().GetConfigByName(gomock.Any(), &api.GetConfigByNameRequest{ConfigName: testName, ConfigType: testMongoType}).Return(&api.GetConfigResponce{Config: byteResMongo}, nil)
	mockGrpcClient.EXPECT().GetConfigByName(gomock.Any(), &api.GetConfigByNameRequest{ConfigName: testName, ConfigType: testTsType}).Return(&api.GetConfigResponce{Config: byteResTs}, nil)
	mockGrpcClient.EXPECT().GetConfigByName(gomock.Any(), &api.GetConfigByNameRequest{ConfigName: testName, ConfigType: testTempType}).Return(&api.GetConfigResponce{Config: byteResTemp}, nil)
	mockGrpcClient.EXPECT().GetConfigByName(gomock.Any(), &api.GetConfigByNameRequest{ConfigName: testName, ConfigType: notPresentedType}).Return(nil, errors.New("config does not exist"))
	mockConfigClient.grpcClient = mockGrpcClient
	result, err := mockConfigClient.retrieveConfig(testName, testMongoType)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	assert.Equal(t, mongoConfig, result)
	result, err = mockConfigClient.retrieveConfig(testName, testTsType)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	assert.Equal(t, tsConfig, result)
	result, err = mockConfigClient.retrieveConfig(testName, testTempType)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	assert.Equal(t, tempConfig, result)
	result, err = mockConfigClient.retrieveConfig(testName, notPresentedType)
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("config does not exist"), err)
	}
}

func TestDeleteConfig(t *testing.T) {
	mockConfigClient := Mock{}
	testMongoType := "mongodb"
	testName := "testName"
	ctrl := gomock.NewController(t)
	mockGrpcClient := NewMockConfigServiceClient(ctrl)
	mockGrpcClient.EXPECT().DeleteConfig(gomock.Any(), &api.DeleteConfigRequest{ConfigType: testMongoType, ConfigName: testName}).Return(&api.Responce{Status: "OK"}, nil)
	mockConfigClient.grpcClient = mockGrpcClient
	c := gin.Context{}
	c.Params = append(c.Params, gin.Param{Key: "type", Value: testMongoType})
	c.Params = append(c.Params, gin.Param{Key: "name", Value: testName})

	result, err := mockConfigClient.deleteConfig(&c)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	assert.Equal(t, &api.Responce{Status: "OK"}, result)
}

//func (m Mock) selectType(cType string) (entities.ConfigInterface, error) {
//	fmt.Println("select")
//	return entities.Mongodb{}, nil
//}

//func TestUpdateConfig(t *testing.T) {
//	mockConfigClient := Mock{}
//	//mockConfigClient.selectType = func(cType string) { return entities.Mongodb{}, nil }
//	testMongoType := "mongodb"
//	testName := "testName"
//	ctrl := gomock.NewController(t)
//	mockGrpcClient := NewMockConfigServiceClient(ctrl)
//	mongoConfig := entities.Mongodb{Domain: testName, Mongodb: true, Host: "testHost", Port: "testPort"}
//	byteResMongo, err := json.Marshal(mongoConfig)
//	if err != nil {
//		t.Error("error during unit testing: ", err)
//	}
//	fmt.Println("-------------")
//	fmt.Println(mockConfigClient.selectType("tsconfig"))
//	fmt.Println("-------------")
//	mockGrpcClient.EXPECT().UpdateConfig(gomock.Any(), &api.Config{ConfigType: testMongoType, Config: byteResMongo}).Return(&api.Responce{Status: "OK"}, nil)
//	fmt.Println("******************")
//	fmt.Println(mockConfigClient.GrpcClient)
//	mockConfigClient.grpcClient = mockGrpcClient
//	fmt.Println(mockConfigClient.grpcClient)
//	c := gin.Context{}
//	c.Params = append(c.Params, gin.Param{Key: "type", Value: testMongoType})
//	c.Params = append(c.Params, gin.Param{Key: "name", Value: testName})
//	//c.Params = append(c.Params, gin.Param{Key: "domain", Value: testName})
//	//c.Params = append(c.Params, gin.Param{Key: "mongodb", Value: testMongoType})
//	//c.Params = append(c.Params, gin.Param{Key: "host", Value: testName})
//	//c.Params = append(c.Params, gin.Param{Key: "port", Value: testName})
//	c.Request.Header.Set("Content-Type", "application/json")
//	c.Request.PostForm.Set("domain", "true")
//	c.Request.PostForm.Set("host", "true")
//	c.Request.PostForm.Set("mongodb", "true")
//	c.Request.PostForm.Set("port", "true")
//	fmt.Println("******************")
//	fmt.Println("c", c)
//	result, err := mockConfigClient.updateConfig(&c)
//	if err != nil {
//		t.Error("error during unit testing: ", err)
//	}
//	assert.Equal(t, &api.Responce{Status: "OK"}, result)
//}

func TestRetrieveConfigsMongo(t *testing.T) {
	mockConfigClient := Mock{}
	testMongoType := "mongodb"
	testName := "testName"
	ctrl := gomock.NewController(t)
	stream := NewMockConfigService_GetConfigsByTypeClient(ctrl)
	mockGrpcClient := NewMockConfigServiceClient(ctrl)
	mongoConfig := entities.Mongodb{Domain: testName, Mongodb: true, Host: "testHost", Port: "testPort"}
	byteResMongo, err := json.Marshal(mongoConfig)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	stream.EXPECT().Recv().Return(&api.GetConfigResponce{Config: byteResMongo}, nil).Times(1)
	stream.EXPECT().Recv().Return(nil, io.EOF)
	mockGrpcClient.EXPECT().GetConfigsByType(gomock.Any(), &api.GetConfigsByTypeRequest{ConfigType: testMongoType}).Return(stream, nil)
	mockConfigClient.grpcClient = mockGrpcClient
	result, err := mockConfigClient.retrieveConfigs(&testMongoType)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	assert.ElementsMatch(t, []entities.Mongodb{mongoConfig}, result)
}

func TestRetrieveConfigsTs(t *testing.T) {
	mockConfigClient := Mock{}
	testTsType := "tsconfig"
	testName := "testName"
	ctrl := gomock.NewController(t)
	stream := NewMockConfigService_GetConfigsByTypeClient(ctrl)
	mockGrpcClient := NewMockConfigServiceClient(ctrl)
	tsConfig := entities.Tsconfig{Module: testName, Target: "testTarget", SourceMap: true, Excluding: 1}
	byteResTs, err := json.Marshal(tsConfig)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}

	stream.EXPECT().Recv().Return(&api.GetConfigResponce{Config: byteResTs}, nil).Times(1)
	stream.EXPECT().Recv().Return(nil, io.EOF)
	mockGrpcClient.EXPECT().GetConfigsByType(gomock.Any(), &api.GetConfigsByTypeRequest{ConfigType: testTsType}).Return(stream, nil)

	mockConfigClient.grpcClient = mockGrpcClient
	result, err := mockConfigClient.retrieveConfigs(&testTsType)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	assert.ElementsMatch(t, []entities.Tsconfig{tsConfig}, result)
}
func TestRetrieveConfigsTemp(t *testing.T) {
	mockConfigClient := Mock{}
	testTempType := "tempconfig"
	testName := "testName"
	ctrl := gomock.NewController(t)
	stream := NewMockConfigService_GetConfigsByTypeClient(ctrl)
	mockGrpcClient := NewMockConfigServiceClient(ctrl)
	tempConfig := entities.Tempconfig{RestApiRoot: testName, Host: "testHost", Port: "testPort", Remoting: "testRemoting", LegasyExplorer: true}
	byteResTemp, err := json.Marshal(tempConfig)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}

	stream.EXPECT().Recv().Return(&api.GetConfigResponce{Config: byteResTemp}, nil).Times(1)
	stream.EXPECT().Recv().Return(nil, io.EOF)
	mockGrpcClient.EXPECT().GetConfigsByType(gomock.Any(), &api.GetConfigsByTypeRequest{ConfigType: testTempType}).Return(stream, nil)

	mockConfigClient.grpcClient = mockGrpcClient
	result, err := mockConfigClient.retrieveConfigs(&testTempType)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	assert.ElementsMatch(t, []entities.Tempconfig{tempConfig}, result)
}
func TestRetrieveConfigsNotExisting(t *testing.T) {
	mockConfigClient := Mock{}
	notPresentedType := "errType"
	ctrl := gomock.NewController(t)
	stream := NewMockConfigService_GetConfigsByTypeClient(ctrl)
	mockGrpcClient := NewMockConfigServiceClient(ctrl)
	stream.EXPECT().Recv().Return(nil, nil).Times(1)
	stream.EXPECT().Recv().Return(nil, io.EOF)
	mockGrpcClient.EXPECT().GetConfigsByType(gomock.Any(), &api.GetConfigsByTypeRequest{ConfigType: notPresentedType}).Return(stream, errors.New("config does not exist"))
	mockConfigClient.grpcClient = mockGrpcClient
	_, err := mockConfigClient.retrieveConfigs(&notPresentedType)
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("config does not exist"), err)
	}
}

func TestCreateConfig(t *testing.T) {
	mockConfigClient := Mock{}
	testMongoType := "mongodb"
	testTsType := "tsconfig"
	testTempType := "tempconfig"
	testNotExistingType := "someType"
	ctrl := gomock.NewController(t)
	mockGrpcClient := NewMockConfigServiceClient(ctrl)

	mongoConfig := entities.Mongodb{Domain: testMongoType, Mongodb: true, Host: "testHost", Port: "testPort"}
	byteResMongo, err := json.Marshal(mongoConfig)
	tsConfig := entities.Tsconfig{Module: testTsType, Target: "testTarget", SourceMap: true, Excluding: 1}
	byteResTs, err := json.Marshal(tsConfig)
	tempConfig := entities.Tempconfig{RestApiRoot: testTempType, LegasyExplorer: true, Remoting: "tempRemoting", Port: "testPort", Host: "testHost"}
	byteResTemp, err := json.Marshal(tempConfig)
	mockGrpcClient.EXPECT().CreateConfig(gomock.Any(), &api.Config{ConfigType: testTsType, Config: byteResTs}).Return(&api.Responce{Status: "OK"}, nil)
	mockGrpcClient.EXPECT().CreateConfig(gomock.Any(), &api.Config{ConfigType: testMongoType, Config: byteResMongo}).Return(&api.Responce{Status: "OK"}, nil)
	mockGrpcClient.EXPECT().CreateConfig(gomock.Any(), &api.Config{ConfigType: testTempType, Config: byteResTemp}).Return(&api.Responce{Status: "OK"}, nil)
	mockGrpcClient.EXPECT().CreateConfig(gomock.Any(), &api.Config{ConfigType: testNotExistingType, Config: byteResMongo}).Return(nil, errors.New("not ex "))

	mockConfigClient.grpcClient = mockGrpcClient

	gin.SetMode(gin.TestMode)

	c := &gin.Context{}
	c.Params = append(c.Params, gin.Param{Key: "type", Value: testMongoType})
	form := url.Values{
		"domain":  {"mongodb"},
		"mongodb": {"true"},
		"host":    {"testHost"},
		"port":    {"testPort"},
	}
	body := bytes.NewBufferString(form.Encode())
	req, err := http.NewRequest("POST", "", body)
	c.Request = req
	c.Request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	result, err := mockConfigClient.createConfig(c)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	assert.Equal(t, &api.Responce{Status: "OK"}, result)

	c = &gin.Context{}
	c.Params = append(c.Params, gin.Param{Key: "type", Value: testTsType})
	form = url.Values{
		"module":    {"tsconfig"},
		"target":    {"testTarget"},
		"sourceMap": {"true"},
		"excluding": {"1"},
	}
	body = bytes.NewBufferString(form.Encode())
	req, err = http.NewRequest("POST", "", body)
	c.Request = req
	c.Request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	result, err = mockConfigClient.createConfig(c)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	assert.Equal(t, &api.Responce{Status: "OK"}, result)

	c = &gin.Context{}
	c.Params = append(c.Params, gin.Param{Key: "type", Value: testTempType})
	form = url.Values{
		"restApiRoot":    {"tempconfig"},
		"legasyExplorer": {"true"},
		"remoting":       {"tempRemoting"},
		"port":           {"testPort"},
		"host":           {"testHost"},
	}
	body = bytes.NewBufferString(form.Encode())
	req, err = http.NewRequest("POST", "", body)
	c.Request = req
	c.Request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	result, err = mockConfigClient.createConfig(c)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	assert.Equal(t, &api.Responce{Status: "OK"}, result)

	c = &gin.Context{}
	c.Params = append(c.Params, gin.Param{Key: "type", Value: testNotExistingType})
	form = url.Values{
		"domain":  {"mongodb"},
		"mongodb": {"true"},
		"host":    {"testHost"},
		"port":    {"testPort"},
	}
	body = bytes.NewBufferString(form.Encode())
	req, err = http.NewRequest("POST", "", body)
	c.Request = req
	c.Request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	_, err = mockConfigClient.createConfig(c)
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("config does not exist"), err)
	}
}

func TestUpdateConfig(t *testing.T) {
	mockConfigClient := Mock{}
	testMongoType := "mongodb"
	testTsType := "tsconfig"
	testTempType := "tempconfig"
	testNotExistingType := "someType"
	ctrl := gomock.NewController(t)
	mockGrpcClient := NewMockConfigServiceClient(ctrl)

	mongoConfig := entities.Mongodb{Domain: testMongoType, Mongodb: true, Host: "testHost", Port: "testPort"}
	byteResMongo, err := json.Marshal(mongoConfig)
	tsConfig := entities.Tsconfig{Module: testTsType, Target: "testTarget", SourceMap: true, Excluding: 1}
	byteResTs, err := json.Marshal(tsConfig)
	tempConfig := entities.Tempconfig{RestApiRoot: testTempType, LegasyExplorer: true, Remoting: "tempRemoting", Port: "testPort", Host: "testHost"}
	byteResTemp, err := json.Marshal(tempConfig)
	mockGrpcClient.EXPECT().UpdateConfig(gomock.Any(), &api.Config{ConfigType: testTsType, Config: byteResTs}).Return(&api.Responce{Status: "OK"}, nil)
	mockGrpcClient.EXPECT().UpdateConfig(gomock.Any(), &api.Config{ConfigType: testMongoType, Config: byteResMongo}).Return(&api.Responce{Status: "OK"}, nil)
	mockGrpcClient.EXPECT().UpdateConfig(gomock.Any(), &api.Config{ConfigType: testTempType, Config: byteResTemp}).Return(&api.Responce{Status: "OK"}, nil)
	mockGrpcClient.EXPECT().UpdateConfig(gomock.Any(), &api.Config{ConfigType: testNotExistingType, Config: byteResMongo}).Return(nil, errors.New("not ex "))

	mockConfigClient.grpcClient = mockGrpcClient

	gin.SetMode(gin.TestMode)

	c := &gin.Context{}
	c.Params = append(c.Params, gin.Param{Key: "type", Value: testMongoType})
	form := url.Values{
		"domain":  {"mongodb"},
		"mongodb": {"true"},
		"host":    {"testHost"},
		"port":    {"testPort"},
	}
	body := bytes.NewBufferString(form.Encode())
	req, err := http.NewRequest("POST", "", body)
	c.Request = req
	c.Request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	result, err := mockConfigClient.updateConfig(c)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	assert.Equal(t, &api.Responce{Status: "OK"}, result)

	c = &gin.Context{}
	c.Params = append(c.Params, gin.Param{Key: "type", Value: testTsType})
	form = url.Values{
		"module":    {"tsconfig"},
		"target":    {"testTarget"},
		"sourceMap": {"true"},
		"excluding": {"1"},
	}
	body = bytes.NewBufferString(form.Encode())
	req, err = http.NewRequest("POST", "", body)
	c.Request = req
	c.Request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	result, err = mockConfigClient.updateConfig(c)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	assert.Equal(t, &api.Responce{Status: "OK"}, result)

	c = &gin.Context{}
	c.Params = append(c.Params, gin.Param{Key: "type", Value: testTempType})
	form = url.Values{
		"restApiRoot":    {"tempconfig"},
		"legasyExplorer": {"true"},
		"remoting":       {"tempRemoting"},
		"port":           {"testPort"},
		"host":           {"testHost"},
	}
	body = bytes.NewBufferString(form.Encode())
	req, err = http.NewRequest("POST", "", body)
	c.Request = req
	c.Request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	result, err = mockConfigClient.updateConfig(c)
	if err != nil {
		t.Error("error during unit testing: ", err)
	}
	assert.Equal(t, &api.Responce{Status: "OK"}, result)

	c = &gin.Context{}
	c.Params = append(c.Params, gin.Param{Key: "type", Value: testNotExistingType})
	form = url.Values{
		"domain":  {"mongodb"},
		"mongodb": {"true"},
		"host":    {"testHost"},
		"port":    {"testPort"},
	}
	body = bytes.NewBufferString(form.Encode())
	req, err = http.NewRequest("POST", "", body)
	c.Request = req
	c.Request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	_, err = mockConfigClient.updateConfig(c)
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("config does not exist"), err)
	}
}
