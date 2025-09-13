package utils

// func SendOTP(email, otpCode string) error {

// 	apiKey := os.Getenv("BREVO_API_KEY")

// 	if apiKey == "" {
// 		return fmt.Errorf("BREVO_API_KEY not set")
// 	}

// 	var ctx context.Context
// 	cfg := sendinblue.NewConfiguration()
// 	//Configure API key authorization: api-key
// 	cfg.AddDefaultHeader("api-key", "YOUR_API_KEY")
// 	//Configure API key authorization: partner-key
// 	cfg.AddDefaultHeader("partner-key", "YOUR_API_KEY")

// 	sib := sendinblue.NewAPIClient(cfg)
// 	result, resp, err := sib.AccountApi.GetAccount(ctx)
// 	if err != nil {
// 		fmt.Println("Error when calling AccountApi->get_account: ", err.Error())
// 		return
// 	}
// 	fmt.Println("GetAccount Object:", result, " GetAccount Response: ", resp)
// 	return

// }
