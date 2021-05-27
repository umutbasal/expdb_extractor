package extract

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// UnmarshalJSON implements json.Unmarshaler.
func (j *ModifiedScopeType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ModifiedScopeType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ModifiedScopeType, v)
	}
	*j = ModifiedScopeType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AttackComplexityType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AttackComplexityType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AttackComplexityType_1, v)
	}
	*j = AttackComplexityType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DefCpeName) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["cpe23Uri"]; !ok || v == nil {
		return fmt.Errorf("field cpe23Uri: required")
	}
	type Plain DefCpeName
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = DefCpeName(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *UserInteractionType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_UserInteractionType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_UserInteractionType, v)
	}
	*j = UserInteractionType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DefCpeMatch) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["cpe23Uri"]; !ok || v == nil {
		return fmt.Errorf("field cpe23Uri: required")
	}
	if v, ok := raw["vulnerable"]; !ok || v == nil {
		return fmt.Errorf("field vulnerable: required")
	}
	type Plain DefCpeMatch
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = DefCpeMatch(plain)
	return nil
}

const AttackComplexityType_1_HIGH AttackComplexityType_1 = "HIGH"

type AccessComplexityType string

const AttackComplexityType_1_LOW AttackComplexityType_1 = "LOW"

// UnmarshalJSON implements json.Unmarshaler.
func (j *AccessComplexityType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AccessComplexityType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AccessComplexityType, v)
	}
	*j = AccessComplexityType(v)
	return nil
}

const AccessComplexityTypeHIGH AccessComplexityType = "HIGH"
const AccessComplexityTypeMEDIUM AccessComplexityType = "MEDIUM"
const AccessComplexityTypeLOW AccessComplexityType = "LOW"

type AccessVectorType string

type AttackVectorType_1 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *AccessVectorType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AccessVectorType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AccessVectorType, v)
	}
	*j = AccessVectorType(v)
	return nil
}

const AccessVectorTypeNETWORK AccessVectorType = "NETWORK"
const AccessVectorTypeADJACENTNETWORK AccessVectorType = "ADJACENT_NETWORK"
const AccessVectorTypeLOCAL AccessVectorType = "LOCAL"

type AuthenticationType string

// UnmarshalJSON implements json.Unmarshaler.
func (j *ModifiedPrivilegesRequiredType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ModifiedPrivilegesRequiredType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ModifiedPrivilegesRequiredType_1, v)
	}
	*j = ModifiedPrivilegesRequiredType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuthenticationType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AuthenticationType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AuthenticationType, v)
	}
	*j = AuthenticationType(v)
	return nil
}

const AuthenticationTypeMULTIPLE AuthenticationType = "MULTIPLE"

type AttackComplexityType string

const AuthenticationTypeNONE AuthenticationType = "NONE"
const AttackComplexityTypeHIGH AttackComplexityType = "HIGH"

// UnmarshalJSON implements json.Unmarshaler.
func (j *AttackVectorType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AttackVectorType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AttackVectorType_1, v)
	}
	*j = AttackVectorType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CiaRequirementType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CiaRequirementType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CiaRequirementType, v)
	}
	*j = CiaRequirementType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CvssV20Json) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["baseScore"]; !ok || v == nil {
		return fmt.Errorf("field baseScore: required")
	}
	if v, ok := raw["vectorString"]; !ok || v == nil {
		return fmt.Errorf("field vectorString: required")
	}
	if v, ok := raw["version"]; !ok || v == nil {
		return fmt.Errorf("field version: required")
	}
	type Plain CvssV20Json
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CvssV20Json(plain)
	return nil
}

const AttackComplexityTypeLOW AttackComplexityType = "LOW"

type AttackVectorType string

// UnmarshalJSON implements json.Unmarshaler.
func (j *CvssV20JsonVersion) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CvssV20JsonVersion {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CvssV20JsonVersion, v)
	}
	*j = CvssV20JsonVersion(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AttackVectorType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AttackVectorType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AttackVectorType, v)
	}
	*j = AttackVectorType(v)
	return nil
}

const AttackVectorType_1_NETWORK AttackVectorType_1 = "NETWORK"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CiaType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CiaType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CiaType, v)
	}
	*j = CiaType(v)
	return nil
}

const AttackVectorTypeNETWORK AttackVectorType = "NETWORK"
const AttackVectorTypeADJACENTNETWORK AttackVectorType = "ADJACENT_NETWORK"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CiaType_7) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CiaType_7 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CiaType_7, v)
	}
	*j = CiaType_7(v)
	return nil
}

type AttackComplexityType_1 string

const AttackVectorType_1_ADJACENTNETWORK AttackVectorType_1 = "ADJACENT_NETWORK"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CollateralDamagePotentialType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CollateralDamagePotentialType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CollateralDamagePotentialType, v)
	}
	*j = CollateralDamagePotentialType(v)
	return nil
}

const AttackVectorType_1_LOCAL AttackVectorType_1 = "LOCAL"

// UnmarshalJSON implements json.Unmarshaler.
func (j *SeverityType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_SeverityType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_SeverityType, v)
	}
	*j = SeverityType(v)
	return nil
}

const AttackVectorType_1_PHYSICAL AttackVectorType_1 = "PHYSICAL"

// UnmarshalJSON implements json.Unmarshaler.
func (j *TargetDistributionType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_TargetDistributionType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_TargetDistributionType_1, v)
	}
	*j = TargetDistributionType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CiaType_5) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CiaType_5 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CiaType_5, v)
	}
	*j = CiaType_5(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SeverityType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_SeverityType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_SeverityType_3, v)
	}
	*j = SeverityType_3(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ModifiedUserInteractionType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ModifiedUserInteractionType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ModifiedUserInteractionType_1, v)
	}
	*j = ModifiedUserInteractionType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CvssV3XJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["baseScore"]; !ok || v == nil {
		return fmt.Errorf("field baseScore: required")
	}
	if v, ok := raw["baseSeverity"]; !ok || v == nil {
		return fmt.Errorf("field baseSeverity: required")
	}
	if v, ok := raw["vectorString"]; !ok || v == nil {
		return fmt.Errorf("field vectorString: required")
	}
	if v, ok := raw["version"]; !ok || v == nil {
		return fmt.Errorf("field version: required")
	}
	type Plain CvssV3XJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CvssV3XJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ExploitabilityType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ExploitabilityType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ExploitabilityType, v)
	}
	*j = ExploitabilityType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ScopeType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ScopeType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ScopeType, v)
	}
	*j = ScopeType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CiaRequirementType_7) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CiaRequirementType_7 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CiaRequirementType_7, v)
	}
	*j = CiaRequirementType_7(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ReportConfidenceType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ReportConfidenceType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ReportConfidenceType_1, v)
	}
	*j = ReportConfidenceType_1(v)
	return nil
}

const AttackVectorTypeLOCAL AttackVectorType = "LOCAL"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CiaRequirementType_5) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CiaRequirementType_5 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CiaRequirementType_5, v)
	}
	*j = CiaRequirementType_5(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ModifiedAttackComplexityType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ModifiedAttackComplexityType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ModifiedAttackComplexityType_1, v)
	}
	*j = ModifiedAttackComplexityType_1(v)
	return nil
}

const AttackVectorTypePHYSICAL AttackVectorType = "PHYSICAL"

