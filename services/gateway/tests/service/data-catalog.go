package service

import (
	"context"
	"github.com/datafabric/gateway/protobuf"
)

type DataCatalogService struct {
	protobuf.UnimplementedDataCatalogServiceServer
}

func (s *DataCatalogService) Preview(ctx context.Context, in *protobuf.ReqId) (*protobuf.DataCatalogPreview, error) {
	if in.Id == "err-id" {
		res := &protobuf.DataCatalogPreview{
			Code:   "1234",
			ErrMsg: "not found data catalog",
		}
		return res, nil
	}
	res := &protobuf.DataCatalogPreview{
		Code: "200",
		Data: &protobuf.DataCatalogPreview_Data{
			DataPreview: &protobuf.DataCatalog{
				Id:          "data-id-01",
				Name:        "Test Data",
				Description: "Test Data Description",
				Status:      "CONNECTED",
				DataType:    "STRUCTURED",
				DataFormat:  "TABLE",
				Row:         10,
				Size:        20,
				DataLocation: []*protobuf.DataLocation{
					{
						StorageId:    "storageId",
						DatabaseName: "database",
						// DataPath:     "",
						TableName: "tableName",
						// FileName:     "",
						// Scope:        "",
						// SheetName:    "",
						// CallRange:    "",
						// Separator:    "",
						// BeginTime:    "",
						// EndTime:      "",
					},
				},
				DataRefine: &protobuf.DataRefine{
					Json:  "create data refine json",
					Query: `select * from data`,
				},
				DataStructure: []*protobuf.DataStructure{
					{
						Order:        1,
						Name:         "id",
						ColType:      "number",
						Length:       4,
						DefaultValue: "null",
						Description:  "id",
					},
					{
						Order:        2,
						Name:         "name",
						ColType:      "string",
						Length:       32,
						DefaultValue: "null",
						Description:  "name",
					},
					{
						Order:        3,
						Name:         "desc",
						ColType:      "string",
						Length:       32,
						DefaultValue: "null",
						Description:  "desc",
					},
				},
				Category: []*protobuf.Category{
					{Name: "category01"},
					{Name: "category02"},
				},
				SystemMeta: []*protobuf.Meta{
					{
						Key:   "meta",
						Value: "metaValue",
					},
					{
						Key:   "meta02",
						Value: "metaValue02",
					},
				},
				UserMeta: []*protobuf.Meta{
					{
						Key:   "user meta 01",
						Value: "user meta value 01",
					},
					{
						Key:   "user meta 02",
						Value: "user meta value 02",
					},
				},
				Tag: []string{"tag01", "tag02"},
				Permission: &protobuf.Permission{
					Read:  true,
					Write: true,
				},
				DownloadInfo: &protobuf.DownloadInfo{
					Status: 1,
					Url:    "http://localhost:8080/download/wawawawaw",
				},
				RatingAndComment: []*protobuf.RatingAndComment{
					{
						User: &protobuf.User{
							Id:       "id",
							Name:     "name",
							Nickname: "nickName",
							Phone:    "012341234",
							Email:    "01234@1234.123",
						},
						LastModifiedAt: &protobuf.DateTime{
							StrDateTime: "2023-01-01 00:00:00.123",
							UtcTime:     123456789123,
						},
						Rating:  10,
						Comment: "comment good",
					},
				},
				Statistics: &protobuf.DataCatalogStatistics{
					Time:            "2022-11-11 11:11:11.123",
					Id:              "data id",
					Name:            "data name",
					Access:          1,
					Bookmark:        2,
					Download:        3,
					Rating:          4.5,
					AvgResponseTime: 1.5,
				},
				Creator: &protobuf.User{
					Id:       "create user id",
					Name:     "creator name",
					Nickname: "creator nickname",
					Phone:    "012345679",
					Email:    "012345@012345.com",
				},
				CreatedAt: &protobuf.DateTime{
					StrDateTime: "2022-11-11 11:11:11.123",
					UtcTime:     123456789123,
				},
				LastModifier: &protobuf.User{
					Id:       "modify user id",
					Name:     "modify userName",
					Nickname: "modify user nickname",
					Phone:    "012345123451",
					Email:    "012345@129123.123",
				},
				LastModifiedAt: &protobuf.DateTime{
					StrDateTime: "2022-12-12 12:12:12.123",
					UtcTime:     123456789123,
				},
			},
		},
	}
	return res, nil

}
func (s *DataCatalogService) Default(ctx context.Context, in *protobuf.ReqId) (*protobuf.DataCatalogDefault, error) {
	if in.Id == "err-id" {
		res := &protobuf.DataCatalogDefault{
			Code:   "1234",
			ErrMsg: "not found data catalog",
		}
		return res, nil
	}
	res := &protobuf.DataCatalogDefault{
		Code: "200",
		Data: &protobuf.DataCatalogDefault_Data{
			DataCatalog: &protobuf.DataCatalog{
				Id:          "data-id-01",
				Name:        "Test Data",
				Description: "Test Data Description",
				Status:      "CONNECTED",
				DataType:    "STRUCTURED",
				DataFormat:  "TABLE",
				Row:         10,
				Size:        20,
				DataLocation: []*protobuf.DataLocation{
					{
						StorageId:    "storageId",
						DatabaseName: "database",
						// DataPath:     "",
						TableName: "tableName",
						// FileName:     "",
						// Scope:        "",
						// SheetName:    "",
						// CallRange:    "",
						// Separator:    "",
						// BeginTime:    "",
						// EndTime:      "",
					},
				},
				DataRefine: &protobuf.DataRefine{
					Json:  "create data refine json",
					Query: `select * from data`,
				},
				DataStructure: []*protobuf.DataStructure{
					{
						Order:        1,
						Name:         "id",
						ColType:      "number",
						Length:       4,
						DefaultValue: "null",
						Description:  "id",
					},
					{
						Order:        2,
						Name:         "name",
						ColType:      "string",
						Length:       32,
						DefaultValue: "null",
						Description:  "name",
					},
					{
						Order:        3,
						Name:         "desc",
						ColType:      "string",
						Length:       32,
						DefaultValue: "null",
						Description:  "desc",
					},
				},
				Category: []*protobuf.Category{
					{Name: "category01"},
					{Name: "category02"},
				},
				SystemMeta: []*protobuf.Meta{
					{
						Key:   "meta",
						Value: "metaValue",
					},
					{
						Key:   "meta02",
						Value: "metaValue02",
					},
				},
				UserMeta: []*protobuf.Meta{
					{
						Key:   "user meta 01",
						Value: "user meta value 01",
					},
					{
						Key:   "user meta 02",
						Value: "user meta value 02",
					},
				},
				Tag: []string{"tag01", "tag02"},
				Permission: &protobuf.Permission{
					Read:  true,
					Write: true,
				},
				DownloadInfo: &protobuf.DownloadInfo{
					Status: 1,
					Url:    "http://localhost:8080/download/wawawawaw",
				},
				RatingAndComment: []*protobuf.RatingAndComment{
					{
						User: &protobuf.User{
							Id:       "id",
							Name:     "name",
							Nickname: "nickName",
							Phone:    "012341234",
							Email:    "01234@1234.123",
						},
						LastModifiedAt: &protobuf.DateTime{
							StrDateTime: "2023-01-01 00:00:00.123",
							UtcTime:     123456789123,
						},
						Rating:  10,
						Comment: "comment good",
					},
				},
				Statistics: &protobuf.DataCatalogStatistics{
					Time:            "2022-11-11 11:11:11.123",
					Id:              "data id",
					Name:            "data name",
					Access:          1,
					Bookmark:        2,
					Download:        3,
					Rating:          4.5,
					AvgResponseTime: 1.5,
				},
				Creator: &protobuf.User{
					Id:       "create user id",
					Name:     "creator name",
					Nickname: "creator nickname",
					Phone:    "012345679",
					Email:    "012345@012345.com",
				},
				CreatedAt: &protobuf.DateTime{
					StrDateTime: "2022-11-11 11:11:11.123",
					UtcTime:     123456789123,
				},
				LastModifier: &protobuf.User{
					Id:       "modify user id",
					Name:     "modify userName",
					Nickname: "modify user nickname",
					Phone:    "012345123451",
					Email:    "012345@129123.123",
				},
				LastModifiedAt: &protobuf.DateTime{
					StrDateTime: "2022-12-12 12:12:12.123",
					UtcTime:     123456789123,
				},
			},
		},
	}
	return res, nil
}

