package templates

//func PaymentEmail(params models.PaymentEmailStruct) string {
//	email := hermes.Email{
//		Body: hermes.Body{
//			Table: hermes.Table{
//				Data: [][]hermes.Entry{
//					{
//						{Key: "Premium", Value: params.Premium},
//						{Key: "Description", Value: params.Description},
//						{Key: "Price", Value: params.Price},
//					},
//				},
//				Columns: hermes.Columns{
//					// Custom style for each rows
//					CustomWidth: map[string]string{
//						"Premium":  "20%",
//						"Price": "15%",
//					},
//					CustomAlignment: map[string]string{
//						"Price": "right",
//					},
//				},
//			},
//		},
//	}
//
//	return StringParsedHTML(email)
//}