// UnmarshalJSON implements json.Unmarshaler.
func (j *RemediationLevelType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_RemediationLevelType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_RemediationLevelType, v)
	}
	*j = RemediationLevelType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RemediationLevelType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_RemediationLevelType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_RemediationLevelType_1, v)
	}
	*j = RemediationLevelType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CiaRequirementType_4) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CiaRequirementType_4 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CiaRequirementType_4, v)
	}
	*j = CiaRequirementType_4(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CiaRequirementType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CiaRequirementType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CiaRequirementType_3, v)
	}
	*j = CiaRequirementType_3(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RemediationLevelType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_RemediationLevelType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_RemediationLevelType_2, v)
	}
	*j = RemediationLevelType_2(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SeverityType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_SeverityType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_SeverityType_1, v)
	}
	*j = SeverityType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ModifiedCiaType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ModifiedCiaType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ModifiedCiaType_3, v)
	}
	*j = ModifiedCiaType_3(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CiaType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CiaType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CiaType_3, v)
	}
	*j = CiaType_3(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ReportConfidenceType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ReportConfidenceType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ReportConfidenceType, v)
	}
	*j = ReportConfidenceType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ModifiedScopeType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ModifiedScopeType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ModifiedScopeType_1, v)
	}
	*j = ModifiedScopeType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PrivilegesRequiredType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_PrivilegesRequiredType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_PrivilegesRequiredType_1, v)
	}
	*j = PrivilegesRequiredType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CiaType_6) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CiaType_6 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CiaType_6, v)
	}
	*j = CiaType_6(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ModifiedAttackVectorType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ModifiedAttackVectorType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ModifiedAttackVectorType_1, v)
	}
	*j = ModifiedAttackVectorType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PrivilegesRequiredType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_PrivilegesRequiredType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_PrivilegesRequiredType, v)
	}
	*j = PrivilegesRequiredType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ScopeType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ScopeType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ScopeType_1, v)
	}
	*j = ScopeType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ExploitabilityType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ExploitabilityType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ExploitabilityType_1, v)
	}
	*j = ExploitabilityType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TargetDistributionType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_TargetDistributionType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_TargetDistributionType, v)
	}
	*j = TargetDistributionType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ModifiedCiaType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ModifiedCiaType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ModifiedCiaType_2, v)
	}
	*j = ModifiedCiaType_2(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CiaType_4) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CiaType_4 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CiaType_4, v)
	}
	*j = CiaType_4(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CiaRequirementType_6) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CiaRequirementType_6 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CiaRequirementType_6, v)
	}
	*j = CiaRequirementType_6(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CiaRequirementType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CiaRequirementType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CiaRequirementType_2, v)
	}
	*j = CiaRequirementType_2(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ModifiedUserInteractionType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ModifiedUserInteractionType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ModifiedUserInteractionType, v)
	}
	*j = ModifiedUserInteractionType(v)
	return nil
}

type AccessComplexityType_1 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *CiaType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CiaType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CiaType_2, v)
	}
	*j = CiaType_2(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AccessComplexityType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AccessComplexityType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AccessComplexityType_1, v)
	}
	*j = AccessComplexityType_1(v)
	return nil
}

const AccessComplexityType_1_HIGH AccessComplexityType_1 = "HIGH"
const AccessComplexityType_1_MEDIUM AccessComplexityType_1 = "MEDIUM"
const AccessComplexityType_1_LOW AccessComplexityType_1 = "LOW"

type AccessVectorType_1 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *ModifiedCiaType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ModifiedCiaType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ModifiedCiaType_1, v)
	}
	*j = ModifiedCiaType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AccessVectorType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AccessVectorType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AccessVectorType_1, v)
	}
	*j = AccessVectorType_1(v)
	return nil
}

const AccessVectorType_1_NETWORK AccessVectorType_1 = "NETWORK"
const AccessVectorType_1_ADJACENTNETWORK AccessVectorType_1 = "ADJACENT_NETWORK"
const AccessVectorType_1_LOCAL AccessVectorType_1 = "LOCAL"

// UnmarshalJSON implements json.Unmarshaler.
func (j *UserInteractionType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_UserInteractionType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_UserInteractionType_1, v)
	}
	*j = UserInteractionType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CvssV3XJsonVersion) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CvssV3XJsonVersion {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CvssV3XJsonVersion, v)
	}
	*j = CvssV3XJsonVersion(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuthenticationType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AuthenticationType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AuthenticationType_1, v)
	}
	*j = AuthenticationType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ExploitCodeMaturityType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ExploitCodeMaturityType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ExploitCodeMaturityType_1, v)
	}
	*j = ExploitCodeMaturityType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ConfidenceType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ConfidenceType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ConfidenceType, v)
	}
	*j = ConfidenceType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ExploitCodeMaturityType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ExploitCodeMaturityType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ExploitCodeMaturityType, v)
	}
	*j = ExploitCodeMaturityType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ModifiedAttackComplexityType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ModifiedAttackComplexityType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ModifiedAttackComplexityType, v)
	}
	*j = ModifiedAttackComplexityType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ModifiedAttackVectorType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ModifiedAttackVectorType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ModifiedAttackVectorType, v)
	}
	*j = ModifiedAttackVectorType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CiaType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CiaType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CiaType_1, v)
	}
	*j = CiaType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ModifiedCiaType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ModifiedCiaType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ModifiedCiaType, v)
	}
	*j = ModifiedCiaType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ModifiedPrivilegesRequiredType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ModifiedPrivilegesRequiredType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ModifiedPrivilegesRequiredType, v)
	}
	*j = ModifiedPrivilegesRequiredType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CollateralDamagePotentialType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CollateralDamagePotentialType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CollateralDamagePotentialType_1, v)
	}
	*j = CollateralDamagePotentialType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ConfidenceType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ConfidenceType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ConfidenceType_1, v)
	}
	*j = ConfidenceType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SeverityType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_SeverityType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_SeverityType_2, v)
	}
	*j = SeverityType_2(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CiaRequirementType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CiaRequirementType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CiaRequirementType_1, v)
	}
	*j = CiaRequirementType_1(v)
	return nil
}

const AuthenticationTypeSINGLE AuthenticationType = "SINGLE"

type AuthenticationType_1 string

const AuthenticationType_1_MULTIPLE AuthenticationType_1 = "MULTIPLE"
const AuthenticationType_1_NONE AuthenticationType_1 = "NONE"
const AuthenticationType_1_SINGLE AuthenticationType_1 = "SINGLE"

type CiaRequirementType string

const CiaRequirementTypeHIGH CiaRequirementType = "HIGH"
const CiaRequirementTypeLOW CiaRequirementType = "LOW"
const CiaRequirementTypeMEDIUM CiaRequirementType = "MEDIUM"
const CiaRequirementTypeNOTDEFINED CiaRequirementType = "NOT_DEFINED"

type CiaRequirementType_1 string

const CiaRequirementType_1_HIGH CiaRequirementType_1 = "HIGH"
const CiaRequirementType_1_LOW CiaRequirementType_1 = "LOW"
const CiaRequirementType_1_MEDIUM CiaRequirementType_1 = "MEDIUM"
const CiaRequirementType_1_NOTDEFINED CiaRequirementType_1 = "NOT_DEFINED"

type CiaRequirementType_2 string

const CiaRequirementType_2_HIGH CiaRequirementType_2 = "HIGH"
const CiaRequirementType_2_LOW CiaRequirementType_2 = "LOW"
const CiaRequirementType_2_MEDIUM CiaRequirementType_2 = "MEDIUM"
const CiaRequirementType_2_NOTDEFINED CiaRequirementType_2 = "NOT_DEFINED"

type CiaRequirementType_3 string

const CiaRequirementType_3_HIGH CiaRequirementType_3 = "HIGH"
const CiaRequirementType_3_LOW CiaRequirementType_3 = "LOW"
const CiaRequirementType_3_MEDIUM CiaRequirementType_3 = "MEDIUM"
const CiaRequirementType_3_NOTDEFINED CiaRequirementType_3 = "NOT_DEFINED"

type CiaRequirementType_4 string

const CiaRequirementType_4_HIGH CiaRequirementType_4 = "HIGH"
const CiaRequirementType_4_LOW CiaRequirementType_4 = "LOW"
const CiaRequirementType_4_MEDIUM CiaRequirementType_4 = "MEDIUM"
const CiaRequirementType_4_NOTDEFINED CiaRequirementType_4 = "NOT_DEFINED"

type CiaRequirementType_5 string

const CiaRequirementType_5_HIGH CiaRequirementType_5 = "HIGH"
const CiaRequirementType_5_LOW CiaRequirementType_5 = "LOW"
const CiaRequirementType_5_MEDIUM CiaRequirementType_5 = "MEDIUM"
const CiaRequirementType_5_NOTDEFINED CiaRequirementType_5 = "NOT_DEFINED"

type CiaRequirementType_6 string