func (s *DataCatalogService) UserMetadata(ctx context.Context, in *protobuf.ReqMetaUpdate) (*protobuf.CommonResponse, error) {
	if in.Id == "0" {
		res := &protobuf.CommonResponse{
			Code:   "200",
			ErrMsg: "Not Found Data Catalog for update user metadata",
		}
		return res, nil
	}
	res := &protobuf.CommonResponse{
		Code: "200",
	}
	return res, nil
}
func (s *DataCatalogService) Tag(ctx context.Context, in *protobuf.ReqTagUpdate) (*protobuf.CommonResponse, error) {
	if in.Id == "0" {
		res := &protobuf.CommonResponse{
			Code:   "200",
			ErrMsg: "Not Found Data Catalog for tag",
		}
		return res, nil
	}
	res := &protobuf.CommonResponse{
		Code: "200",
	}
	return res, nil
}
func (s *DataCatalogService) DownloadRequest(ctx context.Context, in *protobuf.ReqId) (*protobuf.CommonResponse, error) {
	if in.Id == "0" {
		res := &protobuf.CommonResponse{
			Code:   "200",
			ErrMsg: "Not Found Data Catalog for download request",
		}
		return res, nil
	}
	res := &protobuf.CommonResponse{
		Code: "200",
	}
	return res, nil
}
func (s *DataCatalogService) AddComment(ctx context.Context, in *protobuf.ReqRatingAndComment) (*protobuf.CommonResponse, error) {
	if in.Id == "0" {
		res := &protobuf.CommonResponse{
			Code:   "200",
			ErrMsg: "Not Found Data Catalog for add comment",
		}
		return res, nil
	}
	res := &protobuf.CommonResponse{
		Code: "200",
	}
	return res, nil
}
func (s *DataCatalogService) UpdateComment(ctx context.Context, in *protobuf.ReqRatingAndComment) (*protobuf.CommonResponse, error) {
	if in.Id == "0" {
		res := &protobuf.CommonResponse{
			Code:   "200",
			ErrMsg: "Not Found Data Catalog for Update Comment",
		}
		return res, nil
	}
	res := &protobuf.CommonResponse{
		Code: "200",
	}
	return res, nil
}
func (s *DataCatalogService) DelComment(ctx context.Context, in *protobuf.ReqId) (*protobuf.CommonResponse, error) {
	if in.Id == "0" {
		res := &protobuf.CommonResponse{
			Code:   "200",
			ErrMsg: "Not Found Data Catalog for Delete Comment",
		}
		return res, nil
	}
	res := &protobuf.CommonResponse{
		Code: "200",
	}
	return res, nil
}

