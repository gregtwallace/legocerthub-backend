package private_keys

import (
	"encoding/json"
	"errors"
	"legocerthub-backend/pkg/utils"
	"net/http"
)

// PostPayload is a struct for posting a new key
type PostPayload struct {
	ID             *int    `json:"id"`
	Name           *string `json:"name"`
	Description    *string `json:"description"`
	AlgorithmValue *string `json:"algorithm_value"`
	PemContent     *string `json:"pem"`
}

// PostNewKey creates a new private key and saves it to storage
func (service *Service) PostNewKey(w http.ResponseWriter, r *http.Request) {
	var payload PostPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		service.logger.Printf("keys: PostNew: failed to decode json -- err: %s", err)
		utils.WriteErrorJSON(w, err)
		return
	}

	/// do validation
	// id
	err = utils.IsIdNew(payload.ID)
	if err != nil {
		service.logger.Printf("keys: PostNew: invalid id -- err: %s", err)
		utils.WriteErrorJSON(w, err)
		return
	}
	// name
	err = utils.IsNameValid(*payload.Name)
	if err != nil {
		service.logger.Printf("keys: PostNew: invalid name -- err: %s", err)
		utils.WriteErrorJSON(w, err)
		return
	}
	// check db for duplicate name? probably unneeded as sql will error on insert

	/// key add method
	// error if no method specified
	if (payload.AlgorithmValue == nil) && (payload.PemContent == nil) {
		err = errors.New("keys: PostNew: no add method specified")
		service.logger.Println(err)
		utils.WriteErrorJSON(w, err)
		return
	}
	// error if more than one method specified
	if (payload.AlgorithmValue != nil && *payload.AlgorithmValue != "") && (payload.PemContent != nil && *payload.PemContent != "") {
		err = errors.New("keys: PostNew: multiple add methods specified")
		service.logger.Println(err)
		utils.WriteErrorJSON(w, err)
		return
	}
	// generate or verify the key
	// generate with algorithm, error if fails
	if payload.AlgorithmValue != nil && *payload.AlgorithmValue != "" {
		// must initialize to avoid invalid address
		payload.PemContent = new(string)
		*payload.PemContent, err = utils.GeneratePrivateKeyPem(*payload.AlgorithmValue)
		if err != nil {
			service.logger.Printf("keys: PostNew: failed to generate key pem -- err: %s", err)
			utils.WriteErrorJSON(w, err)
			return
		}
	} else if payload.PemContent != nil && *payload.PemContent != "" {
		// pem inputted - verify pem and determine algorithm
		// must initialize to avoid invalid address
		payload.AlgorithmValue = new(string)
		*payload.PemContent, *payload.AlgorithmValue, err = utils.ValidatePrivateKeyPem(*payload.PemContent)
		if err != nil {
			service.logger.Printf("keys: PostNew: failed to verify pem -- err: %s", err)
			utils.WriteErrorJSON(w, err)
			return
		}
	}
	///

	err = service.storage.PostNewKey(payload)
	if err != nil {
		service.logger.Printf("keys: PostNew: failed to write to db -- err: %s", err)
		utils.WriteErrorJSON(w, err)
		return
	}

	response := utils.JsonResp{
		OK: true,
	}
	err = utils.WriteJSON(w, http.StatusOK, response, "response")
	if err != nil {
		service.logger.Printf("keys: PostNew: write response json failed -- err: %s", err)
		utils.WriteErrorJSON(w, err)
		return
	}
}