const CiaRequirementType_6_HIGH CiaRequirementType_6 = "HIGH"
const CiaRequirementType_6_LOW CiaRequirementType_6 = "LOW"
const CiaRequirementType_6_MEDIUM CiaRequirementType_6 = "MEDIUM"
const CiaRequirementType_6_NOTDEFINED CiaRequirementType_6 = "NOT_DEFINED"

type CiaRequirementType_7 string

const CiaRequirementType_7_HIGH CiaRequirementType_7 = "HIGH"
const CiaRequirementType_7_LOW CiaRequirementType_7 = "LOW"
const CiaRequirementType_7_MEDIUM CiaRequirementType_7 = "MEDIUM"
const CiaRequirementType_7_NOTDEFINED CiaRequirementType_7 = "NOT_DEFINED"

type CiaType string

const CiaTypeCOMPLETE CiaType = "COMPLETE"
const CiaTypeNONE CiaType = "NONE"
const CiaTypePARTIAL CiaType = "PARTIAL"

type CiaType_1 string

const CiaType_1_COMPLETE CiaType_1 = "COMPLETE"
const CiaType_1_NONE CiaType_1 = "NONE"
const CiaType_1_PARTIAL CiaType_1 = "PARTIAL"

type CiaType_2 string

const CiaType_2_COMPLETE CiaType_2 = "COMPLETE"
const CiaType_2_NONE CiaType_2 = "NONE"
const CiaType_2_PARTIAL CiaType_2 = "PARTIAL"

type CiaType_3 string

const CiaType_3_COMPLETE CiaType_3 = "COMPLETE"
const CiaType_3_NONE CiaType_3 = "NONE"
const CiaType_3_PARTIAL CiaType_3 = "PARTIAL"

type CiaType_4 string

const CiaType_4_HIGH CiaType_4 = "HIGH"
const CiaType_4_LOW CiaType_4 = "LOW"
const CiaType_4_NONE CiaType_4 = "NONE"

type CiaType_5 string

const CiaType_5_HIGH CiaType_5 = "HIGH"
const CiaType_5_LOW CiaType_5 = "LOW"
const CiaType_5_NONE CiaType_5 = "NONE"

type CiaType_6 string

const CiaType_6_HIGH CiaType_6 = "HIGH"
const CiaType_6_LOW CiaType_6 = "LOW"
const CiaType_6_NONE CiaType_6 = "NONE"

type CiaType_7 string

const CiaType_7_HIGH CiaType_7 = "HIGH"
const CiaType_7_LOW CiaType_7 = "LOW"
const CiaType_7_NONE CiaType_7 = "NONE"

// UnmarshalJSON implements json.Unmarshaler.
func (j *AttackComplexityType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AttackComplexityType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AttackComplexityType, v)
	}
	*j = AttackComplexityType(v)
	return nil
}

type CollateralDamagePotentialType string

const CollateralDamagePotentialTypeHIGH CollateralDamagePotentialType = "HIGH"
const CollateralDamagePotentialTypeLOW CollateralDamagePotentialType = "LOW"
const CollateralDamagePotentialTypeLOWMEDIUM CollateralDamagePotentialType = "LOW_MEDIUM"
const CollateralDamagePotentialTypeMEDIUMHIGH CollateralDamagePotentialType = "MEDIUM_HIGH"
const CollateralDamagePotentialTypeNONE CollateralDamagePotentialType = "NONE"
const CollateralDamagePotentialTypeNOTDEFINED CollateralDamagePotentialType = "NOT_DEFINED"

type CollateralDamagePotentialType_1 string

const CollateralDamagePotentialType_1_HIGH CollateralDamagePotentialType_1 = "HIGH"
const CollateralDamagePotentialType_1_LOW CollateralDamagePotentialType_1 = "LOW"
const CollateralDamagePotentialType_1_LOWMEDIUM CollateralDamagePotentialType_1 = "LOW_MEDIUM"
const CollateralDamagePotentialType_1_MEDIUMHIGH CollateralDamagePotentialType_1 = "MEDIUM_HIGH"
const CollateralDamagePotentialType_1_NONE CollateralDamagePotentialType_1 = "NONE"
const CollateralDamagePotentialType_1_NOTDEFINED CollateralDamagePotentialType_1 = "NOT_DEFINED"

type ConfidenceType string

const ConfidenceTypeCONFIRMED ConfidenceType = "CONFIRMED"
const ConfidenceTypeNOTDEFINED ConfidenceType = "NOT_DEFINED"
const ConfidenceTypeREASONABLE ConfidenceType = "REASONABLE"
const ConfidenceTypeUNKNOWN ConfidenceType = "UNKNOWN"

type ConfidenceType_1 string

const ConfidenceType_1_CONFIRMED ConfidenceType_1 = "CONFIRMED"
const ConfidenceType_1_NOTDEFINED ConfidenceType_1 = "NOT_DEFINED"
const ConfidenceType_1_REASONABLE ConfidenceType_1 = "REASONABLE"
const ConfidenceType_1_UNKNOWN ConfidenceType_1 = "UNKNOWN"

type CvssV20Json struct {
	// AccessComplexity corresponds to the JSON schema field "accessComplexity".
	AccessComplexity *AccessComplexityType_1 `json:"accessComplexity,omitempty"`

	// AccessVector corresponds to the JSON schema field "accessVector".
	AccessVector *AccessVectorType_1 `json:"accessVector,omitempty"`

	// Authentication corresponds to the JSON schema field "authentication".
	Authentication *AuthenticationType_1 `json:"authentication,omitempty"`

	// AvailabilityImpact corresponds to the JSON schema field "availabilityImpact".
	AvailabilityImpact *CiaType_1 `json:"availabilityImpact,omitempty"`

	// AvailabilityRequirement corresponds to the JSON schema field
	// "availabilityRequirement".
	AvailabilityRequirement *CiaRequirementType_1 `json:"availabilityRequirement,omitempty"`

	// BaseScore corresponds to the JSON schema field "baseScore".
	BaseScore ScoreType `json:"baseScore"`

	// CollateralDamagePotential corresponds to the JSON schema field
	// "collateralDamagePotential".
	CollateralDamagePotential *CollateralDamagePotentialType_1 `json:"collateralDamagePotential,omitempty"`

	// ConfidentialityImpact corresponds to the JSON schema field
	// "confidentialityImpact".
	ConfidentialityImpact *CiaType_2 `json:"confidentialityImpact,omitempty"`

	// ConfidentialityRequirement corresponds to the JSON schema field
	// "confidentialityRequirement".
	ConfidentialityRequirement *CiaRequirementType_2 `json:"confidentialityRequirement,omitempty"`

	// EnvironmentalScore corresponds to the JSON schema field "environmentalScore".
	EnvironmentalScore *ScoreType `json:"environmentalScore,omitempty"`

	// Exploitability corresponds to the JSON schema field "exploitability".
	Exploitability *ExploitabilityType_1 `json:"exploitability,omitempty"`

	// IntegrityImpact corresponds to the JSON schema field "integrityImpact".
	IntegrityImpact *CiaType_3 `json:"integrityImpact,omitempty"`

	// IntegrityRequirement corresponds to the JSON schema field
	// "integrityRequirement".
	IntegrityRequirement *CiaRequirementType_3 `json:"integrityRequirement,omitempty"`

	// RemediationLevel corresponds to the JSON schema field "remediationLevel".
	RemediationLevel *RemediationLevelType_1 `json:"remediationLevel,omitempty"`

	// ReportConfidence corresponds to the JSON schema field "reportConfidence".
	ReportConfidence *ReportConfidenceType_1 `json:"reportConfidence,omitempty"`

	// TargetDistribution corresponds to the JSON schema field "targetDistribution".
	TargetDistribution *TargetDistributionType_1 `json:"targetDistribution,omitempty"`

	// TemporalScore corresponds to the JSON schema field "temporalScore".
	TemporalScore *ScoreType `json:"temporalScore,omitempty"`

	// VectorString corresponds to the JSON schema field "vectorString".
	VectorString string `json:"vectorString"`

	// CVSS Version
	Version CvssV20JsonVersion `json:"version"`
}

type CvssV20JsonVersion string

const CvssV20JsonVersionA20 CvssV20JsonVersion = "2.0"