func (s *DataCatalogService) AllDataSummary(ctx context.Context, in *protobuf.DataCatalogSearch) (*protobuf.ResDataCatalogs, error) {
	res := &protobuf.ResDataCatalogs{
		Code: "200",
		Data: &protobuf.ResDataCatalogs_Data{
			Pageable: &protobuf.Pageable{
				Page: &protobuf.Page{
					Size:       in.Pageable.Page.Size,
					TotalSize:  1000,
					SelectPage: in.Pageable.Page.SelectPage,
					TotalPage:  10,
				},
				Sort: []*protobuf.Sort{
					{
						Order:     1,
						Field:     "name",
						Direction: protobuf.Direction_ASC,
					},
					{
						Order:     2,
						Field:     "createdAt",
						Direction: protobuf.Direction_DESC,
					},
				},
			},
			DataCatalogs: []*protobuf.DataCatalog{
				{
					Id:          "data-id-01",
					Name:        "Test Data",
					Description: "Test Data Description",
					Status:      "CONNECTED",
					DataType:    "STRUCTURED",
					DataFormat:  "TABLE",
					Row:         10,
					Size:        20,
					DataLocation: []*protobuf.DataLocation{
						{
							StorageId:    "storageId",
							DatabaseName: "database",
							// DataPath:     "",
							TableName: "tableName",
							// FileName:     "",
							// Scope:        "",
							// SheetName:    "",
							// CallRange:    "",
							// Separator:    "",
							// BeginTime:    "",
							// EndTime:      "",
						},
					},
					DataRefine: &protobuf.DataRefine{
						Json:  "create data refine json",
						Query: `select * from data`,
					},
					DataStructure: []*protobuf.DataStructure{
						{
							Order:        1,
							Name:         "id",
							ColType:      "number",
							Length:       4,
							DefaultValue: "null",
							Description:  "id",
						},
						{
							Order:        2,
							Name:         "name",
							ColType:      "string",
							Length:       32,
							DefaultValue: "null",
							Description:  "name",
						},
						{
							Order:        3,
							Name:         "desc",
							ColType:      "string",
							Length:       32,
							DefaultValue: "null",
							Description:  "desc",
						},
					},
					Category: []*protobuf.Category{
						{Name: "category01"},
						{Name: "category02"},
					},
					SystemMeta: []*protobuf.Meta{
						{
							Key:   "meta",
							Value: "metaValue",
						},
						{
							Key:   "meta02",
							Value: "metaValue02",
						},
					},
					UserMeta: []*protobuf.Meta{
						{
							Key:   "user meta 01",
							Value: "user meta value 01",
						},
						{
							Key:   "user meta 02",
							Value: "user meta value 02",
						},
					},
					Tag: []string{"tag01", "tag02"},
					Permission: &protobuf.Permission{
						Read:  true,
						Write: true,
					},
					DownloadInfo: &protobuf.DownloadInfo{
						Status: 1,
						Url:    "http://localhost:8080/download/wawawawaw",
					},
					RatingAndComment: []*protobuf.RatingAndComment{
						{
							User: &protobuf.User{
								Id:       "id",
								Name:     "name",
								Nickname: "nickName",
								Phone:    "012341234",
								Email:    "01234@1234.123",
							},
							LastModifiedAt: &protobuf.DateTime{
								StrDateTime: "2023-01-01 00:00:00.123",
								UtcTime:     123456789123,
							},
							Rating:  10,
							Comment: "comment good",
						},
					},
					Statistics: &protobuf.DataCatalogStatistics{
						Time:            "2022-11-11 11:11:11.123",
						Id:              "data id",
						Name:            "data name",
						Access:          1,
						Bookmark:        2,
						Download:        3,
						Rating:          4.5,
						AvgResponseTime: 1.5,
					},
					Creator: &protobuf.User{
						Id:       "create user id",
						Name:     "creator name",
						Nickname: "creator nickname",
						Phone:    "012345679",
						Email:    "012345@012345.com",
					},
					CreatedAt: &protobuf.DateTime{
						StrDateTime: "2022-11-11 11:11:11.123",
						UtcTime:     123456789123,
					},
					LastModifier: &protobuf.User{
						Id:       "modify user id",
						Name:     "modify userName",
						Nickname: "modify user nickname",
						Phone:    "012345123451",
						Email:    "012345@129123.123",
					},
					LastModifiedAt: &protobuf.DateTime{
						StrDateTime: "2022-12-12 12:12:12.123",
						UtcTime:     123456789123,
					},
				},
				{
					Id:          "data-id-02",
					Name:        "Test Data 2",
					Description: "Test Data Description 2",
					Status:      "DISCONNECTED",
					DataType:    "STRUCTURED",
					DataFormat:  "TABLE",
					Row:         10,
					Size:        20,
					DataLocation: []*protobuf.DataLocation{
						{
							StorageId:    "storageId 02",
							DatabaseName: "database-test",
							// DataPath:     "",
							TableName: "tableName-test",
							// FileName:     "",
							// Scope:        "",
							// SheetName:    "",
							// CallRange:    "",
							// Separator:    "",
							// BeginTime:    "",
							// EndTime:      "",
						},
					},
					DataRefine: &protobuf.DataRefine{
						Json:  "create data refine json",
						Query: `select * from tableName-test where id = 1`,
					},
					DataStructure: []*protobuf.DataStructure{
						{
							Order:        1,
							Name:         "id",
							ColType:      "number",
							Length:       4,
							DefaultValue: "null",
							Description:  "id",
						},
					},
					SystemMeta: []*protobuf.Meta{
						{
							Key:   "meta",
							Value: "metaValue",
						},
						{
							Key:   "meta02",
							Value: "metaValue02",
						},
					},
					UserMeta: []*protobuf.Meta{
						{
							Key:   "user meta 01",
							Value: "user meta value 01",
						},
						{
							Key:   "user meta 02",
							Value: "user meta value 02",
						},
					},
					Tag: []string{"tag01", "tag02"},
					Permission: &protobuf.Permission{
						Read:  true,
						Write: true,
					},
					DownloadInfo: &protobuf.DownloadInfo{
						Status: protobuf.DownloadInfo_READY,
					},
					RatingAndComment: []*protobuf.RatingAndComment{
						{
							User: &protobuf.User{
								Id:       "id",
								Name:     "name",
								Nickname: "nickName",
								Phone:    "012341234",
								Email:    "01234@1234.123",
							},
							LastModifiedAt: &protobuf.DateTime{
								StrDateTime: "2023-01-01 00:00:00.123",
								UtcTime:     123456789123,
							},
							Rating:  10,
							Comment: "comment good",
						},
						{
							User: &protobuf.User{
								Id:       "id",
								Name:     "name",
								Nickname: "nickName",
								Phone:    "012341234",
								Email:    "01234@1234.123",
							},
							LastModifiedAt: &protobuf.DateTime{
								StrDateTime: "2023-01-01 00:00:00.123",
								UtcTime:     123456789123,
							},
							Rating:  10,
							Comment: "comment good",
						},
					},
					Statistics: &protobuf.DataCatalogStatistics{
						Time:            "2022-11-11 11:11:11.123",
						Id:              "data id",
						Name:            "data name",
						Access:          1,
						Bookmark:        2,
						Download:        3,
						Rating:          4.5,
						AvgResponseTime: 1.5,
					},
					Creator: &protobuf.User{
						Id:       "create user id",
						Name:     "creator name",
						Nickname: "creator nickname",
						Phone:    "012345679",
						Email:    "012345@012345.com",
					},
					CreatedAt: &protobuf.DateTime{
						StrDateTime: "2022-11-11 11:11:11.123",
						UtcTime:     123456789123,
					},
					LastModifier: &protobuf.User{
						Id:       "modify user id",
						Name:     "modify userName",
						Nickname: "modify user nickname",
						Phone:    "012345123451",
						Email:    "012345@129123.123",
					},
					LastModifiedAt: &protobuf.DateTime{
						StrDateTime: "2022-12-12 12:12:12.123",
						UtcTime:     123456789123,
					},
				},
			},
		},
	}
	return res, nil
}
func (s *DataCatalogService) AllData(ctx context.Context, in *protobuf.DataCatalogSearch) (*protobuf.ResDataCatalogs, error) {
	res := &protobuf.ResDataCatalogs{
		Code: "200",
		Data: &protobuf.ResDataCatalogs_Data{
			Pageable: &protobuf.Pageable{
				Page: &protobuf.Page{
					Size:       in.Pageable.Page.Size,
					TotalSize:  1000,
					SelectPage: in.Pageable.Page.SelectPage,
					TotalPage:  10,
				},
				Sort: []*protobuf.Sort{
					{
						Order:     1,
						Field:     "name",
						Direction: protobuf.Direction_ASC,
					},
					{
						Order:     2,
						Field:     "createdAt",
						Direction: protobuf.Direction_DESC,
					},
				},
			},
			DataCatalogs: []*protobuf.DataCatalog{
				{
					Id:          "data-id-01",
					Name:        "Test Data",
					Description: "Test Data Description",
					Status:      "CONNECTED",
					DataType:    "STRUCTURED",
					DataFormat:  "TABLE",
					Row:         10,
					Size:        20,
					DataLocation: []*protobuf.DataLocation{
						{
							StorageId:    "storageId",
							DatabaseName: "database",
							// DataPath:     "",
							TableName: "tableName",
							// FileName:     "",
							// Scope:        "",
							// SheetName:    "",
							// CallRange:    "",
							// Separator:    "",
							// BeginTime:    "",
							// EndTime:      "",
						},
					},
					DataRefine: &protobuf.DataRefine{
						Json:  "create data refine json",
						Query: `select * from data`,
					},
					DataStructure: []*protobuf.DataStructure{
						{
							Order:        1,
							Name:         "id",
							ColType:      "number",
							Length:       4,
							DefaultValue: "null",
							Description:  "id",
						},
						{
							Order:        2,
							Name:         "name",
							ColType:      "string",
							Length:       32,
							DefaultValue: "null",
							Description:  "name",
						},
						{
							Order:        3,
							Name:         "desc",
							ColType:      "string",
							Length:       32,
							DefaultValue: "null",
							Description:  "desc",
						},
					},
					Category: []*protobuf.Category{
						{Name: "category01"},
						{Name: "category02"},
					},
					SystemMeta: []*protobuf.Meta{
						{
							Key:   "meta",
							Value: "metaValue",
						},
						{
							Key:   "meta02",
							Value: "metaValue02",
						},
					},
					UserMeta: []*protobuf.Meta{
						{
							Key:   "user meta 01",
							Value: "user meta value 01",
						},
						{
							Key:   "user meta 02",
							Value: "user meta value 02",
						},
					},
					Tag: []string{"tag01", "tag02"},
					Permission: &protobuf.Permission{
						Read:  true,
						Write: true,
					},
					DownloadInfo: &protobuf.DownloadInfo{
						Status: 1,
						Url:    "http://localhost:8080/download/wawawawaw",
					},
					RatingAndComment: []*protobuf.RatingAndComment{
						{
							User: &protobuf.User{
								Id:       "id",
								Name:     "name",
								Nickname: "nickName",
								Phone:    "012341234",
								Email:    "01234@1234.123",
							},
							LastModifiedAt: &protobuf.DateTime{
								StrDateTime: "2023-01-01 00:00:00.123",
								UtcTime:     123456789123,
							},
							Rating:  10,
							Comment: "comment good",
						},
					},
					Statistics: &protobuf.DataCatalogStatistics{
						Time:            "2022-11-11 11:11:11.123",
						Id:              "data id",
						Name:            "data name",
						Access:          1,
						Bookmark:        2,
						Download:        3,
						Rating:          4.5,
						AvgResponseTime: 1.5,
					},
					Creator: &protobuf.User{
						Id:       "create user id",
						Name:     "creator name",
						Nickname: "creator nickname",
						Phone:    "012345679",
						Email:    "012345@012345.com",
					},
					CreatedAt: &protobuf.DateTime{
						StrDateTime: "2022-11-11 11:11:11.123",
						UtcTime:     123456789123,
					},
					LastModifier: &protobuf.User{
						Id:       "modify user id",
						Name:     "modify userName",
						Nickname: "modify user nickname",
						Phone:    "012345123451",
						Email:    "012345@129123.123",
					},
					LastModifiedAt: &protobuf.DateTime{
						StrDateTime: "2022-12-12 12:12:12.123",
						UtcTime:     123456789123,
					},
				},
				{
					Id:          "data-id-02",
					Name:        "Test Data 2",
					Description: "Test Data Description 2",
					Status:      "DISCONNECTED",
					DataType:    "STRUCTURED",
					DataFormat:  "TABLE",
					Row:         10,
					Size:        20,
					DataLocation: []*protobuf.DataLocation{
						{
							StorageId:    "storageId 02",
							DatabaseName: "database-test",
							// DataPath:     "",
							TableName: "tableName-test",
							// FileName:     "",
							// Scope:        "",
							// SheetName:    "",
							// CallRange:    "",
							// Separator:    "",
							// BeginTime:    "",
							// EndTime:      "",
						},
					},
					DataRefine: &protobuf.DataRefine{
						Json:  "create data refine json",
						Query: `select * from tableName-test where id = 1`,
					},
					DataStructure: []*protobuf.DataStructure{
						{
							Order:        1,
							Name:         "id",
							ColType:      "number",
							Length:       4,
							DefaultValue: "null",
							Description:  "id",
						},
					},
					SystemMeta: []*protobuf.Meta{
						{
							Key:   "meta",
							Value: "metaValue",
						},
						{
							Key:   "meta02",
							Value: "metaValue02",
						},
					},
					UserMeta: []*protobuf.Meta{
						{
							Key:   "user meta 01",
							Value: "user meta value 01",
						},
						{
							Key:   "user meta 02",
							Value: "user meta value 02",
						},
					},
					Tag: []string{"tag01", "tag02"},
					Permission: &protobuf.Permission{
						Read:  true,
						Write: true,
					},
					DownloadInfo: &protobuf.DownloadInfo{
						Status: protobuf.DownloadInfo_READY,
					},
					RatingAndComment: []*protobuf.RatingAndComment{
						{
							User: &protobuf.User{
								Id:       "id",
								Name:     "name",
								Nickname: "nickName",
								Phone:    "012341234",
								Email:    "01234@1234.123",
							},
							LastModifiedAt: &protobuf.DateTime{
								StrDateTime: "2023-01-01 00:00:00.123",
								UtcTime:     123456789123,
							},
							Rating:  10,
							Comment: "comment good",
						},
						{
							User: &protobuf.User{
								Id:       "id",
								Name:     "name",
								Nickname: "nickName",
								Phone:    "012341234",
								Email:    "01234@1234.123",
							},
							LastModifiedAt: &protobuf.DateTime{
								StrDateTime: "2023-01-01 00:00:00.123",
								UtcTime:     123456789123,
							},
							Rating:  10,
							Comment: "comment good",
						},
					},
					Statistics: &protobuf.DataCatalogStatistics{
						Time:            "2022-11-11 11:11:11.123",
						Id:              "data id",
						Name:            "data name",
						Access:          1,
						Bookmark:        2,
						Download:        3,
						Rating:          4.5,
						AvgResponseTime: 1.5,
					},
					Creator: &protobuf.User{
						Id:       "create user id",
						Name:     "creator name",
						Nickname: "creator nickname",
						Phone:    "012345679",
						Email:    "012345@012345.com",
					},
					CreatedAt: &protobuf.DateTime{
						StrDateTime: "2022-11-11 11:11:11.123",
						UtcTime:     123456789123,
					},
					LastModifier: &protobuf.User{
						Id:       "modify user id",
						Name:     "modify userName",
						Nickname: "modify user nickname",
						Phone:    "012345123451",
						Email:    "012345@129123.123",
					},
					LastModifiedAt: &protobuf.DateTime{
						StrDateTime: "2022-12-12 12:12:12.123",
						UtcTime:     123456789123,
					},
				},
			},
		},
	}
	return res, nil
}
