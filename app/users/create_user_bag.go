package users

//func MakeCreateRequestBag(request io.Reader) (CreateRequestBag, error) {
//	body, err := io.ReadAll(request)
//
//	if err != nil {
//		slog.Error("Error reading request body: %v", err)
//		//http.Error(w, "Invalid request payload", http.StatusBadRequest)
//		return CreateRequestBag{}, err
//	}
//
//	if err = json.Unmarshal(body, &requestBag); err != nil {
//		slog.Error("Error decoding JSON: %v", err)
//		http.Error(w, "Invalid request payload", http.StatusBadRequest)
//		return
//	}
//
//	bag := CreateRequestBag{
//		RawRequest: body,
//	}
//
//	return bag, nil
//}
//
//func (receiver CreateRequestBag) GetRawRequest() []byte {
//	return receiver.RawRequest
//}