type CvssV3XJson struct {
	// AttackComplexity corresponds to the JSON schema field "attackComplexity".
	AttackComplexity *AttackComplexityType_1 `json:"attackComplexity,omitempty"`

	// AttackVector corresponds to the JSON schema field "attackVector".
	AttackVector *AttackVectorType_1 `json:"attackVector,omitempty"`

	// AvailabilityImpact corresponds to the JSON schema field "availabilityImpact".
	AvailabilityImpact *CiaType_5 `json:"availabilityImpact,omitempty"`

	// AvailabilityRequirement corresponds to the JSON schema field
	// "availabilityRequirement".
	AvailabilityRequirement *CiaRequirementType_5 `json:"availabilityRequirement,omitempty"`

	// BaseScore corresponds to the JSON schema field "baseScore".
	BaseScore ScoreType_1 `json:"baseScore"`

	// BaseSeverity corresponds to the JSON schema field "baseSeverity".
	BaseSeverity SeverityType_1 `json:"baseSeverity"`

	// ConfidentialityImpact corresponds to the JSON schema field
	// "confidentialityImpact".
	ConfidentialityImpact *CiaType_6 `json:"confidentialityImpact,omitempty"`

	// ConfidentialityRequirement corresponds to the JSON schema field
	// "confidentialityRequirement".
	ConfidentialityRequirement *CiaRequirementType_6 `json:"confidentialityRequirement,omitempty"`

	// EnvironmentalScore corresponds to the JSON schema field "environmentalScore".
	EnvironmentalScore *ScoreType_1 `json:"environmentalScore,omitempty"`

	// EnvironmentalSeverity corresponds to the JSON schema field
	// "environmentalSeverity".
	EnvironmentalSeverity *SeverityType_2 `json:"environmentalSeverity,omitempty"`

	// ExploitCodeMaturity corresponds to the JSON schema field "exploitCodeMaturity".
	ExploitCodeMaturity *ExploitCodeMaturityType_1 `json:"exploitCodeMaturity,omitempty"`

	// IntegrityImpact corresponds to the JSON schema field "integrityImpact".
	IntegrityImpact *CiaType_7 `json:"integrityImpact,omitempty"`

	// IntegrityRequirement corresponds to the JSON schema field
	// "integrityRequirement".
	IntegrityRequirement *CiaRequirementType_7 `json:"integrityRequirement,omitempty"`

	// ModifiedAttackComplexity corresponds to the JSON schema field
	// "modifiedAttackComplexity".
	ModifiedAttackComplexity *ModifiedAttackComplexityType_1 `json:"modifiedAttackComplexity,omitempty"`

	// ModifiedAttackVector corresponds to the JSON schema field
	// "modifiedAttackVector".
	ModifiedAttackVector *ModifiedAttackVectorType_1 `json:"modifiedAttackVector,omitempty"`

	// ModifiedAvailabilityImpact corresponds to the JSON schema field
	// "modifiedAvailabilityImpact".
	ModifiedAvailabilityImpact *ModifiedCiaType_1 `json:"modifiedAvailabilityImpact,omitempty"`

	// ModifiedConfidentialityImpact corresponds to the JSON schema field
	// "modifiedConfidentialityImpact".
	ModifiedConfidentialityImpact *ModifiedCiaType_2 `json:"modifiedConfidentialityImpact,omitempty"`

	// ModifiedIntegrityImpact corresponds to the JSON schema field
	// "modifiedIntegrityImpact".
	ModifiedIntegrityImpact *ModifiedCiaType_3 `json:"modifiedIntegrityImpact,omitempty"`

	// ModifiedPrivilegesRequired corresponds to the JSON schema field
	// "modifiedPrivilegesRequired".
	ModifiedPrivilegesRequired *ModifiedPrivilegesRequiredType_1 `json:"modifiedPrivilegesRequired,omitempty"`

	// ModifiedScope corresponds to the JSON schema field "modifiedScope".
	ModifiedScope *ModifiedScopeType_1 `json:"modifiedScope,omitempty"`

	// ModifiedUserInteraction corresponds to the JSON schema field
	// "modifiedUserInteraction".
	ModifiedUserInteraction *ModifiedUserInteractionType_1 `json:"modifiedUserInteraction,omitempty"`

	// PrivilegesRequired corresponds to the JSON schema field "privilegesRequired".
	PrivilegesRequired *PrivilegesRequiredType_1 `json:"privilegesRequired,omitempty"`

	// RemediationLevel corresponds to the JSON schema field "remediationLevel".
	RemediationLevel *RemediationLevelType_3 `json:"remediationLevel,omitempty"`

	// ReportConfidence corresponds to the JSON schema field "reportConfidence".
	ReportConfidence *ConfidenceType_1 `json:"reportConfidence,omitempty"`

	// Scope corresponds to the JSON schema field "scope".
	Scope *ScopeType_1 `json:"scope,omitempty"`

	// TemporalScore corresponds to the JSON schema field "temporalScore".
	TemporalScore *ScoreType_1 `json:"temporalScore,omitempty"`

	// TemporalSeverity corresponds to the JSON schema field "temporalSeverity".
	TemporalSeverity *SeverityType_3 `json:"temporalSeverity,omitempty"`

	// UserInteraction corresponds to the JSON schema field "userInteraction".
	UserInteraction *UserInteractionType_1 `json:"userInteraction,omitempty"`

	// VectorString corresponds to the JSON schema field "vectorString".
	VectorString string `json:"vectorString"`

	// CVSS Version
	Version CvssV3XJsonVersion `json:"version"`
}

type CvssV3XJsonVersion string

const CvssV3XJsonVersionA30 CvssV3XJsonVersion = "3.0"
const CvssV3XJsonVersionA31 CvssV3XJsonVersion = "3.1"

// Defines the set of product configurations for a NVD applicability statement.
type DefConfigurations interface{}

// CPE match string or range
type DefCpeMatch struct {
	// Cpe22Uri corresponds to the JSON schema field "cpe22Uri".
	Cpe22Uri *string `json:"cpe22Uri,omitempty"`

	// Cpe23Uri corresponds to the JSON schema field "cpe23Uri".
	Cpe23Uri string `json:"cpe23Uri"`

	// CpeName corresponds to the JSON schema field "cpe_name".
	CpeName []DefCpeName `json:"cpe_name,omitempty"`

	// VersionEndExcluding corresponds to the JSON schema field "versionEndExcluding".
	VersionEndExcluding *string `json:"versionEndExcluding,omitempty"`

	// VersionEndIncluding corresponds to the JSON schema field "versionEndIncluding".
	VersionEndIncluding *string `json:"versionEndIncluding,omitempty"`

	// VersionStartExcluding corresponds to the JSON schema field
	// "versionStartExcluding".
	VersionStartExcluding *string `json:"versionStartExcluding,omitempty"`

	// VersionStartIncluding corresponds to the JSON schema field
	// "versionStartIncluding".
	VersionStartIncluding *string `json:"versionStartIncluding,omitempty"`

	// Vulnerable corresponds to the JSON schema field "vulnerable".
	Vulnerable bool `json:"vulnerable"`
}

// CPE name
type DefCpeName struct {
	// Cpe22Uri corresponds to the JSON schema field "cpe22Uri".
	Cpe22Uri *string `json:"cpe22Uri,omitempty"`

	// Cpe23Uri corresponds to the JSON schema field "cpe23Uri".
	Cpe23Uri string `json:"cpe23Uri"`

	// LastModifiedDate corresponds to the JSON schema field "lastModifiedDate".
	LastModifiedDate *string `json:"lastModifiedDate,omitempty"`
}

// Defines a vulnerability in the NVD data feed.

func UnmarshalDefCveItem(data []byte) (DefCveItem, error) {
	var r DefCveItem
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DefCveItem) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DefCveItem struct {
	CVE struct {
		Affects     *Affects    `json:"affects,omitempty"`
		CVEDataMeta CVEDataMeta `json:"CVE_data_meta"`
		DataFormat  DataFormat  `json:"data_format"`
		DataType    DataType    `json:"data_type"`
		DataVersion DataVersion `json:"data_version"`
		Description Description `json:"description"`
		Problemtype Problemtype `json:"problemtype"`
		References  References  `json:"references"`
	} `json:"cve"`
	Impact DefImpact `json:"impact"`
}

