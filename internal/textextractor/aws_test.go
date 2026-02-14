package textextractor

// TODO: re-enable once aws access is provided to github actions
//func TestAWS(t *testing.T) {
//	data, err := os.ReadFile("./testdata/ashirt.png")
//	if err != nil {
//		t.Fatalf("unable to read test data: %v", err)
//	}
//	extractor := NewAWS()
//	extracted, err := extractor.ExtractText(context.Background(), data)
//	if err != nil {
//		t.Fatalf("unable to extract text: %v", err)
//	}
//
//	if extracted != "ASHIRT" {
//		t.Fatalf("extracted text does not match expected text, expected: %v, got: %v", "ashirt", extracted)
//	}
//}
