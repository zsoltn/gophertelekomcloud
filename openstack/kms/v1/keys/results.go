package keys

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type commonResult struct {
	golangsdk.Result
}

// Key contains all the information associated with a CMK.
type Key struct {
	// Current ID of a CMK
	KeyID string `json:"key_id"`
	// ID of a user domain for the key.
	DomainID string `json:"domain_id"`
	// Alias of a CMK
	KeyAlias string `json:"key_alias"`
	// Region where a CMK resides
	Realm string `json:"realm"`
	// Description of a CMK
	KeyDescription string `json:"key_description"`
	// Creation time (time stamp) of a CMK
	CreationDate string `json:"creation_date"`
	// Scheduled deletion time (time stamp) of a CMK
	ScheduledDeletionDate string `json:"scheduled_deletion_date"`
	// State of a CMK
	KeyState string `json:"key_state"`
	// Identification of a Master Key. The value 1 indicates a Default
	// Master Key, and the value 0 indicates a CMK
	DefaultKeyFlag string `json:"default_key_flag"`
	// Expiration time
	ExpirationTime string `json:"expiration_time"`
	// Origin of a CMK. The default value is kms. The following values
	// are enumerated: kms indicates that the CMK material is generated by KMS.
	Origin string `json:"origin"`
}

type ListKey struct {
	Keys       []string `json:"keys"`
	KeyDetails []Key    `json:"key_details"`
	NextMarker string   `json:"next_marker"`
	Truncated  string   `json:"truncated"`
}

type DataKey struct {
	// Current ID of a CMK
	KeyID      string `json:"key_id"`
	PlainText  string `json:"plain_text"`
	CipherText string `json:"cipher_text"`
}

type EncryptDEK struct {
	// Current ID of a CMK
	KeyID         string `json:"key_id"`
	DataKeyLength string `json:"datakey_length"`
	CipherText    string `json:"cipher_text"`
}

type UpdateKeyState struct {
	// Current ID of a CMK
	KeyID    string `json:"key_id"`
	KeyState string `json:"key_state"`
}

// CreateResult contains the response body and error from a Create request.
type CreateResult struct {
	commonResult
}

// GetResult contains the response body and error from a Get request.
type GetResult struct {
	commonResult
}

// DeleteResult contains the response body and error from a Delete request.
type DeleteResult struct {
	commonResult
}

// UpdateAliasResult contains the response body and error from a UpdateAlias request.
type UpdateAliasResult struct {
	commonResult
}

// UpdateDesResult contains the response body and error from a UpdateDes request.
type UpdateDesResult struct {
	commonResult
}

type DataEncryptResult struct {
	commonResult
}

type EncryptDEKResult struct {
	commonResult
}

type ExtractUpdateKeyStateResult struct {
	commonResult
}

type ListResult struct {
	commonResult
}

func (r commonResult) ExtractListKey() (*ListKey, error) {
	var s *ListKey
	err := r.ExtractInto(&s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r commonResult) Extract() (*Key, error) {
	var s *Key
	err := r.ExtractInto(&s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r commonResult) ExtractKeyInfo() (*Key, error) {
	var s Key
	err := r.ExtractKeyInfoInto(&s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r commonResult) ExtractKeyInfoInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "key_info")
}

func (r commonResult) ExtractDataKey() (*DataKey, error) {
	var s *DataKey
	err := r.ExtractInto(&s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r commonResult) ExtractEncryptDEK() (*EncryptDEK, error) {
	var s *EncryptDEK
	err := r.ExtractInto(&s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

type KeyPage struct {
	pagination.LinkedPageBase
}

func (r KeyPage) IsEmpty() (bool, error) {
	s, err := ExtractKeys(r)
	return len(s) == 0, err
}

func ExtractKeys(r pagination.Page) ([]Key, error) {
	var s struct {
		Keys []Key `json:"keys"`
	}
	err := (r.(KeyPage)).ExtractInto(&s)
	if err != nil {
		return nil, err
	}
	return s.Keys, nil
}