type Affects struct {
	Vendor Vendor `json:"vendor"`
}

type Vendor struct {
	VendorData []VendorDatum `json:"vendor_data"`
}

type VendorDatum struct {
	Product    ProductClass `json:"product"`
	VendorName string       `json:"vendor_name"`
}

type ProductClass struct {
	ProductData []ProductDatumElement `json:"product_data"`
}

type ProductDatumElement struct {
	ProductName string  `json:"product_name"`
	Version     Version `json:"version"`
}

type Version struct {
	VersionData []VersionDatum `json:"version_data"`
}

type VersionDatum struct {
	VersionAffected *string `json:"version_affected,omitempty"`
	VersionValue    string  `json:"version_value"`
}

type CVEDataMeta struct {
	Assigner string  `json:"ASSIGNER"`
	ID       string  `json:"ID"`
	State    *string `json:"STATE,omitempty"`
}

type Description struct {
	DescriptionData []LangString `json:"description_data"`
}

type LangString struct {
	Lang  string `json:"lang"`
	Value string `json:"value"`
}

type Problemtype struct {
	ProblemtypeData []ProblemtypeDatum `json:"problemtype_data"`
}

type ProblemtypeDatum struct {
	Description []LangString `json:"description"`
}

type References struct {
	ReferenceData []Reference `json:"reference_data"`
}

type Reference struct {
	Name      *string  `json:"name,omitempty"`
	Refsource *string  `json:"refsource,omitempty"`
	Tags      []string `json:"tags,omitempty"`
	URL       string   `json:"url"`
}

type DataFormat string

const (
	Mitre DataFormat = "MITRE"
)

type DataType string

const (
	_Cve DataType = "CVE"
)

type DataVersion string

const (
	The40 DataVersion = "4.0"
)

// Impact scores for a vulnerability as found on NVD.
type DefImpact struct {
	// CVSS V2.0 score.
	BaseMetricV2 *DefImpactBaseMetricV2 `json:"baseMetricV2,omitempty"`

	// CVSS V3.x score.
	BaseMetricV3 *DefImpactBaseMetricV3 `json:"baseMetricV3,omitempty"`
}

// CVSS V2.0 score.
type DefImpactBaseMetricV2 struct {
	// AcInsufInfo corresponds to the JSON schema field "acInsufInfo".
	AcInsufInfo *bool `json:"acInsufInfo,omitempty"`

	// CvssV2 corresponds to the JSON schema field "cvssV2".
	CvssV2 *CvssV20Json `json:"cvssV2,omitempty"`

	// ExploitabilityScore corresponds to the JSON schema field "exploitabilityScore".
	ExploitabilityScore *DefSubscore `json:"exploitabilityScore,omitempty"`

	// ImpactScore corresponds to the JSON schema field "impactScore".
	ImpactScore *DefSubscore `json:"impactScore,omitempty"`

	// ObtainAllPrivilege corresponds to the JSON schema field "obtainAllPrivilege".
	ObtainAllPrivilege *bool `json:"obtainAllPrivilege,omitempty"`

	// ObtainOtherPrivilege corresponds to the JSON schema field
	// "obtainOtherPrivilege".
	ObtainOtherPrivilege *bool `json:"obtainOtherPrivilege,omitempty"`

	// ObtainUserPrivilege corresponds to the JSON schema field "obtainUserPrivilege".
	ObtainUserPrivilege *bool `json:"obtainUserPrivilege,omitempty"`

	// Severity corresponds to the JSON schema field "severity".
	Severity *string `json:"severity,omitempty"`

	// UserInteractionRequired corresponds to the JSON schema field
	// "userInteractionRequired".
	UserInteractionRequired *bool `json:"userInteractionRequired,omitempty"`
}

// CVSS V3.x score.
type DefImpactBaseMetricV3 struct {
	// CvssV3 corresponds to the JSON schema field "cvssV3".
	CvssV3 *CvssV3XJson `json:"cvssV3,omitempty"`

	// ExploitabilityScore corresponds to the JSON schema field "exploitabilityScore".
	ExploitabilityScore *DefSubscore `json:"exploitabilityScore,omitempty"`

	// ImpactScore corresponds to the JSON schema field "impactScore".
	ImpactScore *DefSubscore `json:"impactScore,omitempty"`
}

// Defines a node or sub-node in an NVD applicability statement.
type DefNode interface{}

// CVSS subscore.
type DefSubscore float64

type ExploitCodeMaturityType string

const ExploitCodeMaturityTypeFUNCTIONAL ExploitCodeMaturityType = "FUNCTIONAL"
const ExploitCodeMaturityTypeHIGH ExploitCodeMaturityType = "HIGH"
const ExploitCodeMaturityTypeNOTDEFINED ExploitCodeMaturityType = "NOT_DEFINED"
const ExploitCodeMaturityTypePROOFOFCONCEPT ExploitCodeMaturityType = "PROOF_OF_CONCEPT"
const ExploitCodeMaturityTypeUNPROVEN ExploitCodeMaturityType = "UNPROVEN"

type ExploitCodeMaturityType_1 string

const ExploitCodeMaturityType_1_FUNCTIONAL ExploitCodeMaturityType_1 = "FUNCTIONAL"
const ExploitCodeMaturityType_1_HIGH ExploitCodeMaturityType_1 = "HIGH"
const ExploitCodeMaturityType_1_NOTDEFINED ExploitCodeMaturityType_1 = "NOT_DEFINED"
const ExploitCodeMaturityType_1_PROOFOFCONCEPT ExploitCodeMaturityType_1 = "PROOF_OF_CONCEPT"
const ExploitCodeMaturityType_1_UNPROVEN ExploitCodeMaturityType_1 = "UNPROVEN"

type ExploitabilityType string

const ExploitabilityTypeFUNCTIONAL ExploitabilityType = "FUNCTIONAL"
const ExploitabilityTypeHIGH ExploitabilityType = "HIGH"
const ExploitabilityTypeNOTDEFINED ExploitabilityType = "NOT_DEFINED"
const ExploitabilityTypePROOFOFCONCEPT ExploitabilityType = "PROOF_OF_CONCEPT"
const ExploitabilityTypeUNPROVEN ExploitabilityType = "UNPROVEN"

type ExploitabilityType_1 string

const ExploitabilityType_1_FUNCTIONAL ExploitabilityType_1 = "FUNCTIONAL"
const ExploitabilityType_1_HIGH ExploitabilityType_1 = "HIGH"
const ExploitabilityType_1_NOTDEFINED ExploitabilityType_1 = "NOT_DEFINED"
const ExploitabilityType_1_PROOFOFCONCEPT ExploitabilityType_1 = "PROOF_OF_CONCEPT"
const ExploitabilityType_1_UNPROVEN ExploitabilityType_1 = "UNPROVEN"

type ModifiedAttackComplexityType string

const ModifiedAttackComplexityTypeHIGH ModifiedAttackComplexityType = "HIGH"
const ModifiedAttackComplexityTypeLOW ModifiedAttackComplexityType = "LOW"
const ModifiedAttackComplexityTypeNOTDEFINED ModifiedAttackComplexityType = "NOT_DEFINED"

type ModifiedAttackComplexityType_1 string

const ModifiedAttackComplexityType_1_HIGH ModifiedAttackComplexityType_1 = "HIGH"
const ModifiedAttackComplexityType_1_LOW ModifiedAttackComplexityType_1 = "LOW"
const ModifiedAttackComplexityType_1_NOTDEFINED ModifiedAttackComplexityType_1 = "NOT_DEFINED"

type ModifiedAttackVectorType string

const ModifiedAttackVectorTypeADJACENTNETWORK ModifiedAttackVectorType = "ADJACENT_NETWORK"
const ModifiedAttackVectorTypeLOCAL ModifiedAttackVectorType = "LOCAL"
const ModifiedAttackVectorTypeNETWORK ModifiedAttackVectorType = "NETWORK"
const ModifiedAttackVectorTypeNOTDEFINED ModifiedAttackVectorType = "NOT_DEFINED"
const ModifiedAttackVectorTypePHYSICAL ModifiedAttackVectorType = "PHYSICAL"

type ModifiedAttackVectorType_1 string

const ModifiedAttackVectorType_1_ADJACENTNETWORK ModifiedAttackVectorType_1 = "ADJACENT_NETWORK"
const ModifiedAttackVectorType_1_LOCAL ModifiedAttackVectorType_1 = "LOCAL"
const ModifiedAttackVectorType_1_NETWORK ModifiedAttackVectorType_1 = "NETWORK"
const ModifiedAttackVectorType_1_NOTDEFINED ModifiedAttackVectorType_1 = "NOT_DEFINED"
const ModifiedAttackVectorType_1_PHYSICAL ModifiedAttackVectorType_1 = "PHYSICAL"

type ModifiedCiaType string

const ModifiedCiaTypeHIGH ModifiedCiaType = "HIGH"
const ModifiedCiaTypeLOW ModifiedCiaType = "LOW"
const ModifiedCiaTypeNONE ModifiedCiaType = "NONE"
const ModifiedCiaTypeNOTDEFINED ModifiedCiaType = "NOT_DEFINED"

type ModifiedCiaType_1 string

const ModifiedCiaType_1_HIGH ModifiedCiaType_1 = "HIGH"
const ModifiedCiaType_1_LOW ModifiedCiaType_1 = "LOW"
const ModifiedCiaType_1_NONE ModifiedCiaType_1 = "NONE"
const ModifiedCiaType_1_NOTDEFINED ModifiedCiaType_1 = "NOT_DEFINED"

type ModifiedCiaType_2 string

const ModifiedCiaType_2_HIGH ModifiedCiaType_2 = "HIGH"
const ModifiedCiaType_2_LOW ModifiedCiaType_2 = "LOW"
const ModifiedCiaType_2_NONE ModifiedCiaType_2 = "NONE"
const ModifiedCiaType_2_NOTDEFINED ModifiedCiaType_2 = "NOT_DEFINED"

type ModifiedCiaType_3 string

const ModifiedCiaType_3_HIGH ModifiedCiaType_3 = "HIGH"
const ModifiedCiaType_3_LOW ModifiedCiaType_3 = "LOW"
const ModifiedCiaType_3_NONE ModifiedCiaType_3 = "NONE"
const ModifiedCiaType_3_NOTDEFINED ModifiedCiaType_3 = "NOT_DEFINED"

type ModifiedPrivilegesRequiredType string

const ModifiedPrivilegesRequiredTypeHIGH ModifiedPrivilegesRequiredType = "HIGH"
const ModifiedPrivilegesRequiredTypeLOW ModifiedPrivilegesRequiredType = "LOW"
const ModifiedPrivilegesRequiredTypeNONE ModifiedPrivilegesRequiredType = "NONE"
const ModifiedPrivilegesRequiredTypeNOTDEFINED ModifiedPrivilegesRequiredType = "NOT_DEFINED"

type ModifiedPrivilegesRequiredType_1 string

const ModifiedPrivilegesRequiredType_1_HIGH ModifiedPrivilegesRequiredType_1 = "HIGH"
const ModifiedPrivilegesRequiredType_1_LOW ModifiedPrivilegesRequiredType_1 = "LOW"

// UnmarshalJSON implements json.Unmarshaler.
func (j *RemediationLevelType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_RemediationLevelType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_RemediationLevelType_3, v)
	}
	*j = RemediationLevelType_3(v)
	return nil
}

const ModifiedPrivilegesRequiredType_1_NONE ModifiedPrivilegesRequiredType_1 = "NONE"
const ModifiedPrivilegesRequiredType_1_NOTDEFINED ModifiedPrivilegesRequiredType_1 = "NOT_DEFINED"

type ModifiedScopeType string

const ModifiedScopeTypeCHANGED ModifiedScopeType = "CHANGED"
const ModifiedScopeTypeNOTDEFINED ModifiedScopeType = "NOT_DEFINED"
const ModifiedScopeTypeUNCHANGED ModifiedScopeType = "UNCHANGED"

type ModifiedScopeType_1 string

const ModifiedScopeType_1_CHANGED ModifiedScopeType_1 = "CHANGED"
const ModifiedScopeType_1_NOTDEFINED ModifiedScopeType_1 = "NOT_DEFINED"
const ModifiedScopeType_1_UNCHANGED ModifiedScopeType_1 = "UNCHANGED"

type ModifiedUserInteractionType string

const ModifiedUserInteractionTypeNONE ModifiedUserInteractionType = "NONE"
const ModifiedUserInteractionTypeNOTDEFINED ModifiedUserInteractionType = "NOT_DEFINED"
const ModifiedUserInteractionTypeREQUIRED ModifiedUserInteractionType = "REQUIRED"

type ModifiedUserInteractionType_1 string

const ModifiedUserInteractionType_1_NONE ModifiedUserInteractionType_1 = "NONE"
const ModifiedUserInteractionType_1_NOTDEFINED ModifiedUserInteractionType_1 = "NOT_DEFINED"
const ModifiedUserInteractionType_1_REQUIRED ModifiedUserInteractionType_1 = "REQUIRED"

type NvdCveFeedJson11Schema struct {
	// NVD feed array of CVE
	CVEItems []DefCveItem `json:"CVE_Items"`

	// CVEDataFormat corresponds to the JSON schema field "CVE_data_format".
	CVEDataFormat string `json:"CVE_data_format"`

	// NVD adds number of CVE in this feed
	CVEDataNumberOfCVEs *string `json:"CVE_data_numberOfCVEs,omitempty"`

	// NVD adds feed date timestamp
	CVEDataTimestamp *string `json:"CVE_data_timestamp,omitempty"`

	// CVEDataType corresponds to the JSON schema field "CVE_data_type".
	CVEDataType string `json:"CVE_data_type"`

	// CVEDataVersion corresponds to the JSON schema field "CVE_data_version".
	CVEDataVersion string `json:"CVE_data_version"`
}

type PrivilegesRequiredType string

const PrivilegesRequiredTypeHIGH PrivilegesRequiredType = "HIGH"
const PrivilegesRequiredTypeLOW PrivilegesRequiredType = "LOW"
const PrivilegesRequiredTypeNONE PrivilegesRequiredType = "NONE"

type PrivilegesRequiredType_1 string

const PrivilegesRequiredType_1_HIGH PrivilegesRequiredType_1 = "HIGH"
const PrivilegesRequiredType_1_LOW PrivilegesRequiredType_1 = "LOW"
const PrivilegesRequiredType_1_NONE PrivilegesRequiredType_1 = "NONE"

type RemediationLevelType string

const RemediationLevelTypeNOTDEFINED RemediationLevelType = "NOT_DEFINED"
const RemediationLevelTypeOFFICIALFIX RemediationLevelType = "OFFICIAL_FIX"
const RemediationLevelTypeTEMPORARYFIX RemediationLevelType = "TEMPORARY_FIX"
const RemediationLevelTypeUNAVAILABLE RemediationLevelType = "UNAVAILABLE"
const RemediationLevelTypeWORKAROUND RemediationLevelType = "WORKAROUND"

type RemediationLevelType_1 string

const RemediationLevelType_1_NOTDEFINED RemediationLevelType_1 = "NOT_DEFINED"
const RemediationLevelType_1_OFFICIALFIX RemediationLevelType_1 = "OFFICIAL_FIX"
const RemediationLevelType_1_TEMPORARYFIX RemediationLevelType_1 = "TEMPORARY_FIX"
const RemediationLevelType_1_UNAVAILABLE RemediationLevelType_1 = "UNAVAILABLE"
const RemediationLevelType_1_WORKAROUND RemediationLevelType_1 = "WORKAROUND"

type RemediationLevelType_2 string

const RemediationLevelType_2_NOTDEFINED RemediationLevelType_2 = "NOT_DEFINED"
const RemediationLevelType_2_OFFICIALFIX RemediationLevelType_2 = "OFFICIAL_FIX"
const RemediationLevelType_2_TEMPORARYFIX RemediationLevelType_2 = "TEMPORARY_FIX"
const RemediationLevelType_2_UNAVAILABLE RemediationLevelType_2 = "UNAVAILABLE"
const RemediationLevelType_2_WORKAROUND RemediationLevelType_2 = "WORKAROUND"

type RemediationLevelType_3 string

const RemediationLevelType_3_NOTDEFINED RemediationLevelType_3 = "NOT_DEFINED"
const RemediationLevelType_3_OFFICIALFIX RemediationLevelType_3 = "OFFICIAL_FIX"
const RemediationLevelType_3_TEMPORARYFIX RemediationLevelType_3 = "TEMPORARY_FIX"
const RemediationLevelType_3_UNAVAILABLE RemediationLevelType_3 = "UNAVAILABLE"
const RemediationLevelType_3_WORKAROUND RemediationLevelType_3 = "WORKAROUND"

type ReportConfidenceType string

const ReportConfidenceTypeCONFIRMED ReportConfidenceType = "CONFIRMED"
const ReportConfidenceTypeNOTDEFINED ReportConfidenceType = "NOT_DEFINED"
const ReportConfidenceTypeUNCONFIRMED ReportConfidenceType = "UNCONFIRMED"
const ReportConfidenceTypeUNCORROBORATED ReportConfidenceType = "UNCORROBORATED"

type ReportConfidenceType_1 string

const ReportConfidenceType_1_CONFIRMED ReportConfidenceType_1 = "CONFIRMED"
const ReportConfidenceType_1_NOTDEFINED ReportConfidenceType_1 = "NOT_DEFINED"
const ReportConfidenceType_1_UNCONFIRMED ReportConfidenceType_1 = "UNCONFIRMED"
const ReportConfidenceType_1_UNCORROBORATED ReportConfidenceType_1 = "UNCORROBORATED"

type ScopeType string

const ScopeTypeCHANGED ScopeType = "CHANGED"
const ScopeTypeUNCHANGED ScopeType = "UNCHANGED"

type ScopeType_1 string

const ScopeType_1_CHANGED ScopeType_1 = "CHANGED"
const ScopeType_1_UNCHANGED ScopeType_1 = "UNCHANGED"

type ScoreType float64

type ScoreType_1 float64

type SeverityType string

const SeverityTypeCRITICAL SeverityType = "CRITICAL"
const SeverityTypeHIGH SeverityType = "HIGH"
const SeverityTypeLOW SeverityType = "LOW"
const SeverityTypeMEDIUM SeverityType = "MEDIUM"
const SeverityTypeNONE SeverityType = "NONE"

type SeverityType_1 string

const SeverityType_1_CRITICAL SeverityType_1 = "CRITICAL"
const SeverityType_1_HIGH SeverityType_1 = "HIGH"
const SeverityType_1_LOW SeverityType_1 = "LOW"
const SeverityType_1_MEDIUM SeverityType_1 = "MEDIUM"
const SeverityType_1_NONE SeverityType_1 = "NONE"

type SeverityType_2 string

const SeverityType_2_CRITICAL SeverityType_2 = "CRITICAL"
const SeverityType_2_HIGH SeverityType_2 = "HIGH"
const SeverityType_2_LOW SeverityType_2 = "LOW"
const SeverityType_2_MEDIUM SeverityType_2 = "MEDIUM"
const SeverityType_2_NONE SeverityType_2 = "NONE"

type SeverityType_3 string

const SeverityType_3_CRITICAL SeverityType_3 = "CRITICAL"
const SeverityType_3_HIGH SeverityType_3 = "HIGH"
const SeverityType_3_LOW SeverityType_3 = "LOW"
const SeverityType_3_MEDIUM SeverityType_3 = "MEDIUM"
const SeverityType_3_NONE SeverityType_3 = "NONE"

type TargetDistributionType string

const TargetDistributionTypeHIGH TargetDistributionType = "HIGH"
const TargetDistributionTypeLOW TargetDistributionType = "LOW"
const TargetDistributionTypeMEDIUM TargetDistributionType = "MEDIUM"
const TargetDistributionTypeNONE TargetDistributionType = "NONE"
const TargetDistributionTypeNOTDEFINED TargetDistributionType = "NOT_DEFINED"

type TargetDistributionType_1 string

const TargetDistributionType_1_HIGH TargetDistributionType_1 = "HIGH"
const TargetDistributionType_1_LOW TargetDistributionType_1 = "LOW"
const TargetDistributionType_1_MEDIUM TargetDistributionType_1 = "MEDIUM"
const TargetDistributionType_1_NONE TargetDistributionType_1 = "NONE"
const TargetDistributionType_1_NOTDEFINED TargetDistributionType_1 = "NOT_DEFINED"

type UserInteractionType string

const UserInteractionTypeNONE UserInteractionType = "NONE"
const UserInteractionTypeREQUIRED UserInteractionType = "REQUIRED"

type UserInteractionType_1 string

const UserInteractionType_1_NONE UserInteractionType_1 = "NONE"
const UserInteractionType_1_REQUIRED UserInteractionType_1 = "REQUIRED"

var enumValues_AccessComplexityType = []interface{}{
	"HIGH",
	"MEDIUM",
	"LOW",
}
var enumValues_AccessComplexityType_1 = []interface{}{
	"HIGH",
	"MEDIUM",
	"LOW",
}
var enumValues_AccessVectorType = []interface{}{
	"NETWORK",
	"ADJACENT_NETWORK",
	"LOCAL",
}
var enumValues_AccessVectorType_1 = []interface{}{
	"NETWORK",
	"ADJACENT_NETWORK",
	"LOCAL",
}
var enumValues_AttackComplexityType = []interface{}{
	"HIGH",
	"LOW",
}
var enumValues_AttackComplexityType_1 = []interface{}{
	"HIGH",
	"LOW",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NvdCveFeedJson11Schema) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["CVE_Items"]; !ok || v == nil {
		return fmt.Errorf("field CVE_Items: required")
	}
	if v, ok := raw["CVE_data_format"]; !ok || v == nil {
		return fmt.Errorf("field CVE_data_format: required")
	}
	if v, ok := raw["CVE_data_type"]; !ok || v == nil {
		return fmt.Errorf("field CVE_data_type: required")
	}
	if v, ok := raw["CVE_data_version"]; !ok || v == nil {
		return fmt.Errorf("field CVE_data_version: required")
	}
	type Plain NvdCveFeedJson11Schema
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = NvdCveFeedJson11Schema(plain)
	return nil
}

var enumValues_AttackVectorType = []interface{}{
	"NETWORK",
	"ADJACENT_NETWORK",
	"LOCAL",
	"PHYSICAL",
}
var enumValues_AttackVectorType_1 = []interface{}{
	"NETWORK",
	"ADJACENT_NETWORK",
	"LOCAL",
	"PHYSICAL",
}
var enumValues_AuthenticationType = []interface{}{
	"MULTIPLE",
	"SINGLE",
	"NONE",
}
var enumValues_AuthenticationType_1 = []interface{}{
	"MULTIPLE",
	"SINGLE",
	"NONE",
}
var enumValues_CiaRequirementType = []interface{}{
	"LOW",
	"MEDIUM",
	"HIGH",
	"NOT_DEFINED",
}
var enumValues_CiaRequirementType_1 = []interface{}{
	"LOW",
	"MEDIUM",
	"HIGH",
	"NOT_DEFINED",
}
var enumValues_CiaRequirementType_2 = []interface{}{
	"LOW",
	"MEDIUM",
	"HIGH",
	"NOT_DEFINED",
}
var enumValues_CiaRequirementType_3 = []interface{}{
	"LOW",
	"MEDIUM",
	"HIGH",
	"NOT_DEFINED",
}
var enumValues_CiaRequirementType_4 = []interface{}{
	"LOW",
	"MEDIUM",
	"HIGH",
	"NOT_DEFINED",
}
var enumValues_CiaRequirementType_5 = []interface{}{
	"LOW",
	"MEDIUM",
	"HIGH",
	"NOT_DEFINED",
}
var enumValues_CiaRequirementType_6 = []interface{}{
	"LOW",
	"MEDIUM",
	"HIGH",
	"NOT_DEFINED",
}
var enumValues_CiaRequirementType_7 = []interface{}{
	"LOW",
	"MEDIUM",
	"HIGH",
	"NOT_DEFINED",
}
var enumValues_CiaType = []interface{}{
	"NONE",
	"PARTIAL",
	"COMPLETE",
}
var enumValues_CiaType_1 = []interface{}{
	"NONE",
	"PARTIAL",
	"COMPLETE",
}
var enumValues_CiaType_2 = []interface{}{
	"NONE",
	"PARTIAL",
	"COMPLETE",
}
var enumValues_CiaType_3 = []interface{}{
	"NONE",
	"PARTIAL",
	"COMPLETE",
}
var enumValues_CiaType_4 = []interface{}{
	"NONE",
	"LOW",
	"HIGH",
}
var enumValues_CiaType_5 = []interface{}{
	"NONE",
	"LOW",
	"HIGH",
}
var enumValues_CiaType_6 = []interface{}{
	"NONE",
	"LOW",
	"HIGH",
}
var enumValues_CiaType_7 = []interface{}{
	"NONE",
	"LOW",
	"HIGH",
}
var enumValues_CollateralDamagePotentialType = []interface{}{
	"NONE",
	"LOW",
	"LOW_MEDIUM",
	"MEDIUM_HIGH",
	"HIGH",
	"NOT_DEFINED",
}
var enumValues_CollateralDamagePotentialType_1 = []interface{}{
	"NONE",
	"LOW",
	"LOW_MEDIUM",
	"MEDIUM_HIGH",
	"HIGH",
	"NOT_DEFINED",
}
var enumValues_ConfidenceType = []interface{}{
	"UNKNOWN",
	"REASONABLE",
	"CONFIRMED",
	"NOT_DEFINED",
}
var enumValues_ConfidenceType_1 = []interface{}{
	"UNKNOWN",
	"REASONABLE",
	"CONFIRMED",
	"NOT_DEFINED",
}
var enumValues_CvssV20JsonVersion = []interface{}{
	"2.0",
}
var enumValues_CvssV3XJsonVersion = []interface{}{
	"3.0",
	"3.1",
}
var enumValues_ExploitCodeMaturityType = []interface{}{
	"UNPROVEN",
	"PROOF_OF_CONCEPT",
	"FUNCTIONAL",
	"HIGH",
	"NOT_DEFINED",
}
var enumValues_ExploitCodeMaturityType_1 = []interface{}{
	"UNPROVEN",
	"PROOF_OF_CONCEPT",
	"FUNCTIONAL",
	"HIGH",
	"NOT_DEFINED",
}
var enumValues_ExploitabilityType = []interface{}{
	"UNPROVEN",
	"PROOF_OF_CONCEPT",
	"FUNCTIONAL",
	"HIGH",
	"NOT_DEFINED",
}
var enumValues_ExploitabilityType_1 = []interface{}{
	"UNPROVEN",
	"PROOF_OF_CONCEPT",
	"FUNCTIONAL",
	"HIGH",
	"NOT_DEFINED",
}
var enumValues_ModifiedAttackComplexityType = []interface{}{
	"HIGH",
	"LOW",
	"NOT_DEFINED",
}
var enumValues_ModifiedAttackComplexityType_1 = []interface{}{
	"HIGH",
	"LOW",
	"NOT_DEFINED",
}
var enumValues_ModifiedAttackVectorType = []interface{}{
	"NETWORK",
	"ADJACENT_NETWORK",
	"LOCAL",
	"PHYSICAL",
	"NOT_DEFINED",
}
var enumValues_ModifiedAttackVectorType_1 = []interface{}{
	"NETWORK",
	"ADJACENT_NETWORK",
	"LOCAL",
	"PHYSICAL",
	"NOT_DEFINED",
}
var enumValues_ModifiedCiaType = []interface{}{
	"NONE",
	"LOW",
	"HIGH",
	"NOT_DEFINED",
}
var enumValues_ModifiedCiaType_1 = []interface{}{
	"NONE",
	"LOW",
	"HIGH",
	"NOT_DEFINED",
}
var enumValues_ModifiedCiaType_2 = []interface{}{
	"NONE",
	"LOW",
	"HIGH",
	"NOT_DEFINED",
}
var enumValues_ModifiedCiaType_3 = []interface{}{
	"NONE",
	"LOW",
	"HIGH",
	"NOT_DEFINED",
}
var enumValues_ModifiedPrivilegesRequiredType = []interface{}{
	"HIGH",
	"LOW",
	"NONE",
	"NOT_DEFINED",
}
var enumValues_ModifiedPrivilegesRequiredType_1 = []interface{}{
	"HIGH",
	"LOW",
	"NONE",
	"NOT_DEFINED",
}
var enumValues_ModifiedScopeType = []interface{}{
	"UNCHANGED",
	"CHANGED",
	"NOT_DEFINED",
}
var enumValues_ModifiedScopeType_1 = []interface{}{
	"UNCHANGED",
	"CHANGED",
	"NOT_DEFINED",
}
var enumValues_ModifiedUserInteractionType = []interface{}{
	"NONE",
	"REQUIRED",
	"NOT_DEFINED",
}
var enumValues_ModifiedUserInteractionType_1 = []interface{}{
	"NONE",
	"REQUIRED",
	"NOT_DEFINED",
}
var enumValues_PrivilegesRequiredType = []interface{}{
	"HIGH",
	"LOW",
	"NONE",
}
var enumValues_PrivilegesRequiredType_1 = []interface{}{
	"HIGH",
	"LOW",
	"NONE",
}
var enumValues_RemediationLevelType = []interface{}{
	"OFFICIAL_FIX",
	"TEMPORARY_FIX",
	"WORKAROUND",
	"UNAVAILABLE",
	"NOT_DEFINED",
}
var enumValues_RemediationLevelType_1 = []interface{}{
	"OFFICIAL_FIX",
	"TEMPORARY_FIX",
	"WORKAROUND",
	"UNAVAILABLE",
	"NOT_DEFINED",
}
var enumValues_RemediationLevelType_2 = []interface{}{
	"OFFICIAL_FIX",
	"TEMPORARY_FIX",
	"WORKAROUND",
	"UNAVAILABLE",
	"NOT_DEFINED",
}
var enumValues_RemediationLevelType_3 = []interface{}{
	"OFFICIAL_FIX",
	"TEMPORARY_FIX",
	"WORKAROUND",
	"UNAVAILABLE",
	"NOT_DEFINED",
}
var enumValues_ReportConfidenceType = []interface{}{
	"UNCONFIRMED",
	"UNCORROBORATED",
	"CONFIRMED",
	"NOT_DEFINED",
}
var enumValues_ReportConfidenceType_1 = []interface{}{
	"UNCONFIRMED",
	"UNCORROBORATED",
	"CONFIRMED",
	"NOT_DEFINED",
}
var enumValues_ScopeType = []interface{}{
	"UNCHANGED",
	"CHANGED",
}
var enumValues_ScopeType_1 = []interface{}{
	"UNCHANGED",
	"CHANGED",
}
var enumValues_SeverityType = []interface{}{
	"NONE",
	"LOW",
	"MEDIUM",
	"HIGH",
	"CRITICAL",
}
var enumValues_SeverityType_1 = []interface{}{
	"NONE",
	"LOW",
	"MEDIUM",
	"HIGH",
	"CRITICAL",
}
var enumValues_SeverityType_2 = []interface{}{
	"NONE",
	"LOW",
	"MEDIUM",
	"HIGH",
	"CRITICAL",
}
var enumValues_SeverityType_3 = []interface{}{
	"NONE",
	"LOW",
	"MEDIUM",
	"HIGH",
	"CRITICAL",
}
var enumValues_TargetDistributionType = []interface{}{
	"NONE",
	"LOW",
	"MEDIUM",
	"HIGH",
	"NOT_DEFINED",
}
var enumValues_TargetDistributionType_1 = []interface{}{
	"NONE",
	"LOW",
	"MEDIUM",
	"HIGH",
	"NOT_DEFINED",
}
var enumValues_UserInteractionType = []interface{}{
	"NONE",
	"REQUIRED",
}
var enumValues_UserInteractionType_1 = []interface{}{
	"NONE",
	"REQUIRED",
}
