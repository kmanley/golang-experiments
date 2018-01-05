package exigoapi

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type AccountCreditCardType string

const (
	AccountCreditCardTypePrimary AccountCreditCardType = "Primary"

	AccountCreditCardTypeSecondary AccountCreditCardType = "Secondary"
)

type OrderShipCarrier string

const (
	OrderShipCarrierFedEx OrderShipCarrier = "FedEx"

	OrderShipCarrierUPS OrderShipCarrier = "UPS"

	OrderShipCarrierPurolator OrderShipCarrier = "Purolator"

	OrderShipCarrierCanadaPostRegular OrderShipCarrier = "CanadaPostRegular"

	OrderShipCarrierCanadaPostExpress OrderShipCarrier = "CanadaPostExpress"

	OrderShipCarrierDHL OrderShipCarrier = "DHL"

	OrderShipCarrierUSPS OrderShipCarrier = "USPS"

	OrderShipCarrierEstafeta OrderShipCarrier = "Estafeta"

	OrderShipCarrierRoyalMail OrderShipCarrier = "RoyalMail"

	OrderShipCarrierGLSHungary OrderShipCarrier = "GLSHungary"

	OrderShipCarrierPostenNorway OrderShipCarrier = "PostenNorway"

	OrderShipCarrierLandMarkGlobal OrderShipCarrier = "LandMarkGlobal"

	OrderShipCarrierGDEX OrderShipCarrier = "GDEX"
)

type PayableType string

const (
	PayableTypeCheck PayableType = "Check"

	PayableTypeWireTransfer PayableType = "WireTransfer"

	PayableTypePaymentCard PayableType = "PaymentCard"

	PayableTypeDirectDeposit PayableType = "DirectDeposit"

	PayableTypeOnHold PayableType = "OnHold"

	PayableTypeBankWire PayableType = "BankWire"

	PayableTypeDebitCardHold PayableType = "DebitCardHold"

	PayableTypeOther100 PayableType = "Other100"
)

type TaxIDType string

const (
	TaxIDTypeSSN TaxIDType = "SSN"

	TaxIDTypeEIN TaxIDType = "EIN"

	TaxIDTypeOtherType3 TaxIDType = "OtherType3"

	TaxIDTypeOtherType4 TaxIDType = "OtherType4"

	TaxIDTypeOtherType5 TaxIDType = "OtherType5"

	TaxIDTypeOtherType6 TaxIDType = "OtherType6"

	TaxIDTypeOtherType7 TaxIDType = "OtherType7"

	TaxIDTypeOtherType8 TaxIDType = "OtherType8"

	TaxIDTypeOtherType9 TaxIDType = "OtherType9"

	TaxIDTypeOtherType10 TaxIDType = "OtherType10"

	TaxIDTypeOtherType11 TaxIDType = "OtherType11"

	TaxIDTypeOtherType12 TaxIDType = "OtherType12"

	TaxIDTypeOtherType13 TaxIDType = "OtherType13"

	TaxIDTypeOtherType14 TaxIDType = "OtherType14"

	TaxIDTypeOtherType15 TaxIDType = "OtherType15"

	TaxIDTypeOtherType16 TaxIDType = "OtherType16"

	TaxIDTypeOtherType17 TaxIDType = "OtherType17"

	TaxIDTypeOtherType18 TaxIDType = "OtherType18"

	TaxIDTypeOtherType19 TaxIDType = "OtherType19"

	TaxIDTypeOtherType20 TaxIDType = "OtherType20"
)

type Gender string

const (
	GenderUnknown Gender = "Unknown"

	GenderMale Gender = "Male"

	GenderFemale Gender = "Female"
)

type OrderType string

const (
	OrderTypeDefault OrderType = "Default"

	OrderTypeCustomerService OrderType = "CustomerService"

	OrderTypeShoppingCart OrderType = "ShoppingCart"

	OrderTypeWebWizard OrderType = "WebWizard"

	OrderTypeAutoOrder OrderType = "AutoOrder"

	OrderTypeImport OrderType = "Import"

	OrderTypeBackOrder OrderType = "BackOrder"

	OrderTypeReplacementOrder OrderType = "ReplacementOrder"

	OrderTypeReturnOrder OrderType = "ReturnOrder"

	OrderTypeWebAutoOrder OrderType = "WebAutoOrder"

	OrderTypeTicketSystem OrderType = "TicketSystem"

	OrderTypeAPIOrder OrderType = "APIOrder"

	OrderTypeBackOrderParent OrderType = "BackOrderParent"

	OrderTypeChildOrder OrderType = "ChildOrder"

	OrderTypeOther1 OrderType = "Other1"

	OrderTypeOther2 OrderType = "Other2"

	OrderTypeOther3 OrderType = "Other3"

	OrderTypeOther4 OrderType = "Other4"
)

type FrequencyType string

const (
	FrequencyTypeWeekly FrequencyType = "Weekly"

	FrequencyTypeBiWeekly FrequencyType = "BiWeekly"

	FrequencyTypeMonthly FrequencyType = "Monthly"

	FrequencyTypeBiMonthly FrequencyType = "BiMonthly"

	FrequencyTypeQuarterly FrequencyType = "Quarterly"

	FrequencyTypeSemiYearly FrequencyType = "SemiYearly"

	FrequencyTypeYearly FrequencyType = "Yearly"

	FrequencyTypeEveryFourWeeks FrequencyType = "EveryFourWeeks"

	FrequencyTypeEverySixWeeks FrequencyType = "EverySixWeeks"

	FrequencyTypeEveryEightWeeks FrequencyType = "EveryEightWeeks"

	FrequencyTypeEveryTwelveWeeks FrequencyType = "EveryTwelveWeeks"

	FrequencyTypeSpecificDays FrequencyType = "SpecificDays"
)

type AutoOrderPaymentType string

const (
	AutoOrderPaymentTypePrimaryCreditCard AutoOrderPaymentType = "PrimaryCreditCard"

	AutoOrderPaymentTypeSecondaryCreditCard AutoOrderPaymentType = "SecondaryCreditCard"

	AutoOrderPaymentTypeCheckingAccount AutoOrderPaymentType = "CheckingAccount"

	AutoOrderPaymentTypeWillSendPayment AutoOrderPaymentType = "WillSendPayment"

	AutoOrderPaymentTypeBankDraft AutoOrderPaymentType = "BankDraft"

	AutoOrderPaymentTypePrimaryWalletAccount AutoOrderPaymentType = "PrimaryWalletAccount"

	AutoOrderPaymentTypeSecondaryWalletAccount AutoOrderPaymentType = "SecondaryWalletAccount"
)

type AutoOrderProcessType string

const (
	AutoOrderProcessTypeAlwaysProcess AutoOrderProcessType = "AlwaysProcess"

	AutoOrderProcessTypeConditional AutoOrderProcessType = "Conditional"
)

type OrderStatusType string

const (
	OrderStatusTypeIncomplete OrderStatusType = "Incomplete"

	OrderStatusTypePending OrderStatusType = "Pending"

	OrderStatusTypeCCDeclined OrderStatusType = "CCDeclined"

	OrderStatusTypeACHDeclined OrderStatusType = "ACHDeclined"

	OrderStatusTypeCanceled OrderStatusType = "Canceled"

	OrderStatusTypeCCPending OrderStatusType = "CCPending"

	OrderStatusTypeACHPending OrderStatusType = "ACHPending"

	OrderStatusTypeAccepted OrderStatusType = "Accepted"

	OrderStatusTypePrinted OrderStatusType = "Printed"

	OrderStatusTypeShipped OrderStatusType = "Shipped"

	OrderStatusTypePendingInventory OrderStatusType = "PendingInventory"
)

type DepositAccountType string

const (
	DepositAccountTypeChecking DepositAccountType = "Checking"

	DepositAccountTypeSaving DepositAccountType = "Saving"

	DepositAccountTypeBusiness DepositAccountType = "Business"

	DepositAccountTypePersonal DepositAccountType = "Personal"
)

type BankAccountType string

const (
	BankAccountTypeCheckingPersonal BankAccountType = "CheckingPersonal"

	BankAccountTypeCheckingBusiness BankAccountType = "CheckingBusiness"

	BankAccountTypeSavingsPersonal BankAccountType = "SavingsPersonal"

	BankAccountTypeSavingsBusiness BankAccountType = "SavingsBusiness"
)

type AccountWalletType string

const (
	AccountWalletTypePrimary AccountWalletType = "Primary"

	AccountWalletTypeSecondary AccountWalletType = "Secondary"
)

type PaymentType string

const (
	PaymentTypeCash PaymentType = "Cash"

	PaymentTypeCreditCard PaymentType = "CreditCard"

	PaymentTypeCheck PaymentType = "Check"

	PaymentTypeCreditVoucher PaymentType = "CreditVoucher"

	PaymentTypeNet30 PaymentType = "Net30"

	PaymentTypeNet60 PaymentType = "Net60"

	PaymentTypeACHDebit PaymentType = "ACHDebit"

	PaymentTypeUseCredit PaymentType = "UseCredit"

	PaymentTypeBankDraft PaymentType = "BankDraft"

	PaymentTypeBankWire PaymentType = "BankWire"

	PaymentTypePointRedemtion PaymentType = "PointRedemtion"

	PaymentTypeCOD PaymentType = "COD"

	PaymentTypeMoneyOrder PaymentType = "MoneyOrder"

	PaymentTypeBankDeposit PaymentType = "BankDeposit"

	PaymentTypeOther1 PaymentType = "Other1"

	PaymentTypeOther2 PaymentType = "Other2"

	PaymentTypeOther3 PaymentType = "Other3"

	PaymentTypeWallet PaymentType = "Wallet"

	PaymentTypeOther4 PaymentType = "Other4"

	PaymentTypeOther5 PaymentType = "Other5"
)

type MailForderType string

const (
	MailForderTypeSubFolder MailForderType = "SubFolder"

	MailForderTypeInbox MailForderType = "Inbox"

	MailForderTypeSentItems MailForderType = "SentItems"

	MailForderTypeDrafts MailForderType = "Drafts"

	MailForderTypeDeletedItems MailForderType = "DeletedItems"

	MailForderTypeJunkMail MailForderType = "JunkMail"

	MailForderTypeFaxInbox MailForderType = "FaxInbox"

	MailForderTypeSentFaxes MailForderType = "SentFaxes"

	MailForderTypePersonalFolders MailForderType = "PersonalFolders"
)

type MailPriority string

const (
	MailPriorityHigh MailPriority = "High"

	MailPriorityLow MailPriority = "Low"

	MailPriorityNormal MailPriority = "Normal"
)

type MailStatusType string

const (
	MailStatusTypeNew MailStatusType = "New"

	MailStatusTypeRead MailStatusType = "Read"

	MailStatusTypeForwarded MailStatusType = "Forwarded"

	MailStatusTypeReplied MailStatusType = "Replied"
)

type ResultStatus string

const (
	ResultStatusSuccess ResultStatus = "Success"

	ResultStatusFailure ResultStatus = "Failure"
)

type PropertyType string

const (
	PropertyTypeInteger PropertyType = "Integer"

	PropertyTypeDateTime PropertyType = "DateTime"

	PropertyTypeDecimal PropertyType = "Decimal"

	PropertyTypeBoolean PropertyType = "Boolean"

	PropertyTypeString PropertyType = "String"

	PropertyTypeBinary PropertyType = "Binary"

	PropertyTypeGuid PropertyType = "Guid"
)

type SandboxType string

const (
	SandboxTypeSandbox SandboxType = "Sandbox"

	SandboxTypeCommission SandboxType = "Commission"

	SandboxTypePremium SandboxType = "Premium"
)

type CustomerSiteImageType string

const (
	CustomerSiteImageTypePrimary CustomerSiteImageType = "Primary"

	CustomerSiteImageTypeSecondary CustomerSiteImageType = "Secondary"
)

type AutoOrderStatusType string

const (
	AutoOrderStatusTypeActive AutoOrderStatusType = "Active"

	AutoOrderStatusTypeInactive AutoOrderStatusType = "Inactive"

	AutoOrderStatusTypeDeleted AutoOrderStatusType = "Deleted"
)

type BinaryPlacementType string

const (
	BinaryPlacementTypeStrategicPlacement BinaryPlacementType = "StrategicPlacement"

	BinaryPlacementTypeBuildLeft BinaryPlacementType = "BuildLeft"

	BinaryPlacementTypeBuildRight BinaryPlacementType = "BuildRight"

	BinaryPlacementTypeBalancedBuild BinaryPlacementType = "BalancedBuild"

	BinaryPlacementTypeEvenFill BinaryPlacementType = "EvenFill"

	BinaryPlacementTypeWeakLeg BinaryPlacementType = "WeakLeg"

	BinaryPlacementTypeEnrollerPreference BinaryPlacementType = "EnrollerPreference"

	BinaryPlacementTypeLeftEvenFill BinaryPlacementType = "LeftEvenFill"

	BinaryPlacementTypeRightEvenFill BinaryPlacementType = "RightEvenFill"

	BinaryPlacementTypeLesserVolumeLeg BinaryPlacementType = "LesserVolumeLeg"

	BinaryPlacementTypeLesserVolumeLegOutside BinaryPlacementType = "LesserVolumeLegOutside"

	BinaryPlacementTypeStrongLegOutside BinaryPlacementType = "StrongLegOutside"

	BinaryPlacementTypeLesserVolumeLegEvenFill BinaryPlacementType = "LesserVolumeLegEvenFill"

	BinaryPlacementTypeGreaterVolumeLegEvenFill BinaryPlacementType = "GreaterVolumeLegEvenFill"

	BinaryPlacementTypeInsertRightMoveDownline BinaryPlacementType = "InsertRightMoveDownline"

	BinaryPlacementTypeInsertLeftMoveDownline BinaryPlacementType = "InsertLeftMoveDownline"

	BinaryPlacementTypeBuildTeamLeg BinaryPlacementType = "BuildTeamLeg"
)

type TreeType string

const (
	TreeTypeEnroller TreeType = "Enroller"

	TreeTypeUniLevel TreeType = "UniLevel"

	TreeTypeBinary TreeType = "Binary"

	TreeTypeMatrix TreeType = "Matrix"

	TreeTypeStack TreeType = "Stack"
)

type NumericCompareType string

const (
	NumericCompareTypeEquals NumericCompareType = "Equals"

	NumericCompareTypeGreaterThan NumericCompareType = "GreaterThan"

	NumericCompareTypeLessThan NumericCompareType = "LessThan"
)

type PointTransactionType string

const (
	PointTransactionTypeRedemption PointTransactionType = "Redemption"

	PointTransactionTypeAdjustment PointTransactionType = "Adjustment"
)

type SubscriptionStatus string

const (
	SubscriptionStatusNotFound SubscriptionStatus = "NotFound"

	SubscriptionStatusActive SubscriptionStatus = "Active"

	SubscriptionStatusExpired SubscriptionStatus = "Expired"

	SubscriptionStatusCancelled SubscriptionStatus = "Cancelled"
)

type InventoryStatusType string

const (
	InventoryStatusTypeAvailable InventoryStatusType = "Available"

	InventoryStatusTypeOnBackOrder InventoryStatusType = "OnBackOrder"

	InventoryStatusTypeOutOfStock InventoryStatusType = "OutOfStock"

	InventoryStatusTypeDiscontinued InventoryStatusType = "Discontinued"
)

type NewsWebSettings string

const (
	NewsWebSettingsAccessAvailable NewsWebSettings = "AccessAvailable"

	NewsWebSettingsAccessNotAvailable NewsWebSettings = "AccessNotAvailable"
)

type NewsCompanySettings string

const (
	NewsCompanySettingsAccessAllUsers NewsCompanySettings = "AccessAllUsers"

	NewsCompanySettingsAccessByDepartment NewsCompanySettings = "AccessByDepartment"

	NewsCompanySettingsAccessNotAvailable NewsCompanySettings = "AccessNotAvailable"
)

type CalendarItemType string

const (
	CalendarItemTypeAppointment CalendarItemType = "Appointment"

	CalendarItemTypeToDo CalendarItemType = "ToDo"

	CalendarItemTypeAnniversary CalendarItemType = "Anniversary"
)

type CalendarItemStatusType string

const (
	CalendarItemStatusTypeOpen CalendarItemStatusType = "Open"

	CalendarItemStatusTypeClosed CalendarItemStatusType = "Closed"
)

type CalendarItemPriorityType string

const (
	CalendarItemPriorityTypeNone CalendarItemPriorityType = "None"

	CalendarItemPriorityTypeHigh CalendarItemPriorityType = "High"

	CalendarItemPriorityTypeMedium CalendarItemPriorityType = "Medium"

	CalendarItemPriorityTypeLow CalendarItemPriorityType = "Low"
)

type ContactPhoneType string

const (
	ContactPhoneTypeOffice ContactPhoneType = "Office"

	ContactPhoneTypeMobile ContactPhoneType = "Mobile"

	ContactPhoneTypeHome ContactPhoneType = "Home"
)

type CreateEmailRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	MailFolderType *MailForderType `xml:"MailFolderType,omitempty" json:"MailFolderType,omitempty" yaml:"MailFolderType,omitempty"`

	Priority *MailPriority `xml:"Priority,omitempty" json:"Priority,omitempty" yaml:"Priority,omitempty"`

	MailStatusType *MailStatusType `xml:"MailStatusType,omitempty" json:"MailStatusType,omitempty" yaml:"MailStatusType,omitempty"`

	Subject string `xml:"Subject,omitempty" json:"Subject,omitempty" yaml:"Subject,omitempty"`

	MailFrom string `xml:"MailFrom,omitempty" json:"MailFrom,omitempty" yaml:"MailFrom,omitempty"`

	MailTo string `xml:"MailTo,omitempty" json:"MailTo,omitempty" yaml:"MailTo,omitempty"`

	ReplyTo string `xml:"ReplyTo,omitempty" json:"ReplyTo,omitempty" yaml:"ReplyTo,omitempty"`

	MailCC string `xml:"MailCC,omitempty" json:"MailCC,omitempty" yaml:"MailCC,omitempty"`

	MailBCC string `xml:"MailBCC,omitempty" json:"MailBCC,omitempty" yaml:"MailBCC,omitempty"`

	Content string `xml:"Content,omitempty" json:"Content,omitempty" yaml:"Content,omitempty"`

	SmtpServer string `xml:"SmtpServer,omitempty" json:"SmtpServer,omitempty" yaml:"SmtpServer,omitempty"`

	Attachments *ArrayOfEmailAttachment `xml:"Attachments,omitempty" json:"Attachments,omitempty" yaml:"Attachments,omitempty"`

	ForwardedAttachments *ArrayOfForwardedAttachment `xml:"ForwardedAttachments,omitempty" json:"ForwardedAttachments,omitempty" yaml:"ForwardedAttachments,omitempty"`
}

type ApiRequest struct {
}

type DeleteOrderDetailRequest struct {
	*ApiRequest

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	OrderLine int32 `xml:"OrderLine,omitempty" json:"OrderLine,omitempty" yaml:"OrderLine,omitempty"`
}

type UpdateOrderDetailRequest struct {
	*ApiRequest

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	ItemID int32 `xml:"ItemID,omitempty" json:"ItemID,omitempty" yaml:"ItemID,omitempty"`

	ItemCode string `xml:"ItemCode,omitempty" json:"ItemCode,omitempty" yaml:"ItemCode,omitempty"`

	OrderLine int32 `xml:"OrderLine,omitempty" json:"OrderLine,omitempty" yaml:"OrderLine,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	Qty float64 `xml:"Qty,omitempty" json:"Qty,omitempty" yaml:"Qty,omitempty"`

	PriceEach float64 `xml:"PriceEach,omitempty" json:"PriceEach,omitempty" yaml:"PriceEach,omitempty"`

	PriceExt float64 `xml:"PriceExt,omitempty" json:"PriceExt,omitempty" yaml:"PriceExt,omitempty"`

	BVEach float64 `xml:"BVEach,omitempty" json:"BVEach,omitempty" yaml:"BVEach,omitempty"`

	BusinessVolume float64 `xml:"BusinessVolume,omitempty" json:"BusinessVolume,omitempty" yaml:"BusinessVolume,omitempty"`

	CVEach float64 `xml:"CVEach,omitempty" json:"CVEach,omitempty" yaml:"CVEach,omitempty"`

	CommissionableVolume float64 `xml:"CommissionableVolume,omitempty" json:"CommissionableVolume,omitempty" yaml:"CommissionableVolume,omitempty"`

	ShippingPriceEach float64 `xml:"ShippingPriceEach,omitempty" json:"ShippingPriceEach,omitempty" yaml:"ShippingPriceEach,omitempty"`

	ChargeShippingOn float64 `xml:"ChargeShippingOn,omitempty" json:"ChargeShippingOn,omitempty" yaml:"ChargeShippingOn,omitempty"`

	IsTaxedInRegion bool `xml:"IsTaxedInRegion,omitempty" json:"IsTaxedInRegion,omitempty" yaml:"IsTaxedInRegion,omitempty"`

	IsTaxedInRegionFed bool `xml:"IsTaxedInRegionFed,omitempty" json:"IsTaxedInRegionFed,omitempty" yaml:"IsTaxedInRegionFed,omitempty"`

	IsTaxedInRegionState bool `xml:"IsTaxedInRegionState,omitempty" json:"IsTaxedInRegionState,omitempty" yaml:"IsTaxedInRegionState,omitempty"`

	TaxablePriceEach float64 `xml:"TaxablePriceEach,omitempty" json:"TaxablePriceEach,omitempty" yaml:"TaxablePriceEach,omitempty"`

	Taxable float64 `xml:"Taxable,omitempty" json:"Taxable,omitempty" yaml:"Taxable,omitempty"`

	CombinedTax float64 `xml:"CombinedTax,omitempty" json:"CombinedTax,omitempty" yaml:"CombinedTax,omitempty"`

	FedTax float64 `xml:"FedTax,omitempty" json:"FedTax,omitempty" yaml:"FedTax,omitempty"`

	StateTax float64 `xml:"StateTax,omitempty" json:"StateTax,omitempty" yaml:"StateTax,omitempty"`

	CityTax float64 `xml:"CityTax,omitempty" json:"CityTax,omitempty" yaml:"CityTax,omitempty"`

	CityLocalTax float64 `xml:"CityLocalTax,omitempty" json:"CityLocalTax,omitempty" yaml:"CityLocalTax,omitempty"`

	CountyTax float64 `xml:"CountyTax,omitempty" json:"CountyTax,omitempty" yaml:"CountyTax,omitempty"`

	CountyLocalTax float64 `xml:"CountyLocalTax,omitempty" json:"CountyLocalTax,omitempty" yaml:"CountyLocalTax,omitempty"`

	ManualTax float64 `xml:"ManualTax,omitempty" json:"ManualTax,omitempty" yaml:"ManualTax,omitempty"`

	IsBackOrder bool `xml:"IsBackOrder,omitempty" json:"IsBackOrder,omitempty" yaml:"IsBackOrder,omitempty"`

	WeightEach float64 `xml:"WeightEach,omitempty" json:"WeightEach,omitempty" yaml:"WeightEach,omitempty"`

	Other1Each float64 `xml:"Other1Each,omitempty" json:"Other1Each,omitempty" yaml:"Other1Each,omitempty"`

	Other1 float64 `xml:"Other1,omitempty" json:"Other1,omitempty" yaml:"Other1,omitempty"`

	Other2Each float64 `xml:"Other2Each,omitempty" json:"Other2Each,omitempty" yaml:"Other2Each,omitempty"`

	Other2 float64 `xml:"Other2,omitempty" json:"Other2,omitempty" yaml:"Other2,omitempty"`

	Other3Each float64 `xml:"Other3Each,omitempty" json:"Other3Each,omitempty" yaml:"Other3Each,omitempty"`

	Other3 float64 `xml:"Other3,omitempty" json:"Other3,omitempty" yaml:"Other3,omitempty"`

	Other4Each float64 `xml:"Other4Each,omitempty" json:"Other4Each,omitempty" yaml:"Other4Each,omitempty"`

	Other4 float64 `xml:"Other4,omitempty" json:"Other4,omitempty" yaml:"Other4,omitempty"`

	Other5Each float64 `xml:"Other5Each,omitempty" json:"Other5Each,omitempty" yaml:"Other5Each,omitempty"`

	Other5 float64 `xml:"Other5,omitempty" json:"Other5,omitempty" yaml:"Other5,omitempty"`

	Other6Each float64 `xml:"Other6Each,omitempty" json:"Other6Each,omitempty" yaml:"Other6Each,omitempty"`

	Other6 float64 `xml:"Other6,omitempty" json:"Other6,omitempty" yaml:"Other6,omitempty"`

	Other7Each float64 `xml:"Other7Each,omitempty" json:"Other7Each,omitempty" yaml:"Other7Each,omitempty"`

	Other7 float64 `xml:"Other7,omitempty" json:"Other7,omitempty" yaml:"Other7,omitempty"`

	Other8Each float64 `xml:"Other8Each,omitempty" json:"Other8Each,omitempty" yaml:"Other8Each,omitempty"`

	Other8 float64 `xml:"Other8,omitempty" json:"Other8,omitempty" yaml:"Other8,omitempty"`

	Other9Each float64 `xml:"Other9Each,omitempty" json:"Other9Each,omitempty" yaml:"Other9Each,omitempty"`

	Other9 float64 `xml:"Other9,omitempty" json:"Other9,omitempty" yaml:"Other9,omitempty"`

	Other10Each float64 `xml:"Other10Each,omitempty" json:"Other10Each,omitempty" yaml:"Other10Each,omitempty"`

	Other10 float64 `xml:"Other10,omitempty" json:"Other10,omitempty" yaml:"Other10,omitempty"`

	DiscountExt float64 `xml:"DiscountExt,omitempty" json:"DiscountExt,omitempty" yaml:"DiscountExt,omitempty"`

	OriginalTaxableEach float64 `xml:"OriginalTaxableEach,omitempty" json:"OriginalTaxableEach,omitempty" yaml:"OriginalTaxableEach,omitempty"`

	OriginalBVEach float64 `xml:"OriginalBVEach,omitempty" json:"OriginalBVEach,omitempty" yaml:"OriginalBVEach,omitempty"`

	OriginalCVEach float64 `xml:"OriginalCVEach,omitempty" json:"OriginalCVEach,omitempty" yaml:"OriginalCVEach,omitempty"`

	StateTaxable float64 `xml:"StateTaxable,omitempty" json:"StateTaxable,omitempty" yaml:"StateTaxable,omitempty"`

	IsStateTaxOverride bool `xml:"IsStateTaxOverride,omitempty" json:"IsStateTaxOverride,omitempty" yaml:"IsStateTaxOverride,omitempty"`

	DynamicKitItemID int32 `xml:"DynamicKitItemID,omitempty" json:"DynamicKitItemID,omitempty" yaml:"DynamicKitItemID,omitempty"`

	HandlingFee float64 `xml:"HandlingFee,omitempty" json:"HandlingFee,omitempty" yaml:"HandlingFee,omitempty"`

	Reference1 string `xml:"Reference1,omitempty" json:"Reference1,omitempty" yaml:"Reference1,omitempty"`
}

type CreateOrderDetailRequest struct {
	*ApiRequest

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	OrderLine int32 `xml:"OrderLine,omitempty" json:"OrderLine,omitempty" yaml:"OrderLine,omitempty"`

	ItemCode string `xml:"ItemCode,omitempty" json:"ItemCode,omitempty" yaml:"ItemCode,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	Qty float64 `xml:"Qty,omitempty" json:"Qty,omitempty" yaml:"Qty,omitempty"`

	PriceEach float64 `xml:"PriceEach,omitempty" json:"PriceEach,omitempty" yaml:"PriceEach,omitempty"`

	PriceExt float64 `xml:"PriceExt,omitempty" json:"PriceExt,omitempty" yaml:"PriceExt,omitempty"`

	BVEach float64 `xml:"BVEach,omitempty" json:"BVEach,omitempty" yaml:"BVEach,omitempty"`

	BusinessVolume float64 `xml:"BusinessVolume,omitempty" json:"BusinessVolume,omitempty" yaml:"BusinessVolume,omitempty"`

	CVEach float64 `xml:"CVEach,omitempty" json:"CVEach,omitempty" yaml:"CVEach,omitempty"`

	CommissionableVolume float64 `xml:"CommissionableVolume,omitempty" json:"CommissionableVolume,omitempty" yaml:"CommissionableVolume,omitempty"`

	ShippingPriceEach float64 `xml:"ShippingPriceEach,omitempty" json:"ShippingPriceEach,omitempty" yaml:"ShippingPriceEach,omitempty"`

	ChargeShippingOn float64 `xml:"ChargeShippingOn,omitempty" json:"ChargeShippingOn,omitempty" yaml:"ChargeShippingOn,omitempty"`

	IsTaxedInRegion bool `xml:"IsTaxedInRegion,omitempty" json:"IsTaxedInRegion,omitempty" yaml:"IsTaxedInRegion,omitempty"`

	IsTaxedInRegionFed bool `xml:"IsTaxedInRegionFed,omitempty" json:"IsTaxedInRegionFed,omitempty" yaml:"IsTaxedInRegionFed,omitempty"`

	IsTaxedInRegionState bool `xml:"IsTaxedInRegionState,omitempty" json:"IsTaxedInRegionState,omitempty" yaml:"IsTaxedInRegionState,omitempty"`

	TaxablePriceEach float64 `xml:"TaxablePriceEach,omitempty" json:"TaxablePriceEach,omitempty" yaml:"TaxablePriceEach,omitempty"`

	Taxable float64 `xml:"Taxable,omitempty" json:"Taxable,omitempty" yaml:"Taxable,omitempty"`

	CombinedTax float64 `xml:"CombinedTax,omitempty" json:"CombinedTax,omitempty" yaml:"CombinedTax,omitempty"`

	FedTax float64 `xml:"FedTax,omitempty" json:"FedTax,omitempty" yaml:"FedTax,omitempty"`

	StateTax float64 `xml:"StateTax,omitempty" json:"StateTax,omitempty" yaml:"StateTax,omitempty"`

	CityTax float64 `xml:"CityTax,omitempty" json:"CityTax,omitempty" yaml:"CityTax,omitempty"`

	CityLocalTax float64 `xml:"CityLocalTax,omitempty" json:"CityLocalTax,omitempty" yaml:"CityLocalTax,omitempty"`

	CountyTax float64 `xml:"CountyTax,omitempty" json:"CountyTax,omitempty" yaml:"CountyTax,omitempty"`

	CountyLocalTax float64 `xml:"CountyLocalTax,omitempty" json:"CountyLocalTax,omitempty" yaml:"CountyLocalTax,omitempty"`

	ManualTax float64 `xml:"ManualTax,omitempty" json:"ManualTax,omitempty" yaml:"ManualTax,omitempty"`

	IsBackOrder bool `xml:"IsBackOrder,omitempty" json:"IsBackOrder,omitempty" yaml:"IsBackOrder,omitempty"`

	WeightEach float64 `xml:"WeightEach,omitempty" json:"WeightEach,omitempty" yaml:"WeightEach,omitempty"`

	Other1Each float64 `xml:"Other1Each,omitempty" json:"Other1Each,omitempty" yaml:"Other1Each,omitempty"`

	Other1 float64 `xml:"Other1,omitempty" json:"Other1,omitempty" yaml:"Other1,omitempty"`

	Other2Each float64 `xml:"Other2Each,omitempty" json:"Other2Each,omitempty" yaml:"Other2Each,omitempty"`

	Other2 float64 `xml:"Other2,omitempty" json:"Other2,omitempty" yaml:"Other2,omitempty"`

	Other3Each float64 `xml:"Other3Each,omitempty" json:"Other3Each,omitempty" yaml:"Other3Each,omitempty"`

	Other3 float64 `xml:"Other3,omitempty" json:"Other3,omitempty" yaml:"Other3,omitempty"`

	Other4Each float64 `xml:"Other4Each,omitempty" json:"Other4Each,omitempty" yaml:"Other4Each,omitempty"`

	Other4 float64 `xml:"Other4,omitempty" json:"Other4,omitempty" yaml:"Other4,omitempty"`

	Other5Each float64 `xml:"Other5Each,omitempty" json:"Other5Each,omitempty" yaml:"Other5Each,omitempty"`

	Other5 float64 `xml:"Other5,omitempty" json:"Other5,omitempty" yaml:"Other5,omitempty"`

	Other6Each float64 `xml:"Other6Each,omitempty" json:"Other6Each,omitempty" yaml:"Other6Each,omitempty"`

	Other6 float64 `xml:"Other6,omitempty" json:"Other6,omitempty" yaml:"Other6,omitempty"`

	Other7Each float64 `xml:"Other7Each,omitempty" json:"Other7Each,omitempty" yaml:"Other7Each,omitempty"`

	Other7 float64 `xml:"Other7,omitempty" json:"Other7,omitempty" yaml:"Other7,omitempty"`

	Other8Each float64 `xml:"Other8Each,omitempty" json:"Other8Each,omitempty" yaml:"Other8Each,omitempty"`

	Other8 float64 `xml:"Other8,omitempty" json:"Other8,omitempty" yaml:"Other8,omitempty"`

	Other9Each float64 `xml:"Other9Each,omitempty" json:"Other9Each,omitempty" yaml:"Other9Each,omitempty"`

	Other9 float64 `xml:"Other9,omitempty" json:"Other9,omitempty" yaml:"Other9,omitempty"`

	Other10Each float64 `xml:"Other10Each,omitempty" json:"Other10Each,omitempty" yaml:"Other10Each,omitempty"`

	Other10 float64 `xml:"Other10,omitempty" json:"Other10,omitempty" yaml:"Other10,omitempty"`

	DiscountExt float64 `xml:"DiscountExt,omitempty" json:"DiscountExt,omitempty" yaml:"DiscountExt,omitempty"`

	OriginalTaxableEach float64 `xml:"OriginalTaxableEach,omitempty" json:"OriginalTaxableEach,omitempty" yaml:"OriginalTaxableEach,omitempty"`

	OriginalBVEach float64 `xml:"OriginalBVEach,omitempty" json:"OriginalBVEach,omitempty" yaml:"OriginalBVEach,omitempty"`

	OriginalCVEach float64 `xml:"OriginalCVEach,omitempty" json:"OriginalCVEach,omitempty" yaml:"OriginalCVEach,omitempty"`

	StateTaxable float64 `xml:"StateTaxable,omitempty" json:"StateTaxable,omitempty" yaml:"StateTaxable,omitempty"`

	IsStateTaxOverride bool `xml:"IsStateTaxOverride,omitempty" json:"IsStateTaxOverride,omitempty" yaml:"IsStateTaxOverride,omitempty"`

	DynamicKitItemID int32 `xml:"DynamicKitItemID,omitempty" json:"DynamicKitItemID,omitempty" yaml:"DynamicKitItemID,omitempty"`

	HandlingFee float64 `xml:"HandlingFee,omitempty" json:"HandlingFee,omitempty" yaml:"HandlingFee,omitempty"`

	Reference1 string `xml:"Reference1,omitempty" json:"Reference1,omitempty" yaml:"Reference1,omitempty"`
}

type BaseAuthorizeOnlyCreditCardTokenRequest struct {
	*ApiRequest
}

type AuthorizeOnlyCreditCardTokenOnFileRequest struct {
	*BaseAuthorizeOnlyCreditCardTokenRequest

	CreditCardAccountType *AccountCreditCardType `xml:"CreditCardAccountType,omitempty" json:"CreditCardAccountType,omitempty" yaml:"CreditCardAccountType,omitempty"`

	WarehouseID int32 `xml:"WarehouseID,omitempty" json:"WarehouseID,omitempty" yaml:"WarehouseID,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty" yaml:"CurrencyCode,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	Amount float64 `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`
}

type AuthorizeOnlyCreditCardTokenRequest struct {
	*BaseAuthorizeOnlyCreditCardTokenRequest

	CreditCardToken string `xml:"CreditCardToken,omitempty" json:"CreditCardToken,omitempty" yaml:"CreditCardToken,omitempty"`

	Amount float64 `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`

	BillingName string `xml:"BillingName,omitempty" json:"BillingName,omitempty" yaml:"BillingName,omitempty"`

	BillingAddress string `xml:"BillingAddress,omitempty" json:"BillingAddress,omitempty" yaml:"BillingAddress,omitempty"`

	BillingCity string `xml:"BillingCity,omitempty" json:"BillingCity,omitempty" yaml:"BillingCity,omitempty"`

	BillingState string `xml:"BillingState,omitempty" json:"BillingState,omitempty" yaml:"BillingState,omitempty"`

	BillingZip string `xml:"BillingZip,omitempty" json:"BillingZip,omitempty" yaml:"BillingZip,omitempty"`

	BillingCountry string `xml:"BillingCountry,omitempty" json:"BillingCountry,omitempty" yaml:"BillingCountry,omitempty"`

	WarehouseID int32 `xml:"WarehouseID,omitempty" json:"WarehouseID,omitempty" yaml:"WarehouseID,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty" yaml:"CurrencyCode,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	CvcCode string `xml:"CvcCode,omitempty" json:"CvcCode,omitempty" yaml:"CvcCode,omitempty"`
}

type SendEmailRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	MailFrom string `xml:"MailFrom,omitempty" json:"MailFrom,omitempty" yaml:"MailFrom,omitempty"`

	MailTo string `xml:"MailTo,omitempty" json:"MailTo,omitempty" yaml:"MailTo,omitempty"`

	Subject string `xml:"Subject,omitempty" json:"Subject,omitempty" yaml:"Subject,omitempty"`

	Body string `xml:"Body,omitempty" json:"Body,omitempty" yaml:"Body,omitempty"`
}

type SetItemKitMembersRequest struct {
	*ApiRequest

	ParentItemCode string `xml:"ParentItemCode,omitempty" json:"ParentItemCode,omitempty" yaml:"ParentItemCode,omitempty"`

	ItemKitMembers *ArrayOfKitMember `xml:"ItemKitMembers,omitempty" json:"ItemKitMembers,omitempty" yaml:"ItemKitMembers,omitempty"`
}

type ArrayOfKitMember struct {
	KitMember []*KitMember `xml:"KitMember,omitempty" json:"KitMember,omitempty" yaml:"KitMember,omitempty"`
}

type KitMember struct {
	ItemID int32 `xml:"ItemID,omitempty" json:"ItemID,omitempty" yaml:"ItemID,omitempty"`

	ItemCode string `xml:"ItemCode,omitempty" json:"ItemCode,omitempty" yaml:"ItemCode,omitempty"`

	Quantity int32 `xml:"Quantity,omitempty" json:"Quantity,omitempty" yaml:"Quantity,omitempty"`
}

type GetFileContentsRequest struct {
	*ApiRequest

	FileName string `xml:"FileName,omitempty" json:"FileName,omitempty" yaml:"FileName,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`
}

type GetFilesRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`
}

type ChargeGroupOrderCreditCardTokenRequest struct {
	*ApiRequest

	_orders *ArrayOfGroupOrder `xml:"_orders,omitempty" json:"_orders,omitempty" yaml:"_orders,omitempty"`

	CreditCardToken string `xml:"CreditCardToken,omitempty" json:"CreditCardToken,omitempty" yaml:"CreditCardToken,omitempty"`

	BillingName string `xml:"BillingName,omitempty" json:"BillingName,omitempty" yaml:"BillingName,omitempty"`

	BillingAddress string `xml:"BillingAddress,omitempty" json:"BillingAddress,omitempty" yaml:"BillingAddress,omitempty"`

	BillingAddress2 string `xml:"BillingAddress2,omitempty" json:"BillingAddress2,omitempty" yaml:"BillingAddress2,omitempty"`

	BillingCity string `xml:"BillingCity,omitempty" json:"BillingCity,omitempty" yaml:"BillingCity,omitempty"`

	BillingState string `xml:"BillingState,omitempty" json:"BillingState,omitempty" yaml:"BillingState,omitempty"`

	BillingZip string `xml:"BillingZip,omitempty" json:"BillingZip,omitempty" yaml:"BillingZip,omitempty"`

	BillingCountry string `xml:"BillingCountry,omitempty" json:"BillingCountry,omitempty" yaml:"BillingCountry,omitempty"`

	CvcCode string `xml:"CvcCode,omitempty" json:"CvcCode,omitempty" yaml:"CvcCode,omitempty"`

	IssueNumber string `xml:"IssueNumber,omitempty" json:"IssueNumber,omitempty" yaml:"IssueNumber,omitempty"`

	CreditCardType int32 `xml:"CreditCardType,omitempty" json:"CreditCardType,omitempty" yaml:"CreditCardType,omitempty"`

	MasterOrderID int32 `xml:"MasterOrderID,omitempty" json:"MasterOrderID,omitempty" yaml:"MasterOrderID,omitempty"`

	Orders *ArrayOfGroupOrder `xml:"Orders,omitempty" json:"Orders,omitempty" yaml:"Orders,omitempty"`

	MerchantWarehouseIDOverride int32 `xml:"MerchantWarehouseIDOverride,omitempty" json:"MerchantWarehouseIDOverride,omitempty" yaml:"MerchantWarehouseIDOverride,omitempty"`

	ClientIPAddress string `xml:"ClientIPAddress,omitempty" json:"ClientIPAddress,omitempty" yaml:"ClientIPAddress,omitempty"`

	OtherData1 string `xml:"OtherData1,omitempty" json:"OtherData1,omitempty" yaml:"OtherData1,omitempty"`

	OtherData2 string `xml:"OtherData2,omitempty" json:"OtherData2,omitempty" yaml:"OtherData2,omitempty"`

	OtherData3 string `xml:"OtherData3,omitempty" json:"OtherData3,omitempty" yaml:"OtherData3,omitempty"`

	OtherData4 string `xml:"OtherData4,omitempty" json:"OtherData4,omitempty" yaml:"OtherData4,omitempty"`

	OtherData5 string `xml:"OtherData5,omitempty" json:"OtherData5,omitempty" yaml:"OtherData5,omitempty"`

	OtherData6 string `xml:"OtherData6,omitempty" json:"OtherData6,omitempty" yaml:"OtherData6,omitempty"`

	OtherData7 string `xml:"OtherData7,omitempty" json:"OtherData7,omitempty" yaml:"OtherData7,omitempty"`

	OtherData8 string `xml:"OtherData8,omitempty" json:"OtherData8,omitempty" yaml:"OtherData8,omitempty"`

	OtherData9 string `xml:"OtherData9,omitempty" json:"OtherData9,omitempty" yaml:"OtherData9,omitempty"`

	OtherData10 string `xml:"OtherData10,omitempty" json:"OtherData10,omitempty" yaml:"OtherData10,omitempty"`

	PaymentMemo string `xml:"PaymentMemo,omitempty" json:"PaymentMemo,omitempty" yaml:"PaymentMemo,omitempty"`
}

type ArrayOfGroupOrder struct {
	GroupOrder []*GroupOrder `xml:"GroupOrder,omitempty" json:"GroupOrder,omitempty" yaml:"GroupOrder,omitempty"`
}

type GroupOrder struct {
	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	MaxAmount float64 `xml:"MaxAmount,omitempty" json:"MaxAmount,omitempty" yaml:"MaxAmount,omitempty"`
}

type DeleteCustomerExtendedRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	ExtendedGroupID int32 `xml:"ExtendedGroupID,omitempty" json:"ExtendedGroupID,omitempty" yaml:"ExtendedGroupID,omitempty"`

	CustomerExtendedID int32 `xml:"CustomerExtendedID,omitempty" json:"CustomerExtendedID,omitempty" yaml:"CustomerExtendedID,omitempty"`
}

type UpdateCustomerExtendedRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	ExtendedGroupID int32 `xml:"ExtendedGroupID,omitempty" json:"ExtendedGroupID,omitempty" yaml:"ExtendedGroupID,omitempty"`

	CustomerExtendedID int32 `xml:"CustomerExtendedID,omitempty" json:"CustomerExtendedID,omitempty" yaml:"CustomerExtendedID,omitempty"`

	Field1 string `xml:"Field1,omitempty" json:"Field1,omitempty" yaml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty" json:"Field2,omitempty" yaml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty" json:"Field3,omitempty" yaml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty" json:"Field4,omitempty" yaml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty" json:"Field5,omitempty" yaml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty" json:"Field6,omitempty" yaml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty" json:"Field7,omitempty" yaml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty" json:"Field8,omitempty" yaml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty" json:"Field9,omitempty" yaml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty" json:"Field10,omitempty" yaml:"Field10,omitempty"`

	Field11 string `xml:"Field11,omitempty" json:"Field11,omitempty" yaml:"Field11,omitempty"`

	Field12 string `xml:"Field12,omitempty" json:"Field12,omitempty" yaml:"Field12,omitempty"`

	Field13 string `xml:"Field13,omitempty" json:"Field13,omitempty" yaml:"Field13,omitempty"`

	Field14 string `xml:"Field14,omitempty" json:"Field14,omitempty" yaml:"Field14,omitempty"`

	Field15 string `xml:"Field15,omitempty" json:"Field15,omitempty" yaml:"Field15,omitempty"`

	Field16 string `xml:"Field16,omitempty" json:"Field16,omitempty" yaml:"Field16,omitempty"`

	Field17 string `xml:"Field17,omitempty" json:"Field17,omitempty" yaml:"Field17,omitempty"`

	Field18 string `xml:"Field18,omitempty" json:"Field18,omitempty" yaml:"Field18,omitempty"`

	Field19 string `xml:"Field19,omitempty" json:"Field19,omitempty" yaml:"Field19,omitempty"`

	Field20 string `xml:"Field20,omitempty" json:"Field20,omitempty" yaml:"Field20,omitempty"`

	Field21 string `xml:"Field21,omitempty" json:"Field21,omitempty" yaml:"Field21,omitempty"`

	Field22 string `xml:"Field22,omitempty" json:"Field22,omitempty" yaml:"Field22,omitempty"`

	Field23 string `xml:"Field23,omitempty" json:"Field23,omitempty" yaml:"Field23,omitempty"`

	Field24 string `xml:"Field24,omitempty" json:"Field24,omitempty" yaml:"Field24,omitempty"`

	Field25 string `xml:"Field25,omitempty" json:"Field25,omitempty" yaml:"Field25,omitempty"`

	Field26 string `xml:"Field26,omitempty" json:"Field26,omitempty" yaml:"Field26,omitempty"`

	Field27 string `xml:"Field27,omitempty" json:"Field27,omitempty" yaml:"Field27,omitempty"`

	Field28 string `xml:"Field28,omitempty" json:"Field28,omitempty" yaml:"Field28,omitempty"`

	Field29 string `xml:"Field29,omitempty" json:"Field29,omitempty" yaml:"Field29,omitempty"`

	Field30 string `xml:"Field30,omitempty" json:"Field30,omitempty" yaml:"Field30,omitempty"`
}

type CreateCustomerExtendedRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	ExtendedGroupID int32 `xml:"ExtendedGroupID,omitempty" json:"ExtendedGroupID,omitempty" yaml:"ExtendedGroupID,omitempty"`

	Field1 string `xml:"Field1,omitempty" json:"Field1,omitempty" yaml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty" json:"Field2,omitempty" yaml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty" json:"Field3,omitempty" yaml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty" json:"Field4,omitempty" yaml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty" json:"Field5,omitempty" yaml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty" json:"Field6,omitempty" yaml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty" json:"Field7,omitempty" yaml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty" json:"Field8,omitempty" yaml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty" json:"Field9,omitempty" yaml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty" json:"Field10,omitempty" yaml:"Field10,omitempty"`

	Field11 string `xml:"Field11,omitempty" json:"Field11,omitempty" yaml:"Field11,omitempty"`

	Field12 string `xml:"Field12,omitempty" json:"Field12,omitempty" yaml:"Field12,omitempty"`

	Field13 string `xml:"Field13,omitempty" json:"Field13,omitempty" yaml:"Field13,omitempty"`

	Field14 string `xml:"Field14,omitempty" json:"Field14,omitempty" yaml:"Field14,omitempty"`

	Field15 string `xml:"Field15,omitempty" json:"Field15,omitempty" yaml:"Field15,omitempty"`

	Field16 string `xml:"Field16,omitempty" json:"Field16,omitempty" yaml:"Field16,omitempty"`

	Field17 string `xml:"Field17,omitempty" json:"Field17,omitempty" yaml:"Field17,omitempty"`

	Field18 string `xml:"Field18,omitempty" json:"Field18,omitempty" yaml:"Field18,omitempty"`

	Field19 string `xml:"Field19,omitempty" json:"Field19,omitempty" yaml:"Field19,omitempty"`

	Field20 string `xml:"Field20,omitempty" json:"Field20,omitempty" yaml:"Field20,omitempty"`

	Field21 string `xml:"Field21,omitempty" json:"Field21,omitempty" yaml:"Field21,omitempty"`

	Field22 string `xml:"Field22,omitempty" json:"Field22,omitempty" yaml:"Field22,omitempty"`

	Field23 string `xml:"Field23,omitempty" json:"Field23,omitempty" yaml:"Field23,omitempty"`

	Field24 string `xml:"Field24,omitempty" json:"Field24,omitempty" yaml:"Field24,omitempty"`

	Field25 string `xml:"Field25,omitempty" json:"Field25,omitempty" yaml:"Field25,omitempty"`

	Field26 string `xml:"Field26,omitempty" json:"Field26,omitempty" yaml:"Field26,omitempty"`

	Field27 string `xml:"Field27,omitempty" json:"Field27,omitempty" yaml:"Field27,omitempty"`

	Field28 string `xml:"Field28,omitempty" json:"Field28,omitempty" yaml:"Field28,omitempty"`

	Field29 string `xml:"Field29,omitempty" json:"Field29,omitempty" yaml:"Field29,omitempty"`

	Field30 string `xml:"Field30,omitempty" json:"Field30,omitempty" yaml:"Field30,omitempty"`
}

type UpdatePartyRequest struct {
	*ApiRequest

	PartyID int32 `xml:"PartyID,omitempty" json:"PartyID,omitempty" yaml:"PartyID,omitempty"`

	PartyType int32 `xml:"PartyType,omitempty" json:"PartyType,omitempty" yaml:"PartyType,omitempty"`

	PartyStatusType int32 `xml:"PartyStatusType,omitempty" json:"PartyStatusType,omitempty" yaml:"PartyStatusType,omitempty"`

	HostID int32 `xml:"HostID,omitempty" json:"HostID,omitempty" yaml:"HostID,omitempty"`

	DistributorID int32 `xml:"DistributorID,omitempty" json:"DistributorID,omitempty" yaml:"DistributorID,omitempty"`

	StartDate time.Time `xml:"StartDate,omitempty" json:"StartDate,omitempty" yaml:"StartDate,omitempty"`

	CloseDate time.Time `xml:"CloseDate,omitempty" json:"CloseDate,omitempty" yaml:"CloseDate,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	EventStart time.Time `xml:"EventStart,omitempty" json:"EventStart,omitempty" yaml:"EventStart,omitempty"`

	EventEnd time.Time `xml:"EventEnd,omitempty" json:"EventEnd,omitempty" yaml:"EventEnd,omitempty"`

	LanguageID int32 `xml:"LanguageID,omitempty" json:"LanguageID,omitempty" yaml:"LanguageID,omitempty"`

	Information string `xml:"Information,omitempty" json:"Information,omitempty" yaml:"Information,omitempty"`

	Address *PartyAddress `xml:"Address,omitempty" json:"Address,omitempty" yaml:"Address,omitempty"`

	BookingPartyID int32 `xml:"BookingPartyID,omitempty" json:"BookingPartyID,omitempty" yaml:"BookingPartyID,omitempty"`

	Field1 string `xml:"Field1,omitempty" json:"Field1,omitempty" yaml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty" json:"Field2,omitempty" yaml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty" json:"Field3,omitempty" yaml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty" json:"Field4,omitempty" yaml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty" json:"Field5,omitempty" yaml:"Field5,omitempty"`
}

type PartyAddress struct {
	Address1 string `xml:"Address1,omitempty" json:"Address1,omitempty" yaml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty" json:"Address2,omitempty" yaml:"Address2,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty" json:"Zip,omitempty" yaml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`
}

type CreatePartyRequest struct {
	*ApiRequest

	PartyType int32 `xml:"PartyType,omitempty" json:"PartyType,omitempty" yaml:"PartyType,omitempty"`

	PartyStatusType int32 `xml:"PartyStatusType,omitempty" json:"PartyStatusType,omitempty" yaml:"PartyStatusType,omitempty"`

	HostID int32 `xml:"HostID,omitempty" json:"HostID,omitempty" yaml:"HostID,omitempty"`

	DistributorID int32 `xml:"DistributorID,omitempty" json:"DistributorID,omitempty" yaml:"DistributorID,omitempty"`

	StartDate time.Time `xml:"StartDate,omitempty" json:"StartDate,omitempty" yaml:"StartDate,omitempty"`

	CloseDate time.Time `xml:"CloseDate,omitempty" json:"CloseDate,omitempty" yaml:"CloseDate,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	EventStart time.Time `xml:"EventStart,omitempty" json:"EventStart,omitempty" yaml:"EventStart,omitempty"`

	EventEnd time.Time `xml:"EventEnd,omitempty" json:"EventEnd,omitempty" yaml:"EventEnd,omitempty"`

	LanguageID int32 `xml:"LanguageID,omitempty" json:"LanguageID,omitempty" yaml:"LanguageID,omitempty"`

	Information string `xml:"Information,omitempty" json:"Information,omitempty" yaml:"Information,omitempty"`

	Address *PartyAddress `xml:"Address,omitempty" json:"Address,omitempty" yaml:"Address,omitempty"`

	BookingPartyID int32 `xml:"BookingPartyID,omitempty" json:"BookingPartyID,omitempty" yaml:"BookingPartyID,omitempty"`

	Field1 string `xml:"Field1,omitempty" json:"Field1,omitempty" yaml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty" json:"Field2,omitempty" yaml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty" json:"Field3,omitempty" yaml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty" json:"Field4,omitempty" yaml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty" json:"Field5,omitempty" yaml:"Field5,omitempty"`
}

type CreateBillRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty" yaml:"CurrencyCode,omitempty"`

	IsOtherIncome bool `xml:"IsOtherIncome,omitempty" json:"IsOtherIncome,omitempty" yaml:"IsOtherIncome,omitempty"`

	DueDate time.Time `xml:"DueDate,omitempty" json:"DueDate,omitempty" yaml:"DueDate,omitempty"`

	Amount float64 `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`

	Reference string `xml:"Reference,omitempty" json:"Reference,omitempty" yaml:"Reference,omitempty"`

	Notes string `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`

	BillStatusTypeID int32 `xml:"BillStatusTypeID,omitempty" json:"BillStatusTypeID,omitempty" yaml:"BillStatusTypeID,omitempty"`

	PayableTypeIDOverride int32 `xml:"PayableTypeIDOverride,omitempty" json:"PayableTypeIDOverride,omitempty" yaml:"PayableTypeIDOverride,omitempty"`
}

type BaseCreatePayoutRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	BankAccountID int32 `xml:"BankAccountID,omitempty" json:"BankAccountID,omitempty" yaml:"BankAccountID,omitempty"`

	Reference string `xml:"Reference,omitempty" json:"Reference,omitempty" yaml:"Reference,omitempty"`

	TransactionNote string `xml:"TransactionNote,omitempty" json:"TransactionNote,omitempty" yaml:"TransactionNote,omitempty"`

	PaymentDate time.Time `xml:"PaymentDate,omitempty" json:"PaymentDate,omitempty" yaml:"PaymentDate,omitempty"`
}

type CreatePayoutRequest struct {
	*BaseCreatePayoutRequest

	BillIDs_ToPay *ArrayOfInt `xml:"BillIDs_ToPay,omitempty" json:"BillIDs_ToPay,omitempty" yaml:"BillIDs_ToPay,omitempty"`

	VendorPaymentTypeID int32 `xml:"VendorPaymentTypeID,omitempty" json:"VendorPaymentTypeID,omitempty" yaml:"VendorPaymentTypeID,omitempty"`
}

type ArrayOfInt struct {
	Int []int32 `xml:"int,omitempty" json:"int,omitempty" yaml:"int,omitempty"`
}

type CreateCustomerInquiryRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	Detail string `xml:"Detail,omitempty" json:"Detail,omitempty" yaml:"Detail,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	AssignToUser string `xml:"AssignToUser,omitempty" json:"AssignToUser,omitempty" yaml:"AssignToUser,omitempty"`

	CustomerInquiryStatusID int32 `xml:"CustomerInquiryStatusID,omitempty" json:"CustomerInquiryStatusID,omitempty" yaml:"CustomerInquiryStatusID,omitempty"`

	CustomerInquiryCategoryID int32 `xml:"CustomerInquiryCategoryID,omitempty" json:"CustomerInquiryCategoryID,omitempty" yaml:"CustomerInquiryCategoryID,omitempty"`
}

type CreateCustomerFileRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	FileName string `xml:"FileName,omitempty" json:"FileName,omitempty" yaml:"FileName,omitempty"`

	FileData []byte `xml:"FileData,omitempty" json:"FileData,omitempty" yaml:"FileData,omitempty"`

	OverwriteExistingFile bool `xml:"OverwriteExistingFile,omitempty" json:"OverwriteExistingFile,omitempty" yaml:"OverwriteExistingFile,omitempty"`
}

type SetItemCountryRegionRequest struct {
	*ApiRequest

	ItemCode string `xml:"ItemCode,omitempty" json:"ItemCode,omitempty" yaml:"ItemCode,omitempty"`

	CountryCode string `xml:"CountryCode,omitempty" json:"CountryCode,omitempty" yaml:"CountryCode,omitempty"`

	RegionCode string `xml:"RegionCode,omitempty" json:"RegionCode,omitempty" yaml:"RegionCode,omitempty"`

	Taxed bool `xml:"Taxed,omitempty" json:"Taxed,omitempty" yaml:"Taxed,omitempty"`

	TaxedFed bool `xml:"TaxedFed,omitempty" json:"TaxedFed,omitempty" yaml:"TaxedFed,omitempty"`

	TaxedState bool `xml:"TaxedState,omitempty" json:"TaxedState,omitempty" yaml:"TaxedState,omitempty"`

	UseTaxOverride bool `xml:"UseTaxOverride,omitempty" json:"UseTaxOverride,omitempty" yaml:"UseTaxOverride,omitempty"`

	TaxOverridePct float64 `xml:"TaxOverridePct,omitempty" json:"TaxOverridePct,omitempty" yaml:"TaxOverridePct,omitempty"`
}

type SetItemWarehouseRequest struct {
	*ApiRequest

	AllowedUserWarehouses *ArrayOfInt `xml:"AllowedUserWarehouses,omitempty" json:"AllowedUserWarehouses,omitempty" yaml:"AllowedUserWarehouses,omitempty"`

	AllowedWarehouseManagementTypes *ArrayOfInt `xml:"AllowedWarehouseManagementTypes,omitempty" json:"AllowedWarehouseManagementTypes,omitempty" yaml:"AllowedWarehouseManagementTypes,omitempty"`

	ItemCode string `xml:"ItemCode,omitempty" json:"ItemCode,omitempty" yaml:"ItemCode,omitempty"`

	WarehouseID int32 `xml:"WarehouseID,omitempty" json:"WarehouseID,omitempty" yaml:"WarehouseID,omitempty"`

	IsAvailable bool `xml:"IsAvailable,omitempty" json:"IsAvailable,omitempty" yaml:"IsAvailable,omitempty"`

	ItemManageTypeID int32 `xml:"ItemManageTypeID,omitempty" json:"ItemManageTypeID,omitempty" yaml:"ItemManageTypeID,omitempty"`
}

type SetItemPriceRequest struct {
	*ApiRequest

	ItemCode string `xml:"ItemCode,omitempty" json:"ItemCode,omitempty" yaml:"ItemCode,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty" yaml:"CurrencyCode,omitempty"`

	PriceType int32 `xml:"PriceType,omitempty" json:"PriceType,omitempty" yaml:"PriceType,omitempty"`

	Price float64 `xml:"Price,omitempty" json:"Price,omitempty" yaml:"Price,omitempty"`

	BusinessVolume float64 `xml:"BusinessVolume,omitempty" json:"BusinessVolume,omitempty" yaml:"BusinessVolume,omitempty"`

	CommissionableVolume float64 `xml:"CommissionableVolume,omitempty" json:"CommissionableVolume,omitempty" yaml:"CommissionableVolume,omitempty"`

	TaxablePrice float64 `xml:"TaxablePrice,omitempty" json:"TaxablePrice,omitempty" yaml:"TaxablePrice,omitempty"`

	ShippingPrice float64 `xml:"ShippingPrice,omitempty" json:"ShippingPrice,omitempty" yaml:"ShippingPrice,omitempty"`

	Other1Price float64 `xml:"Other1Price,omitempty" json:"Other1Price,omitempty" yaml:"Other1Price,omitempty"`

	Other2Price float64 `xml:"Other2Price,omitempty" json:"Other2Price,omitempty" yaml:"Other2Price,omitempty"`

	Other3Price float64 `xml:"Other3Price,omitempty" json:"Other3Price,omitempty" yaml:"Other3Price,omitempty"`

	Other4Price float64 `xml:"Other4Price,omitempty" json:"Other4Price,omitempty" yaml:"Other4Price,omitempty"`

	Other5Price float64 `xml:"Other5Price,omitempty" json:"Other5Price,omitempty" yaml:"Other5Price,omitempty"`

	Other6Price float64 `xml:"Other6Price,omitempty" json:"Other6Price,omitempty" yaml:"Other6Price,omitempty"`

	Other7Price float64 `xml:"Other7Price,omitempty" json:"Other7Price,omitempty" yaml:"Other7Price,omitempty"`

	Other8Price float64 `xml:"Other8Price,omitempty" json:"Other8Price,omitempty" yaml:"Other8Price,omitempty"`

	Other9Price float64 `xml:"Other9Price,omitempty" json:"Other9Price,omitempty" yaml:"Other9Price,omitempty"`

	Other10Price float64 `xml:"Other10Price,omitempty" json:"Other10Price,omitempty" yaml:"Other10Price,omitempty"`
}

type CreateItemRequest struct {
	*ApiRequest

	ItemCode string `xml:"ItemCode,omitempty" json:"ItemCode,omitempty" yaml:"ItemCode,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	Weight float64 `xml:"Weight,omitempty" json:"Weight,omitempty" yaml:"Weight,omitempty"`

	Notes string `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`

	AvailableInAllCountryRegions bool `xml:"AvailableInAllCountryRegions,omitempty" json:"AvailableInAllCountryRegions,omitempty" yaml:"AvailableInAllCountryRegions,omitempty"`

	TaxedInAllCountryRegions bool `xml:"TaxedInAllCountryRegions,omitempty" json:"TaxedInAllCountryRegions,omitempty" yaml:"TaxedInAllCountryRegions,omitempty"`

	AvailableInAllWarehouses bool `xml:"AvailableInAllWarehouses,omitempty" json:"AvailableInAllWarehouses,omitempty" yaml:"AvailableInAllWarehouses,omitempty"`

	IsVirtual bool `xml:"IsVirtual,omitempty" json:"IsVirtual,omitempty" yaml:"IsVirtual,omitempty"`

	ItemTypeID int32 `xml:"ItemTypeID,omitempty" json:"ItemTypeID,omitempty" yaml:"ItemTypeID,omitempty"`

	OtherCheck1 bool `xml:"OtherCheck1,omitempty" json:"OtherCheck1,omitempty" yaml:"OtherCheck1,omitempty"`

	OtherCheck2 bool `xml:"OtherCheck2,omitempty" json:"OtherCheck2,omitempty" yaml:"OtherCheck2,omitempty"`

	OtherCheck3 bool `xml:"OtherCheck3,omitempty" json:"OtherCheck3,omitempty" yaml:"OtherCheck3,omitempty"`

	OtherCheck4 bool `xml:"OtherCheck4,omitempty" json:"OtherCheck4,omitempty" yaml:"OtherCheck4,omitempty"`

	OtherCheck5 bool `xml:"OtherCheck5,omitempty" json:"OtherCheck5,omitempty" yaml:"OtherCheck5,omitempty"`

	Field1 string `xml:"Field1,omitempty" json:"Field1,omitempty" yaml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty" json:"Field2,omitempty" yaml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty" json:"Field3,omitempty" yaml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty" json:"Field4,omitempty" yaml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty" json:"Field5,omitempty" yaml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty" json:"Field6,omitempty" yaml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty" json:"Field7,omitempty" yaml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty" json:"Field8,omitempty" yaml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty" json:"Field9,omitempty" yaml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty" json:"Field10,omitempty" yaml:"Field10,omitempty"`

	HideFromSearch bool `xml:"HideFromSearch,omitempty" json:"HideFromSearch,omitempty" yaml:"HideFromSearch,omitempty"`
}

type CreateCustomerWallItemRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	Text string `xml:"Text,omitempty" json:"Text,omitempty" yaml:"Text,omitempty"`

	Field1 string `xml:"Field1,omitempty" json:"Field1,omitempty" yaml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty" json:"Field2,omitempty" yaml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty" json:"Field3,omitempty" yaml:"Field3,omitempty"`
}

type DeleteCustomerWallItemRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	WallItemID int32 `xml:"WallItemID,omitempty" json:"WallItemID,omitempty" yaml:"WallItemID,omitempty"`

	OlderThanEntryDate time.Time `xml:"OlderThanEntryDate,omitempty" json:"OlderThanEntryDate,omitempty" yaml:"OlderThanEntryDate,omitempty"`
}

type GetCustomerWallRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	WallItemID int32 `xml:"WallItemID,omitempty" json:"WallItemID,omitempty" yaml:"WallItemID,omitempty"`

	OlderThanEntryDate time.Time `xml:"OlderThanEntryDate,omitempty" json:"OlderThanEntryDate,omitempty" yaml:"OlderThanEntryDate,omitempty"`

	BatchSize int32 `xml:"BatchSize,omitempty" json:"BatchSize,omitempty" yaml:"BatchSize,omitempty"`
}

type CreateCustomerLeadRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty" json:"MiddleName,omitempty" yaml:"MiddleName,omitempty"`

	Company string `xml:"Company,omitempty" json:"Company,omitempty" yaml:"Company,omitempty"`

	Phone string `xml:"Phone,omitempty" json:"Phone,omitempty" yaml:"Phone,omitempty"`

	Phone2 string `xml:"Phone2,omitempty" json:"Phone2,omitempty" yaml:"Phone2,omitempty"`

	MobilePhone string `xml:"MobilePhone,omitempty" json:"MobilePhone,omitempty" yaml:"MobilePhone,omitempty"`

	Fax string `xml:"Fax,omitempty" json:"Fax,omitempty" yaml:"Fax,omitempty"`

	Email string `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`

	Address1 string `xml:"Address1,omitempty" json:"Address1,omitempty" yaml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty" json:"Address2,omitempty" yaml:"Address2,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty" json:"Zip,omitempty" yaml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`

	BirthDate time.Time `xml:"BirthDate,omitempty" json:"BirthDate,omitempty" yaml:"BirthDate,omitempty"`

	Notes string `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`
}

type UpdateCustomerLeadRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	CustomerLeadID int32 `xml:"CustomerLeadID,omitempty" json:"CustomerLeadID,omitempty" yaml:"CustomerLeadID,omitempty"`

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty" json:"MiddleName,omitempty" yaml:"MiddleName,omitempty"`

	Company string `xml:"Company,omitempty" json:"Company,omitempty" yaml:"Company,omitempty"`

	Phone string `xml:"Phone,omitempty" json:"Phone,omitempty" yaml:"Phone,omitempty"`

	Phone2 string `xml:"Phone2,omitempty" json:"Phone2,omitempty" yaml:"Phone2,omitempty"`

	MobilePhone string `xml:"MobilePhone,omitempty" json:"MobilePhone,omitempty" yaml:"MobilePhone,omitempty"`

	Fax string `xml:"Fax,omitempty" json:"Fax,omitempty" yaml:"Fax,omitempty"`

	Email string `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`

	Address1 string `xml:"Address1,omitempty" json:"Address1,omitempty" yaml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty" json:"Address2,omitempty" yaml:"Address2,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty" json:"Zip,omitempty" yaml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`

	BirthDate time.Time `xml:"BirthDate,omitempty" json:"BirthDate,omitempty" yaml:"BirthDate,omitempty"`

	Notes string `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`
}

type GetCustomerLeadsRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	CustomerLeadID int32 `xml:"CustomerLeadID,omitempty" json:"CustomerLeadID,omitempty" yaml:"CustomerLeadID,omitempty"`

	GreaterThanCustomerLeadID int32 `xml:"GreaterThanCustomerLeadID,omitempty" json:"GreaterThanCustomerLeadID,omitempty" yaml:"GreaterThanCustomerLeadID,omitempty"`

	BatchSize int32 `xml:"BatchSize,omitempty" json:"BatchSize,omitempty" yaml:"BatchSize,omitempty"`
}

type SetCustomerLeadSocialNetworksRequest struct {
	*ApiRequest

	CustomerLeadID int32 `xml:"CustomerLeadID,omitempty" json:"CustomerLeadID,omitempty" yaml:"CustomerLeadID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	CustomerLeadSocialNetworks *ArrayOfCustomerLeadSocialNetworkRequest `xml:"CustomerLeadSocialNetworks,omitempty" json:"CustomerLeadSocialNetworks,omitempty" yaml:"CustomerLeadSocialNetworks,omitempty"`
}

type ArrayOfCustomerLeadSocialNetworkRequest struct {
	CustomerLeadSocialNetworkRequest []*CustomerLeadSocialNetworkRequest `xml:"CustomerLeadSocialNetworkRequest,omitempty" json:"CustomerLeadSocialNetworkRequest,omitempty" yaml:"CustomerLeadSocialNetworkRequest,omitempty"`
}

type CustomerLeadSocialNetworkRequest struct {
	SocialNetworkID int32 `xml:"SocialNetworkID,omitempty" json:"SocialNetworkID,omitempty" yaml:"SocialNetworkID,omitempty"`

	Url string `xml:"Url,omitempty" json:"Url,omitempty" yaml:"Url,omitempty"`
}

type GetCustomerLeadSocialNetworksRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	CustomerLeadID int32 `xml:"CustomerLeadID,omitempty" json:"CustomerLeadID,omitempty" yaml:"CustomerLeadID,omitempty"`
}

type GetCustomerSocialNetworksRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`
}

type SetCustomerSocialNetworksRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	CustomerSocialNetworks *ArrayOfCustomerSocialNetworkRequest `xml:"CustomerSocialNetworks,omitempty" json:"CustomerSocialNetworks,omitempty" yaml:"CustomerSocialNetworks,omitempty"`
}

type ArrayOfCustomerSocialNetworkRequest struct {
	CustomerSocialNetworkRequest []*CustomerSocialNetworkRequest `xml:"CustomerSocialNetworkRequest,omitempty" json:"CustomerSocialNetworkRequest,omitempty" yaml:"CustomerSocialNetworkRequest,omitempty"`
}

type CustomerSocialNetworkRequest struct {
	SocialNetworkID int32 `xml:"SocialNetworkID,omitempty" json:"SocialNetworkID,omitempty" yaml:"SocialNetworkID,omitempty"`

	Url string `xml:"Url,omitempty" json:"Url,omitempty" yaml:"Url,omitempty"`
}

type SetCustomerSiteRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetCusotmerSiteRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	WebAlias string `xml:"WebAlias,omitempty" json:"WebAlias,omitempty" yaml:"WebAlias,omitempty"`

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`

	Company string `xml:"Company,omitempty" json:"Company,omitempty" yaml:"Company,omitempty"`

	Address1 string `xml:"Address1,omitempty" json:"Address1,omitempty" yaml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty" json:"Address2,omitempty" yaml:"Address2,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty" json:"Zip,omitempty" yaml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`

	Email string `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`

	Phone string `xml:"Phone,omitempty" json:"Phone,omitempty" yaml:"Phone,omitempty"`

	Phone2 string `xml:"Phone2,omitempty" json:"Phone2,omitempty" yaml:"Phone2,omitempty"`

	Fax string `xml:"Fax,omitempty" json:"Fax,omitempty" yaml:"Fax,omitempty"`

	Notes1 string `xml:"Notes1,omitempty" json:"Notes1,omitempty" yaml:"Notes1,omitempty"`

	Notes2 string `xml:"Notes2,omitempty" json:"Notes2,omitempty" yaml:"Notes2,omitempty"`

	Notes3 string `xml:"Notes3,omitempty" json:"Notes3,omitempty" yaml:"Notes3,omitempty"`

	Notes4 string `xml:"Notes4,omitempty" json:"Notes4,omitempty" yaml:"Notes4,omitempty"`

	Url1 string `xml:"Url1,omitempty" json:"Url1,omitempty" yaml:"Url1,omitempty"`

	Url2 string `xml:"Url2,omitempty" json:"Url2,omitempty" yaml:"Url2,omitempty"`

	Url3 string `xml:"Url3,omitempty" json:"Url3,omitempty" yaml:"Url3,omitempty"`

	Url4 string `xml:"Url4,omitempty" json:"Url4,omitempty" yaml:"Url4,omitempty"`

	Url5 string `xml:"Url5,omitempty" json:"Url5,omitempty" yaml:"Url5,omitempty"`

	Url6 string `xml:"Url6,omitempty" json:"Url6,omitempty" yaml:"Url6,omitempty"`

	Url7 string `xml:"Url7,omitempty" json:"Url7,omitempty" yaml:"Url7,omitempty"`

	Url8 string `xml:"Url8,omitempty" json:"Url8,omitempty" yaml:"Url8,omitempty"`

	Url9 string `xml:"Url9,omitempty" json:"Url9,omitempty" yaml:"Url9,omitempty"`

	Url10 string `xml:"Url10,omitempty" json:"Url10,omitempty" yaml:"Url10,omitempty"`

	Url1Description string `xml:"Url1Description,omitempty" json:"Url1Description,omitempty" yaml:"Url1Description,omitempty"`

	Url2Description string `xml:"Url2Description,omitempty" json:"Url2Description,omitempty" yaml:"Url2Description,omitempty"`

	Url3Description string `xml:"Url3Description,omitempty" json:"Url3Description,omitempty" yaml:"Url3Description,omitempty"`

	Url4Description string `xml:"Url4Description,omitempty" json:"Url4Description,omitempty" yaml:"Url4Description,omitempty"`

	Url5Description string `xml:"Url5Description,omitempty" json:"Url5Description,omitempty" yaml:"Url5Description,omitempty"`

	Url6Description string `xml:"Url6Description,omitempty" json:"Url6Description,omitempty" yaml:"Url6Description,omitempty"`

	Url7Description string `xml:"Url7Description,omitempty" json:"Url7Description,omitempty" yaml:"Url7Description,omitempty"`

	Url8Description string `xml:"Url8Description,omitempty" json:"Url8Description,omitempty" yaml:"Url8Description,omitempty"`

	Url9Description string `xml:"Url9Description,omitempty" json:"Url9Description,omitempty" yaml:"Url9Description,omitempty"`

	Url10Description string `xml:"Url10Description,omitempty" json:"Url10Description,omitempty" yaml:"Url10Description,omitempty"`
}

type OptInSmsRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	PhoneNumber string `xml:"PhoneNumber,omitempty" json:"PhoneNumber,omitempty" yaml:"PhoneNumber,omitempty"`
}

type OptInEmailRequest struct {
	*ApiRequest

	Email string `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`
}

type UpdateOrderRequest struct {
	*ApiRequest

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	OrderDate time.Time `xml:"OrderDate,omitempty" json:"OrderDate,omitempty" yaml:"OrderDate,omitempty"`

	DeclineCount int32 `xml:"DeclineCount,omitempty" json:"DeclineCount,omitempty" yaml:"DeclineCount,omitempty"`

	OrderTy int32 `xml:"OrderTy,omitempty" json:"OrderTy,omitempty" yaml:"OrderTy,omitempty"`

	OrderStatus int32 `xml:"OrderStatus,omitempty" json:"OrderStatus,omitempty" yaml:"OrderStatus,omitempty"`

	PriceTy int32 `xml:"PriceTy,omitempty" json:"PriceTy,omitempty" yaml:"PriceTy,omitempty"`

	Total float64 `xml:"Total,omitempty" json:"Total,omitempty" yaml:"Total,omitempty"`

	SubTotal float64 `xml:"SubTotal,omitempty" json:"SubTotal,omitempty" yaml:"SubTotal,omitempty"`

	Shipping float64 `xml:"Shipping,omitempty" json:"Shipping,omitempty" yaml:"Shipping,omitempty"`

	OrderTax float64 `xml:"OrderTax,omitempty" json:"OrderTax,omitempty" yaml:"OrderTax,omitempty"`

	ShippingTax float64 `xml:"ShippingTax,omitempty" json:"ShippingTax,omitempty" yaml:"ShippingTax,omitempty"`

	FedShippingTax float64 `xml:"FedShippingTax,omitempty" json:"FedShippingTax,omitempty" yaml:"FedShippingTax,omitempty"`

	StateShippingTax float64 `xml:"StateShippingTax,omitempty" json:"StateShippingTax,omitempty" yaml:"StateShippingTax,omitempty"`

	CityShippingTax float64 `xml:"CityShippingTax,omitempty" json:"CityShippingTax,omitempty" yaml:"CityShippingTax,omitempty"`

	CityLocalShippingTax float64 `xml:"CityLocalShippingTax,omitempty" json:"CityLocalShippingTax,omitempty" yaml:"CityLocalShippingTax,omitempty"`

	CountyShippingTax float64 `xml:"CountyShippingTax,omitempty" json:"CountyShippingTax,omitempty" yaml:"CountyShippingTax,omitempty"`

	CountyLocalShippingTax float64 `xml:"CountyLocalShippingTax,omitempty" json:"CountyLocalShippingTax,omitempty" yaml:"CountyLocalShippingTax,omitempty"`

	ManualTaxRate float64 `xml:"ManualTaxRate,omitempty" json:"ManualTaxRate,omitempty" yaml:"ManualTaxRate,omitempty"`

	TotalTax float64 `xml:"TotalTax,omitempty" json:"TotalTax,omitempty" yaml:"TotalTax,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty" yaml:"CurrencyCode,omitempty"`

	PaymentMethod int32 `xml:"PaymentMethod,omitempty" json:"PaymentMethod,omitempty" yaml:"PaymentMethod,omitempty"`

	WarehouseID int32 `xml:"WarehouseID,omitempty" json:"WarehouseID,omitempty" yaml:"WarehouseID,omitempty"`

	BatchID int32 `xml:"BatchID,omitempty" json:"BatchID,omitempty" yaml:"BatchID,omitempty"`

	PreviousBalance float64 `xml:"PreviousBalance,omitempty" json:"PreviousBalance,omitempty" yaml:"PreviousBalance,omitempty"`

	OverrideShipping bool `xml:"OverrideShipping,omitempty" json:"OverrideShipping,omitempty" yaml:"OverrideShipping,omitempty"`

	OverrideTax bool `xml:"OverrideTax,omitempty" json:"OverrideTax,omitempty" yaml:"OverrideTax,omitempty"`

	BusinessVolume float64 `xml:"BusinessVolume,omitempty" json:"BusinessVolume,omitempty" yaml:"BusinessVolume,omitempty"`

	CommissionableVolume float64 `xml:"CommissionableVolume,omitempty" json:"CommissionableVolume,omitempty" yaml:"CommissionableVolume,omitempty"`

	Other1 float64 `xml:"Other1,omitempty" json:"Other1,omitempty" yaml:"Other1,omitempty"`

	Other2 float64 `xml:"Other2,omitempty" json:"Other2,omitempty" yaml:"Other2,omitempty"`

	Other3 float64 `xml:"Other3,omitempty" json:"Other3,omitempty" yaml:"Other3,omitempty"`

	Other4 float64 `xml:"Other4,omitempty" json:"Other4,omitempty" yaml:"Other4,omitempty"`

	Other5 float64 `xml:"Other5,omitempty" json:"Other5,omitempty" yaml:"Other5,omitempty"`

	Discount float64 `xml:"Discount,omitempty" json:"Discount,omitempty" yaml:"Discount,omitempty"`

	DiscountPercent float64 `xml:"DiscountPercent,omitempty" json:"DiscountPercent,omitempty" yaml:"DiscountPercent,omitempty"`

	Weight float64 `xml:"Weight,omitempty" json:"Weight,omitempty" yaml:"Weight,omitempty"`

	Sourcety int32 `xml:"Sourcety,omitempty" json:"Sourcety,omitempty" yaml:"Sourcety,omitempty"`

	Notes string `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`

	Usr string `xml:"Usr,omitempty" json:"Usr,omitempty" yaml:"Usr,omitempty"`

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`

	Company string `xml:"Company,omitempty" json:"Company,omitempty" yaml:"Company,omitempty"`

	Address1 string `xml:"Address1,omitempty" json:"Address1,omitempty" yaml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty" json:"Address2,omitempty" yaml:"Address2,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty" json:"Zip,omitempty" yaml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`

	Email string `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`

	Phone string `xml:"Phone,omitempty" json:"Phone,omitempty" yaml:"Phone,omitempty"`

	SuppressPackSlipPrice bool `xml:"SuppressPackSlipPrice,omitempty" json:"SuppressPackSlipPrice,omitempty" yaml:"SuppressPackSlipPrice,omitempty"`

	ShipMethodID int32 `xml:"ShipMethodID,omitempty" json:"ShipMethodID,omitempty" yaml:"ShipMethodID,omitempty"`

	AutoOrderID int32 `xml:"AutoOrderID,omitempty" json:"AutoOrderID,omitempty" yaml:"AutoOrderID,omitempty"`

	CreatedBy string `xml:"CreatedBy,omitempty" json:"CreatedBy,omitempty" yaml:"CreatedBy,omitempty"`

	ReturnOrderID int32 `xml:"ReturnOrderID,omitempty" json:"ReturnOrderID,omitempty" yaml:"ReturnOrderID,omitempty"`

	OrderRankID int32 `xml:"OrderRankID,omitempty" json:"OrderRankID,omitempty" yaml:"OrderRankID,omitempty"`

	OrderPayRankID int32 `xml:"OrderPayRankID,omitempty" json:"OrderPayRankID,omitempty" yaml:"OrderPayRankID,omitempty"`

	AddressIsVerified bool `xml:"AddressIsVerified,omitempty" json:"AddressIsVerified,omitempty" yaml:"AddressIsVerified,omitempty"`

	County string `xml:"County,omitempty" json:"County,omitempty" yaml:"County,omitempty"`

	IsRMA bool `xml:"IsRMA,omitempty" json:"IsRMA,omitempty" yaml:"IsRMA,omitempty"`

	BackOrderFromID int32 `xml:"BackOrderFromID,omitempty" json:"BackOrderFromID,omitempty" yaml:"BackOrderFromID,omitempty"`

	CreditsEarned float64 `xml:"CreditsEarned,omitempty" json:"CreditsEarned,omitempty" yaml:"CreditsEarned,omitempty"`

	TotalFedTax float64 `xml:"TotalFedTax,omitempty" json:"TotalFedTax,omitempty" yaml:"TotalFedTax,omitempty"`

	TotalStateTax float64 `xml:"TotalStateTax,omitempty" json:"TotalStateTax,omitempty" yaml:"TotalStateTax,omitempty"`

	ManualShippingTax float64 `xml:"ManualShippingTax,omitempty" json:"ManualShippingTax,omitempty" yaml:"ManualShippingTax,omitempty"`

	ReplacementOrderID int32 `xml:"ReplacementOrderID,omitempty" json:"ReplacementOrderID,omitempty" yaml:"ReplacementOrderID,omitempty"`

	LockedDate time.Time `xml:"LockedDate,omitempty" json:"LockedDate,omitempty" yaml:"LockedDate,omitempty"`

	CommissionedDate time.Time `xml:"CommissionedDate,omitempty" json:"CommissionedDate,omitempty" yaml:"CommissionedDate,omitempty"`

	Flag1 bool `xml:"Flag1,omitempty" json:"Flag1,omitempty" yaml:"Flag1,omitempty"`

	Flag2 bool `xml:"Flag2,omitempty" json:"Flag2,omitempty" yaml:"Flag2,omitempty"`

	Flag3 bool `xml:"Flag3,omitempty" json:"Flag3,omitempty" yaml:"Flag3,omitempty"`

	Other6 float64 `xml:"Other6,omitempty" json:"Other6,omitempty" yaml:"Other6,omitempty"`

	Other7 float64 `xml:"Other7,omitempty" json:"Other7,omitempty" yaml:"Other7,omitempty"`

	Other8 float64 `xml:"Other8,omitempty" json:"Other8,omitempty" yaml:"Other8,omitempty"`

	Other9 float64 `xml:"Other9,omitempty" json:"Other9,omitempty" yaml:"Other9,omitempty"`

	Other10 float64 `xml:"Other10,omitempty" json:"Other10,omitempty" yaml:"Other10,omitempty"`

	OriginalWarehouseID int32 `xml:"OriginalWarehouseID,omitempty" json:"OriginalWarehouseID,omitempty" yaml:"OriginalWarehouseID,omitempty"`

	PickupName string `xml:"PickupName,omitempty" json:"PickupName,omitempty" yaml:"PickupName,omitempty"`

	TransferToID int32 `xml:"TransferToID,omitempty" json:"TransferToID,omitempty" yaml:"TransferToID,omitempty"`

	IsCommissionable bool `xml:"IsCommissionable,omitempty" json:"IsCommissionable,omitempty" yaml:"IsCommissionable,omitempty"`

	FulfilledBy string `xml:"FulfilledBy,omitempty" json:"FulfilledBy,omitempty" yaml:"FulfilledBy,omitempty"`

	CreditApplied float64 `xml:"CreditApplied,omitempty" json:"CreditApplied,omitempty" yaml:"CreditApplied,omitempty"`

	ShippedDate time.Time `xml:"ShippedDate,omitempty" json:"ShippedDate,omitempty" yaml:"ShippedDate,omitempty"`

	TaxLockDate time.Time `xml:"TaxLockDate,omitempty" json:"TaxLockDate,omitempty" yaml:"TaxLockDate,omitempty"`

	TotalTaxable float64 `xml:"TotalTaxable,omitempty" json:"TotalTaxable,omitempty" yaml:"TotalTaxable,omitempty"`

	ReturnCategoryID int32 `xml:"ReturnCategoryID,omitempty" json:"ReturnCategoryID,omitempty" yaml:"ReturnCategoryID,omitempty"`

	ReplacementCategoryID int32 `xml:"ReplacementCategoryID,omitempty" json:"ReplacementCategoryID,omitempty" yaml:"ReplacementCategoryID,omitempty"`

	CalculatedShipping float64 `xml:"CalculatedShipping,omitempty" json:"CalculatedShipping,omitempty" yaml:"CalculatedShipping,omitempty"`

	HandlingFee float64 `xml:"HandlingFee,omitempty" json:"HandlingFee,omitempty" yaml:"HandlingFee,omitempty"`

	OrderProcessTy int32 `xml:"OrderProcessTy,omitempty" json:"OrderProcessTy,omitempty" yaml:"OrderProcessTy,omitempty"`

	ActualCarrier int32 `xml:"ActualCarrier,omitempty" json:"ActualCarrier,omitempty" yaml:"ActualCarrier,omitempty"`

	ParentOrderID int32 `xml:"ParentOrderID,omitempty" json:"ParentOrderID,omitempty" yaml:"ParentOrderID,omitempty"`

	CustomerTy int32 `xml:"CustomerTy,omitempty" json:"CustomerTy,omitempty" yaml:"CustomerTy,omitempty"`

	Reference string `xml:"Reference,omitempty" json:"Reference,omitempty" yaml:"Reference,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty" json:"MiddleName,omitempty" yaml:"MiddleName,omitempty"`

	NameSuffix string `xml:"NameSuffix,omitempty" json:"NameSuffix,omitempty" yaml:"NameSuffix,omitempty"`

	Address3 string `xml:"Address3,omitempty" json:"Address3,omitempty" yaml:"Address3,omitempty"`

	PartyID int32 `xml:"PartyID,omitempty" json:"PartyID,omitempty" yaml:"PartyID,omitempty"`

	TrackingNumber1 string `xml:"TrackingNumber1,omitempty" json:"TrackingNumber1,omitempty" yaml:"TrackingNumber1,omitempty"`

	TrackingNumber2 string `xml:"TrackingNumber2,omitempty" json:"TrackingNumber2,omitempty" yaml:"TrackingNumber2,omitempty"`

	TrackingNumber3 string `xml:"TrackingNumber3,omitempty" json:"TrackingNumber3,omitempty" yaml:"TrackingNumber3,omitempty"`

	TrackingNumber4 string `xml:"TrackingNumber4,omitempty" json:"TrackingNumber4,omitempty" yaml:"TrackingNumber4,omitempty"`

	TrackingNumber5 string `xml:"TrackingNumber5,omitempty" json:"TrackingNumber5,omitempty" yaml:"TrackingNumber5,omitempty"`

	WebCarrierID *OrderShipCarrier `xml:"WebCarrierID,omitempty" json:"WebCarrierID,omitempty" yaml:"WebCarrierID,omitempty"`

	WebCarrierID2 *OrderShipCarrier `xml:"WebCarrierID2,omitempty" json:"WebCarrierID2,omitempty" yaml:"WebCarrierID2,omitempty"`

	WebCarrierID3 *OrderShipCarrier `xml:"WebCarrierID3,omitempty" json:"WebCarrierID3,omitempty" yaml:"WebCarrierID3,omitempty"`

	WebCarrierID4 *OrderShipCarrier `xml:"WebCarrierID4,omitempty" json:"WebCarrierID4,omitempty" yaml:"WebCarrierID4,omitempty"`

	WebCarrierID5 *OrderShipCarrier `xml:"WebCarrierID5,omitempty" json:"WebCarrierID5,omitempty" yaml:"WebCarrierID5,omitempty"`

	Other11 string `xml:"Other11,omitempty" json:"Other11,omitempty" yaml:"Other11,omitempty"`

	Other12 string `xml:"Other12,omitempty" json:"Other12,omitempty" yaml:"Other12,omitempty"`

	Other13 string `xml:"Other13,omitempty" json:"Other13,omitempty" yaml:"Other13,omitempty"`

	Other14 string `xml:"Other14,omitempty" json:"Other14,omitempty" yaml:"Other14,omitempty"`

	Other15 string `xml:"Other15,omitempty" json:"Other15,omitempty" yaml:"Other15,omitempty"`

	Other16 string `xml:"Other16,omitempty" json:"Other16,omitempty" yaml:"Other16,omitempty"`

	Other17 string `xml:"Other17,omitempty" json:"Other17,omitempty" yaml:"Other17,omitempty"`

	Other18 string `xml:"Other18,omitempty" json:"Other18,omitempty" yaml:"Other18,omitempty"`

	Other19 string `xml:"Other19,omitempty" json:"Other19,omitempty" yaml:"Other19,omitempty"`

	Other20 string `xml:"Other20,omitempty" json:"Other20,omitempty" yaml:"Other20,omitempty"`
}

type UpdateCustomerRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`

	Company string `xml:"Company,omitempty" json:"Company,omitempty" yaml:"Company,omitempty"`

	CustomerType int32 `xml:"CustomerType,omitempty" json:"CustomerType,omitempty" yaml:"CustomerType,omitempty"`

	CustomerStatus int32 `xml:"CustomerStatus,omitempty" json:"CustomerStatus,omitempty" yaml:"CustomerStatus,omitempty"`

	Email string `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`

	Phone string `xml:"Phone,omitempty" json:"Phone,omitempty" yaml:"Phone,omitempty"`

	Phone2 string `xml:"Phone2,omitempty" json:"Phone2,omitempty" yaml:"Phone2,omitempty"`

	MobilePhone string `xml:"MobilePhone,omitempty" json:"MobilePhone,omitempty" yaml:"MobilePhone,omitempty"`

	Fax string `xml:"Fax,omitempty" json:"Fax,omitempty" yaml:"Fax,omitempty"`

	MainAddress1 string `xml:"MainAddress1,omitempty" json:"MainAddress1,omitempty" yaml:"MainAddress1,omitempty"`

	MainAddress2 string `xml:"MainAddress2,omitempty" json:"MainAddress2,omitempty" yaml:"MainAddress2,omitempty"`

	MainAddress3 string `xml:"MainAddress3,omitempty" json:"MainAddress3,omitempty" yaml:"MainAddress3,omitempty"`

	MainCity string `xml:"MainCity,omitempty" json:"MainCity,omitempty" yaml:"MainCity,omitempty"`

	MainState string `xml:"MainState,omitempty" json:"MainState,omitempty" yaml:"MainState,omitempty"`

	MainZip string `xml:"MainZip,omitempty" json:"MainZip,omitempty" yaml:"MainZip,omitempty"`

	MainCountry string `xml:"MainCountry,omitempty" json:"MainCountry,omitempty" yaml:"MainCountry,omitempty"`

	MainCounty string `xml:"MainCounty,omitempty" json:"MainCounty,omitempty" yaml:"MainCounty,omitempty"`

	MailAddress1 string `xml:"MailAddress1,omitempty" json:"MailAddress1,omitempty" yaml:"MailAddress1,omitempty"`

	MailAddress2 string `xml:"MailAddress2,omitempty" json:"MailAddress2,omitempty" yaml:"MailAddress2,omitempty"`

	MailAddress3 string `xml:"MailAddress3,omitempty" json:"MailAddress3,omitempty" yaml:"MailAddress3,omitempty"`

	MailCity string `xml:"MailCity,omitempty" json:"MailCity,omitempty" yaml:"MailCity,omitempty"`

	MailState string `xml:"MailState,omitempty" json:"MailState,omitempty" yaml:"MailState,omitempty"`

	MailZip string `xml:"MailZip,omitempty" json:"MailZip,omitempty" yaml:"MailZip,omitempty"`

	MailCountry string `xml:"MailCountry,omitempty" json:"MailCountry,omitempty" yaml:"MailCountry,omitempty"`

	MailCounty string `xml:"MailCounty,omitempty" json:"MailCounty,omitempty" yaml:"MailCounty,omitempty"`

	OtherAddress1 string `xml:"OtherAddress1,omitempty" json:"OtherAddress1,omitempty" yaml:"OtherAddress1,omitempty"`

	OtherAddress2 string `xml:"OtherAddress2,omitempty" json:"OtherAddress2,omitempty" yaml:"OtherAddress2,omitempty"`

	OtherAddress3 string `xml:"OtherAddress3,omitempty" json:"OtherAddress3,omitempty" yaml:"OtherAddress3,omitempty"`

	OtherCity string `xml:"OtherCity,omitempty" json:"OtherCity,omitempty" yaml:"OtherCity,omitempty"`

	OtherState string `xml:"OtherState,omitempty" json:"OtherState,omitempty" yaml:"OtherState,omitempty"`

	OtherZip string `xml:"OtherZip,omitempty" json:"OtherZip,omitempty" yaml:"OtherZip,omitempty"`

	OtherCountry string `xml:"OtherCountry,omitempty" json:"OtherCountry,omitempty" yaml:"OtherCountry,omitempty"`

	OtherCounty string `xml:"OtherCounty,omitempty" json:"OtherCounty,omitempty" yaml:"OtherCounty,omitempty"`

	CanLogin bool `xml:"CanLogin,omitempty" json:"CanLogin,omitempty" yaml:"CanLogin,omitempty"`

	LoginName string `xml:"LoginName,omitempty" json:"LoginName,omitempty" yaml:"LoginName,omitempty"`

	LoginPassword string `xml:"LoginPassword,omitempty" json:"LoginPassword,omitempty" yaml:"LoginPassword,omitempty"`

	TaxID string `xml:"TaxID,omitempty" json:"TaxID,omitempty" yaml:"TaxID,omitempty"`

	SalesTaxID string `xml:"SalesTaxID,omitempty" json:"SalesTaxID,omitempty" yaml:"SalesTaxID,omitempty"`

	SalesTaxExemptExpireDate time.Time `xml:"SalesTaxExemptExpireDate,omitempty" json:"SalesTaxExemptExpireDate,omitempty" yaml:"SalesTaxExemptExpireDate,omitempty"`

	IsSalesTaxExempt bool `xml:"IsSalesTaxExempt,omitempty" json:"IsSalesTaxExempt,omitempty" yaml:"IsSalesTaxExempt,omitempty"`

	BirthDate time.Time `xml:"BirthDate,omitempty" json:"BirthDate,omitempty" yaml:"BirthDate,omitempty"`

	Field1 string `xml:"Field1,omitempty" json:"Field1,omitempty" yaml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty" json:"Field2,omitempty" yaml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty" json:"Field3,omitempty" yaml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty" json:"Field4,omitempty" yaml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty" json:"Field5,omitempty" yaml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty" json:"Field6,omitempty" yaml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty" json:"Field7,omitempty" yaml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty" json:"Field8,omitempty" yaml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty" json:"Field9,omitempty" yaml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty" json:"Field10,omitempty" yaml:"Field10,omitempty"`

	Field11 string `xml:"Field11,omitempty" json:"Field11,omitempty" yaml:"Field11,omitempty"`

	Field12 string `xml:"Field12,omitempty" json:"Field12,omitempty" yaml:"Field12,omitempty"`

	Field13 string `xml:"Field13,omitempty" json:"Field13,omitempty" yaml:"Field13,omitempty"`

	Field14 string `xml:"Field14,omitempty" json:"Field14,omitempty" yaml:"Field14,omitempty"`

	Field15 string `xml:"Field15,omitempty" json:"Field15,omitempty" yaml:"Field15,omitempty"`

	SubscribeToBroadcasts bool `xml:"SubscribeToBroadcasts,omitempty" json:"SubscribeToBroadcasts,omitempty" yaml:"SubscribeToBroadcasts,omitempty"`

	SubscribeFromIPAddress string `xml:"SubscribeFromIPAddress,omitempty" json:"SubscribeFromIPAddress,omitempty" yaml:"SubscribeFromIPAddress,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty" yaml:"CurrencyCode,omitempty"`

	PayableToName string `xml:"PayableToName,omitempty" json:"PayableToName,omitempty" yaml:"PayableToName,omitempty"`

	PayableType *PayableType `xml:"PayableType,omitempty" json:"PayableType,omitempty" yaml:"PayableType,omitempty"`

	DefaultWarehouseID int32 `xml:"DefaultWarehouseID,omitempty" json:"DefaultWarehouseID,omitempty" yaml:"DefaultWarehouseID,omitempty"`

	CheckThreshold float64 `xml:"CheckThreshold,omitempty" json:"CheckThreshold,omitempty" yaml:"CheckThreshold,omitempty"`

	CreatedDate time.Time `xml:"CreatedDate,omitempty" json:"CreatedDate,omitempty" yaml:"CreatedDate,omitempty"`

	TaxIDType *TaxIDType `xml:"TaxIDType,omitempty" json:"TaxIDType,omitempty" yaml:"TaxIDType,omitempty"`

	LanguageID int32 `xml:"LanguageID,omitempty" json:"LanguageID,omitempty" yaml:"LanguageID,omitempty"`

	Gender *Gender `xml:"Gender,omitempty" json:"Gender,omitempty" yaml:"Gender,omitempty"`

	VatRegistration string `xml:"VatRegistration,omitempty" json:"VatRegistration,omitempty" yaml:"VatRegistration,omitempty"`

	Date1 time.Time `xml:"Date1,omitempty" json:"Date1,omitempty" yaml:"Date1,omitempty"`

	Date2 time.Time `xml:"Date2,omitempty" json:"Date2,omitempty" yaml:"Date2,omitempty"`

	Date3 time.Time `xml:"Date3,omitempty" json:"Date3,omitempty" yaml:"Date3,omitempty"`

	Date4 time.Time `xml:"Date4,omitempty" json:"Date4,omitempty" yaml:"Date4,omitempty"`

	Date5 time.Time `xml:"Date5,omitempty" json:"Date5,omitempty" yaml:"Date5,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty" json:"MiddleName,omitempty" yaml:"MiddleName,omitempty"`

	NameSuffix string `xml:"NameSuffix,omitempty" json:"NameSuffix,omitempty" yaml:"NameSuffix,omitempty"`

	BinaryPlacementPreference int32 `xml:"BinaryPlacementPreference,omitempty" json:"BinaryPlacementPreference,omitempty" yaml:"BinaryPlacementPreference,omitempty"`

	UseBinaryHoldingTank bool `xml:"UseBinaryHoldingTank,omitempty" json:"UseBinaryHoldingTank,omitempty" yaml:"UseBinaryHoldingTank,omitempty"`

	MainAddressVerified bool `xml:"MainAddressVerified,omitempty" json:"MainAddressVerified,omitempty" yaml:"MainAddressVerified,omitempty"`

	MailAddressVerified bool `xml:"MailAddressVerified,omitempty" json:"MailAddressVerified,omitempty" yaml:"MailAddressVerified,omitempty"`

	OtherAddressVerified bool `xml:"OtherAddressVerified,omitempty" json:"OtherAddressVerified,omitempty" yaml:"OtherAddressVerified,omitempty"`
}

type BaseCalculateOrderRequest struct {
	*ApiRequest
}

type CalculateOrderRequest struct {
	*BaseCalculateOrderRequest

	CurrencyCode string `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty" yaml:"CurrencyCode,omitempty"`

	WarehouseID int32 `xml:"WarehouseID,omitempty" json:"WarehouseID,omitempty" yaml:"WarehouseID,omitempty"`

	ShipMethodID int32 `xml:"ShipMethodID,omitempty" json:"ShipMethodID,omitempty" yaml:"ShipMethodID,omitempty"`

	PriceType int32 `xml:"PriceType,omitempty" json:"PriceType,omitempty" yaml:"PriceType,omitempty"`

	Address1 string `xml:"Address1,omitempty" json:"Address1,omitempty" yaml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty" json:"Address2,omitempty" yaml:"Address2,omitempty"`

	Address3 string `xml:"Address3,omitempty" json:"Address3,omitempty" yaml:"Address3,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty" json:"Zip,omitempty" yaml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`

	County string `xml:"County,omitempty" json:"County,omitempty" yaml:"County,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	OrderType *OrderType `xml:"OrderType,omitempty" json:"OrderType,omitempty" yaml:"OrderType,omitempty"`

	TaxRateOverride float64 `xml:"TaxRateOverride,omitempty" json:"TaxRateOverride,omitempty" yaml:"TaxRateOverride,omitempty"`

	ShippingAmountOverride float64 `xml:"ShippingAmountOverride,omitempty" json:"ShippingAmountOverride,omitempty" yaml:"ShippingAmountOverride,omitempty"`

	ReturnOrderID int32 `xml:"ReturnOrderID,omitempty" json:"ReturnOrderID,omitempty" yaml:"ReturnOrderID,omitempty"`

	Other11 string `xml:"Other11,omitempty" json:"Other11,omitempty" yaml:"Other11,omitempty"`

	Other12 string `xml:"Other12,omitempty" json:"Other12,omitempty" yaml:"Other12,omitempty"`

	Other13 string `xml:"Other13,omitempty" json:"Other13,omitempty" yaml:"Other13,omitempty"`

	Other14 string `xml:"Other14,omitempty" json:"Other14,omitempty" yaml:"Other14,omitempty"`

	Other15 string `xml:"Other15,omitempty" json:"Other15,omitempty" yaml:"Other15,omitempty"`

	Other16 string `xml:"Other16,omitempty" json:"Other16,omitempty" yaml:"Other16,omitempty"`

	Other17 string `xml:"Other17,omitempty" json:"Other17,omitempty" yaml:"Other17,omitempty"`

	Other18 string `xml:"Other18,omitempty" json:"Other18,omitempty" yaml:"Other18,omitempty"`

	Other19 string `xml:"Other19,omitempty" json:"Other19,omitempty" yaml:"Other19,omitempty"`

	Other20 string `xml:"Other20,omitempty" json:"Other20,omitempty" yaml:"Other20,omitempty"`

	Details *ArrayOfOrderDetailRequest `xml:"Details,omitempty" json:"Details,omitempty" yaml:"Details,omitempty"`

	ReturnShipMethods bool `xml:"ReturnShipMethods,omitempty" json:"ReturnShipMethods,omitempty" yaml:"ReturnShipMethods,omitempty"`

	PartyID int32 `xml:"PartyID,omitempty" json:"PartyID,omitempty" yaml:"PartyID,omitempty"`
}

type ArrayOfOrderDetailRequest struct {
	OrderDetailRequest []*OrderDetailRequest `xml:"OrderDetailRequest,omitempty" json:"OrderDetailRequest,omitempty" yaml:"OrderDetailRequest,omitempty"`
}

type OrderDetailRequest struct {
	ItemCode string `xml:"ItemCode,omitempty" json:"ItemCode,omitempty" yaml:"ItemCode,omitempty"`

	Quantity float64 `xml:"Quantity,omitempty" json:"Quantity,omitempty" yaml:"Quantity,omitempty"`

	ParentItemCode string `xml:"ParentItemCode,omitempty" json:"ParentItemCode,omitempty" yaml:"ParentItemCode,omitempty"`

	PriceEachOverride float64 `xml:"PriceEachOverride,omitempty" json:"PriceEachOverride,omitempty" yaml:"PriceEachOverride,omitempty"`

	TaxableEachOverride float64 `xml:"TaxableEachOverride,omitempty" json:"TaxableEachOverride,omitempty" yaml:"TaxableEachOverride,omitempty"`

	ShippingPriceEachOverride float64 `xml:"ShippingPriceEachOverride,omitempty" json:"ShippingPriceEachOverride,omitempty" yaml:"ShippingPriceEachOverride,omitempty"`

	BusinessVolumeEachOverride float64 `xml:"BusinessVolumeEachOverride,omitempty" json:"BusinessVolumeEachOverride,omitempty" yaml:"BusinessVolumeEachOverride,omitempty"`

	CommissionableVolumeEachOverride float64 `xml:"CommissionableVolumeEachOverride,omitempty" json:"CommissionableVolumeEachOverride,omitempty" yaml:"CommissionableVolumeEachOverride,omitempty"`

	Other1EachOverride float64 `xml:"Other1EachOverride,omitempty" json:"Other1EachOverride,omitempty" yaml:"Other1EachOverride,omitempty"`

	Other2EachOverride float64 `xml:"Other2EachOverride,omitempty" json:"Other2EachOverride,omitempty" yaml:"Other2EachOverride,omitempty"`

	Other3EachOverride float64 `xml:"Other3EachOverride,omitempty" json:"Other3EachOverride,omitempty" yaml:"Other3EachOverride,omitempty"`

	Other4EachOverride float64 `xml:"Other4EachOverride,omitempty" json:"Other4EachOverride,omitempty" yaml:"Other4EachOverride,omitempty"`

	Other5EachOverride float64 `xml:"Other5EachOverride,omitempty" json:"Other5EachOverride,omitempty" yaml:"Other5EachOverride,omitempty"`

	Other6EachOverride float64 `xml:"Other6EachOverride,omitempty" json:"Other6EachOverride,omitempty" yaml:"Other6EachOverride,omitempty"`

	Other7EachOverride float64 `xml:"Other7EachOverride,omitempty" json:"Other7EachOverride,omitempty" yaml:"Other7EachOverride,omitempty"`

	Other8EachOverride float64 `xml:"Other8EachOverride,omitempty" json:"Other8EachOverride,omitempty" yaml:"Other8EachOverride,omitempty"`

	Other9EachOverride float64 `xml:"Other9EachOverride,omitempty" json:"Other9EachOverride,omitempty" yaml:"Other9EachOverride,omitempty"`

	Other10EachOverride float64 `xml:"Other10EachOverride,omitempty" json:"Other10EachOverride,omitempty" yaml:"Other10EachOverride,omitempty"`

	DescriptionOverride string `xml:"DescriptionOverride,omitempty" json:"DescriptionOverride,omitempty" yaml:"DescriptionOverride,omitempty"`

	Reference1 string `xml:"Reference1,omitempty" json:"Reference1,omitempty" yaml:"Reference1,omitempty"`

	AdvancedAutoOptions *AdvancedAutoOptionsRequest `xml:"AdvancedAutoOptions,omitempty" json:"AdvancedAutoOptions,omitempty" yaml:"AdvancedAutoOptions,omitempty"`
}

type AdvancedAutoOptionsRequest struct {
	ProcessWhileDate time.Time `xml:"ProcessWhileDate,omitempty" json:"ProcessWhileDate,omitempty" yaml:"ProcessWhileDate,omitempty"`

	SkipUntilDate time.Time `xml:"SkipUntilDate,omitempty" json:"SkipUntilDate,omitempty" yaml:"SkipUntilDate,omitempty"`
}

type CreateAutoOrderRequest struct {
	*BaseCalculateOrderRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	Frequency *FrequencyType `xml:"Frequency,omitempty" json:"Frequency,omitempty" yaml:"Frequency,omitempty"`

	StartDate time.Time `xml:"StartDate,omitempty" json:"StartDate,omitempty" yaml:"StartDate,omitempty"`

	StopDate time.Time `xml:"StopDate,omitempty" json:"StopDate,omitempty" yaml:"StopDate,omitempty"`

	SpecificDayInterval int32 `xml:"SpecificDayInterval,omitempty" json:"SpecificDayInterval,omitempty" yaml:"SpecificDayInterval,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty" yaml:"CurrencyCode,omitempty"`

	WarehouseID int32 `xml:"WarehouseID,omitempty" json:"WarehouseID,omitempty" yaml:"WarehouseID,omitempty"`

	ShipMethodID int32 `xml:"ShipMethodID,omitempty" json:"ShipMethodID,omitempty" yaml:"ShipMethodID,omitempty"`

	PriceType int32 `xml:"PriceType,omitempty" json:"PriceType,omitempty" yaml:"PriceType,omitempty"`

	PaymentType *AutoOrderPaymentType `xml:"PaymentType,omitempty" json:"PaymentType,omitempty" yaml:"PaymentType,omitempty"`

	ProcessType *AutoOrderProcessType `xml:"ProcessType,omitempty" json:"ProcessType,omitempty" yaml:"ProcessType,omitempty"`

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty" json:"MiddleName,omitempty" yaml:"MiddleName,omitempty"`

	LastName string `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`

	NameSuffix string `xml:"NameSuffix,omitempty" json:"NameSuffix,omitempty" yaml:"NameSuffix,omitempty"`

	Company string `xml:"Company,omitempty" json:"Company,omitempty" yaml:"Company,omitempty"`

	Address1 string `xml:"Address1,omitempty" json:"Address1,omitempty" yaml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty" json:"Address2,omitempty" yaml:"Address2,omitempty"`

	Address3 string `xml:"Address3,omitempty" json:"Address3,omitempty" yaml:"Address3,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty" json:"Zip,omitempty" yaml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`

	County string `xml:"County,omitempty" json:"County,omitempty" yaml:"County,omitempty"`

	Email string `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`

	Phone string `xml:"Phone,omitempty" json:"Phone,omitempty" yaml:"Phone,omitempty"`

	Notes string `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`

	Other11 string `xml:"Other11,omitempty" json:"Other11,omitempty" yaml:"Other11,omitempty"`

	Other12 string `xml:"Other12,omitempty" json:"Other12,omitempty" yaml:"Other12,omitempty"`

	Other13 string `xml:"Other13,omitempty" json:"Other13,omitempty" yaml:"Other13,omitempty"`

	Other14 string `xml:"Other14,omitempty" json:"Other14,omitempty" yaml:"Other14,omitempty"`

	Other15 string `xml:"Other15,omitempty" json:"Other15,omitempty" yaml:"Other15,omitempty"`

	Other16 string `xml:"Other16,omitempty" json:"Other16,omitempty" yaml:"Other16,omitempty"`

	Other17 string `xml:"Other17,omitempty" json:"Other17,omitempty" yaml:"Other17,omitempty"`

	Other18 string `xml:"Other18,omitempty" json:"Other18,omitempty" yaml:"Other18,omitempty"`

	Other19 string `xml:"Other19,omitempty" json:"Other19,omitempty" yaml:"Other19,omitempty"`

	Other20 string `xml:"Other20,omitempty" json:"Other20,omitempty" yaml:"Other20,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	OverwriteExistingAutoOrder bool `xml:"OverwriteExistingAutoOrder,omitempty" json:"OverwriteExistingAutoOrder,omitempty" yaml:"OverwriteExistingAutoOrder,omitempty"`

	ExistingAutoOrderID int32 `xml:"ExistingAutoOrderID,omitempty" json:"ExistingAutoOrderID,omitempty" yaml:"ExistingAutoOrderID,omitempty"`

	Details *ArrayOfOrderDetailRequest `xml:"Details,omitempty" json:"Details,omitempty" yaml:"Details,omitempty"`
}

type CreateOrderRequest struct {
	*BaseCalculateOrderRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	OrderStatus *OrderStatusType `xml:"OrderStatus,omitempty" json:"OrderStatus,omitempty" yaml:"OrderStatus,omitempty"`

	OrderDate time.Time `xml:"OrderDate,omitempty" json:"OrderDate,omitempty" yaml:"OrderDate,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty" yaml:"CurrencyCode,omitempty"`

	WarehouseID int32 `xml:"WarehouseID,omitempty" json:"WarehouseID,omitempty" yaml:"WarehouseID,omitempty"`

	ShipMethodID int32 `xml:"ShipMethodID,omitempty" json:"ShipMethodID,omitempty" yaml:"ShipMethodID,omitempty"`

	PriceType int32 `xml:"PriceType,omitempty" json:"PriceType,omitempty" yaml:"PriceType,omitempty"`

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty" json:"MiddleName,omitempty" yaml:"MiddleName,omitempty"`

	LastName string `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`

	NameSuffix string `xml:"NameSuffix,omitempty" json:"NameSuffix,omitempty" yaml:"NameSuffix,omitempty"`

	Company string `xml:"Company,omitempty" json:"Company,omitempty" yaml:"Company,omitempty"`

	Address1 string `xml:"Address1,omitempty" json:"Address1,omitempty" yaml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty" json:"Address2,omitempty" yaml:"Address2,omitempty"`

	Address3 string `xml:"Address3,omitempty" json:"Address3,omitempty" yaml:"Address3,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty" json:"Zip,omitempty" yaml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`

	County string `xml:"County,omitempty" json:"County,omitempty" yaml:"County,omitempty"`

	Email string `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`

	Phone string `xml:"Phone,omitempty" json:"Phone,omitempty" yaml:"Phone,omitempty"`

	Notes string `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`

	Other11 string `xml:"Other11,omitempty" json:"Other11,omitempty" yaml:"Other11,omitempty"`

	Other12 string `xml:"Other12,omitempty" json:"Other12,omitempty" yaml:"Other12,omitempty"`

	Other13 string `xml:"Other13,omitempty" json:"Other13,omitempty" yaml:"Other13,omitempty"`

	Other14 string `xml:"Other14,omitempty" json:"Other14,omitempty" yaml:"Other14,omitempty"`

	Other15 string `xml:"Other15,omitempty" json:"Other15,omitempty" yaml:"Other15,omitempty"`

	Other16 string `xml:"Other16,omitempty" json:"Other16,omitempty" yaml:"Other16,omitempty"`

	Other17 string `xml:"Other17,omitempty" json:"Other17,omitempty" yaml:"Other17,omitempty"`

	Other18 string `xml:"Other18,omitempty" json:"Other18,omitempty" yaml:"Other18,omitempty"`

	Other19 string `xml:"Other19,omitempty" json:"Other19,omitempty" yaml:"Other19,omitempty"`

	Other20 string `xml:"Other20,omitempty" json:"Other20,omitempty" yaml:"Other20,omitempty"`

	OrderType *OrderType `xml:"OrderType,omitempty" json:"OrderType,omitempty" yaml:"OrderType,omitempty"`

	TaxRateOverride float64 `xml:"TaxRateOverride,omitempty" json:"TaxRateOverride,omitempty" yaml:"TaxRateOverride,omitempty"`

	ShippingAmountOverride float64 `xml:"ShippingAmountOverride,omitempty" json:"ShippingAmountOverride,omitempty" yaml:"ShippingAmountOverride,omitempty"`

	UseManualOrderID bool `xml:"UseManualOrderID,omitempty" json:"UseManualOrderID,omitempty" yaml:"UseManualOrderID,omitempty"`

	ManualOrderID int32 `xml:"ManualOrderID,omitempty" json:"ManualOrderID,omitempty" yaml:"ManualOrderID,omitempty"`

	TransferVolumeToID int32 `xml:"TransferVolumeToID,omitempty" json:"TransferVolumeToID,omitempty" yaml:"TransferVolumeToID,omitempty"`

	ReturnOrderID int32 `xml:"ReturnOrderID,omitempty" json:"ReturnOrderID,omitempty" yaml:"ReturnOrderID,omitempty"`

	OverwriteExistingOrder bool `xml:"OverwriteExistingOrder,omitempty" json:"OverwriteExistingOrder,omitempty" yaml:"OverwriteExistingOrder,omitempty"`

	ExistingOrderID int32 `xml:"ExistingOrderID,omitempty" json:"ExistingOrderID,omitempty" yaml:"ExistingOrderID,omitempty"`

	PartyID int32 `xml:"PartyID,omitempty" json:"PartyID,omitempty" yaml:"PartyID,omitempty"`

	Details *ArrayOfOrderDetailRequest `xml:"Details,omitempty" json:"Details,omitempty" yaml:"Details,omitempty"`

	SuppressPackSlipPrice bool `xml:"SuppressPackSlipPrice,omitempty" json:"SuppressPackSlipPrice,omitempty" yaml:"SuppressPackSlipPrice,omitempty"`
}

type SetAccountDirectDepositRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	NameOnAccount string `xml:"NameOnAccount,omitempty" json:"NameOnAccount,omitempty" yaml:"NameOnAccount,omitempty"`

	BankAccountNumber string `xml:"BankAccountNumber,omitempty" json:"BankAccountNumber,omitempty" yaml:"BankAccountNumber,omitempty"`

	BankRoutingNumber string `xml:"BankRoutingNumber,omitempty" json:"BankRoutingNumber,omitempty" yaml:"BankRoutingNumber,omitempty"`

	DepositAccountType *DepositAccountType `xml:"DepositAccountType,omitempty" json:"DepositAccountType,omitempty" yaml:"DepositAccountType,omitempty"`

	BankName string `xml:"BankName,omitempty" json:"BankName,omitempty" yaml:"BankName,omitempty"`

	BankAddress string `xml:"BankAddress,omitempty" json:"BankAddress,omitempty" yaml:"BankAddress,omitempty"`

	BankCity string `xml:"BankCity,omitempty" json:"BankCity,omitempty" yaml:"BankCity,omitempty"`

	BankState string `xml:"BankState,omitempty" json:"BankState,omitempty" yaml:"BankState,omitempty"`

	BankZip string `xml:"BankZip,omitempty" json:"BankZip,omitempty" yaml:"BankZip,omitempty"`

	BankCountry string `xml:"BankCountry,omitempty" json:"BankCountry,omitempty" yaml:"BankCountry,omitempty"`

	Iban string `xml:"Iban,omitempty" json:"Iban,omitempty" yaml:"Iban,omitempty"`

	SwiftCode string `xml:"SwiftCode,omitempty" json:"SwiftCode,omitempty" yaml:"SwiftCode,omitempty"`

	BankAccountType *BankAccountType `xml:"BankAccountType,omitempty" json:"BankAccountType,omitempty" yaml:"BankAccountType,omitempty"`
}

type SetAccountWalletRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	WalletAccountType *AccountWalletType `xml:"WalletAccountType,omitempty" json:"WalletAccountType,omitempty" yaml:"WalletAccountType,omitempty"`

	WalletType int32 `xml:"WalletType,omitempty" json:"WalletType,omitempty" yaml:"WalletType,omitempty"`

	WalletAccount string `xml:"WalletAccount,omitempty" json:"WalletAccount,omitempty" yaml:"WalletAccount,omitempty"`
}

type SetAccountCreditCardTokenRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	CreditCardAccountType *AccountCreditCardType `xml:"CreditCardAccountType,omitempty" json:"CreditCardAccountType,omitempty" yaml:"CreditCardAccountType,omitempty"`

	CreditCardToken string `xml:"CreditCardToken,omitempty" json:"CreditCardToken,omitempty" yaml:"CreditCardToken,omitempty"`

	ExpirationMonth int32 `xml:"ExpirationMonth,omitempty" json:"ExpirationMonth,omitempty" yaml:"ExpirationMonth,omitempty"`

	ExpirationYear int32 `xml:"ExpirationYear,omitempty" json:"ExpirationYear,omitempty" yaml:"ExpirationYear,omitempty"`

	CreditCardType int32 `xml:"CreditCardType,omitempty" json:"CreditCardType,omitempty" yaml:"CreditCardType,omitempty"`

	BillingName string `xml:"BillingName,omitempty" json:"BillingName,omitempty" yaml:"BillingName,omitempty"`

	UseMainAddress bool `xml:"UseMainAddress,omitempty" json:"UseMainAddress,omitempty" yaml:"UseMainAddress,omitempty"`

	BillingAddress string `xml:"BillingAddress,omitempty" json:"BillingAddress,omitempty" yaml:"BillingAddress,omitempty"`

	BillingAddress2 string `xml:"BillingAddress2,omitempty" json:"BillingAddress2,omitempty" yaml:"BillingAddress2,omitempty"`

	BillingCity string `xml:"BillingCity,omitempty" json:"BillingCity,omitempty" yaml:"BillingCity,omitempty"`

	BillingState string `xml:"BillingState,omitempty" json:"BillingState,omitempty" yaml:"BillingState,omitempty"`

	BillingZip string `xml:"BillingZip,omitempty" json:"BillingZip,omitempty" yaml:"BillingZip,omitempty"`

	BillingCountry string `xml:"BillingCountry,omitempty" json:"BillingCountry,omitempty" yaml:"BillingCountry,omitempty"`

	HideFromWeb bool `xml:"HideFromWeb,omitempty" json:"HideFromWeb,omitempty" yaml:"HideFromWeb,omitempty"`
}

type SetAccountCreditCardRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	CreditCardAccountType *AccountCreditCardType `xml:"CreditCardAccountType,omitempty" json:"CreditCardAccountType,omitempty" yaml:"CreditCardAccountType,omitempty"`

	CreditCardNumber string `xml:"CreditCardNumber,omitempty" json:"CreditCardNumber,omitempty" yaml:"CreditCardNumber,omitempty"`

	ExpirationMonth int32 `xml:"ExpirationMonth,omitempty" json:"ExpirationMonth,omitempty" yaml:"ExpirationMonth,omitempty"`

	ExpirationYear int32 `xml:"ExpirationYear,omitempty" json:"ExpirationYear,omitempty" yaml:"ExpirationYear,omitempty"`

	CvcCode string `xml:"CvcCode,omitempty" json:"CvcCode,omitempty" yaml:"CvcCode,omitempty"`

	IssueCode string `xml:"IssueCode,omitempty" json:"IssueCode,omitempty" yaml:"IssueCode,omitempty"`

	CreditCardType int32 `xml:"CreditCardType,omitempty" json:"CreditCardType,omitempty" yaml:"CreditCardType,omitempty"`

	BillingName string `xml:"BillingName,omitempty" json:"BillingName,omitempty" yaml:"BillingName,omitempty"`

	UseMainAddress bool `xml:"UseMainAddress,omitempty" json:"UseMainAddress,omitempty" yaml:"UseMainAddress,omitempty"`

	BillingAddress string `xml:"BillingAddress,omitempty" json:"BillingAddress,omitempty" yaml:"BillingAddress,omitempty"`

	BillingCity string `xml:"BillingCity,omitempty" json:"BillingCity,omitempty" yaml:"BillingCity,omitempty"`

	BillingState string `xml:"BillingState,omitempty" json:"BillingState,omitempty" yaml:"BillingState,omitempty"`

	BillingZip string `xml:"BillingZip,omitempty" json:"BillingZip,omitempty" yaml:"BillingZip,omitempty"`

	BillingCountry string `xml:"BillingCountry,omitempty" json:"BillingCountry,omitempty" yaml:"BillingCountry,omitempty"`

	HideFromWeb bool `xml:"HideFromWeb,omitempty" json:"HideFromWeb,omitempty" yaml:"HideFromWeb,omitempty"`
}

type SetAccountCheckingRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	BankAccountNumber string `xml:"BankAccountNumber,omitempty" json:"BankAccountNumber,omitempty" yaml:"BankAccountNumber,omitempty"`

	BankRoutingNumber string `xml:"BankRoutingNumber,omitempty" json:"BankRoutingNumber,omitempty" yaml:"BankRoutingNumber,omitempty"`

	BankName string `xml:"BankName,omitempty" json:"BankName,omitempty" yaml:"BankName,omitempty"`

	BankAccountType *BankAccountType `xml:"BankAccountType,omitempty" json:"BankAccountType,omitempty" yaml:"BankAccountType,omitempty"`

	NameOnAccount string `xml:"NameOnAccount,omitempty" json:"NameOnAccount,omitempty" yaml:"NameOnAccount,omitempty"`

	UseMainAddress bool `xml:"UseMainAddress,omitempty" json:"UseMainAddress,omitempty" yaml:"UseMainAddress,omitempty"`

	BillingAddress string `xml:"BillingAddress,omitempty" json:"BillingAddress,omitempty" yaml:"BillingAddress,omitempty"`

	BillingCity string `xml:"BillingCity,omitempty" json:"BillingCity,omitempty" yaml:"BillingCity,omitempty"`

	BillingState string `xml:"BillingState,omitempty" json:"BillingState,omitempty" yaml:"BillingState,omitempty"`

	BillingZip string `xml:"BillingZip,omitempty" json:"BillingZip,omitempty" yaml:"BillingZip,omitempty"`

	BillingCountry string `xml:"BillingCountry,omitempty" json:"BillingCountry,omitempty" yaml:"BillingCountry,omitempty"`

	DriversLicenseNumber string `xml:"DriversLicenseNumber,omitempty" json:"DriversLicenseNumber,omitempty" yaml:"DriversLicenseNumber,omitempty"`
}

type CreateOrderImportRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`

	Company string `xml:"Company,omitempty" json:"Company,omitempty" yaml:"Company,omitempty"`

	Phone string `xml:"Phone,omitempty" json:"Phone,omitempty" yaml:"Phone,omitempty"`

	Email string `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`

	ShipMethodID int32 `xml:"ShipMethodID,omitempty" json:"ShipMethodID,omitempty" yaml:"ShipMethodID,omitempty"`

	Address1 string `xml:"Address1,omitempty" json:"Address1,omitempty" yaml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty" json:"Address2,omitempty" yaml:"Address2,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty" json:"Zip,omitempty" yaml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`

	County string `xml:"County,omitempty" json:"County,omitempty" yaml:"County,omitempty"`

	Notes string `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`

	WarehouseID int32 `xml:"WarehouseID,omitempty" json:"WarehouseID,omitempty" yaml:"WarehouseID,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty" yaml:"CurrencyCode,omitempty"`

	ShippingStateTax float64 `xml:"ShippingStateTax,omitempty" json:"ShippingStateTax,omitempty" yaml:"ShippingStateTax,omitempty"`

	ShippingFedTax float64 `xml:"ShippingFedTax,omitempty" json:"ShippingFedTax,omitempty" yaml:"ShippingFedTax,omitempty"`

	ShippingCountyLocalTax float64 `xml:"ShippingCountyLocalTax,omitempty" json:"ShippingCountyLocalTax,omitempty" yaml:"ShippingCountyLocalTax,omitempty"`

	ShippingCountyTax float64 `xml:"ShippingCountyTax,omitempty" json:"ShippingCountyTax,omitempty" yaml:"ShippingCountyTax,omitempty"`

	ShippingCityLocalTax float64 `xml:"ShippingCityLocalTax,omitempty" json:"ShippingCityLocalTax,omitempty" yaml:"ShippingCityLocalTax,omitempty"`

	ShippingCityTax float64 `xml:"ShippingCityTax,omitempty" json:"ShippingCityTax,omitempty" yaml:"ShippingCityTax,omitempty"`

	Shipping float64 `xml:"Shipping,omitempty" json:"Shipping,omitempty" yaml:"Shipping,omitempty"`

	PriceType int32 `xml:"PriceType,omitempty" json:"PriceType,omitempty" yaml:"PriceType,omitempty"`

	OrderStatus *OrderStatusType `xml:"OrderStatus,omitempty" json:"OrderStatus,omitempty" yaml:"OrderStatus,omitempty"`

	OrderDate time.Time `xml:"OrderDate,omitempty" json:"OrderDate,omitempty" yaml:"OrderDate,omitempty"`

	Other11 string `xml:"Other11,omitempty" json:"Other11,omitempty" yaml:"Other11,omitempty"`

	Other12 string `xml:"Other12,omitempty" json:"Other12,omitempty" yaml:"Other12,omitempty"`

	Other13 string `xml:"Other13,omitempty" json:"Other13,omitempty" yaml:"Other13,omitempty"`

	Other14 string `xml:"Other14,omitempty" json:"Other14,omitempty" yaml:"Other14,omitempty"`

	Other15 string `xml:"Other15,omitempty" json:"Other15,omitempty" yaml:"Other15,omitempty"`

	Other16 string `xml:"Other16,omitempty" json:"Other16,omitempty" yaml:"Other16,omitempty"`

	Other17 string `xml:"Other17,omitempty" json:"Other17,omitempty" yaml:"Other17,omitempty"`

	Other18 string `xml:"Other18,omitempty" json:"Other18,omitempty" yaml:"Other18,omitempty"`

	Other19 string `xml:"Other19,omitempty" json:"Other19,omitempty" yaml:"Other19,omitempty"`

	Other20 string `xml:"Other20,omitempty" json:"Other20,omitempty" yaml:"Other20,omitempty"`

	OrderType *OrderType `xml:"OrderType,omitempty" json:"OrderType,omitempty" yaml:"OrderType,omitempty"`

	UseManualOrderID bool `xml:"UseManualOrderID,omitempty" json:"UseManualOrderID,omitempty" yaml:"UseManualOrderID,omitempty"`

	ManualOrderID int32 `xml:"ManualOrderID,omitempty" json:"ManualOrderID,omitempty" yaml:"ManualOrderID,omitempty"`

	ReturnOrderID int32 `xml:"ReturnOrderID,omitempty" json:"ReturnOrderID,omitempty" yaml:"ReturnOrderID,omitempty"`

	OrderDetails *ArrayOfOrderImportDetail `xml:"OrderDetails,omitempty" json:"OrderDetails,omitempty" yaml:"OrderDetails,omitempty"`

	PartyID int32 `xml:"PartyID,omitempty" json:"PartyID,omitempty" yaml:"PartyID,omitempty"`
}

type ArrayOfOrderImportDetail struct {
	OrderImportDetail []*OrderImportDetail `xml:"OrderImportDetail,omitempty" json:"OrderImportDetail,omitempty" yaml:"OrderImportDetail,omitempty"`
}

type OrderImportDetail struct {
	ParentItemCode string `xml:"ParentItemCode,omitempty" json:"ParentItemCode,omitempty" yaml:"ParentItemCode,omitempty"`

	ItemCode string `xml:"ItemCode,omitempty" json:"ItemCode,omitempty" yaml:"ItemCode,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	Qty float64 `xml:"Qty,omitempty" json:"Qty,omitempty" yaml:"Qty,omitempty"`

	WeightEach float64 `xml:"WeightEach,omitempty" json:"WeightEach,omitempty" yaml:"WeightEach,omitempty"`

	CountyLocalTax float64 `xml:"CountyLocalTax,omitempty" json:"CountyLocalTax,omitempty" yaml:"CountyLocalTax,omitempty"`

	CountyTax float64 `xml:"CountyTax,omitempty" json:"CountyTax,omitempty" yaml:"CountyTax,omitempty"`

	CityTax float64 `xml:"CityTax,omitempty" json:"CityTax,omitempty" yaml:"CityTax,omitempty"`

	StateTax float64 `xml:"StateTax,omitempty" json:"StateTax,omitempty" yaml:"StateTax,omitempty"`

	FedTax float64 `xml:"FedTax,omitempty" json:"FedTax,omitempty" yaml:"FedTax,omitempty"`

	TaxablePriceEach float64 `xml:"TaxablePriceEach,omitempty" json:"TaxablePriceEach,omitempty" yaml:"TaxablePriceEach,omitempty"`

	CVEach float64 `xml:"CVEach,omitempty" json:"CVEach,omitempty" yaml:"CVEach,omitempty"`

	BVEach float64 `xml:"BVEach,omitempty" json:"BVEach,omitempty" yaml:"BVEach,omitempty"`

	PriceEach float64 `xml:"PriceEach,omitempty" json:"PriceEach,omitempty" yaml:"PriceEach,omitempty"`

	Other10Each float64 `xml:"Other10Each,omitempty" json:"Other10Each,omitempty" yaml:"Other10Each,omitempty"`

	Other9Each float64 `xml:"Other9Each,omitempty" json:"Other9Each,omitempty" yaml:"Other9Each,omitempty"`

	Other8Each float64 `xml:"Other8Each,omitempty" json:"Other8Each,omitempty" yaml:"Other8Each,omitempty"`

	Other7Each float64 `xml:"Other7Each,omitempty" json:"Other7Each,omitempty" yaml:"Other7Each,omitempty"`

	Other6Each float64 `xml:"Other6Each,omitempty" json:"Other6Each,omitempty" yaml:"Other6Each,omitempty"`

	Other5Each float64 `xml:"Other5Each,omitempty" json:"Other5Each,omitempty" yaml:"Other5Each,omitempty"`

	Other4Each float64 `xml:"Other4Each,omitempty" json:"Other4Each,omitempty" yaml:"Other4Each,omitempty"`

	Other3Each float64 `xml:"Other3Each,omitempty" json:"Other3Each,omitempty" yaml:"Other3Each,omitempty"`

	Other2Each float64 `xml:"Other2Each,omitempty" json:"Other2Each,omitempty" yaml:"Other2Each,omitempty"`

	Other1Each float64 `xml:"Other1Each,omitempty" json:"Other1Each,omitempty" yaml:"Other1Each,omitempty"`
}

type CreateCustomerRequest struct {
	*ApiRequest

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`

	Company string `xml:"Company,omitempty" json:"Company,omitempty" yaml:"Company,omitempty"`

	CustomerType int32 `xml:"CustomerType,omitempty" json:"CustomerType,omitempty" yaml:"CustomerType,omitempty"`

	CustomerStatus int32 `xml:"CustomerStatus,omitempty" json:"CustomerStatus,omitempty" yaml:"CustomerStatus,omitempty"`

	Email string `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`

	Phone string `xml:"Phone,omitempty" json:"Phone,omitempty" yaml:"Phone,omitempty"`

	Phone2 string `xml:"Phone2,omitempty" json:"Phone2,omitempty" yaml:"Phone2,omitempty"`

	MobilePhone string `xml:"MobilePhone,omitempty" json:"MobilePhone,omitempty" yaml:"MobilePhone,omitempty"`

	Fax string `xml:"Fax,omitempty" json:"Fax,omitempty" yaml:"Fax,omitempty"`

	Notes string `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`

	MainAddress1 string `xml:"MainAddress1,omitempty" json:"MainAddress1,omitempty" yaml:"MainAddress1,omitempty"`

	MainAddress2 string `xml:"MainAddress2,omitempty" json:"MainAddress2,omitempty" yaml:"MainAddress2,omitempty"`

	MainAddress3 string `xml:"MainAddress3,omitempty" json:"MainAddress3,omitempty" yaml:"MainAddress3,omitempty"`

	MainCity string `xml:"MainCity,omitempty" json:"MainCity,omitempty" yaml:"MainCity,omitempty"`

	MainState string `xml:"MainState,omitempty" json:"MainState,omitempty" yaml:"MainState,omitempty"`

	MainZip string `xml:"MainZip,omitempty" json:"MainZip,omitempty" yaml:"MainZip,omitempty"`

	MainCountry string `xml:"MainCountry,omitempty" json:"MainCountry,omitempty" yaml:"MainCountry,omitempty"`

	MainCounty string `xml:"MainCounty,omitempty" json:"MainCounty,omitempty" yaml:"MainCounty,omitempty"`

	MailAddress1 string `xml:"MailAddress1,omitempty" json:"MailAddress1,omitempty" yaml:"MailAddress1,omitempty"`

	MailAddress2 string `xml:"MailAddress2,omitempty" json:"MailAddress2,omitempty" yaml:"MailAddress2,omitempty"`

	MailAddress3 string `xml:"MailAddress3,omitempty" json:"MailAddress3,omitempty" yaml:"MailAddress3,omitempty"`

	MailCity string `xml:"MailCity,omitempty" json:"MailCity,omitempty" yaml:"MailCity,omitempty"`

	MailState string `xml:"MailState,omitempty" json:"MailState,omitempty" yaml:"MailState,omitempty"`

	MailZip string `xml:"MailZip,omitempty" json:"MailZip,omitempty" yaml:"MailZip,omitempty"`

	MailCountry string `xml:"MailCountry,omitempty" json:"MailCountry,omitempty" yaml:"MailCountry,omitempty"`

	MailCounty string `xml:"MailCounty,omitempty" json:"MailCounty,omitempty" yaml:"MailCounty,omitempty"`

	OtherAddress1 string `xml:"OtherAddress1,omitempty" json:"OtherAddress1,omitempty" yaml:"OtherAddress1,omitempty"`

	OtherAddress2 string `xml:"OtherAddress2,omitempty" json:"OtherAddress2,omitempty" yaml:"OtherAddress2,omitempty"`

	OtherAddress3 string `xml:"OtherAddress3,omitempty" json:"OtherAddress3,omitempty" yaml:"OtherAddress3,omitempty"`

	OtherCity string `xml:"OtherCity,omitempty" json:"OtherCity,omitempty" yaml:"OtherCity,omitempty"`

	OtherState string `xml:"OtherState,omitempty" json:"OtherState,omitempty" yaml:"OtherState,omitempty"`

	OtherZip string `xml:"OtherZip,omitempty" json:"OtherZip,omitempty" yaml:"OtherZip,omitempty"`

	OtherCountry string `xml:"OtherCountry,omitempty" json:"OtherCountry,omitempty" yaml:"OtherCountry,omitempty"`

	OtherCounty string `xml:"OtherCounty,omitempty" json:"OtherCounty,omitempty" yaml:"OtherCounty,omitempty"`

	CanLogin bool `xml:"CanLogin,omitempty" json:"CanLogin,omitempty" yaml:"CanLogin,omitempty"`

	LoginName string `xml:"LoginName,omitempty" json:"LoginName,omitempty" yaml:"LoginName,omitempty"`

	LoginPassword string `xml:"LoginPassword,omitempty" json:"LoginPassword,omitempty" yaml:"LoginPassword,omitempty"`

	InsertEnrollerTree bool `xml:"InsertEnrollerTree,omitempty" json:"InsertEnrollerTree,omitempty" yaml:"InsertEnrollerTree,omitempty"`

	EnrollerID int32 `xml:"EnrollerID,omitempty" json:"EnrollerID,omitempty" yaml:"EnrollerID,omitempty"`

	InsertUnilevelTree bool `xml:"InsertUnilevelTree,omitempty" json:"InsertUnilevelTree,omitempty" yaml:"InsertUnilevelTree,omitempty"`

	SponsorID int32 `xml:"SponsorID,omitempty" json:"SponsorID,omitempty" yaml:"SponsorID,omitempty"`

	UseManualCustomerID bool `xml:"UseManualCustomerID,omitempty" json:"UseManualCustomerID,omitempty" yaml:"UseManualCustomerID,omitempty"`

	ManualCustomerID int32 `xml:"ManualCustomerID,omitempty" json:"ManualCustomerID,omitempty" yaml:"ManualCustomerID,omitempty"`

	TaxID string `xml:"TaxID,omitempty" json:"TaxID,omitempty" yaml:"TaxID,omitempty"`

	SalesTaxID string `xml:"SalesTaxID,omitempty" json:"SalesTaxID,omitempty" yaml:"SalesTaxID,omitempty"`

	SalesTaxExemptExpireDate time.Time `xml:"SalesTaxExemptExpireDate,omitempty" json:"SalesTaxExemptExpireDate,omitempty" yaml:"SalesTaxExemptExpireDate,omitempty"`

	IsSalesTaxExempt bool `xml:"IsSalesTaxExempt,omitempty" json:"IsSalesTaxExempt,omitempty" yaml:"IsSalesTaxExempt,omitempty"`

	BirthDate time.Time `xml:"BirthDate,omitempty" json:"BirthDate,omitempty" yaml:"BirthDate,omitempty"`

	Field1 string `xml:"Field1,omitempty" json:"Field1,omitempty" yaml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty" json:"Field2,omitempty" yaml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty" json:"Field3,omitempty" yaml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty" json:"Field4,omitempty" yaml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty" json:"Field5,omitempty" yaml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty" json:"Field6,omitempty" yaml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty" json:"Field7,omitempty" yaml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty" json:"Field8,omitempty" yaml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty" json:"Field9,omitempty" yaml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty" json:"Field10,omitempty" yaml:"Field10,omitempty"`

	Field11 string `xml:"Field11,omitempty" json:"Field11,omitempty" yaml:"Field11,omitempty"`

	Field12 string `xml:"Field12,omitempty" json:"Field12,omitempty" yaml:"Field12,omitempty"`

	Field13 string `xml:"Field13,omitempty" json:"Field13,omitempty" yaml:"Field13,omitempty"`

	Field14 string `xml:"Field14,omitempty" json:"Field14,omitempty" yaml:"Field14,omitempty"`

	Field15 string `xml:"Field15,omitempty" json:"Field15,omitempty" yaml:"Field15,omitempty"`

	SubscribeToBroadcasts bool `xml:"SubscribeToBroadcasts,omitempty" json:"SubscribeToBroadcasts,omitempty" yaml:"SubscribeToBroadcasts,omitempty"`

	SubscribeFromIPAddress string `xml:"SubscribeFromIPAddress,omitempty" json:"SubscribeFromIPAddress,omitempty" yaml:"SubscribeFromIPAddress,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty" yaml:"CurrencyCode,omitempty"`

	PayableToName string `xml:"PayableToName,omitempty" json:"PayableToName,omitempty" yaml:"PayableToName,omitempty"`

	EntryDate time.Time `xml:"EntryDate,omitempty" json:"EntryDate,omitempty" yaml:"EntryDate,omitempty"`

	DefaultWarehouseID int32 `xml:"DefaultWarehouseID,omitempty" json:"DefaultWarehouseID,omitempty" yaml:"DefaultWarehouseID,omitempty"`

	PayableType *PayableType `xml:"PayableType,omitempty" json:"PayableType,omitempty" yaml:"PayableType,omitempty"`

	CheckThreshold float64 `xml:"CheckThreshold,omitempty" json:"CheckThreshold,omitempty" yaml:"CheckThreshold,omitempty"`

	TaxIDType *TaxIDType `xml:"TaxIDType,omitempty" json:"TaxIDType,omitempty" yaml:"TaxIDType,omitempty"`

	LanguageID int32 `xml:"LanguageID,omitempty" json:"LanguageID,omitempty" yaml:"LanguageID,omitempty"`

	Gender *Gender `xml:"Gender,omitempty" json:"Gender,omitempty" yaml:"Gender,omitempty"`

	VatRegistration string `xml:"VatRegistration,omitempty" json:"VatRegistration,omitempty" yaml:"VatRegistration,omitempty"`

	Date1 time.Time `xml:"Date1,omitempty" json:"Date1,omitempty" yaml:"Date1,omitempty"`

	Date2 time.Time `xml:"Date2,omitempty" json:"Date2,omitempty" yaml:"Date2,omitempty"`

	Date3 time.Time `xml:"Date3,omitempty" json:"Date3,omitempty" yaml:"Date3,omitempty"`

	Date4 time.Time `xml:"Date4,omitempty" json:"Date4,omitempty" yaml:"Date4,omitempty"`

	Date5 time.Time `xml:"Date5,omitempty" json:"Date5,omitempty" yaml:"Date5,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty" json:"MiddleName,omitempty" yaml:"MiddleName,omitempty"`

	NameSuffix string `xml:"NameSuffix,omitempty" json:"NameSuffix,omitempty" yaml:"NameSuffix,omitempty"`

	BinaryPlacementPreference int32 `xml:"BinaryPlacementPreference,omitempty" json:"BinaryPlacementPreference,omitempty" yaml:"BinaryPlacementPreference,omitempty"`

	UseBinaryHoldingTank bool `xml:"UseBinaryHoldingTank,omitempty" json:"UseBinaryHoldingTank,omitempty" yaml:"UseBinaryHoldingTank,omitempty"`

	MainAddressVerified bool `xml:"MainAddressVerified,omitempty" json:"MainAddressVerified,omitempty" yaml:"MainAddressVerified,omitempty"`

	MailAddressVerified bool `xml:"MailAddressVerified,omitempty" json:"MailAddressVerified,omitempty" yaml:"MailAddressVerified,omitempty"`

	OtherAddressVerified bool `xml:"OtherAddressVerified,omitempty" json:"OtherAddressVerified,omitempty" yaml:"OtherAddressVerified,omitempty"`
}

type TransactionalRequest struct {
	*ApiRequest

	TransactionRequests *ArrayOfApiRequest `xml:"TransactionRequests,omitempty" json:"TransactionRequests,omitempty" yaml:"TransactionRequests,omitempty"`
}

type ArrayOfApiRequest struct {
	ApiRequest []*ApiRequest `xml:"ApiRequest,omitempty" json:"ApiRequest,omitempty" yaml:"ApiRequest,omitempty"`
}

type BaseCreateExpectedPaymentRequest struct {
	*ApiRequest
}

type CreateExpectedBankWireRequest struct {
	*BaseCreateExpectedPaymentRequest

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	PaymentDate time.Time `xml:"PaymentDate,omitempty" json:"PaymentDate,omitempty" yaml:"PaymentDate,omitempty"`

	Amount float64 `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`

	BankName string `xml:"BankName,omitempty" json:"BankName,omitempty" yaml:"BankName,omitempty"`

	NameOnAccount string `xml:"NameOnAccount,omitempty" json:"NameOnAccount,omitempty" yaml:"NameOnAccount,omitempty"`

	Memo string `xml:"Memo,omitempty" json:"Memo,omitempty" yaml:"Memo,omitempty"`

	AuthorizationCode string `xml:"AuthorizationCode,omitempty" json:"AuthorizationCode,omitempty" yaml:"AuthorizationCode,omitempty"`
}

type CreateExpectedCODRequest struct {
	*BaseCreateExpectedPaymentRequest

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	PaymentDate time.Time `xml:"PaymentDate,omitempty" json:"PaymentDate,omitempty" yaml:"PaymentDate,omitempty"`

	Amount float64 `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`
}

type CreateExpectedPaymentRequest struct {
	*BaseCreateExpectedPaymentRequest

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	PaymentDate time.Time `xml:"PaymentDate,omitempty" json:"PaymentDate,omitempty" yaml:"PaymentDate,omitempty"`

	Amount float64 `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`

	PaymentType *PaymentType `xml:"PaymentType,omitempty" json:"PaymentType,omitempty" yaml:"PaymentType,omitempty"`

	AuthorizationCode string `xml:"AuthorizationCode,omitempty" json:"AuthorizationCode,omitempty" yaml:"AuthorizationCode,omitempty"`
}

type BaseCreatePaymentRequest struct {
	*ApiRequest
}

type ChargePriorAuthorizationRequest struct {
	*BaseCreatePaymentRequest

	MerchantTransactionKey string `xml:"MerchantTransactionKey,omitempty" json:"MerchantTransactionKey,omitempty" yaml:"MerchantTransactionKey,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	MaxAmount float64 `xml:"MaxAmount,omitempty" json:"MaxAmount,omitempty" yaml:"MaxAmount,omitempty"`
}

type BaseChargeWalletAccountRequest struct {
	*BaseCreatePaymentRequest
}

type ChargeWalletAccountRequest struct {
	*BaseChargeWalletAccountRequest

	WalletAccountNumber string `xml:"WalletAccountNumber,omitempty" json:"WalletAccountNumber,omitempty" yaml:"WalletAccountNumber,omitempty"`

	WalletTy int32 `xml:"WalletTy,omitempty" json:"WalletTy,omitempty" yaml:"WalletTy,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	Memo string `xml:"Memo,omitempty" json:"Memo,omitempty" yaml:"Memo,omitempty"`

	MaxAmount float64 `xml:"MaxAmount,omitempty" json:"MaxAmount,omitempty" yaml:"MaxAmount,omitempty"`
}

type ChargeWalletAccountOnFileRequest struct {
	*BaseCreatePaymentRequest

	WalletAccountType *AccountWalletType `xml:"WalletAccountType,omitempty" json:"WalletAccountType,omitempty" yaml:"WalletAccountType,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	Memo string `xml:"Memo,omitempty" json:"Memo,omitempty" yaml:"Memo,omitempty"`

	MaxAmount float64 `xml:"MaxAmount,omitempty" json:"MaxAmount,omitempty" yaml:"MaxAmount,omitempty"`
}

type BaseDebitBankAccountRequest struct {
	*BaseCreatePaymentRequest
}

type DebitBankAccountRequest struct {
	*BaseDebitBankAccountRequest

	BankAccountNumber string `xml:"BankAccountNumber,omitempty" json:"BankAccountNumber,omitempty" yaml:"BankAccountNumber,omitempty"`

	BankRoutingNumber string `xml:"BankRoutingNumber,omitempty" json:"BankRoutingNumber,omitempty" yaml:"BankRoutingNumber,omitempty"`

	BankName string `xml:"BankName,omitempty" json:"BankName,omitempty" yaml:"BankName,omitempty"`

	BankAccountType *BankAccountType `xml:"BankAccountType,omitempty" json:"BankAccountType,omitempty" yaml:"BankAccountType,omitempty"`

	CheckNumber string `xml:"CheckNumber,omitempty" json:"CheckNumber,omitempty" yaml:"CheckNumber,omitempty"`

	NameOnAccount string `xml:"NameOnAccount,omitempty" json:"NameOnAccount,omitempty" yaml:"NameOnAccount,omitempty"`

	BillingAddress string `xml:"BillingAddress,omitempty" json:"BillingAddress,omitempty" yaml:"BillingAddress,omitempty"`

	BillingAddress2 string `xml:"BillingAddress2,omitempty" json:"BillingAddress2,omitempty" yaml:"BillingAddress2,omitempty"`

	BillingCity string `xml:"BillingCity,omitempty" json:"BillingCity,omitempty" yaml:"BillingCity,omitempty"`

	BillingState string `xml:"BillingState,omitempty" json:"BillingState,omitempty" yaml:"BillingState,omitempty"`

	BillingZip string `xml:"BillingZip,omitempty" json:"BillingZip,omitempty" yaml:"BillingZip,omitempty"`

	BillingCountry string `xml:"BillingCountry,omitempty" json:"BillingCountry,omitempty" yaml:"BillingCountry,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	MaxAmount float64 `xml:"MaxAmount,omitempty" json:"MaxAmount,omitempty" yaml:"MaxAmount,omitempty"`
}

type DebitBankAccountOnFileRequest struct {
	*BaseDebitBankAccountRequest

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	MaxAmount float64 `xml:"MaxAmount,omitempty" json:"MaxAmount,omitempty" yaml:"MaxAmount,omitempty"`
}

type RefundPriorWalletChargeRequest struct {
	*BaseCreatePaymentRequest

	ReturnPaymentID int32 `xml:"ReturnPaymentID,omitempty" json:"ReturnPaymentID,omitempty" yaml:"ReturnPaymentID,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	MaxAmount float64 `xml:"MaxAmount,omitempty" json:"MaxAmount,omitempty" yaml:"MaxAmount,omitempty"`
}

type CreatePaymentRequest struct {
	*BaseCreatePaymentRequest

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	PaymentDate time.Time `xml:"PaymentDate,omitempty" json:"PaymentDate,omitempty" yaml:"PaymentDate,omitempty"`

	Amount float64 `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`

	PaymentType *PaymentType `xml:"PaymentType,omitempty" json:"PaymentType,omitempty" yaml:"PaymentType,omitempty"`

	AuthorizationCode string `xml:"AuthorizationCode,omitempty" json:"AuthorizationCode,omitempty" yaml:"AuthorizationCode,omitempty"`

	Memo string `xml:"Memo,omitempty" json:"Memo,omitempty" yaml:"Memo,omitempty"`

	BillingName string `xml:"BillingName,omitempty" json:"BillingName,omitempty" yaml:"BillingName,omitempty"`

	BillingAddress string `xml:"BillingAddress,omitempty" json:"BillingAddress,omitempty" yaml:"BillingAddress,omitempty"`

	BillingAddress2 string `xml:"BillingAddress2,omitempty" json:"BillingAddress2,omitempty" yaml:"BillingAddress2,omitempty"`

	BillingCity string `xml:"BillingCity,omitempty" json:"BillingCity,omitempty" yaml:"BillingCity,omitempty"`

	BillingState string `xml:"BillingState,omitempty" json:"BillingState,omitempty" yaml:"BillingState,omitempty"`

	BillingZip string `xml:"BillingZip,omitempty" json:"BillingZip,omitempty" yaml:"BillingZip,omitempty"`

	BillingCountry string `xml:"BillingCountry,omitempty" json:"BillingCountry,omitempty" yaml:"BillingCountry,omitempty"`
}

type BaseCreatePaymentCreditCardRequest struct {
	*BaseCreatePaymentRequest
}

type RefundPriorCreditCardChargeRequest struct {
	*BaseCreatePaymentCreditCardRequest

	ReturnPaymentID int32 `xml:"ReturnPaymentID,omitempty" json:"ReturnPaymentID,omitempty" yaml:"ReturnPaymentID,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	MaxAmount float64 `xml:"MaxAmount,omitempty" json:"MaxAmount,omitempty" yaml:"MaxAmount,omitempty"`
}

type CreatePaymentCreditCardRequest struct {
	*BaseCreatePaymentCreditCardRequest

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	PaymentDate time.Time `xml:"PaymentDate,omitempty" json:"PaymentDate,omitempty" yaml:"PaymentDate,omitempty"`

	Amount float64 `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`

	CreditCardNumber string `xml:"CreditCardNumber,omitempty" json:"CreditCardNumber,omitempty" yaml:"CreditCardNumber,omitempty"`

	ExpirationMonth int32 `xml:"ExpirationMonth,omitempty" json:"ExpirationMonth,omitempty" yaml:"ExpirationMonth,omitempty"`

	ExpirationYear int32 `xml:"ExpirationYear,omitempty" json:"ExpirationYear,omitempty" yaml:"ExpirationYear,omitempty"`

	BillingName string `xml:"BillingName,omitempty" json:"BillingName,omitempty" yaml:"BillingName,omitempty"`

	BillingAddress string `xml:"BillingAddress,omitempty" json:"BillingAddress,omitempty" yaml:"BillingAddress,omitempty"`

	BillingAddress2 string `xml:"BillingAddress2,omitempty" json:"BillingAddress2,omitempty" yaml:"BillingAddress2,omitempty"`

	BillingCity string `xml:"BillingCity,omitempty" json:"BillingCity,omitempty" yaml:"BillingCity,omitempty"`

	BillingState string `xml:"BillingState,omitempty" json:"BillingState,omitempty" yaml:"BillingState,omitempty"`

	BillingZip string `xml:"BillingZip,omitempty" json:"BillingZip,omitempty" yaml:"BillingZip,omitempty"`

	BillingCountry string `xml:"BillingCountry,omitempty" json:"BillingCountry,omitempty" yaml:"BillingCountry,omitempty"`

	CreditCardType int32 `xml:"CreditCardType,omitempty" json:"CreditCardType,omitempty" yaml:"CreditCardType,omitempty"`

	AuthorizationCode string `xml:"AuthorizationCode,omitempty" json:"AuthorizationCode,omitempty" yaml:"AuthorizationCode,omitempty"`

	Memo string `xml:"Memo,omitempty" json:"Memo,omitempty" yaml:"Memo,omitempty"`
}

type BaseChargeCreditCardRequest struct {
	*BaseCreatePaymentCreditCardRequest
}

type ChargeCreditCardTokenOnFileRequest struct {
	*BaseChargeCreditCardRequest

	CreditCardAccountType *AccountCreditCardType `xml:"CreditCardAccountType,omitempty" json:"CreditCardAccountType,omitempty" yaml:"CreditCardAccountType,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	CvcCode string `xml:"CvcCode,omitempty" json:"CvcCode,omitempty" yaml:"CvcCode,omitempty"`

	MaxAmount float64 `xml:"MaxAmount,omitempty" json:"MaxAmount,omitempty" yaml:"MaxAmount,omitempty"`

	MerchantWarehouseIDOverride int32 `xml:"MerchantWarehouseIDOverride,omitempty" json:"MerchantWarehouseIDOverride,omitempty" yaml:"MerchantWarehouseIDOverride,omitempty"`

	ClientIPAddress string `xml:"ClientIPAddress,omitempty" json:"ClientIPAddress,omitempty" yaml:"ClientIPAddress,omitempty"`

	OtherData1 string `xml:"OtherData1,omitempty" json:"OtherData1,omitempty" yaml:"OtherData1,omitempty"`

	OtherData2 string `xml:"OtherData2,omitempty" json:"OtherData2,omitempty" yaml:"OtherData2,omitempty"`

	OtherData3 string `xml:"OtherData3,omitempty" json:"OtherData3,omitempty" yaml:"OtherData3,omitempty"`

	OtherData4 string `xml:"OtherData4,omitempty" json:"OtherData4,omitempty" yaml:"OtherData4,omitempty"`

	OtherData5 string `xml:"OtherData5,omitempty" json:"OtherData5,omitempty" yaml:"OtherData5,omitempty"`

	OtherData6 string `xml:"OtherData6,omitempty" json:"OtherData6,omitempty" yaml:"OtherData6,omitempty"`

	OtherData7 string `xml:"OtherData7,omitempty" json:"OtherData7,omitempty" yaml:"OtherData7,omitempty"`

	OtherData8 string `xml:"OtherData8,omitempty" json:"OtherData8,omitempty" yaml:"OtherData8,omitempty"`

	OtherData9 string `xml:"OtherData9,omitempty" json:"OtherData9,omitempty" yaml:"OtherData9,omitempty"`

	OtherData10 string `xml:"OtherData10,omitempty" json:"OtherData10,omitempty" yaml:"OtherData10,omitempty"`
}

type ChargeCreditCardTokenRequest struct {
	*BaseChargeCreditCardRequest

	CreditCardToken string `xml:"CreditCardToken,omitempty" json:"CreditCardToken,omitempty" yaml:"CreditCardToken,omitempty"`

	BillingName string `xml:"BillingName,omitempty" json:"BillingName,omitempty" yaml:"BillingName,omitempty"`

	BillingAddress string `xml:"BillingAddress,omitempty" json:"BillingAddress,omitempty" yaml:"BillingAddress,omitempty"`

	BillingAddress2 string `xml:"BillingAddress2,omitempty" json:"BillingAddress2,omitempty" yaml:"BillingAddress2,omitempty"`

	BillingCity string `xml:"BillingCity,omitempty" json:"BillingCity,omitempty" yaml:"BillingCity,omitempty"`

	BillingState string `xml:"BillingState,omitempty" json:"BillingState,omitempty" yaml:"BillingState,omitempty"`

	BillingZip string `xml:"BillingZip,omitempty" json:"BillingZip,omitempty" yaml:"BillingZip,omitempty"`

	BillingCountry string `xml:"BillingCountry,omitempty" json:"BillingCountry,omitempty" yaml:"BillingCountry,omitempty"`

	CvcCode string `xml:"CvcCode,omitempty" json:"CvcCode,omitempty" yaml:"CvcCode,omitempty"`

	IssueNumber string `xml:"IssueNumber,omitempty" json:"IssueNumber,omitempty" yaml:"IssueNumber,omitempty"`

	CreditCardType int32 `xml:"CreditCardType,omitempty" json:"CreditCardType,omitempty" yaml:"CreditCardType,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	MaxAmount float64 `xml:"MaxAmount,omitempty" json:"MaxAmount,omitempty" yaml:"MaxAmount,omitempty"`

	MerchantWarehouseIDOverride int32 `xml:"MerchantWarehouseIDOverride,omitempty" json:"MerchantWarehouseIDOverride,omitempty" yaml:"MerchantWarehouseIDOverride,omitempty"`

	ClientIPAddress string `xml:"ClientIPAddress,omitempty" json:"ClientIPAddress,omitempty" yaml:"ClientIPAddress,omitempty"`

	OtherData1 string `xml:"OtherData1,omitempty" json:"OtherData1,omitempty" yaml:"OtherData1,omitempty"`

	OtherData2 string `xml:"OtherData2,omitempty" json:"OtherData2,omitempty" yaml:"OtherData2,omitempty"`

	OtherData3 string `xml:"OtherData3,omitempty" json:"OtherData3,omitempty" yaml:"OtherData3,omitempty"`

	OtherData4 string `xml:"OtherData4,omitempty" json:"OtherData4,omitempty" yaml:"OtherData4,omitempty"`

	OtherData5 string `xml:"OtherData5,omitempty" json:"OtherData5,omitempty" yaml:"OtherData5,omitempty"`

	OtherData6 string `xml:"OtherData6,omitempty" json:"OtherData6,omitempty" yaml:"OtherData6,omitempty"`

	OtherData7 string `xml:"OtherData7,omitempty" json:"OtherData7,omitempty" yaml:"OtherData7,omitempty"`

	OtherData8 string `xml:"OtherData8,omitempty" json:"OtherData8,omitempty" yaml:"OtherData8,omitempty"`

	OtherData9 string `xml:"OtherData9,omitempty" json:"OtherData9,omitempty" yaml:"OtherData9,omitempty"`

	OtherData10 string `xml:"OtherData10,omitempty" json:"OtherData10,omitempty" yaml:"OtherData10,omitempty"`
}

type ChargeCreditCardOnFileRequest struct {
	*BaseChargeCreditCardRequest

	CreditCardAccountType *AccountCreditCardType `xml:"CreditCardAccountType,omitempty" json:"CreditCardAccountType,omitempty" yaml:"CreditCardAccountType,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	MaxAmount float64 `xml:"MaxAmount,omitempty" json:"MaxAmount,omitempty" yaml:"MaxAmount,omitempty"`
}

type ChargeCreditCardRequest struct {
	*BaseChargeCreditCardRequest

	CreditCardNumber string `xml:"CreditCardNumber,omitempty" json:"CreditCardNumber,omitempty" yaml:"CreditCardNumber,omitempty"`

	ExpirationMonth int32 `xml:"ExpirationMonth,omitempty" json:"ExpirationMonth,omitempty" yaml:"ExpirationMonth,omitempty"`

	ExpirationYear int32 `xml:"ExpirationYear,omitempty" json:"ExpirationYear,omitempty" yaml:"ExpirationYear,omitempty"`

	BillingName string `xml:"BillingName,omitempty" json:"BillingName,omitempty" yaml:"BillingName,omitempty"`

	BillingAddress string `xml:"BillingAddress,omitempty" json:"BillingAddress,omitempty" yaml:"BillingAddress,omitempty"`

	BillingAddress2 string `xml:"BillingAddress2,omitempty" json:"BillingAddress2,omitempty" yaml:"BillingAddress2,omitempty"`

	BillingCity string `xml:"BillingCity,omitempty" json:"BillingCity,omitempty" yaml:"BillingCity,omitempty"`

	BillingState string `xml:"BillingState,omitempty" json:"BillingState,omitempty" yaml:"BillingState,omitempty"`

	BillingZip string `xml:"BillingZip,omitempty" json:"BillingZip,omitempty" yaml:"BillingZip,omitempty"`

	BillingCountry string `xml:"BillingCountry,omitempty" json:"BillingCountry,omitempty" yaml:"BillingCountry,omitempty"`

	CvcCode string `xml:"CvcCode,omitempty" json:"CvcCode,omitempty" yaml:"CvcCode,omitempty"`

	IssueNumber string `xml:"IssueNumber,omitempty" json:"IssueNumber,omitempty" yaml:"IssueNumber,omitempty"`

	CreditCardType int32 `xml:"CreditCardType,omitempty" json:"CreditCardType,omitempty" yaml:"CreditCardType,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	MaxAmount float64 `xml:"MaxAmount,omitempty" json:"MaxAmount,omitempty" yaml:"MaxAmount,omitempty"`
}

type CreatePaymentWalletRequest struct {
	*BaseCreatePaymentRequest

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	PaymentDate time.Time `xml:"PaymentDate,omitempty" json:"PaymentDate,omitempty" yaml:"PaymentDate,omitempty"`

	Amount float64 `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`

	WalletType int32 `xml:"WalletType,omitempty" json:"WalletType,omitempty" yaml:"WalletType,omitempty"`

	WalletAccount string `xml:"WalletAccount,omitempty" json:"WalletAccount,omitempty" yaml:"WalletAccount,omitempty"`

	AuthorizationCode string `xml:"AuthorizationCode,omitempty" json:"AuthorizationCode,omitempty" yaml:"AuthorizationCode,omitempty"`

	Memo string `xml:"Memo,omitempty" json:"Memo,omitempty" yaml:"Memo,omitempty"`

	BillingName string `xml:"BillingName,omitempty" json:"BillingName,omitempty" yaml:"BillingName,omitempty"`
}

type DeleteCustomerLeadRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	CustomerLeadID int32 `xml:"CustomerLeadID,omitempty" json:"CustomerLeadID,omitempty" yaml:"CustomerLeadID,omitempty"`
}

type ArrayOfEmailAttachment struct {
	EmailAttachment []*EmailAttachment `xml:"EmailAttachment,omitempty" json:"EmailAttachment,omitempty" yaml:"EmailAttachment,omitempty"`
}

type EmailAttachment struct {
	BinaryData []byte `xml:"BinaryData,omitempty" json:"BinaryData,omitempty" yaml:"BinaryData,omitempty"`

	FileName string `xml:"FileName,omitempty" json:"FileName,omitempty" yaml:"FileName,omitempty"`

	ContentLength int32 `xml:"ContentLength,omitempty" json:"ContentLength,omitempty" yaml:"ContentLength,omitempty"`
}

type ArrayOfForwardedAttachment struct {
	ForwardedAttachment []*ForwardedAttachment `xml:"ForwardedAttachment,omitempty" json:"ForwardedAttachment,omitempty" yaml:"ForwardedAttachment,omitempty"`
}

type ForwardedAttachment struct {
	MailID int32 `xml:"MailID,omitempty" json:"MailID,omitempty" yaml:"MailID,omitempty"`

	AttachmentID int32 `xml:"AttachmentID,omitempty" json:"AttachmentID,omitempty" yaml:"AttachmentID,omitempty"`
}

type CreateEmailResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateEmailResult"`

	*ApiResponse
}

type ApiResponse struct {
	Result *ApiResult `xml:"Result,omitempty" json:"Result,omitempty" yaml:"Result,omitempty"`
}

type ApiResult struct {
	Status *ResultStatus `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`

	Errors *ArrayOfString `xml:"Errors,omitempty" json:"Errors,omitempty" yaml:"Errors,omitempty"`

	TransactionKey string `xml:"TransactionKey,omitempty" json:"TransactionKey,omitempty" yaml:"TransactionKey,omitempty"`
}

type ArrayOfString struct {
	String []string `xml:"string,omitempty" json:"string,omitempty" yaml:"string,omitempty"`
}

type UpdateOrderDetailResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateOrderDetailResult"`

	*ApiResponse
}

type CreateOrderDetailResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateOrderDetailResult"`

	*ApiResponse

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	OrderLine int32 `xml:"OrderLine,omitempty" json:"OrderLine,omitempty" yaml:"OrderLine,omitempty"`
}

type SendEmailResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SendEmailResult"`

	*ApiResponse
}

type SetItemKitMembersResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetItemKitMembersResult"`

	*ApiResponse
}

type GetFileContentsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetFileContentsResult"`

	*ApiResponse

	File []byte `xml:"File,omitempty" json:"File,omitempty" yaml:"File,omitempty"`
}

type GetFilesResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetFilesResult"`

	*ApiResponse

	CustomerFileList *ArrayOfCustomerFilesResponse `xml:"CustomerFileList,omitempty" json:"CustomerFileList,omitempty" yaml:"CustomerFileList,omitempty"`
}

type ArrayOfCustomerFilesResponse struct {
	CustomerFilesResponse []*CustomerFilesResponse `xml:"CustomerFilesResponse,omitempty" json:"CustomerFilesResponse,omitempty" yaml:"CustomerFilesResponse,omitempty"`
}

type CustomerFilesResponse struct {
	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	FileName string `xml:"FileName,omitempty" json:"FileName,omitempty" yaml:"FileName,omitempty"`
}

type ChargeGroupOrderCreditCardTokenResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ChargeGroupOrderCreditCardTokenResult"`

	*ApiResponse

	_paymentIDs *ArrayOfPaymentsResponse `xml:"_paymentIDs,omitempty" json:"_paymentIDs,omitempty" yaml:"_paymentIDs,omitempty"`

	Amount float64 `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`

	AuthorizationCode string `xml:"AuthorizationCode,omitempty" json:"AuthorizationCode,omitempty" yaml:"AuthorizationCode,omitempty"`

	Payments *ArrayOfPaymentsResponse `xml:"Payments,omitempty" json:"Payments,omitempty" yaml:"Payments,omitempty"`
}

type ArrayOfPaymentsResponse struct {
	PaymentsResponse []*PaymentsResponse `xml:"PaymentsResponse,omitempty" json:"PaymentsResponse,omitempty" yaml:"PaymentsResponse,omitempty"`
}

type PaymentsResponse struct {
	_PaymentID int32 `xml:"_PaymentID,omitempty" json:"_PaymentID,omitempty" yaml:"_PaymentID,omitempty"`

	_OrderID int32 `xml:"_OrderID,omitempty" json:"_OrderID,omitempty" yaml:"_OrderID,omitempty"`
}

type DeleteCustomerExtendedResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteCustomerExtendedResult"`

	*ApiResponse

	CustomerExtendedID int32 `xml:"CustomerExtendedID,omitempty" json:"CustomerExtendedID,omitempty" yaml:"CustomerExtendedID,omitempty"`
}

type UpdateCustomerExtendedResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateCustomerExtendedResult"`

	*ApiResponse
}

type CreateCustomerExtendedResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateCustomerExtendedResult"`

	*ApiResponse

	CustomerExtendedID int32 `xml:"CustomerExtendedID,omitempty" json:"CustomerExtendedID,omitempty" yaml:"CustomerExtendedID,omitempty"`
}

type CreateCustomerLeadResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateCustomerLeadResult"`

	*ApiResponse

	CustomerLeadID int32 `xml:"CustomerLeadID,omitempty" json:"CustomerLeadID,omitempty" yaml:"CustomerLeadID,omitempty"`
}

type CreatePartyResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreatePartyResult"`

	*ApiResponse

	PartyID int32 `xml:"PartyID,omitempty" json:"PartyID,omitempty" yaml:"PartyID,omitempty"`
}

type CreateBillResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateBillResult"`

	*ApiResponse

	BillID int32 `xml:"BillID,omitempty" json:"BillID,omitempty" yaml:"BillID,omitempty"`
}

type CreatePayoutResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreatePayoutResult"`

	*ApiResponse

	PayoutID int32 `xml:"PayoutID,omitempty" json:"PayoutID,omitempty" yaml:"PayoutID,omitempty"`

	TotalDollarAmount float64 `xml:"TotalDollarAmount,omitempty" json:"TotalDollarAmount,omitempty" yaml:"TotalDollarAmount,omitempty"`
}

type CreateCustomerInquiryResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateCustomerInquiryResult"`

	*ApiResponse

	NewCustomerHistoryID int32 `xml:"NewCustomerHistoryID,omitempty" json:"NewCustomerHistoryID,omitempty" yaml:"NewCustomerHistoryID,omitempty"`
}

type CreateCustomerFileResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateCustomerFileResult"`

	*ApiResponse

	FolderID int32 `xml:"FolderID,omitempty" json:"FolderID,omitempty" yaml:"FolderID,omitempty"`

	FileID int32 `xml:"FileID,omitempty" json:"FileID,omitempty" yaml:"FileID,omitempty"`
}

type SetItemCountryRegionResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetItemCountryRegionResult"`

	*ApiResponse
}

type GetItemCountryRegionResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetItemCountryRegionResult"`

	*ApiResponse

	ItemCountryRegions *ArrayOfItemCountryRegionResponse `xml:"ItemCountryRegions,omitempty" json:"ItemCountryRegions,omitempty" yaml:"ItemCountryRegions,omitempty"`
}

type ArrayOfItemCountryRegionResponse struct {
	ItemCountryRegionResponse []*ItemCountryRegionResponse `xml:"ItemCountryRegionResponse,omitempty" json:"ItemCountryRegionResponse,omitempty" yaml:"ItemCountryRegionResponse,omitempty"`
}

type ItemCountryRegionResponse struct {
	ItemCode string `xml:"ItemCode,omitempty" json:"ItemCode,omitempty" yaml:"ItemCode,omitempty"`

	CountryCode string `xml:"CountryCode,omitempty" json:"CountryCode,omitempty" yaml:"CountryCode,omitempty"`

	RegionCode string `xml:"RegionCode,omitempty" json:"RegionCode,omitempty" yaml:"RegionCode,omitempty"`

	Taxed bool `xml:"Taxed,omitempty" json:"Taxed,omitempty" yaml:"Taxed,omitempty"`

	TaxedFed bool `xml:"TaxedFed,omitempty" json:"TaxedFed,omitempty" yaml:"TaxedFed,omitempty"`

	TaxedState bool `xml:"TaxedState,omitempty" json:"TaxedState,omitempty" yaml:"TaxedState,omitempty"`

	UseTaxOverride bool `xml:"UseTaxOverride,omitempty" json:"UseTaxOverride,omitempty" yaml:"UseTaxOverride,omitempty"`

	TaxOverridePct float64 `xml:"TaxOverridePct,omitempty" json:"TaxOverridePct,omitempty" yaml:"TaxOverridePct,omitempty"`
}

type CreateItemResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateItemResult"`

	*ApiResponse

	ItemCode string `xml:"ItemCode,omitempty" json:"ItemCode,omitempty" yaml:"ItemCode,omitempty"`
}

type GetCustomerWallResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomerWallResult"`

	*ApiResponse

	CustomerWallItems *ArrayOfCustomerWallItemResponse `xml:"CustomerWallItems,omitempty" json:"CustomerWallItems,omitempty" yaml:"CustomerWallItems,omitempty"`
}

type ArrayOfCustomerWallItemResponse struct {
	CustomerWallItemResponse []*CustomerWallItemResponse `xml:"CustomerWallItemResponse,omitempty" json:"CustomerWallItemResponse,omitempty" yaml:"CustomerWallItemResponse,omitempty"`
}

type CustomerWallItemResponse struct {
	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	WallItemID int32 `xml:"WallItemID,omitempty" json:"WallItemID,omitempty" yaml:"WallItemID,omitempty"`

	Text string `xml:"Text,omitempty" json:"Text,omitempty" yaml:"Text,omitempty"`

	EntryDate time.Time `xml:"EntryDate,omitempty" json:"EntryDate,omitempty" yaml:"EntryDate,omitempty"`

	Field1 string `xml:"Field1,omitempty" json:"Field1,omitempty" yaml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty" json:"Field2,omitempty" yaml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty" json:"Field3,omitempty" yaml:"Field3,omitempty"`
}

type DeleteCustomerWallItemResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteCustomerWallItemResult"`

	*ApiResponse

	CountOfDeletedRows int32 `xml:"CountOfDeletedRows,omitempty" json:"CountOfDeletedRows,omitempty" yaml:"CountOfDeletedRows,omitempty"`
}

type CreateCustomerWallItemResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateCustomerWallItemResult"`

	*ApiResponse

	WallItemID int32 `xml:"WallItemID,omitempty" json:"WallItemID,omitempty" yaml:"WallItemID,omitempty"`
}

type GetCustomerLeadsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomerLeadsResult"`

	*ApiResponse

	CustomerLeads *ArrayOfCustomerLeadsResponse `xml:"CustomerLeads,omitempty" json:"CustomerLeads,omitempty" yaml:"CustomerLeads,omitempty"`
}

type ArrayOfCustomerLeadsResponse struct {
	CustomerLeadsResponse []*CustomerLeadsResponse `xml:"CustomerLeadsResponse,omitempty" json:"CustomerLeadsResponse,omitempty" yaml:"CustomerLeadsResponse,omitempty"`
}

type CustomerLeadsResponse struct {
	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	CustomerLeadID int32 `xml:"CustomerLeadID,omitempty" json:"CustomerLeadID,omitempty" yaml:"CustomerLeadID,omitempty"`

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty" json:"MiddleName,omitempty" yaml:"MiddleName,omitempty"`

	Company string `xml:"Company,omitempty" json:"Company,omitempty" yaml:"Company,omitempty"`

	Phone string `xml:"Phone,omitempty" json:"Phone,omitempty" yaml:"Phone,omitempty"`

	Phone2 string `xml:"Phone2,omitempty" json:"Phone2,omitempty" yaml:"Phone2,omitempty"`

	MobilePhone string `xml:"MobilePhone,omitempty" json:"MobilePhone,omitempty" yaml:"MobilePhone,omitempty"`

	Fax string `xml:"Fax,omitempty" json:"Fax,omitempty" yaml:"Fax,omitempty"`

	Email string `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`

	Address1 string `xml:"Address1,omitempty" json:"Address1,omitempty" yaml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty" json:"Address2,omitempty" yaml:"Address2,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty" json:"Zip,omitempty" yaml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`

	BirthDate time.Time `xml:"BirthDate,omitempty" json:"BirthDate,omitempty" yaml:"BirthDate,omitempty"`

	Notes string `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`
}

type DeleteCustomerLeadResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteCustomerLeadResult"`

	*ApiResponse
}

type DeleteOrderDetailResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteOrderDetailResult"`

	*ApiResponse
}

type UpdatePartyResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdatePartyResult"`

	*ApiResponse
}

type GetCustomerLeadSocialNetworksResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomerLeadSocialNetworksResult"`

	*ApiResponse

	CustomerLeadSocialNetwork *ArrayOfCustomerLeadSocialNetworksResponse `xml:"CustomerLeadSocialNetwork,omitempty" json:"CustomerLeadSocialNetwork,omitempty" yaml:"CustomerLeadSocialNetwork,omitempty"`
}

type ArrayOfCustomerLeadSocialNetworksResponse struct {
	CustomerLeadSocialNetworksResponse []*CustomerLeadSocialNetworksResponse `xml:"CustomerLeadSocialNetworksResponse,omitempty" json:"CustomerLeadSocialNetworksResponse,omitempty" yaml:"CustomerLeadSocialNetworksResponse,omitempty"`
}

type CustomerLeadSocialNetworksResponse struct {
	SocialNetworkID int32 `xml:"SocialNetworkID,omitempty" json:"SocialNetworkID,omitempty" yaml:"SocialNetworkID,omitempty"`

	SocialNetworkDescription string `xml:"SocialNetworkDescription,omitempty" json:"SocialNetworkDescription,omitempty" yaml:"SocialNetworkDescription,omitempty"`

	Url string `xml:"Url,omitempty" json:"Url,omitempty" yaml:"Url,omitempty"`
}

type SetCustomerSocialNetworksResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetCustomerSocialNetworksResult"`

	*ApiResponse
}

type GetCustomerSocialNetworksResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomerSocialNetworksResult"`

	*ApiResponse

	CustomerSocialNetwork *ArrayOfCustomerSocialNetworksResponse `xml:"CustomerSocialNetwork,omitempty" json:"CustomerSocialNetwork,omitempty" yaml:"CustomerSocialNetwork,omitempty"`
}

type ArrayOfCustomerSocialNetworksResponse struct {
	CustomerSocialNetworksResponse []*CustomerSocialNetworksResponse `xml:"CustomerSocialNetworksResponse,omitempty" json:"CustomerSocialNetworksResponse,omitempty" yaml:"CustomerSocialNetworksResponse,omitempty"`
}

type CustomerSocialNetworksResponse struct {
	SocialNetworkID int32 `xml:"SocialNetworkID,omitempty" json:"SocialNetworkID,omitempty" yaml:"SocialNetworkID,omitempty"`

	SocialNetworkDescription string `xml:"SocialNetworkDescription,omitempty" json:"SocialNetworkDescription,omitempty" yaml:"SocialNetworkDescription,omitempty"`

	Url string `xml:"Url,omitempty" json:"Url,omitempty" yaml:"Url,omitempty"`
}

type SetCustomerSiteResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetCustomerSiteResult"`

	*ApiResponse
}

type UpdateCustomerResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateCustomerResult"`

	*ApiResponse
}

type UpdateOrderResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateOrderResult"`

	*ApiResponse
}

type BaseCalculateOrderResponse struct {
	*ApiResponse
}

type CalculateOrderResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CalculateOrderResult"`

	*BaseCalculateOrderResponse

	Total float64 `xml:"Total,omitempty" json:"Total,omitempty" yaml:"Total,omitempty"`

	SubTotal float64 `xml:"SubTotal,omitempty" json:"SubTotal,omitempty" yaml:"SubTotal,omitempty"`

	TaxTotal float64 `xml:"TaxTotal,omitempty" json:"TaxTotal,omitempty" yaml:"TaxTotal,omitempty"`

	ShippingTotal float64 `xml:"ShippingTotal,omitempty" json:"ShippingTotal,omitempty" yaml:"ShippingTotal,omitempty"`

	DiscountTotal float64 `xml:"DiscountTotal,omitempty" json:"DiscountTotal,omitempty" yaml:"DiscountTotal,omitempty"`

	DiscountPercent float64 `xml:"DiscountPercent,omitempty" json:"DiscountPercent,omitempty" yaml:"DiscountPercent,omitempty"`

	WeightTotal float64 `xml:"WeightTotal,omitempty" json:"WeightTotal,omitempty" yaml:"WeightTotal,omitempty"`

	BusinessVolumeTotal float64 `xml:"BusinessVolumeTotal,omitempty" json:"BusinessVolumeTotal,omitempty" yaml:"BusinessVolumeTotal,omitempty"`

	CommissionableVolumeTotal float64 `xml:"CommissionableVolumeTotal,omitempty" json:"CommissionableVolumeTotal,omitempty" yaml:"CommissionableVolumeTotal,omitempty"`

	Other1Total float64 `xml:"Other1Total,omitempty" json:"Other1Total,omitempty" yaml:"Other1Total,omitempty"`

	Other2Total float64 `xml:"Other2Total,omitempty" json:"Other2Total,omitempty" yaml:"Other2Total,omitempty"`

	Other3Total float64 `xml:"Other3Total,omitempty" json:"Other3Total,omitempty" yaml:"Other3Total,omitempty"`

	Other4Total float64 `xml:"Other4Total,omitempty" json:"Other4Total,omitempty" yaml:"Other4Total,omitempty"`

	Other5Total float64 `xml:"Other5Total,omitempty" json:"Other5Total,omitempty" yaml:"Other5Total,omitempty"`

	Other6Total float64 `xml:"Other6Total,omitempty" json:"Other6Total,omitempty" yaml:"Other6Total,omitempty"`

	Other7Total float64 `xml:"Other7Total,omitempty" json:"Other7Total,omitempty" yaml:"Other7Total,omitempty"`

	Other8Total float64 `xml:"Other8Total,omitempty" json:"Other8Total,omitempty" yaml:"Other8Total,omitempty"`

	Other9Total float64 `xml:"Other9Total,omitempty" json:"Other9Total,omitempty" yaml:"Other9Total,omitempty"`

	Other10Total float64 `xml:"Other10Total,omitempty" json:"Other10Total,omitempty" yaml:"Other10Total,omitempty"`

	Other11 string `xml:"Other11,omitempty" json:"Other11,omitempty" yaml:"Other11,omitempty"`

	Other12 string `xml:"Other12,omitempty" json:"Other12,omitempty" yaml:"Other12,omitempty"`

	Other13 string `xml:"Other13,omitempty" json:"Other13,omitempty" yaml:"Other13,omitempty"`

	Other14 string `xml:"Other14,omitempty" json:"Other14,omitempty" yaml:"Other14,omitempty"`

	Other15 string `xml:"Other15,omitempty" json:"Other15,omitempty" yaml:"Other15,omitempty"`

	Other16 string `xml:"Other16,omitempty" json:"Other16,omitempty" yaml:"Other16,omitempty"`

	Other17 string `xml:"Other17,omitempty" json:"Other17,omitempty" yaml:"Other17,omitempty"`

	Other18 string `xml:"Other18,omitempty" json:"Other18,omitempty" yaml:"Other18,omitempty"`

	Other19 string `xml:"Other19,omitempty" json:"Other19,omitempty" yaml:"Other19,omitempty"`

	Other20 string `xml:"Other20,omitempty" json:"Other20,omitempty" yaml:"Other20,omitempty"`

	ShippingTax float64 `xml:"ShippingTax,omitempty" json:"ShippingTax,omitempty" yaml:"ShippingTax,omitempty"`

	OrderTax float64 `xml:"OrderTax,omitempty" json:"OrderTax,omitempty" yaml:"OrderTax,omitempty"`

	FedTaxTotal float64 `xml:"FedTaxTotal,omitempty" json:"FedTaxTotal,omitempty" yaml:"FedTaxTotal,omitempty"`

	StateTaxTotal float64 `xml:"StateTaxTotal,omitempty" json:"StateTaxTotal,omitempty" yaml:"StateTaxTotal,omitempty"`

	Details *ArrayOfOrderDetailResponse `xml:"Details,omitempty" json:"Details,omitempty" yaml:"Details,omitempty"`

	ShipMethods *ArrayOfShipMethodResponse `xml:"ShipMethods,omitempty" json:"ShipMethods,omitempty" yaml:"ShipMethods,omitempty"`

	Warnings *ArrayOfString `xml:"Warnings,omitempty" json:"Warnings,omitempty" yaml:"Warnings,omitempty"`
}

type ArrayOfOrderDetailResponse struct {
	OrderDetailResponse []*OrderDetailResponse `xml:"OrderDetailResponse,omitempty" json:"OrderDetailResponse,omitempty" yaml:"OrderDetailResponse,omitempty"`
}

type OrderDetailResponse struct {
	ItemCode string `xml:"ItemCode,omitempty" json:"ItemCode,omitempty" yaml:"ItemCode,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	Quantity float64 `xml:"Quantity,omitempty" json:"Quantity,omitempty" yaml:"Quantity,omitempty"`

	PriceEach float64 `xml:"PriceEach,omitempty" json:"PriceEach,omitempty" yaml:"PriceEach,omitempty"`

	PriceTotal float64 `xml:"PriceTotal,omitempty" json:"PriceTotal,omitempty" yaml:"PriceTotal,omitempty"`

	Tax float64 `xml:"Tax,omitempty" json:"Tax,omitempty" yaml:"Tax,omitempty"`

	WeightEach float64 `xml:"WeightEach,omitempty" json:"WeightEach,omitempty" yaml:"WeightEach,omitempty"`

	Weight float64 `xml:"Weight,omitempty" json:"Weight,omitempty" yaml:"Weight,omitempty"`

	BusinessVolumeEach float64 `xml:"BusinessVolumeEach,omitempty" json:"BusinessVolumeEach,omitempty" yaml:"BusinessVolumeEach,omitempty"`

	BusinesVolume float64 `xml:"BusinesVolume,omitempty" json:"BusinesVolume,omitempty" yaml:"BusinesVolume,omitempty"`

	CommissionableVolumeEach float64 `xml:"CommissionableVolumeEach,omitempty" json:"CommissionableVolumeEach,omitempty" yaml:"CommissionableVolumeEach,omitempty"`

	CommissionableVolume float64 `xml:"CommissionableVolume,omitempty" json:"CommissionableVolume,omitempty" yaml:"CommissionableVolume,omitempty"`

	Other1Each float64 `xml:"Other1Each,omitempty" json:"Other1Each,omitempty" yaml:"Other1Each,omitempty"`

	Other1 float64 `xml:"Other1,omitempty" json:"Other1,omitempty" yaml:"Other1,omitempty"`

	Other2Each float64 `xml:"Other2Each,omitempty" json:"Other2Each,omitempty" yaml:"Other2Each,omitempty"`

	Other2 float64 `xml:"Other2,omitempty" json:"Other2,omitempty" yaml:"Other2,omitempty"`

	Other3Each float64 `xml:"Other3Each,omitempty" json:"Other3Each,omitempty" yaml:"Other3Each,omitempty"`

	Other3 float64 `xml:"Other3,omitempty" json:"Other3,omitempty" yaml:"Other3,omitempty"`

	Other4Each float64 `xml:"Other4Each,omitempty" json:"Other4Each,omitempty" yaml:"Other4Each,omitempty"`

	Other4 float64 `xml:"Other4,omitempty" json:"Other4,omitempty" yaml:"Other4,omitempty"`

	Other5Each float64 `xml:"Other5Each,omitempty" json:"Other5Each,omitempty" yaml:"Other5Each,omitempty"`

	Other5 float64 `xml:"Other5,omitempty" json:"Other5,omitempty" yaml:"Other5,omitempty"`

	Other6Each float64 `xml:"Other6Each,omitempty" json:"Other6Each,omitempty" yaml:"Other6Each,omitempty"`

	Other6 float64 `xml:"Other6,omitempty" json:"Other6,omitempty" yaml:"Other6,omitempty"`

	Other7Each float64 `xml:"Other7Each,omitempty" json:"Other7Each,omitempty" yaml:"Other7Each,omitempty"`

	Other7 float64 `xml:"Other7,omitempty" json:"Other7,omitempty" yaml:"Other7,omitempty"`

	Other8Each float64 `xml:"Other8Each,omitempty" json:"Other8Each,omitempty" yaml:"Other8Each,omitempty"`

	Other8 float64 `xml:"Other8,omitempty" json:"Other8,omitempty" yaml:"Other8,omitempty"`

	Other9Each float64 `xml:"Other9Each,omitempty" json:"Other9Each,omitempty" yaml:"Other9Each,omitempty"`

	Other9 float64 `xml:"Other9,omitempty" json:"Other9,omitempty" yaml:"Other9,omitempty"`

	Other10Each float64 `xml:"Other10Each,omitempty" json:"Other10Each,omitempty" yaml:"Other10Each,omitempty"`

	Other10 float64 `xml:"Other10,omitempty" json:"Other10,omitempty" yaml:"Other10,omitempty"`

	ParentItemCode string `xml:"ParentItemCode,omitempty" json:"ParentItemCode,omitempty" yaml:"ParentItemCode,omitempty"`

	Taxable float64 `xml:"Taxable,omitempty" json:"Taxable,omitempty" yaml:"Taxable,omitempty"`

	FedTax float64 `xml:"FedTax,omitempty" json:"FedTax,omitempty" yaml:"FedTax,omitempty"`

	StateTax float64 `xml:"StateTax,omitempty" json:"StateTax,omitempty" yaml:"StateTax,omitempty"`

	CityTax float64 `xml:"CityTax,omitempty" json:"CityTax,omitempty" yaml:"CityTax,omitempty"`

	CityLocalTax float64 `xml:"CityLocalTax,omitempty" json:"CityLocalTax,omitempty" yaml:"CityLocalTax,omitempty"`

	CountyTax float64 `xml:"CountyTax,omitempty" json:"CountyTax,omitempty" yaml:"CountyTax,omitempty"`

	CountyLocalTax float64 `xml:"CountyLocalTax,omitempty" json:"CountyLocalTax,omitempty" yaml:"CountyLocalTax,omitempty"`

	ManualTax float64 `xml:"ManualTax,omitempty" json:"ManualTax,omitempty" yaml:"ManualTax,omitempty"`

	IsStateTaxOverride bool `xml:"IsStateTaxOverride,omitempty" json:"IsStateTaxOverride,omitempty" yaml:"IsStateTaxOverride,omitempty"`

	OrderLine int32 `xml:"OrderLine,omitempty" json:"OrderLine,omitempty" yaml:"OrderLine,omitempty"`

	Reference1 string `xml:"Reference1,omitempty" json:"Reference1,omitempty" yaml:"Reference1,omitempty"`

	ShippingPriceEach float64 `xml:"ShippingPriceEach,omitempty" json:"ShippingPriceEach,omitempty" yaml:"ShippingPriceEach,omitempty"`
}

type ArrayOfShipMethodResponse struct {
	ShipMethodResponse []*ShipMethodResponse `xml:"ShipMethodResponse,omitempty" json:"ShipMethodResponse,omitempty" yaml:"ShipMethodResponse,omitempty"`
}

type ShipMethodResponse struct {
	ShipMethodID int32 `xml:"ShipMethodID,omitempty" json:"ShipMethodID,omitempty" yaml:"ShipMethodID,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	ShippingAmount float64 `xml:"ShippingAmount,omitempty" json:"ShippingAmount,omitempty" yaml:"ShippingAmount,omitempty"`
}

type CreateAutoOrderResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateAutoOrderResult"`

	*BaseCalculateOrderResponse

	AutoOrderID int32 `xml:"AutoOrderID,omitempty" json:"AutoOrderID,omitempty" yaml:"AutoOrderID,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	Total float64 `xml:"Total,omitempty" json:"Total,omitempty" yaml:"Total,omitempty"`

	SubTotal float64 `xml:"SubTotal,omitempty" json:"SubTotal,omitempty" yaml:"SubTotal,omitempty"`

	TaxTotal float64 `xml:"TaxTotal,omitempty" json:"TaxTotal,omitempty" yaml:"TaxTotal,omitempty"`

	ShippingTotal float64 `xml:"ShippingTotal,omitempty" json:"ShippingTotal,omitempty" yaml:"ShippingTotal,omitempty"`

	DiscountTotal float64 `xml:"DiscountTotal,omitempty" json:"DiscountTotal,omitempty" yaml:"DiscountTotal,omitempty"`

	WeightTotal float64 `xml:"WeightTotal,omitempty" json:"WeightTotal,omitempty" yaml:"WeightTotal,omitempty"`

	BusinessVolumeTotal float64 `xml:"BusinessVolumeTotal,omitempty" json:"BusinessVolumeTotal,omitempty" yaml:"BusinessVolumeTotal,omitempty"`

	CommissionableVolumeTotal float64 `xml:"CommissionableVolumeTotal,omitempty" json:"CommissionableVolumeTotal,omitempty" yaml:"CommissionableVolumeTotal,omitempty"`

	Other1Total float64 `xml:"Other1Total,omitempty" json:"Other1Total,omitempty" yaml:"Other1Total,omitempty"`

	Other2Total float64 `xml:"Other2Total,omitempty" json:"Other2Total,omitempty" yaml:"Other2Total,omitempty"`

	Other3Total float64 `xml:"Other3Total,omitempty" json:"Other3Total,omitempty" yaml:"Other3Total,omitempty"`

	Other4Total float64 `xml:"Other4Total,omitempty" json:"Other4Total,omitempty" yaml:"Other4Total,omitempty"`

	Other5Total float64 `xml:"Other5Total,omitempty" json:"Other5Total,omitempty" yaml:"Other5Total,omitempty"`

	Other6Total float64 `xml:"Other6Total,omitempty" json:"Other6Total,omitempty" yaml:"Other6Total,omitempty"`

	Other7Total float64 `xml:"Other7Total,omitempty" json:"Other7Total,omitempty" yaml:"Other7Total,omitempty"`

	Other8Total float64 `xml:"Other8Total,omitempty" json:"Other8Total,omitempty" yaml:"Other8Total,omitempty"`

	Other9Total float64 `xml:"Other9Total,omitempty" json:"Other9Total,omitempty" yaml:"Other9Total,omitempty"`

	Other10Total float64 `xml:"Other10Total,omitempty" json:"Other10Total,omitempty" yaml:"Other10Total,omitempty"`

	ShippingTax float64 `xml:"ShippingTax,omitempty" json:"ShippingTax,omitempty" yaml:"ShippingTax,omitempty"`

	OrderTax float64 `xml:"OrderTax,omitempty" json:"OrderTax,omitempty" yaml:"OrderTax,omitempty"`

	Details *ArrayOfOrderDetailResponse `xml:"Details,omitempty" json:"Details,omitempty" yaml:"Details,omitempty"`
}

type CreateOrderResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateOrderResult"`

	*BaseCalculateOrderResponse

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	Total float64 `xml:"Total,omitempty" json:"Total,omitempty" yaml:"Total,omitempty"`

	SubTotal float64 `xml:"SubTotal,omitempty" json:"SubTotal,omitempty" yaml:"SubTotal,omitempty"`

	TaxTotal float64 `xml:"TaxTotal,omitempty" json:"TaxTotal,omitempty" yaml:"TaxTotal,omitempty"`

	ShippingTotal float64 `xml:"ShippingTotal,omitempty" json:"ShippingTotal,omitempty" yaml:"ShippingTotal,omitempty"`

	DiscountTotal float64 `xml:"DiscountTotal,omitempty" json:"DiscountTotal,omitempty" yaml:"DiscountTotal,omitempty"`

	WeightTotal float64 `xml:"WeightTotal,omitempty" json:"WeightTotal,omitempty" yaml:"WeightTotal,omitempty"`

	BusinessVolumeTotal float64 `xml:"BusinessVolumeTotal,omitempty" json:"BusinessVolumeTotal,omitempty" yaml:"BusinessVolumeTotal,omitempty"`

	CommissionableVolumeTotal float64 `xml:"CommissionableVolumeTotal,omitempty" json:"CommissionableVolumeTotal,omitempty" yaml:"CommissionableVolumeTotal,omitempty"`

	Other1Total float64 `xml:"Other1Total,omitempty" json:"Other1Total,omitempty" yaml:"Other1Total,omitempty"`

	Other2Total float64 `xml:"Other2Total,omitempty" json:"Other2Total,omitempty" yaml:"Other2Total,omitempty"`

	Other3Total float64 `xml:"Other3Total,omitempty" json:"Other3Total,omitempty" yaml:"Other3Total,omitempty"`

	Other4Total float64 `xml:"Other4Total,omitempty" json:"Other4Total,omitempty" yaml:"Other4Total,omitempty"`

	Other5Total float64 `xml:"Other5Total,omitempty" json:"Other5Total,omitempty" yaml:"Other5Total,omitempty"`

	Other6Total float64 `xml:"Other6Total,omitempty" json:"Other6Total,omitempty" yaml:"Other6Total,omitempty"`

	Other7Total float64 `xml:"Other7Total,omitempty" json:"Other7Total,omitempty" yaml:"Other7Total,omitempty"`

	Other8Total float64 `xml:"Other8Total,omitempty" json:"Other8Total,omitempty" yaml:"Other8Total,omitempty"`

	Other9Total float64 `xml:"Other9Total,omitempty" json:"Other9Total,omitempty" yaml:"Other9Total,omitempty"`

	Other10Total float64 `xml:"Other10Total,omitempty" json:"Other10Total,omitempty" yaml:"Other10Total,omitempty"`

	ShippingTax float64 `xml:"ShippingTax,omitempty" json:"ShippingTax,omitempty" yaml:"ShippingTax,omitempty"`

	OrderTax float64 `xml:"OrderTax,omitempty" json:"OrderTax,omitempty" yaml:"OrderTax,omitempty"`

	Details *ArrayOfOrderDetailResponse `xml:"Details,omitempty" json:"Details,omitempty" yaml:"Details,omitempty"`

	Warnings *ArrayOfString `xml:"Warnings,omitempty" json:"Warnings,omitempty" yaml:"Warnings,omitempty"`
}

type CreateCustomerResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateCustomerResult"`

	*ApiResponse

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`
}

type SetAccountResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetAccountCreditCardResult"`

	*ApiResponse
}

type CreateOrderImportResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateOrderImportResult"`

	*ApiResponse

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`
}

type TransactionalResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ProcessTransactionResult"`

	*ApiResponse

	TransactionResponses *ArrayOfApiResponse `xml:"TransactionResponses,omitempty" json:"TransactionResponses,omitempty" yaml:"TransactionResponses,omitempty"`
}

type ArrayOfApiResponse struct {
	ApiResponse []*ApiResponse `xml:"ApiResponse,omitempty" json:"ApiResponse,omitempty" yaml:"ApiResponse,omitempty"`
}

type BaseCreateExpectedPaymentResponse struct {
	*ApiResponse
}

type CreateExpectedBankWireResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateExpectedBankWireResult"`

	*BaseCreateExpectedPaymentResponse

	ExpectedPaymentID int32 `xml:"ExpectedPaymentID,omitempty" json:"ExpectedPaymentID,omitempty" yaml:"ExpectedPaymentID,omitempty"`
}

type CreateExpectedCODResponse struct {
	*BaseCreateExpectedPaymentResponse

	ExpectedPaymentID int32 `xml:"ExpectedPaymentID,omitempty" json:"ExpectedPaymentID,omitempty" yaml:"ExpectedPaymentID,omitempty"`
}

type CreateExpectedPaymentResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateExpectedPaymentResult"`

	*BaseCreateExpectedPaymentResponse

	ExpectedPaymentID int32 `xml:"ExpectedPaymentID,omitempty" json:"ExpectedPaymentID,omitempty" yaml:"ExpectedPaymentID,omitempty"`
}

type AuthorizeOnlyCreditCardResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ AuthorizeOnlyCreditCardTokenResult"`

	*ApiResponse

	AuthorizationCode string `xml:"AuthorizationCode,omitempty" json:"AuthorizationCode,omitempty" yaml:"AuthorizationCode,omitempty"`

	MerchantTransactionKey string `xml:"MerchantTransactionKey,omitempty" json:"MerchantTransactionKey,omitempty" yaml:"MerchantTransactionKey,omitempty"`
}

type BaseCreatePaymentResponse struct {
	*ApiResponse
}

type CreatePaymentCreditCardResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreatePaymentCreditCardResult"`

	*BaseCreatePaymentResponse

	PaymentID int32 `xml:"PaymentID,omitempty" json:"PaymentID,omitempty" yaml:"PaymentID,omitempty"`
}

type ChargeCreditCardResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ChargeCreditCardResult"`

	*CreatePaymentCreditCardResponse

	Amount float64 `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`

	AuthorizationCode string `xml:"AuthorizationCode,omitempty" json:"AuthorizationCode,omitempty" yaml:"AuthorizationCode,omitempty"`
}

type RefundPriorCreditCardChargeResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ RefundPriorCreditCardChargeResult"`

	*CreatePaymentCreditCardResponse

	Amount float64 `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`

	AuthorizationCode string `xml:"AuthorizationCode,omitempty" json:"AuthorizationCode,omitempty" yaml:"AuthorizationCode,omitempty"`
}

type CreatePaymentResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreatePaymentResult"`

	*BaseCreatePaymentResponse

	PaymentID int32 `xml:"PaymentID,omitempty" json:"PaymentID,omitempty" yaml:"PaymentID,omitempty"`
}

type DebitBankAccountResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DebitBankAccountResult"`

	*CreatePaymentResponse

	Amount float64 `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`

	AuthorizationCode string `xml:"AuthorizationCode,omitempty" json:"AuthorizationCode,omitempty" yaml:"AuthorizationCode,omitempty"`
}

type RefundPriorWalletChargeResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ RefundPriorWalletChargeResult"`

	*CreatePaymentResponse

	Amount float64 `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`

	AuthorizationCode string `xml:"AuthorizationCode,omitempty" json:"AuthorizationCode,omitempty" yaml:"AuthorizationCode,omitempty"`
}

type UpdateCustomerLeadResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateCustomerLeadResult"`

	*ApiResponse
}

type UpdateItemResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateItemResult"`

	*ApiResponse
}

type SetCustomerLeadSocialNetworksResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetCustomerLeadSocialNetworksResult"`

	*ApiResponse
}

type ApiAuthentication struct {
	LoginName string `xml:"LoginName,omitempty" json:"LoginName,omitempty" yaml:"LoginName,omitempty"`

	Password string `xml:"Password,omitempty" json:"Password,omitempty" yaml:"Password,omitempty"`

	Company string `xml:"Company,omitempty" json:"Company,omitempty" yaml:"Company,omitempty"`

	Identity string `xml:"Identity,omitempty" json:"Identity,omitempty" yaml:"Identity,omitempty"`

	RequestTimeUtc time.Time `xml:"RequestTimeUtc,omitempty" json:"RequestTimeUtc,omitempty" yaml:"RequestTimeUtc,omitempty"`

	Signature string `xml:"Signature,omitempty" json:"Signature,omitempty" yaml:"Signature,omitempty"`
}

type MoveEmailRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	MailID int32 `xml:"MailID,omitempty" json:"MailID,omitempty" yaml:"MailID,omitempty"`

	ToMailFolderID int32 `xml:"ToMailFolderID,omitempty" json:"ToMailFolderID,omitempty" yaml:"ToMailFolderID,omitempty"`
}

type MoveEmailResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ MoveEmailResult"`

	*ApiResponse
}

type UpdateEmailStatusRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	MailID int32 `xml:"MailID,omitempty" json:"MailID,omitempty" yaml:"MailID,omitempty"`

	MailStatusType *MailStatusType `xml:"MailStatusType,omitempty" json:"MailStatusType,omitempty" yaml:"MailStatusType,omitempty"`
}

type UpdateEmailStatusResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateEmailStatusResult"`

	*ApiResponse
}

type GetEmailAttachmentRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	MailID int32 `xml:"MailID,omitempty" json:"MailID,omitempty" yaml:"MailID,omitempty"`

	AttachmentID int32 `xml:"AttachmentID,omitempty" json:"AttachmentID,omitempty" yaml:"AttachmentID,omitempty"`
}

type GetEmailAttachmentResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetEmailAttachmentResult"`

	*ApiResponse

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	MailID int32 `xml:"MailID,omitempty" json:"MailID,omitempty" yaml:"MailID,omitempty"`

	Attachment *EmailAttachment `xml:"Attachment,omitempty" json:"Attachment,omitempty" yaml:"Attachment,omitempty"`
}

type DeleteEmailRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	MailID int32 `xml:"MailID,omitempty" json:"MailID,omitempty" yaml:"MailID,omitempty"`
}

type DeleteEmailResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteEmailResult"`

	*ApiResponse
}

type CreateEmailTemplateRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	Content string `xml:"Content,omitempty" json:"Content,omitempty" yaml:"Content,omitempty"`
}

type CreateEmailTemplateResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateEmailTemplateResult"`

	*ApiResponse

	TemplateID int32 `xml:"TemplateID,omitempty" json:"TemplateID,omitempty" yaml:"TemplateID,omitempty"`
}

type UpdateEmailTemplateRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	TemplateID int32 `xml:"TemplateID,omitempty" json:"TemplateID,omitempty" yaml:"TemplateID,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	Content string `xml:"Content,omitempty" json:"Content,omitempty" yaml:"Content,omitempty"`
}

type UpdateEmailTemplateResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateEmailTemplateResult"`

	*ApiResponse
}

type DeleteEmailTemplateRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	TemplateID int32 `xml:"TemplateID,omitempty" json:"TemplateID,omitempty" yaml:"TemplateID,omitempty"`
}

type DeleteEmailTemplateResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteEmailTemplateResult"`

	*ApiResponse
}

type EnsureMailFoldersRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`
}

type EnsureMailFoldersResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ EnsureMailFoldersResult"`

	*ApiResponse
}

type CreateMailFolderRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	MailFolderName string `xml:"MailFolderName,omitempty" json:"MailFolderName,omitempty" yaml:"MailFolderName,omitempty"`
}

type CreateMailFolderResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateMailFolderResult"`

	*ApiResponse
}

type UpdateMailFolderRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	MailFolderID int32 `xml:"MailFolderID,omitempty" json:"MailFolderID,omitempty" yaml:"MailFolderID,omitempty"`

	MailFolderName string `xml:"MailFolderName,omitempty" json:"MailFolderName,omitempty" yaml:"MailFolderName,omitempty"`
}

type UpdateMailFolderResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateMailFolderResult"`

	*ApiResponse
}

type DeleteMailFolderRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	MailFolderID int32 `xml:"MailFolderID,omitempty" json:"MailFolderID,omitempty" yaml:"MailFolderID,omitempty"`
}

type DeleteMailFolderResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteMailFolderResult"`

	*ApiResponse
}

type EmptyMailFolderRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	MailFolderID int32 `xml:"MailFolderID,omitempty" json:"MailFolderID,omitempty" yaml:"MailFolderID,omitempty"`
}

type EmptyMailFolderResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ EmptyMailFolderResult"`

	*ApiResponse
}

type UpdateItemRequest struct {
	*ApiRequest

	ItemCode string `xml:"ItemCode,omitempty" json:"ItemCode,omitempty" yaml:"ItemCode,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	Weight float64 `xml:"Weight,omitempty" json:"Weight,omitempty" yaml:"Weight,omitempty"`

	Notes string `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`

	AvailableInAllCountryRegions bool `xml:"AvailableInAllCountryRegions,omitempty" json:"AvailableInAllCountryRegions,omitempty" yaml:"AvailableInAllCountryRegions,omitempty"`

	TaxedInAllCountryRegions bool `xml:"TaxedInAllCountryRegions,omitempty" json:"TaxedInAllCountryRegions,omitempty" yaml:"TaxedInAllCountryRegions,omitempty"`

	AvailableInAllWarehouses bool `xml:"AvailableInAllWarehouses,omitempty" json:"AvailableInAllWarehouses,omitempty" yaml:"AvailableInAllWarehouses,omitempty"`

	IsVirtual bool `xml:"IsVirtual,omitempty" json:"IsVirtual,omitempty" yaml:"IsVirtual,omitempty"`

	ItemTypeID int32 `xml:"ItemTypeID,omitempty" json:"ItemTypeID,omitempty" yaml:"ItemTypeID,omitempty"`

	ShortDetail string `xml:"ShortDetail,omitempty" json:"ShortDetail,omitempty" yaml:"ShortDetail,omitempty"`

	ShortDetail2 string `xml:"ShortDetail2,omitempty" json:"ShortDetail2,omitempty" yaml:"ShortDetail2,omitempty"`

	ShortDetail3 string `xml:"ShortDetail3,omitempty" json:"ShortDetail3,omitempty" yaml:"ShortDetail3,omitempty"`

	ShortDetail4 string `xml:"ShortDetail4,omitempty" json:"ShortDetail4,omitempty" yaml:"ShortDetail4,omitempty"`

	LongDetail string `xml:"LongDetail,omitempty" json:"LongDetail,omitempty" yaml:"LongDetail,omitempty"`

	LongDetail2 string `xml:"LongDetail2,omitempty" json:"LongDetail2,omitempty" yaml:"LongDetail2,omitempty"`

	LongDetail3 string `xml:"LongDetail3,omitempty" json:"LongDetail3,omitempty" yaml:"LongDetail3,omitempty"`

	LongDetail4 string `xml:"LongDetail4,omitempty" json:"LongDetail4,omitempty" yaml:"LongDetail4,omitempty"`

	OtherCheck1 bool `xml:"OtherCheck1,omitempty" json:"OtherCheck1,omitempty" yaml:"OtherCheck1,omitempty"`

	OtherCheck2 bool `xml:"OtherCheck2,omitempty" json:"OtherCheck2,omitempty" yaml:"OtherCheck2,omitempty"`

	OtherCheck3 bool `xml:"OtherCheck3,omitempty" json:"OtherCheck3,omitempty" yaml:"OtherCheck3,omitempty"`

	OtherCheck4 bool `xml:"OtherCheck4,omitempty" json:"OtherCheck4,omitempty" yaml:"OtherCheck4,omitempty"`

	OtherCheck5 bool `xml:"OtherCheck5,omitempty" json:"OtherCheck5,omitempty" yaml:"OtherCheck5,omitempty"`

	Field1 string `xml:"Field1,omitempty" json:"Field1,omitempty" yaml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty" json:"Field2,omitempty" yaml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty" json:"Field3,omitempty" yaml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty" json:"Field4,omitempty" yaml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty" json:"Field5,omitempty" yaml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty" json:"Field6,omitempty" yaml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty" json:"Field7,omitempty" yaml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty" json:"Field8,omitempty" yaml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty" json:"Field9,omitempty" yaml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty" json:"Field10,omitempty" yaml:"Field10,omitempty"`

	HideFromSearch bool `xml:"HideFromSearch,omitempty" json:"HideFromSearch,omitempty" yaml:"HideFromSearch,omitempty"`
}

type SetItemPriceResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetItemPriceResult"`

	*ApiResponse
}

type SetItemWarehouseResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetItemWarehouseResult"`

	*ApiResponse
}

type GetItemCountryRegionRequest struct {
	*ApiRequest

	ItemCode string `xml:"ItemCode,omitempty" json:"ItemCode,omitempty" yaml:"ItemCode,omitempty"`
}

type SetItemImageRequest struct {
	*ApiRequest

	ItemCode string `xml:"ItemCode,omitempty" json:"ItemCode,omitempty" yaml:"ItemCode,omitempty"`

	TinyImageName string `xml:"TinyImageName,omitempty" json:"TinyImageName,omitempty" yaml:"TinyImageName,omitempty"`

	TinyImageData []byte `xml:"TinyImageData,omitempty" json:"TinyImageData,omitempty" yaml:"TinyImageData,omitempty"`

	SmallImageName string `xml:"SmallImageName,omitempty" json:"SmallImageName,omitempty" yaml:"SmallImageName,omitempty"`

	SmallImageData []byte `xml:"SmallImageData,omitempty" json:"SmallImageData,omitempty" yaml:"SmallImageData,omitempty"`

	LargeImageName string `xml:"LargeImageName,omitempty" json:"LargeImageName,omitempty" yaml:"LargeImageName,omitempty"`

	LargeImageData []byte `xml:"LargeImageData,omitempty" json:"LargeImageData,omitempty" yaml:"LargeImageData,omitempty"`
}

type SetItemImageResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetItemImageResult"`

	*ApiResponse
}

type SetImageFileRequest struct {
	*ApiRequest

	Path string `xml:"Path,omitempty" json:"Path,omitempty" yaml:"Path,omitempty"`

	Name string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`

	ImageData []byte `xml:"ImageData,omitempty" json:"ImageData,omitempty" yaml:"ImageData,omitempty"`
}

type SetImageFileResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetImageFileResult"`

	*ApiResponse
}

type GetCustomerBalancesRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`
}

type GetCustomerBalancesResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomerBalancesResult"`

	*ApiResponse

	CustomerBalances *ArrayOfCustomerBalanceResponse `xml:"CustomerBalances,omitempty" json:"CustomerBalances,omitempty" yaml:"CustomerBalances,omitempty"`
}

type ArrayOfCustomerBalanceResponse struct {
	CustomerBalanceResponse []*CustomerBalanceResponse `xml:"CustomerBalanceResponse,omitempty" json:"CustomerBalanceResponse,omitempty" yaml:"CustomerBalanceResponse,omitempty"`
}

type CustomerBalanceResponse struct {
	CurrencyCode string `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty" yaml:"CurrencyCode,omitempty"`

	CurrencyDescription string `xml:"CurrencyDescription,omitempty" json:"CurrencyDescription,omitempty" yaml:"CurrencyDescription,omitempty"`

	Balance float64 `xml:"Balance,omitempty" json:"Balance,omitempty" yaml:"Balance,omitempty"`
}

type CreateCustomerBalanceAdjustmentRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	CustomerTransactionTypeID int32 `xml:"CustomerTransactionTypeID,omitempty" json:"CustomerTransactionTypeID,omitempty" yaml:"CustomerTransactionTypeID,omitempty"`

	TransactionDate time.Time `xml:"TransactionDate,omitempty" json:"TransactionDate,omitempty" yaml:"TransactionDate,omitempty"`

	Amount float64 `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty" yaml:"CurrencyCode,omitempty"`

	Notes string `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`
}

type CreateCustomerBalanceAdjustmentResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateCustomerBalanceAdjustmentResult"`

	*ApiResponse

	TransactionID int32 `xml:"TransactionID,omitempty" json:"TransactionID,omitempty" yaml:"TransactionID,omitempty"`
}

type GetPartiesRequest struct {
	*ApiRequest

	PartyID int32 `xml:"PartyID,omitempty" json:"PartyID,omitempty" yaml:"PartyID,omitempty"`

	HostID int32 `xml:"HostID,omitempty" json:"HostID,omitempty" yaml:"HostID,omitempty"`

	DistributorID int32 `xml:"DistributorID,omitempty" json:"DistributorID,omitempty" yaml:"DistributorID,omitempty"`

	PartyStatusType int32 `xml:"PartyStatusType,omitempty" json:"PartyStatusType,omitempty" yaml:"PartyStatusType,omitempty"`

	BookingPartyID int32 `xml:"BookingPartyID,omitempty" json:"BookingPartyID,omitempty" yaml:"BookingPartyID,omitempty"`

	Field1 string `xml:"Field1,omitempty" json:"Field1,omitempty" yaml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty" json:"Field2,omitempty" yaml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty" json:"Field3,omitempty" yaml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty" json:"Field4,omitempty" yaml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty" json:"Field5,omitempty" yaml:"Field5,omitempty"`
}

type GetPartiesResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetPartiesResult"`

	*ApiResponse

	Parties *ArrayOfPartyResponse `xml:"Parties,omitempty" json:"Parties,omitempty" yaml:"Parties,omitempty"`
}

type ArrayOfPartyResponse struct {
	PartyResponse []*PartyResponse `xml:"PartyResponse,omitempty" json:"PartyResponse,omitempty" yaml:"PartyResponse,omitempty"`
}

type PartyResponse struct {
	PartyID int32 `xml:"PartyID,omitempty" json:"PartyID,omitempty" yaml:"PartyID,omitempty"`

	PartyType int32 `xml:"PartyType,omitempty" json:"PartyType,omitempty" yaml:"PartyType,omitempty"`

	PartyStatusType int32 `xml:"PartyStatusType,omitempty" json:"PartyStatusType,omitempty" yaml:"PartyStatusType,omitempty"`

	HostID int32 `xml:"HostID,omitempty" json:"HostID,omitempty" yaml:"HostID,omitempty"`

	DistributorID int32 `xml:"DistributorID,omitempty" json:"DistributorID,omitempty" yaml:"DistributorID,omitempty"`

	StartDate time.Time `xml:"StartDate,omitempty" json:"StartDate,omitempty" yaml:"StartDate,omitempty"`

	CloseDate time.Time `xml:"CloseDate,omitempty" json:"CloseDate,omitempty" yaml:"CloseDate,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	EventStart time.Time `xml:"EventStart,omitempty" json:"EventStart,omitempty" yaml:"EventStart,omitempty"`

	EventEnd time.Time `xml:"EventEnd,omitempty" json:"EventEnd,omitempty" yaml:"EventEnd,omitempty"`

	LanguageID int32 `xml:"LanguageID,omitempty" json:"LanguageID,omitempty" yaml:"LanguageID,omitempty"`

	Information string `xml:"Information,omitempty" json:"Information,omitempty" yaml:"Information,omitempty"`

	Address *PartyAddress `xml:"Address,omitempty" json:"Address,omitempty" yaml:"Address,omitempty"`

	BookingPartyID int32 `xml:"BookingPartyID,omitempty" json:"BookingPartyID,omitempty" yaml:"BookingPartyID,omitempty"`

	Field1 string `xml:"Field1,omitempty" json:"Field1,omitempty" yaml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty" json:"Field2,omitempty" yaml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty" json:"Field3,omitempty" yaml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty" json:"Field4,omitempty" yaml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty" json:"Field5,omitempty" yaml:"Field5,omitempty"`
}

type GetGuestsRequest struct {
	*ApiRequest

	GuestID int32 `xml:"GuestID,omitempty" json:"GuestID,omitempty" yaml:"GuestID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	HostID int32 `xml:"HostID,omitempty" json:"HostID,omitempty" yaml:"HostID,omitempty"`

	GuestStatuses *ArrayOfInt `xml:"GuestStatuses,omitempty" json:"GuestStatuses,omitempty" yaml:"GuestStatuses,omitempty"`

	LanguageID int32 `xml:"LanguageID,omitempty" json:"LanguageID,omitempty" yaml:"LanguageID,omitempty"`

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty" json:"MiddleName,omitempty" yaml:"MiddleName,omitempty"`

	LastName string `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`

	NameSuffix string `xml:"NameSuffix,omitempty" json:"NameSuffix,omitempty" yaml:"NameSuffix,omitempty"`

	Company string `xml:"Company,omitempty" json:"Company,omitempty" yaml:"Company,omitempty"`

	Email string `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`

	Phone string `xml:"Phone,omitempty" json:"Phone,omitempty" yaml:"Phone,omitempty"`

	Phone2 string `xml:"Phone2,omitempty" json:"Phone2,omitempty" yaml:"Phone2,omitempty"`

	MobilePhone string `xml:"MobilePhone,omitempty" json:"MobilePhone,omitempty" yaml:"MobilePhone,omitempty"`

	Address1 string `xml:"Address1,omitempty" json:"Address1,omitempty" yaml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty" json:"Address2,omitempty" yaml:"Address2,omitempty"`

	Address3 string `xml:"Address3,omitempty" json:"Address3,omitempty" yaml:"Address3,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`

	County string `xml:"County,omitempty" json:"County,omitempty" yaml:"County,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty" json:"Zip,omitempty" yaml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`

	Field1 string `xml:"Field1,omitempty" json:"Field1,omitempty" yaml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty" json:"Field2,omitempty" yaml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty" json:"Field3,omitempty" yaml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty" json:"Field4,omitempty" yaml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty" json:"Field5,omitempty" yaml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty" json:"Field6,omitempty" yaml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty" json:"Field7,omitempty" yaml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty" json:"Field8,omitempty" yaml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty" json:"Field9,omitempty" yaml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty" json:"Field10,omitempty" yaml:"Field10,omitempty"`

	Field11 string `xml:"Field11,omitempty" json:"Field11,omitempty" yaml:"Field11,omitempty"`

	Field12 string `xml:"Field12,omitempty" json:"Field12,omitempty" yaml:"Field12,omitempty"`

	Field13 string `xml:"Field13,omitempty" json:"Field13,omitempty" yaml:"Field13,omitempty"`

	Field14 string `xml:"Field14,omitempty" json:"Field14,omitempty" yaml:"Field14,omitempty"`

	Field15 string `xml:"Field15,omitempty" json:"Field15,omitempty" yaml:"Field15,omitempty"`

	Date1 time.Time `xml:"Date1,omitempty" json:"Date1,omitempty" yaml:"Date1,omitempty"`

	Date2 time.Time `xml:"Date2,omitempty" json:"Date2,omitempty" yaml:"Date2,omitempty"`

	Date3 time.Time `xml:"Date3,omitempty" json:"Date3,omitempty" yaml:"Date3,omitempty"`

	Date4 time.Time `xml:"Date4,omitempty" json:"Date4,omitempty" yaml:"Date4,omitempty"`

	Date5 time.Time `xml:"Date5,omitempty" json:"Date5,omitempty" yaml:"Date5,omitempty"`

	CreatedDateStart time.Time `xml:"CreatedDateStart,omitempty" json:"CreatedDateStart,omitempty" yaml:"CreatedDateStart,omitempty"`

	CreatedDateEnd time.Time `xml:"CreatedDateEnd,omitempty" json:"CreatedDateEnd,omitempty" yaml:"CreatedDateEnd,omitempty"`
}

type GetGuestsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetGuestsResult"`

	*ApiResponse

	Guests *ArrayOfGuestResponse `xml:"Guests,omitempty" json:"Guests,omitempty" yaml:"Guests,omitempty"`
}

type ArrayOfGuestResponse struct {
	GuestResponse []*GuestResponse `xml:"GuestResponse,omitempty" json:"GuestResponse,omitempty" yaml:"GuestResponse,omitempty"`
}

type GuestResponse struct {
	GuestID int32 `xml:"GuestID,omitempty" json:"GuestID,omitempty" yaml:"GuestID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	HostID int32 `xml:"HostID,omitempty" json:"HostID,omitempty" yaml:"HostID,omitempty"`

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty" json:"MiddleName,omitempty" yaml:"MiddleName,omitempty"`

	LastName string `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`

	NameSuffix string `xml:"NameSuffix,omitempty" json:"NameSuffix,omitempty" yaml:"NameSuffix,omitempty"`

	Company string `xml:"Company,omitempty" json:"Company,omitempty" yaml:"Company,omitempty"`

	Gender *Gender `xml:"Gender,omitempty" json:"Gender,omitempty" yaml:"Gender,omitempty"`

	GuestStatus int32 `xml:"GuestStatus,omitempty" json:"GuestStatus,omitempty" yaml:"GuestStatus,omitempty"`

	LanguageID int32 `xml:"LanguageID,omitempty" json:"LanguageID,omitempty" yaml:"LanguageID,omitempty"`

	Address1 string `xml:"Address1,omitempty" json:"Address1,omitempty" yaml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty" json:"Address2,omitempty" yaml:"Address2,omitempty"`

	Address3 string `xml:"Address3,omitempty" json:"Address3,omitempty" yaml:"Address3,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`

	County string `xml:"County,omitempty" json:"County,omitempty" yaml:"County,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty" json:"Zip,omitempty" yaml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`

	Phone string `xml:"Phone,omitempty" json:"Phone,omitempty" yaml:"Phone,omitempty"`

	Phone2 string `xml:"Phone2,omitempty" json:"Phone2,omitempty" yaml:"Phone2,omitempty"`

	MobilePhone string `xml:"MobilePhone,omitempty" json:"MobilePhone,omitempty" yaml:"MobilePhone,omitempty"`

	Email string `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`

	Field1 string `xml:"Field1,omitempty" json:"Field1,omitempty" yaml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty" json:"Field2,omitempty" yaml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty" json:"Field3,omitempty" yaml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty" json:"Field4,omitempty" yaml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty" json:"Field5,omitempty" yaml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty" json:"Field6,omitempty" yaml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty" json:"Field7,omitempty" yaml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty" json:"Field8,omitempty" yaml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty" json:"Field9,omitempty" yaml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty" json:"Field10,omitempty" yaml:"Field10,omitempty"`

	Field11 string `xml:"Field11,omitempty" json:"Field11,omitempty" yaml:"Field11,omitempty"`

	Field12 string `xml:"Field12,omitempty" json:"Field12,omitempty" yaml:"Field12,omitempty"`

	Field13 string `xml:"Field13,omitempty" json:"Field13,omitempty" yaml:"Field13,omitempty"`

	Field14 string `xml:"Field14,omitempty" json:"Field14,omitempty" yaml:"Field14,omitempty"`

	Field15 string `xml:"Field15,omitempty" json:"Field15,omitempty" yaml:"Field15,omitempty"`

	Date1 time.Time `xml:"Date1,omitempty" json:"Date1,omitempty" yaml:"Date1,omitempty"`

	Date2 time.Time `xml:"Date2,omitempty" json:"Date2,omitempty" yaml:"Date2,omitempty"`

	Date3 time.Time `xml:"Date3,omitempty" json:"Date3,omitempty" yaml:"Date3,omitempty"`

	Date4 time.Time `xml:"Date4,omitempty" json:"Date4,omitempty" yaml:"Date4,omitempty"`

	Date5 time.Time `xml:"Date5,omitempty" json:"Date5,omitempty" yaml:"Date5,omitempty"`

	CreatedDate time.Time `xml:"CreatedDate,omitempty" json:"CreatedDate,omitempty" yaml:"CreatedDate,omitempty"`

	ModifiedDate time.Time `xml:"ModifiedDate,omitempty" json:"ModifiedDate,omitempty" yaml:"ModifiedDate,omitempty"`
}

type GetPartyGuestsRequest struct {
	*ApiRequest

	PartyID int32 `xml:"PartyID,omitempty" json:"PartyID,omitempty" yaml:"PartyID,omitempty"`
}

type GetPartyGuestsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetPartyGuestsResult"`

	*ApiResponse

	Guests *ArrayOfGuestResponse `xml:"Guests,omitempty" json:"Guests,omitempty" yaml:"Guests,omitempty"`
}

type GetGuestSocialNetworksRequest struct {
	*ApiRequest

	GuestID int32 `xml:"GuestID,omitempty" json:"GuestID,omitempty" yaml:"GuestID,omitempty"`
}

type GetGuestSocialNetworksResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetGuestSocialNetworksResult"`

	*ApiResponse

	GuestSocialNetworks *ArrayOfGuestSocialNetworksResponse `xml:"GuestSocialNetworks,omitempty" json:"GuestSocialNetworks,omitempty" yaml:"GuestSocialNetworks,omitempty"`
}

type ArrayOfGuestSocialNetworksResponse struct {
	GuestSocialNetworksResponse []*GuestSocialNetworksResponse `xml:"GuestSocialNetworksResponse,omitempty" json:"GuestSocialNetworksResponse,omitempty" yaml:"GuestSocialNetworksResponse,omitempty"`
}

type GuestSocialNetworksResponse struct {
	SocialNetworkID int32 `xml:"SocialNetworkID,omitempty" json:"SocialNetworkID,omitempty" yaml:"SocialNetworkID,omitempty"`

	SocialNetworkDescription string `xml:"SocialNetworkDescription,omitempty" json:"SocialNetworkDescription,omitempty" yaml:"SocialNetworkDescription,omitempty"`

	Url string `xml:"Url,omitempty" json:"Url,omitempty" yaml:"Url,omitempty"`
}

type SetGuestSocialNetworksRequest struct {
	*ApiRequest

	GuestID int32 `xml:"GuestID,omitempty" json:"GuestID,omitempty" yaml:"GuestID,omitempty"`

	GuestSocialNetworks *ArrayOfGuestSocialNetworkRequest `xml:"GuestSocialNetworks,omitempty" json:"GuestSocialNetworks,omitempty" yaml:"GuestSocialNetworks,omitempty"`
}

type ArrayOfGuestSocialNetworkRequest struct {
	GuestSocialNetworkRequest []*GuestSocialNetworkRequest `xml:"GuestSocialNetworkRequest,omitempty" json:"GuestSocialNetworkRequest,omitempty" yaml:"GuestSocialNetworkRequest,omitempty"`
}

type GuestSocialNetworkRequest struct {
	SocialNetworkID int32 `xml:"SocialNetworkID,omitempty" json:"SocialNetworkID,omitempty" yaml:"SocialNetworkID,omitempty"`

	Url string `xml:"Url,omitempty" json:"Url,omitempty" yaml:"Url,omitempty"`
}

type SetGuestSocialNetworksResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetGuestSocialNetworksResult"`

	*ApiResponse
}

type CreateGuestRequest struct {
	*ApiRequest

	Date1 time.Time `xml:"Date1,omitempty" json:"Date1,omitempty" yaml:"Date1,omitempty"`

	Date2 time.Time `xml:"Date2,omitempty" json:"Date2,omitempty" yaml:"Date2,omitempty"`

	Date3 time.Time `xml:"Date3,omitempty" json:"Date3,omitempty" yaml:"Date3,omitempty"`

	Date4 time.Time `xml:"Date4,omitempty" json:"Date4,omitempty" yaml:"Date4,omitempty"`

	Date5 time.Time `xml:"Date5,omitempty" json:"Date5,omitempty" yaml:"Date5,omitempty"`

	HostID int32 `xml:"HostID,omitempty" json:"HostID,omitempty" yaml:"HostID,omitempty"`

	PartyID int32 `xml:"PartyID,omitempty" json:"PartyID,omitempty" yaml:"PartyID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty" json:"MiddleName,omitempty" yaml:"MiddleName,omitempty"`

	LastName string `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`

	NameSuffix string `xml:"NameSuffix,omitempty" json:"NameSuffix,omitempty" yaml:"NameSuffix,omitempty"`

	Company string `xml:"Company,omitempty" json:"Company,omitempty" yaml:"Company,omitempty"`

	Gender *Gender `xml:"Gender,omitempty" json:"Gender,omitempty" yaml:"Gender,omitempty"`

	LanguageID int32 `xml:"LanguageID,omitempty" json:"LanguageID,omitempty" yaml:"LanguageID,omitempty"`

	GuestStatus int32 `xml:"GuestStatus,omitempty" json:"GuestStatus,omitempty" yaml:"GuestStatus,omitempty"`

	Address1 string `xml:"Address1,omitempty" json:"Address1,omitempty" yaml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty" json:"Address2,omitempty" yaml:"Address2,omitempty"`

	Address3 string `xml:"Address3,omitempty" json:"Address3,omitempty" yaml:"Address3,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`

	County string `xml:"County,omitempty" json:"County,omitempty" yaml:"County,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty" json:"Zip,omitempty" yaml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`

	Phone string `xml:"Phone,omitempty" json:"Phone,omitempty" yaml:"Phone,omitempty"`

	Phone2 string `xml:"Phone2,omitempty" json:"Phone2,omitempty" yaml:"Phone2,omitempty"`

	MobilePhone string `xml:"MobilePhone,omitempty" json:"MobilePhone,omitempty" yaml:"MobilePhone,omitempty"`

	Email string `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`

	Field1 string `xml:"Field1,omitempty" json:"Field1,omitempty" yaml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty" json:"Field2,omitempty" yaml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty" json:"Field3,omitempty" yaml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty" json:"Field4,omitempty" yaml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty" json:"Field5,omitempty" yaml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty" json:"Field6,omitempty" yaml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty" json:"Field7,omitempty" yaml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty" json:"Field8,omitempty" yaml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty" json:"Field9,omitempty" yaml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty" json:"Field10,omitempty" yaml:"Field10,omitempty"`

	Field11 string `xml:"Field11,omitempty" json:"Field11,omitempty" yaml:"Field11,omitempty"`

	Field12 string `xml:"Field12,omitempty" json:"Field12,omitempty" yaml:"Field12,omitempty"`

	Field13 string `xml:"Field13,omitempty" json:"Field13,omitempty" yaml:"Field13,omitempty"`

	Field14 string `xml:"Field14,omitempty" json:"Field14,omitempty" yaml:"Field14,omitempty"`

	Field15 string `xml:"Field15,omitempty" json:"Field15,omitempty" yaml:"Field15,omitempty"`

	Notes string `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`

	EntryDate time.Time `xml:"EntryDate,omitempty" json:"EntryDate,omitempty" yaml:"EntryDate,omitempty"`
}

type CreateGuestResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateGuestResult"`

	*ApiResponse

	GuestID int32 `xml:"GuestID,omitempty" json:"GuestID,omitempty" yaml:"GuestID,omitempty"`
}

type UpdateGuestRequest struct {
	*ApiRequest

	GuestID int32 `xml:"GuestID,omitempty" json:"GuestID,omitempty" yaml:"GuestID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty" json:"MiddleName,omitempty" yaml:"MiddleName,omitempty"`

	LastName string `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`

	NameSuffix string `xml:"NameSuffix,omitempty" json:"NameSuffix,omitempty" yaml:"NameSuffix,omitempty"`

	Company string `xml:"Company,omitempty" json:"Company,omitempty" yaml:"Company,omitempty"`

	Gender *Gender `xml:"Gender,omitempty" json:"Gender,omitempty" yaml:"Gender,omitempty"`

	GuestStatus int32 `xml:"GuestStatus,omitempty" json:"GuestStatus,omitempty" yaml:"GuestStatus,omitempty"`

	LanguageID int32 `xml:"LanguageID,omitempty" json:"LanguageID,omitempty" yaml:"LanguageID,omitempty"`

	Address1 string `xml:"Address1,omitempty" json:"Address1,omitempty" yaml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty" json:"Address2,omitempty" yaml:"Address2,omitempty"`

	Address3 string `xml:"Address3,omitempty" json:"Address3,omitempty" yaml:"Address3,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`

	County string `xml:"County,omitempty" json:"County,omitempty" yaml:"County,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty" json:"Zip,omitempty" yaml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`

	Phone string `xml:"Phone,omitempty" json:"Phone,omitempty" yaml:"Phone,omitempty"`

	Phone2 string `xml:"Phone2,omitempty" json:"Phone2,omitempty" yaml:"Phone2,omitempty"`

	MobilePhone string `xml:"MobilePhone,omitempty" json:"MobilePhone,omitempty" yaml:"MobilePhone,omitempty"`

	Email string `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`

	Field1 string `xml:"Field1,omitempty" json:"Field1,omitempty" yaml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty" json:"Field2,omitempty" yaml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty" json:"Field3,omitempty" yaml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty" json:"Field4,omitempty" yaml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty" json:"Field5,omitempty" yaml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty" json:"Field6,omitempty" yaml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty" json:"Field7,omitempty" yaml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty" json:"Field8,omitempty" yaml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty" json:"Field9,omitempty" yaml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty" json:"Field10,omitempty" yaml:"Field10,omitempty"`

	Field11 string `xml:"Field11,omitempty" json:"Field11,omitempty" yaml:"Field11,omitempty"`

	Field12 string `xml:"Field12,omitempty" json:"Field12,omitempty" yaml:"Field12,omitempty"`

	Field13 string `xml:"Field13,omitempty" json:"Field13,omitempty" yaml:"Field13,omitempty"`

	Field14 string `xml:"Field14,omitempty" json:"Field14,omitempty" yaml:"Field14,omitempty"`

	Field15 string `xml:"Field15,omitempty" json:"Field15,omitempty" yaml:"Field15,omitempty"`

	Date1 time.Time `xml:"Date1,omitempty" json:"Date1,omitempty" yaml:"Date1,omitempty"`

	Date2 time.Time `xml:"Date2,omitempty" json:"Date2,omitempty" yaml:"Date2,omitempty"`

	Date3 time.Time `xml:"Date3,omitempty" json:"Date3,omitempty" yaml:"Date3,omitempty"`

	Date4 time.Time `xml:"Date4,omitempty" json:"Date4,omitempty" yaml:"Date4,omitempty"`

	Date5 time.Time `xml:"Date5,omitempty" json:"Date5,omitempty" yaml:"Date5,omitempty"`

	Notes string `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`
}

type UpdateGuestResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateGuestResult"`

	*ApiResponse
}

type AddPartyGuestsRequest struct {
	*ApiRequest

	PartyID int32 `xml:"PartyID,omitempty" json:"PartyID,omitempty" yaml:"PartyID,omitempty"`

	GuestIDs *ArrayOfInt `xml:"GuestIDs,omitempty" json:"GuestIDs,omitempty" yaml:"GuestIDs,omitempty"`
}

type AddPartyGuestsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ AddPartyGuestsResult"`

	*ApiResponse
}

type RemovePartyGuestsRequest struct {
	*ApiRequest

	PartyID int32 `xml:"PartyID,omitempty" json:"PartyID,omitempty" yaml:"PartyID,omitempty"`

	GuestIDs *ArrayOfInt `xml:"GuestIDs,omitempty" json:"GuestIDs,omitempty" yaml:"GuestIDs,omitempty"`
}

type RemovePartyGuestsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ RemovePartyGuestsResult"`

	*ApiResponse
}

type CreateExtendedDbSchemeRequest struct {
	*ApiRequest

	Schema *Schema `xml:"Schema,omitempty" json:"Schema,omitempty" yaml:"Schema,omitempty"`
}

type Schema struct {
	Name string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`

	Entities *ArrayOfEntity `xml:"Entities,omitempty" json:"Entities,omitempty" yaml:"Entities,omitempty"`
}

type ArrayOfEntity struct {
	Entity []*Entity `xml:"Entity,omitempty" json:"Entity,omitempty" yaml:"Entity,omitempty"`
}

type Entity struct {
	SchemaName string `xml:"SchemaName,omitempty" json:"SchemaName,omitempty" yaml:"SchemaName,omitempty"`

	DbSchema string `xml:"DbSchema,omitempty" json:"DbSchema,omitempty" yaml:"DbSchema,omitempty"`

	EntityName string `xml:"EntityName,omitempty" json:"EntityName,omitempty" yaml:"EntityName,omitempty"`

	EntitySetName string `xml:"EntitySetName,omitempty" json:"EntitySetName,omitempty" yaml:"EntitySetName,omitempty"`

	Properties *ArrayOfProperty `xml:"Properties,omitempty" json:"Properties,omitempty" yaml:"Properties,omitempty"`

	Navigations *ArrayOfNavigation `xml:"Navigations,omitempty" json:"Navigations,omitempty" yaml:"Navigations,omitempty"`

	SyncTypeID int32 `xml:"SyncTypeID,omitempty" json:"SyncTypeID,omitempty" yaml:"SyncTypeID,omitempty"`
}

type ArrayOfProperty struct {
	Property []*Property `xml:"Property,omitempty" json:"Property,omitempty" yaml:"Property,omitempty"`
}

type Property struct {
	Name string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`

	IsKey bool `xml:"IsKey,omitempty" json:"IsKey,omitempty" yaml:"IsKey,omitempty"`

	IsNew bool `xml:"IsNew,omitempty" json:"IsNew,omitempty" yaml:"IsNew,omitempty"`

	IsAutoNumber bool `xml:"IsAutoNumber,omitempty" json:"IsAutoNumber,omitempty" yaml:"IsAutoNumber,omitempty"`

	AllowDbNull bool `xml:"AllowDbNull,omitempty" json:"AllowDbNull,omitempty" yaml:"AllowDbNull,omitempty"`

	Type *PropertyType `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`

	DefaultName string `xml:"DefaultName,omitempty" json:"DefaultName,omitempty" yaml:"DefaultName,omitempty"`

	DefaultValue string `xml:"DefaultValue,omitempty" json:"DefaultValue,omitempty" yaml:"DefaultValue,omitempty"`

	Size int32 `xml:"Size,omitempty" json:"Size,omitempty" yaml:"Size,omitempty"`
}

type ArrayOfNavigation struct {
	Navigation []*Navigation `xml:"Navigation,omitempty" json:"Navigation,omitempty" yaml:"Navigation,omitempty"`
}

type Navigation struct {
	Name string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`
}

type CreateExtendedDbSchemaResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateExtendedDbSchemaResult"`

	*ApiResponse

	Schema *Schema `xml:"Schema,omitempty" json:"Schema,omitempty" yaml:"Schema,omitempty"`
}

type GetSchemaRequest struct {
	*ApiRequest

	SchemaName string `xml:"SchemaName,omitempty" json:"SchemaName,omitempty" yaml:"SchemaName,omitempty"`
}

type GetSchemaResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetExtendedDbSchemaResult"`

	*ApiResponse

	Schema *Schema `xml:"Schema,omitempty" json:"Schema,omitempty" yaml:"Schema,omitempty"`
}

type DeleteSchemaRequest struct {
	*ApiRequest

	SchemaName string `xml:"SchemaName,omitempty" json:"SchemaName,omitempty" yaml:"SchemaName,omitempty"`
}

type DeleteSchemaResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteExtendedDbSchemaResult"`

	*ApiResponse

	SchemaName string `xml:"SchemaName,omitempty" json:"SchemaName,omitempty" yaml:"SchemaName,omitempty"`
}

type CreateEntityRequest struct {
	*ApiRequest

	Entity *Entity `xml:"Entity,omitempty" json:"Entity,omitempty" yaml:"Entity,omitempty"`
}

type CreateEntityResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateExtendedDbEntityResult"`

	*ApiResponse

	Entity *Entity `xml:"Entity,omitempty" json:"Entity,omitempty" yaml:"Entity,omitempty"`
}

type GetEntityRequest struct {
	*ApiRequest

	SchemaName string `xml:"SchemaName,omitempty" json:"SchemaName,omitempty" yaml:"SchemaName,omitempty"`

	EntityName string `xml:"EntityName,omitempty" json:"EntityName,omitempty" yaml:"EntityName,omitempty"`
}

type GetEntityResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetExtendedDbEntityResult"`

	*ApiResponse

	Entity *Entity `xml:"Entity,omitempty" json:"Entity,omitempty" yaml:"Entity,omitempty"`
}

type UpdateEntityRequest struct {
	*ApiRequest

	Entity *Entity `xml:"Entity,omitempty" json:"Entity,omitempty" yaml:"Entity,omitempty"`

	EntityName string `xml:"EntityName,omitempty" json:"EntityName,omitempty" yaml:"EntityName,omitempty"`
}

type UpdateEntityResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateExtendedDbEntityResult"`

	*ApiResponse

	Entity *Entity `xml:"Entity,omitempty" json:"Entity,omitempty" yaml:"Entity,omitempty"`
}

type DeleteEntityRequest struct {
	*ApiRequest

	SchemaName string `xml:"SchemaName,omitempty" json:"SchemaName,omitempty" yaml:"SchemaName,omitempty"`

	EntityName string `xml:"EntityName,omitempty" json:"EntityName,omitempty" yaml:"EntityName,omitempty"`
}

type DeleteEntityResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteExtendedDbEntityResult"`

	*ApiResponse

	SchemaName string `xml:"schemaName,omitempty" json:"schemaName,omitempty" yaml:"schemaName,omitempty"`

	EntityName string `xml:"EntityName,omitempty" json:"EntityName,omitempty" yaml:"EntityName,omitempty"`
}

type StartSandboxRequest struct {
	*ApiRequest

	SandboxID int32 `xml:"SandboxID,omitempty" json:"SandboxID,omitempty" yaml:"SandboxID,omitempty"`

	EnableRevolvingCommissionRun bool `xml:"EnableRevolvingCommissionRun,omitempty" json:"EnableRevolvingCommissionRun,omitempty" yaml:"EnableRevolvingCommissionRun,omitempty"`

	EnableBiSync bool `xml:"EnableBiSync,omitempty" json:"EnableBiSync,omitempty" yaml:"EnableBiSync,omitempty"`

	UseRealTimeBackup bool `xml:"UseRealTimeBackup,omitempty" json:"UseRealTimeBackup,omitempty" yaml:"UseRealTimeBackup,omitempty"`

	SyncFilterDays int32 `xml:"SyncFilterDays,omitempty" json:"SyncFilterDays,omitempty" yaml:"SyncFilterDays,omitempty"`

	SyncSettingsEnable string `xml:"SyncSettingsEnable,omitempty" json:"SyncSettingsEnable,omitempty" yaml:"SyncSettingsEnable,omitempty"`

	PremiumSandbox bool `xml:"PremiumSandbox,omitempty" json:"PremiumSandbox,omitempty" yaml:"PremiumSandbox,omitempty"`
}

type StartSandboxResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ StartSandboxResult"`

	*ApiResponse

	Sandbox *Sandbox `xml:"Sandbox,omitempty" json:"Sandbox,omitempty" yaml:"Sandbox,omitempty"`
}

type Sandbox struct {
	CompanyID int32 `xml:"CompanyID,omitempty" json:"CompanyID,omitempty" yaml:"CompanyID,omitempty"`

	SandboxID int32 `xml:"SandboxID,omitempty" json:"SandboxID,omitempty" yaml:"SandboxID,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	Type *SandboxType `xml:"Type,omitempty" json:"Type,omitempty" yaml:"Type,omitempty"`

	Status string `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`

	StartDate time.Time `xml:"StartDate,omitempty" json:"StartDate,omitempty" yaml:"StartDate,omitempty"`

	PercentComplete float64 `xml:"PercentComplete,omitempty" json:"PercentComplete,omitempty" yaml:"PercentComplete,omitempty"`

	Hours float64 `xml:"Hours,omitempty" json:"Hours,omitempty" yaml:"Hours,omitempty"`

	AllowVolumePush bool `xml:"AllowVolumePush,omitempty" json:"AllowVolumePush,omitempty" yaml:"AllowVolumePush,omitempty"`

	AllowBiSync bool `xml:"AllowBiSync,omitempty" json:"AllowBiSync,omitempty" yaml:"AllowBiSync,omitempty"`

	SyncFilterDays int32 `xml:"SyncFilterDays,omitempty" json:"SyncFilterDays,omitempty" yaml:"SyncFilterDays,omitempty"`

	UseRealTimeBackup bool `xml:"UseRealTimeBackup,omitempty" json:"UseRealTimeBackup,omitempty" yaml:"UseRealTimeBackup,omitempty"`
}

type StopSandboxRequest struct {
	*ApiRequest

	SandboxID int32 `xml:"SandboxID,omitempty" json:"SandboxID,omitempty" yaml:"SandboxID,omitempty"`
}

type StopSandboxResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ StopSandboxResult"`

	*ApiResponse

	SandboxID int32 `xml:"SandboxID,omitempty" json:"SandboxID,omitempty" yaml:"SandboxID,omitempty"`
}

type RefreshSandboxRequest struct {
	*ApiRequest

	SandboxID int32 `xml:"SandboxID,omitempty" json:"SandboxID,omitempty" yaml:"SandboxID,omitempty"`
}

type RefreshSandboxResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ RefreshSandboxResult"`

	*ApiResponse

	Sandbox *Sandbox `xml:"Sandbox,omitempty" json:"Sandbox,omitempty" yaml:"Sandbox,omitempty"`
}

type GetSandboxRequest struct {
	*ApiRequest

	SandboxID int32 `xml:"SandboxID,omitempty" json:"SandboxID,omitempty" yaml:"SandboxID,omitempty"`
}

type GetSandboxResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetSandboxResult"`

	*ApiResponse

	Sandbox *Sandbox `xml:"Sandbox,omitempty" json:"Sandbox,omitempty" yaml:"Sandbox,omitempty"`

	Sandboxes *ArrayOfSandbox `xml:"Sandboxes,omitempty" json:"Sandboxes,omitempty" yaml:"Sandboxes,omitempty"`
}

type ArrayOfSandbox struct {
	Sandbox []*Sandbox `xml:"Sandbox,omitempty" json:"Sandbox,omitempty" yaml:"Sandbox,omitempty"`
}

type GetCustomersRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	LoginName string `xml:"LoginName,omitempty" json:"LoginName,omitempty" yaml:"LoginName,omitempty"`

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`

	Company string `xml:"Company,omitempty" json:"Company,omitempty" yaml:"Company,omitempty"`

	Email string `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`

	Phone string `xml:"Phone,omitempty" json:"Phone,omitempty" yaml:"Phone,omitempty"`

	Phone2 string `xml:"Phone2,omitempty" json:"Phone2,omitempty" yaml:"Phone2,omitempty"`

	MobilePhone string `xml:"MobilePhone,omitempty" json:"MobilePhone,omitempty" yaml:"MobilePhone,omitempty"`

	Fax string `xml:"Fax,omitempty" json:"Fax,omitempty" yaml:"Fax,omitempty"`

	MainAddress1 string `xml:"MainAddress1,omitempty" json:"MainAddress1,omitempty" yaml:"MainAddress1,omitempty"`

	MainAddress2 string `xml:"MainAddress2,omitempty" json:"MainAddress2,omitempty" yaml:"MainAddress2,omitempty"`

	MainAddress3 string `xml:"MainAddress3,omitempty" json:"MainAddress3,omitempty" yaml:"MainAddress3,omitempty"`

	MainCity string `xml:"MainCity,omitempty" json:"MainCity,omitempty" yaml:"MainCity,omitempty"`

	MainState string `xml:"MainState,omitempty" json:"MainState,omitempty" yaml:"MainState,omitempty"`

	MainZip string `xml:"MainZip,omitempty" json:"MainZip,omitempty" yaml:"MainZip,omitempty"`

	MainCountry string `xml:"MainCountry,omitempty" json:"MainCountry,omitempty" yaml:"MainCountry,omitempty"`

	TaxID string `xml:"TaxID,omitempty" json:"TaxID,omitempty" yaml:"TaxID,omitempty"`

	CustomerTypes *ArrayOfInt `xml:"CustomerTypes,omitempty" json:"CustomerTypes,omitempty" yaml:"CustomerTypes,omitempty"`

	CustomerStatuses *ArrayOfInt `xml:"CustomerStatuses,omitempty" json:"CustomerStatuses,omitempty" yaml:"CustomerStatuses,omitempty"`

	EnrollerID int32 `xml:"EnrollerID,omitempty" json:"EnrollerID,omitempty" yaml:"EnrollerID,omitempty"`

	SponsorID int32 `xml:"SponsorID,omitempty" json:"SponsorID,omitempty" yaml:"SponsorID,omitempty"`

	Field1 string `xml:"Field1,omitempty" json:"Field1,omitempty" yaml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty" json:"Field2,omitempty" yaml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty" json:"Field3,omitempty" yaml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty" json:"Field4,omitempty" yaml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty" json:"Field5,omitempty" yaml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty" json:"Field6,omitempty" yaml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty" json:"Field7,omitempty" yaml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty" json:"Field8,omitempty" yaml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty" json:"Field9,omitempty" yaml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty" json:"Field10,omitempty" yaml:"Field10,omitempty"`

	Field11 string `xml:"Field11,omitempty" json:"Field11,omitempty" yaml:"Field11,omitempty"`

	Field12 string `xml:"Field12,omitempty" json:"Field12,omitempty" yaml:"Field12,omitempty"`

	Field13 string `xml:"Field13,omitempty" json:"Field13,omitempty" yaml:"Field13,omitempty"`

	Field14 string `xml:"Field14,omitempty" json:"Field14,omitempty" yaml:"Field14,omitempty"`

	Field15 string `xml:"Field15,omitempty" json:"Field15,omitempty" yaml:"Field15,omitempty"`

	CreatedDateStart time.Time `xml:"CreatedDateStart,omitempty" json:"CreatedDateStart,omitempty" yaml:"CreatedDateStart,omitempty"`

	CreatedDateEnd time.Time `xml:"CreatedDateEnd,omitempty" json:"CreatedDateEnd,omitempty" yaml:"CreatedDateEnd,omitempty"`

	GreaterThanCustomerID int32 `xml:"GreaterThanCustomerID,omitempty" json:"GreaterThanCustomerID,omitempty" yaml:"GreaterThanCustomerID,omitempty"`

	GreaterThanModifiedDate time.Time `xml:"GreaterThanModifiedDate,omitempty" json:"GreaterThanModifiedDate,omitempty" yaml:"GreaterThanModifiedDate,omitempty"`

	BatchSize int32 `xml:"BatchSize,omitempty" json:"BatchSize,omitempty" yaml:"BatchSize,omitempty"`
}

type GetCustomersResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomersResult"`

	*ApiResponse

	Customers *ArrayOfCustomerResponse `xml:"Customers,omitempty" json:"Customers,omitempty" yaml:"Customers,omitempty"`

	RecordCount int32 `xml:"RecordCount,omitempty" json:"RecordCount,omitempty" yaml:"RecordCount,omitempty"`
}

type ArrayOfCustomerResponse struct {
	CustomerResponse []*CustomerResponse `xml:"CustomerResponse,omitempty" json:"CustomerResponse,omitempty" yaml:"CustomerResponse,omitempty"`
}

type CustomerResponse struct {
	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`

	Company string `xml:"Company,omitempty" json:"Company,omitempty" yaml:"Company,omitempty"`

	CustomerType int32 `xml:"CustomerType,omitempty" json:"CustomerType,omitempty" yaml:"CustomerType,omitempty"`

	CustomerStatus int32 `xml:"CustomerStatus,omitempty" json:"CustomerStatus,omitempty" yaml:"CustomerStatus,omitempty"`

	Email string `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`

	Phone string `xml:"Phone,omitempty" json:"Phone,omitempty" yaml:"Phone,omitempty"`

	Phone2 string `xml:"Phone2,omitempty" json:"Phone2,omitempty" yaml:"Phone2,omitempty"`

	MobilePhone string `xml:"MobilePhone,omitempty" json:"MobilePhone,omitempty" yaml:"MobilePhone,omitempty"`

	Fax string `xml:"Fax,omitempty" json:"Fax,omitempty" yaml:"Fax,omitempty"`

	MainAddress1 string `xml:"MainAddress1,omitempty" json:"MainAddress1,omitempty" yaml:"MainAddress1,omitempty"`

	MainAddress2 string `xml:"MainAddress2,omitempty" json:"MainAddress2,omitempty" yaml:"MainAddress2,omitempty"`

	MainCity string `xml:"MainCity,omitempty" json:"MainCity,omitempty" yaml:"MainCity,omitempty"`

	MainState string `xml:"MainState,omitempty" json:"MainState,omitempty" yaml:"MainState,omitempty"`

	MainZip string `xml:"MainZip,omitempty" json:"MainZip,omitempty" yaml:"MainZip,omitempty"`

	MainCountry string `xml:"MainCountry,omitempty" json:"MainCountry,omitempty" yaml:"MainCountry,omitempty"`

	MainCounty string `xml:"MainCounty,omitempty" json:"MainCounty,omitempty" yaml:"MainCounty,omitempty"`

	MailAddress1 string `xml:"MailAddress1,omitempty" json:"MailAddress1,omitempty" yaml:"MailAddress1,omitempty"`

	MailAddress2 string `xml:"MailAddress2,omitempty" json:"MailAddress2,omitempty" yaml:"MailAddress2,omitempty"`

	MailCity string `xml:"MailCity,omitempty" json:"MailCity,omitempty" yaml:"MailCity,omitempty"`

	MailState string `xml:"MailState,omitempty" json:"MailState,omitempty" yaml:"MailState,omitempty"`

	MailZip string `xml:"MailZip,omitempty" json:"MailZip,omitempty" yaml:"MailZip,omitempty"`

	MailCountry string `xml:"MailCountry,omitempty" json:"MailCountry,omitempty" yaml:"MailCountry,omitempty"`

	MailCounty string `xml:"MailCounty,omitempty" json:"MailCounty,omitempty" yaml:"MailCounty,omitempty"`

	OtherAddress1 string `xml:"OtherAddress1,omitempty" json:"OtherAddress1,omitempty" yaml:"OtherAddress1,omitempty"`

	OtherAddress2 string `xml:"OtherAddress2,omitempty" json:"OtherAddress2,omitempty" yaml:"OtherAddress2,omitempty"`

	OtherCity string `xml:"OtherCity,omitempty" json:"OtherCity,omitempty" yaml:"OtherCity,omitempty"`

	OtherState string `xml:"OtherState,omitempty" json:"OtherState,omitempty" yaml:"OtherState,omitempty"`

	OtherZip string `xml:"OtherZip,omitempty" json:"OtherZip,omitempty" yaml:"OtherZip,omitempty"`

	OtherCountry string `xml:"OtherCountry,omitempty" json:"OtherCountry,omitempty" yaml:"OtherCountry,omitempty"`

	OtherCounty string `xml:"OtherCounty,omitempty" json:"OtherCounty,omitempty" yaml:"OtherCounty,omitempty"`

	LoginName string `xml:"LoginName,omitempty" json:"LoginName,omitempty" yaml:"LoginName,omitempty"`

	EnrollerID int32 `xml:"EnrollerID,omitempty" json:"EnrollerID,omitempty" yaml:"EnrollerID,omitempty"`

	SponsorID int32 `xml:"SponsorID,omitempty" json:"SponsorID,omitempty" yaml:"SponsorID,omitempty"`

	RankID int32 `xml:"RankID,omitempty" json:"RankID,omitempty" yaml:"RankID,omitempty"`

	BirthDate time.Time `xml:"BirthDate,omitempty" json:"BirthDate,omitempty" yaml:"BirthDate,omitempty"`

	Field1 string `xml:"Field1,omitempty" json:"Field1,omitempty" yaml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty" json:"Field2,omitempty" yaml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty" json:"Field3,omitempty" yaml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty" json:"Field4,omitempty" yaml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty" json:"Field5,omitempty" yaml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty" json:"Field6,omitempty" yaml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty" json:"Field7,omitempty" yaml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty" json:"Field8,omitempty" yaml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty" json:"Field9,omitempty" yaml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty" json:"Field10,omitempty" yaml:"Field10,omitempty"`

	Field11 string `xml:"Field11,omitempty" json:"Field11,omitempty" yaml:"Field11,omitempty"`

	Field12 string `xml:"Field12,omitempty" json:"Field12,omitempty" yaml:"Field12,omitempty"`

	Field13 string `xml:"Field13,omitempty" json:"Field13,omitempty" yaml:"Field13,omitempty"`

	Field14 string `xml:"Field14,omitempty" json:"Field14,omitempty" yaml:"Field14,omitempty"`

	Field15 string `xml:"Field15,omitempty" json:"Field15,omitempty" yaml:"Field15,omitempty"`

	Date1 time.Time `xml:"Date1,omitempty" json:"Date1,omitempty" yaml:"Date1,omitempty"`

	Date2 time.Time `xml:"Date2,omitempty" json:"Date2,omitempty" yaml:"Date2,omitempty"`

	Date3 time.Time `xml:"Date3,omitempty" json:"Date3,omitempty" yaml:"Date3,omitempty"`

	Date4 time.Time `xml:"Date4,omitempty" json:"Date4,omitempty" yaml:"Date4,omitempty"`

	Date5 time.Time `xml:"Date5,omitempty" json:"Date5,omitempty" yaml:"Date5,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty" yaml:"CurrencyCode,omitempty"`

	PayableToName string `xml:"PayableToName,omitempty" json:"PayableToName,omitempty" yaml:"PayableToName,omitempty"`

	DefaultWarehouseID int32 `xml:"DefaultWarehouseID,omitempty" json:"DefaultWarehouseID,omitempty" yaml:"DefaultWarehouseID,omitempty"`

	PayableType *PayableType `xml:"PayableType,omitempty" json:"PayableType,omitempty" yaml:"PayableType,omitempty"`

	CheckThreshold float64 `xml:"CheckThreshold,omitempty" json:"CheckThreshold,omitempty" yaml:"CheckThreshold,omitempty"`

	PriceType int32 `xml:"PriceType,omitempty" json:"PriceType,omitempty" yaml:"PriceType,omitempty"`

	LanguageID int32 `xml:"LanguageID,omitempty" json:"LanguageID,omitempty" yaml:"LanguageID,omitempty"`

	Gender *Gender `xml:"Gender,omitempty" json:"Gender,omitempty" yaml:"Gender,omitempty"`

	SalesTaxID string `xml:"SalesTaxID,omitempty" json:"SalesTaxID,omitempty" yaml:"SalesTaxID,omitempty"`

	VatRegistration string `xml:"VatRegistration,omitempty" json:"VatRegistration,omitempty" yaml:"VatRegistration,omitempty"`

	IsSalesTaxExempt bool `xml:"IsSalesTaxExempt,omitempty" json:"IsSalesTaxExempt,omitempty" yaml:"IsSalesTaxExempt,omitempty"`

	IsSubscribedToBroadcasts bool `xml:"IsSubscribedToBroadcasts,omitempty" json:"IsSubscribedToBroadcasts,omitempty" yaml:"IsSubscribedToBroadcasts,omitempty"`

	CreatedDate time.Time `xml:"CreatedDate,omitempty" json:"CreatedDate,omitempty" yaml:"CreatedDate,omitempty"`

	ModifiedDate time.Time `xml:"ModifiedDate,omitempty" json:"ModifiedDate,omitempty" yaml:"ModifiedDate,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty" json:"MiddleName,omitempty" yaml:"MiddleName,omitempty"`

	NameSuffix string `xml:"NameSuffix,omitempty" json:"NameSuffix,omitempty" yaml:"NameSuffix,omitempty"`

	MainAddress3 string `xml:"MainAddress3,omitempty" json:"MainAddress3,omitempty" yaml:"MainAddress3,omitempty"`

	MailAddress3 string `xml:"MailAddress3,omitempty" json:"MailAddress3,omitempty" yaml:"MailAddress3,omitempty"`

	OtherAddress3 string `xml:"OtherAddress3,omitempty" json:"OtherAddress3,omitempty" yaml:"OtherAddress3,omitempty"`

	BinaryPlacementPreference int32 `xml:"BinaryPlacementPreference,omitempty" json:"BinaryPlacementPreference,omitempty" yaml:"BinaryPlacementPreference,omitempty"`

	UseBinaryHoldingTank bool `xml:"UseBinaryHoldingTank,omitempty" json:"UseBinaryHoldingTank,omitempty" yaml:"UseBinaryHoldingTank,omitempty"`

	MainAddressVerified bool `xml:"MainAddressVerified,omitempty" json:"MainAddressVerified,omitempty" yaml:"MainAddressVerified,omitempty"`

	MailAddressVerified bool `xml:"MailAddressVerified,omitempty" json:"MailAddressVerified,omitempty" yaml:"MailAddressVerified,omitempty"`

	OtherAddressVerified bool `xml:"OtherAddressVerified,omitempty" json:"OtherAddressVerified,omitempty" yaml:"OtherAddressVerified,omitempty"`
}

type CreateWarehouseRequest struct {
	*ApiRequest

	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	Address1 string `xml:"Address1,omitempty" json:"Address1,omitempty" yaml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty" json:"Address2,omitempty" yaml:"Address2,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty" json:"Zip,omitempty" yaml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`

	Currencies *ArrayOfString `xml:"Currencies,omitempty" json:"Currencies,omitempty" yaml:"Currencies,omitempty"`
}

type CreateWarehouseResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateWarehouseResult"`

	*ApiResponse

	Warehouse *WarehouseResponse `xml:"Warehouse,omitempty" json:"Warehouse,omitempty" yaml:"Warehouse,omitempty"`
}

type WarehouseResponse struct {
	WarehouseID int32 `xml:"WarehouseID,omitempty" json:"WarehouseID,omitempty" yaml:"WarehouseID,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	Address1 string `xml:"Address1,omitempty" json:"Address1,omitempty" yaml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty" json:"Address2,omitempty" yaml:"Address2,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty" json:"Zip,omitempty" yaml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`
}

type GetCustomerNotesRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`
}

type GetCustomerNotesResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomerNotesResult"`

	*ApiResponse

	CustomerNotes *ArrayOfCustomerNotesResponse `xml:"CustomerNotes,omitempty" json:"CustomerNotes,omitempty" yaml:"CustomerNotes,omitempty"`
}

type ArrayOfCustomerNotesResponse struct {
	CustomerNotesResponse []*CustomerNotesResponse `xml:"CustomerNotesResponse,omitempty" json:"CustomerNotesResponse,omitempty" yaml:"CustomerNotesResponse,omitempty"`
}

type CustomerNotesResponse struct {
	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	Notes string `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`
}

type AppendCustomerNotesRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	Notes string `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`
}

type AppendCustomerNotesResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ AppendCustomerNotesResult"`

	*ApiResponse

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`
}

type GetVolumesRequest struct {
	*ApiRequest

	PeriodType int32 `xml:"PeriodType,omitempty" json:"PeriodType,omitempty" yaml:"PeriodType,omitempty"`

	PeriodID int32 `xml:"PeriodID,omitempty" json:"PeriodID,omitempty" yaml:"PeriodID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`
}

type GetVolumesResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetVolumesResult"`

	*ApiResponse

	Volumes *ArrayOfVolumeResponse `xml:"Volumes,omitempty" json:"Volumes,omitempty" yaml:"Volumes,omitempty"`
}

type ArrayOfVolumeResponse struct {
	VolumeResponse []*VolumeResponse `xml:"VolumeResponse,omitempty" json:"VolumeResponse,omitempty" yaml:"VolumeResponse,omitempty"`
}

type VolumeResponse struct {
	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	PeriodType int32 `xml:"PeriodType,omitempty" json:"PeriodType,omitempty" yaml:"PeriodType,omitempty"`

	PeriodID int32 `xml:"PeriodID,omitempty" json:"PeriodID,omitempty" yaml:"PeriodID,omitempty"`

	PeriodDescription string `xml:"PeriodDescription,omitempty" json:"PeriodDescription,omitempty" yaml:"PeriodDescription,omitempty"`

	Volume1 float64 `xml:"Volume1,omitempty" json:"Volume1,omitempty" yaml:"Volume1,omitempty"`

	Volume2 float64 `xml:"Volume2,omitempty" json:"Volume2,omitempty" yaml:"Volume2,omitempty"`

	Volume3 float64 `xml:"Volume3,omitempty" json:"Volume3,omitempty" yaml:"Volume3,omitempty"`

	Volume4 float64 `xml:"Volume4,omitempty" json:"Volume4,omitempty" yaml:"Volume4,omitempty"`

	Volume5 float64 `xml:"Volume5,omitempty" json:"Volume5,omitempty" yaml:"Volume5,omitempty"`

	Volume6 float64 `xml:"Volume6,omitempty" json:"Volume6,omitempty" yaml:"Volume6,omitempty"`

	Volume7 float64 `xml:"Volume7,omitempty" json:"Volume7,omitempty" yaml:"Volume7,omitempty"`

	Volume8 float64 `xml:"Volume8,omitempty" json:"Volume8,omitempty" yaml:"Volume8,omitempty"`

	Volume9 float64 `xml:"Volume9,omitempty" json:"Volume9,omitempty" yaml:"Volume9,omitempty"`

	Volume10 float64 `xml:"Volume10,omitempty" json:"Volume10,omitempty" yaml:"Volume10,omitempty"`

	Volume11 float64 `xml:"Volume11,omitempty" json:"Volume11,omitempty" yaml:"Volume11,omitempty"`

	Volume12 float64 `xml:"Volume12,omitempty" json:"Volume12,omitempty" yaml:"Volume12,omitempty"`

	Volume13 float64 `xml:"Volume13,omitempty" json:"Volume13,omitempty" yaml:"Volume13,omitempty"`

	Volume14 float64 `xml:"Volume14,omitempty" json:"Volume14,omitempty" yaml:"Volume14,omitempty"`

	Volume15 float64 `xml:"Volume15,omitempty" json:"Volume15,omitempty" yaml:"Volume15,omitempty"`

	Volume16 float64 `xml:"Volume16,omitempty" json:"Volume16,omitempty" yaml:"Volume16,omitempty"`

	Volume17 float64 `xml:"Volume17,omitempty" json:"Volume17,omitempty" yaml:"Volume17,omitempty"`

	Volume18 float64 `xml:"Volume18,omitempty" json:"Volume18,omitempty" yaml:"Volume18,omitempty"`

	Volume19 float64 `xml:"Volume19,omitempty" json:"Volume19,omitempty" yaml:"Volume19,omitempty"`

	Volume20 float64 `xml:"Volume20,omitempty" json:"Volume20,omitempty" yaml:"Volume20,omitempty"`

	Volume21 float64 `xml:"Volume21,omitempty" json:"Volume21,omitempty" yaml:"Volume21,omitempty"`

	Volume22 float64 `xml:"Volume22,omitempty" json:"Volume22,omitempty" yaml:"Volume22,omitempty"`

	Volume23 float64 `xml:"Volume23,omitempty" json:"Volume23,omitempty" yaml:"Volume23,omitempty"`

	Volume24 float64 `xml:"Volume24,omitempty" json:"Volume24,omitempty" yaml:"Volume24,omitempty"`

	Volume25 float64 `xml:"Volume25,omitempty" json:"Volume25,omitempty" yaml:"Volume25,omitempty"`

	Volume26 float64 `xml:"Volume26,omitempty" json:"Volume26,omitempty" yaml:"Volume26,omitempty"`

	Volume27 float64 `xml:"Volume27,omitempty" json:"Volume27,omitempty" yaml:"Volume27,omitempty"`

	Volume28 float64 `xml:"Volume28,omitempty" json:"Volume28,omitempty" yaml:"Volume28,omitempty"`

	Volume29 float64 `xml:"Volume29,omitempty" json:"Volume29,omitempty" yaml:"Volume29,omitempty"`

	Volume30 float64 `xml:"Volume30,omitempty" json:"Volume30,omitempty" yaml:"Volume30,omitempty"`

	Volume31 float64 `xml:"Volume31,omitempty" json:"Volume31,omitempty" yaml:"Volume31,omitempty"`

	Volume32 float64 `xml:"Volume32,omitempty" json:"Volume32,omitempty" yaml:"Volume32,omitempty"`

	Volume33 float64 `xml:"Volume33,omitempty" json:"Volume33,omitempty" yaml:"Volume33,omitempty"`

	Volume34 float64 `xml:"Volume34,omitempty" json:"Volume34,omitempty" yaml:"Volume34,omitempty"`

	Volume35 float64 `xml:"Volume35,omitempty" json:"Volume35,omitempty" yaml:"Volume35,omitempty"`

	Volume36 float64 `xml:"Volume36,omitempty" json:"Volume36,omitempty" yaml:"Volume36,omitempty"`

	Volume37 float64 `xml:"Volume37,omitempty" json:"Volume37,omitempty" yaml:"Volume37,omitempty"`

	Volume38 float64 `xml:"Volume38,omitempty" json:"Volume38,omitempty" yaml:"Volume38,omitempty"`

	Volume39 float64 `xml:"Volume39,omitempty" json:"Volume39,omitempty" yaml:"Volume39,omitempty"`

	Volume40 float64 `xml:"Volume40,omitempty" json:"Volume40,omitempty" yaml:"Volume40,omitempty"`

	Volume41 float64 `xml:"Volume41,omitempty" json:"Volume41,omitempty" yaml:"Volume41,omitempty"`

	Volume42 float64 `xml:"Volume42,omitempty" json:"Volume42,omitempty" yaml:"Volume42,omitempty"`

	Volume43 float64 `xml:"Volume43,omitempty" json:"Volume43,omitempty" yaml:"Volume43,omitempty"`

	Volume44 float64 `xml:"Volume44,omitempty" json:"Volume44,omitempty" yaml:"Volume44,omitempty"`

	Volume45 float64 `xml:"Volume45,omitempty" json:"Volume45,omitempty" yaml:"Volume45,omitempty"`

	Volume46 float64 `xml:"Volume46,omitempty" json:"Volume46,omitempty" yaml:"Volume46,omitempty"`

	Volume47 float64 `xml:"Volume47,omitempty" json:"Volume47,omitempty" yaml:"Volume47,omitempty"`

	Volume48 float64 `xml:"Volume48,omitempty" json:"Volume48,omitempty" yaml:"Volume48,omitempty"`

	Volume49 float64 `xml:"Volume49,omitempty" json:"Volume49,omitempty" yaml:"Volume49,omitempty"`

	Volume50 float64 `xml:"Volume50,omitempty" json:"Volume50,omitempty" yaml:"Volume50,omitempty"`

	Volume51 float64 `xml:"Volume51,omitempty" json:"Volume51,omitempty" yaml:"Volume51,omitempty"`

	Volume52 float64 `xml:"Volume52,omitempty" json:"Volume52,omitempty" yaml:"Volume52,omitempty"`

	Volume53 float64 `xml:"Volume53,omitempty" json:"Volume53,omitempty" yaml:"Volume53,omitempty"`

	Volume54 float64 `xml:"Volume54,omitempty" json:"Volume54,omitempty" yaml:"Volume54,omitempty"`

	Volume55 float64 `xml:"Volume55,omitempty" json:"Volume55,omitempty" yaml:"Volume55,omitempty"`

	Volume56 float64 `xml:"Volume56,omitempty" json:"Volume56,omitempty" yaml:"Volume56,omitempty"`

	Volume57 float64 `xml:"Volume57,omitempty" json:"Volume57,omitempty" yaml:"Volume57,omitempty"`

	Volume58 float64 `xml:"Volume58,omitempty" json:"Volume58,omitempty" yaml:"Volume58,omitempty"`

	Volume59 float64 `xml:"Volume59,omitempty" json:"Volume59,omitempty" yaml:"Volume59,omitempty"`

	Volume60 float64 `xml:"Volume60,omitempty" json:"Volume60,omitempty" yaml:"Volume60,omitempty"`

	Volume61 float64 `xml:"Volume61,omitempty" json:"Volume61,omitempty" yaml:"Volume61,omitempty"`

	Volume62 float64 `xml:"Volume62,omitempty" json:"Volume62,omitempty" yaml:"Volume62,omitempty"`

	Volume63 float64 `xml:"Volume63,omitempty" json:"Volume63,omitempty" yaml:"Volume63,omitempty"`

	Volume64 float64 `xml:"Volume64,omitempty" json:"Volume64,omitempty" yaml:"Volume64,omitempty"`

	Volume65 float64 `xml:"Volume65,omitempty" json:"Volume65,omitempty" yaml:"Volume65,omitempty"`

	Volume66 float64 `xml:"Volume66,omitempty" json:"Volume66,omitempty" yaml:"Volume66,omitempty"`

	Volume67 float64 `xml:"Volume67,omitempty" json:"Volume67,omitempty" yaml:"Volume67,omitempty"`

	Volume68 float64 `xml:"Volume68,omitempty" json:"Volume68,omitempty" yaml:"Volume68,omitempty"`

	Volume69 float64 `xml:"Volume69,omitempty" json:"Volume69,omitempty" yaml:"Volume69,omitempty"`

	Volume70 float64 `xml:"Volume70,omitempty" json:"Volume70,omitempty" yaml:"Volume70,omitempty"`

	Volume71 float64 `xml:"Volume71,omitempty" json:"Volume71,omitempty" yaml:"Volume71,omitempty"`

	Volume72 float64 `xml:"Volume72,omitempty" json:"Volume72,omitempty" yaml:"Volume72,omitempty"`

	Volume73 float64 `xml:"Volume73,omitempty" json:"Volume73,omitempty" yaml:"Volume73,omitempty"`

	Volume74 float64 `xml:"Volume74,omitempty" json:"Volume74,omitempty" yaml:"Volume74,omitempty"`

	Volume75 float64 `xml:"Volume75,omitempty" json:"Volume75,omitempty" yaml:"Volume75,omitempty"`

	Volume76 float64 `xml:"Volume76,omitempty" json:"Volume76,omitempty" yaml:"Volume76,omitempty"`

	Volume77 float64 `xml:"Volume77,omitempty" json:"Volume77,omitempty" yaml:"Volume77,omitempty"`

	Volume78 float64 `xml:"Volume78,omitempty" json:"Volume78,omitempty" yaml:"Volume78,omitempty"`

	Volume79 float64 `xml:"Volume79,omitempty" json:"Volume79,omitempty" yaml:"Volume79,omitempty"`

	Volume80 float64 `xml:"Volume80,omitempty" json:"Volume80,omitempty" yaml:"Volume80,omitempty"`

	Volume81 float64 `xml:"Volume81,omitempty" json:"Volume81,omitempty" yaml:"Volume81,omitempty"`

	Volume82 float64 `xml:"Volume82,omitempty" json:"Volume82,omitempty" yaml:"Volume82,omitempty"`

	Volume83 float64 `xml:"Volume83,omitempty" json:"Volume83,omitempty" yaml:"Volume83,omitempty"`

	Volume84 float64 `xml:"Volume84,omitempty" json:"Volume84,omitempty" yaml:"Volume84,omitempty"`

	Volume85 float64 `xml:"Volume85,omitempty" json:"Volume85,omitempty" yaml:"Volume85,omitempty"`

	Volume86 float64 `xml:"Volume86,omitempty" json:"Volume86,omitempty" yaml:"Volume86,omitempty"`

	Volume87 float64 `xml:"Volume87,omitempty" json:"Volume87,omitempty" yaml:"Volume87,omitempty"`

	Volume88 float64 `xml:"Volume88,omitempty" json:"Volume88,omitempty" yaml:"Volume88,omitempty"`

	Volume89 float64 `xml:"Volume89,omitempty" json:"Volume89,omitempty" yaml:"Volume89,omitempty"`

	Volume90 float64 `xml:"Volume90,omitempty" json:"Volume90,omitempty" yaml:"Volume90,omitempty"`

	Volume91 float64 `xml:"Volume91,omitempty" json:"Volume91,omitempty" yaml:"Volume91,omitempty"`

	Volume92 float64 `xml:"Volume92,omitempty" json:"Volume92,omitempty" yaml:"Volume92,omitempty"`

	Volume93 float64 `xml:"Volume93,omitempty" json:"Volume93,omitempty" yaml:"Volume93,omitempty"`

	Volume94 float64 `xml:"Volume94,omitempty" json:"Volume94,omitempty" yaml:"Volume94,omitempty"`

	Volume95 float64 `xml:"Volume95,omitempty" json:"Volume95,omitempty" yaml:"Volume95,omitempty"`

	Volume96 float64 `xml:"Volume96,omitempty" json:"Volume96,omitempty" yaml:"Volume96,omitempty"`

	Volume97 float64 `xml:"Volume97,omitempty" json:"Volume97,omitempty" yaml:"Volume97,omitempty"`

	Volume98 float64 `xml:"Volume98,omitempty" json:"Volume98,omitempty" yaml:"Volume98,omitempty"`

	Volume99 float64 `xml:"Volume99,omitempty" json:"Volume99,omitempty" yaml:"Volume99,omitempty"`

	Volume100 float64 `xml:"Volume100,omitempty" json:"Volume100,omitempty" yaml:"Volume100,omitempty"`

	RankID int32 `xml:"RankID,omitempty" json:"RankID,omitempty" yaml:"RankID,omitempty"`

	PaidRankID int32 `xml:"PaidRankID,omitempty" json:"PaidRankID,omitempty" yaml:"PaidRankID,omitempty"`

	Volume101 float64 `xml:"Volume101,omitempty" json:"Volume101,omitempty" yaml:"Volume101,omitempty"`

	Volume102 float64 `xml:"Volume102,omitempty" json:"Volume102,omitempty" yaml:"Volume102,omitempty"`

	Volume103 float64 `xml:"Volume103,omitempty" json:"Volume103,omitempty" yaml:"Volume103,omitempty"`

	Volume104 float64 `xml:"Volume104,omitempty" json:"Volume104,omitempty" yaml:"Volume104,omitempty"`

	Volume105 float64 `xml:"Volume105,omitempty" json:"Volume105,omitempty" yaml:"Volume105,omitempty"`

	Volume106 float64 `xml:"Volume106,omitempty" json:"Volume106,omitempty" yaml:"Volume106,omitempty"`

	Volume107 float64 `xml:"Volume107,omitempty" json:"Volume107,omitempty" yaml:"Volume107,omitempty"`

	Volume108 float64 `xml:"Volume108,omitempty" json:"Volume108,omitempty" yaml:"Volume108,omitempty"`

	Volume109 float64 `xml:"Volume109,omitempty" json:"Volume109,omitempty" yaml:"Volume109,omitempty"`

	Volume110 float64 `xml:"Volume110,omitempty" json:"Volume110,omitempty" yaml:"Volume110,omitempty"`

	Volume111 float64 `xml:"Volume111,omitempty" json:"Volume111,omitempty" yaml:"Volume111,omitempty"`

	Volume112 float64 `xml:"Volume112,omitempty" json:"Volume112,omitempty" yaml:"Volume112,omitempty"`

	Volume113 float64 `xml:"Volume113,omitempty" json:"Volume113,omitempty" yaml:"Volume113,omitempty"`

	Volume114 float64 `xml:"Volume114,omitempty" json:"Volume114,omitempty" yaml:"Volume114,omitempty"`

	Volume115 float64 `xml:"Volume115,omitempty" json:"Volume115,omitempty" yaml:"Volume115,omitempty"`

	Volume116 float64 `xml:"Volume116,omitempty" json:"Volume116,omitempty" yaml:"Volume116,omitempty"`

	Volume117 float64 `xml:"Volume117,omitempty" json:"Volume117,omitempty" yaml:"Volume117,omitempty"`

	Volume118 float64 `xml:"Volume118,omitempty" json:"Volume118,omitempty" yaml:"Volume118,omitempty"`

	Volume119 float64 `xml:"Volume119,omitempty" json:"Volume119,omitempty" yaml:"Volume119,omitempty"`

	Volume120 float64 `xml:"Volume120,omitempty" json:"Volume120,omitempty" yaml:"Volume120,omitempty"`

	Volume121 float64 `xml:"Volume121,omitempty" json:"Volume121,omitempty" yaml:"Volume121,omitempty"`

	Volume122 float64 `xml:"Volume122,omitempty" json:"Volume122,omitempty" yaml:"Volume122,omitempty"`

	Volume123 float64 `xml:"Volume123,omitempty" json:"Volume123,omitempty" yaml:"Volume123,omitempty"`

	Volume124 float64 `xml:"Volume124,omitempty" json:"Volume124,omitempty" yaml:"Volume124,omitempty"`

	Volume125 float64 `xml:"Volume125,omitempty" json:"Volume125,omitempty" yaml:"Volume125,omitempty"`

	Volume126 float64 `xml:"Volume126,omitempty" json:"Volume126,omitempty" yaml:"Volume126,omitempty"`

	Volume127 float64 `xml:"Volume127,omitempty" json:"Volume127,omitempty" yaml:"Volume127,omitempty"`

	Volume128 float64 `xml:"Volume128,omitempty" json:"Volume128,omitempty" yaml:"Volume128,omitempty"`

	Volume129 float64 `xml:"Volume129,omitempty" json:"Volume129,omitempty" yaml:"Volume129,omitempty"`

	Volume130 float64 `xml:"Volume130,omitempty" json:"Volume130,omitempty" yaml:"Volume130,omitempty"`

	Volume131 float64 `xml:"Volume131,omitempty" json:"Volume131,omitempty" yaml:"Volume131,omitempty"`

	Volume132 float64 `xml:"Volume132,omitempty" json:"Volume132,omitempty" yaml:"Volume132,omitempty"`

	Volume133 float64 `xml:"Volume133,omitempty" json:"Volume133,omitempty" yaml:"Volume133,omitempty"`

	Volume134 float64 `xml:"Volume134,omitempty" json:"Volume134,omitempty" yaml:"Volume134,omitempty"`

	Volume135 float64 `xml:"Volume135,omitempty" json:"Volume135,omitempty" yaml:"Volume135,omitempty"`

	Volume136 float64 `xml:"Volume136,omitempty" json:"Volume136,omitempty" yaml:"Volume136,omitempty"`

	Volume137 float64 `xml:"Volume137,omitempty" json:"Volume137,omitempty" yaml:"Volume137,omitempty"`

	Volume138 float64 `xml:"Volume138,omitempty" json:"Volume138,omitempty" yaml:"Volume138,omitempty"`

	Volume139 float64 `xml:"Volume139,omitempty" json:"Volume139,omitempty" yaml:"Volume139,omitempty"`

	Volume140 float64 `xml:"Volume140,omitempty" json:"Volume140,omitempty" yaml:"Volume140,omitempty"`

	Volume141 float64 `xml:"Volume141,omitempty" json:"Volume141,omitempty" yaml:"Volume141,omitempty"`

	Volume142 float64 `xml:"Volume142,omitempty" json:"Volume142,omitempty" yaml:"Volume142,omitempty"`

	Volume143 float64 `xml:"Volume143,omitempty" json:"Volume143,omitempty" yaml:"Volume143,omitempty"`

	Volume144 float64 `xml:"Volume144,omitempty" json:"Volume144,omitempty" yaml:"Volume144,omitempty"`

	Volume145 float64 `xml:"Volume145,omitempty" json:"Volume145,omitempty" yaml:"Volume145,omitempty"`

	Volume146 float64 `xml:"Volume146,omitempty" json:"Volume146,omitempty" yaml:"Volume146,omitempty"`

	Volume147 float64 `xml:"Volume147,omitempty" json:"Volume147,omitempty" yaml:"Volume147,omitempty"`

	Volume148 float64 `xml:"Volume148,omitempty" json:"Volume148,omitempty" yaml:"Volume148,omitempty"`

	Volume149 float64 `xml:"Volume149,omitempty" json:"Volume149,omitempty" yaml:"Volume149,omitempty"`

	Volume150 float64 `xml:"Volume150,omitempty" json:"Volume150,omitempty" yaml:"Volume150,omitempty"`

	Volume151 float64 `xml:"Volume151,omitempty" json:"Volume151,omitempty" yaml:"Volume151,omitempty"`

	Volume152 float64 `xml:"Volume152,omitempty" json:"Volume152,omitempty" yaml:"Volume152,omitempty"`

	Volume153 float64 `xml:"Volume153,omitempty" json:"Volume153,omitempty" yaml:"Volume153,omitempty"`

	Volume154 float64 `xml:"Volume154,omitempty" json:"Volume154,omitempty" yaml:"Volume154,omitempty"`

	Volume155 float64 `xml:"Volume155,omitempty" json:"Volume155,omitempty" yaml:"Volume155,omitempty"`

	Volume156 float64 `xml:"Volume156,omitempty" json:"Volume156,omitempty" yaml:"Volume156,omitempty"`

	Volume157 float64 `xml:"Volume157,omitempty" json:"Volume157,omitempty" yaml:"Volume157,omitempty"`

	Volume158 float64 `xml:"Volume158,omitempty" json:"Volume158,omitempty" yaml:"Volume158,omitempty"`

	Volume159 float64 `xml:"Volume159,omitempty" json:"Volume159,omitempty" yaml:"Volume159,omitempty"`

	Volume160 float64 `xml:"Volume160,omitempty" json:"Volume160,omitempty" yaml:"Volume160,omitempty"`

	Volume161 float64 `xml:"Volume161,omitempty" json:"Volume161,omitempty" yaml:"Volume161,omitempty"`

	Volume162 float64 `xml:"Volume162,omitempty" json:"Volume162,omitempty" yaml:"Volume162,omitempty"`

	Volume163 float64 `xml:"Volume163,omitempty" json:"Volume163,omitempty" yaml:"Volume163,omitempty"`

	Volume164 float64 `xml:"Volume164,omitempty" json:"Volume164,omitempty" yaml:"Volume164,omitempty"`

	Volume165 float64 `xml:"Volume165,omitempty" json:"Volume165,omitempty" yaml:"Volume165,omitempty"`

	Volume166 float64 `xml:"Volume166,omitempty" json:"Volume166,omitempty" yaml:"Volume166,omitempty"`

	Volume167 float64 `xml:"Volume167,omitempty" json:"Volume167,omitempty" yaml:"Volume167,omitempty"`

	Volume168 float64 `xml:"Volume168,omitempty" json:"Volume168,omitempty" yaml:"Volume168,omitempty"`

	Volume169 float64 `xml:"Volume169,omitempty" json:"Volume169,omitempty" yaml:"Volume169,omitempty"`

	Volume170 float64 `xml:"Volume170,omitempty" json:"Volume170,omitempty" yaml:"Volume170,omitempty"`

	Volume171 float64 `xml:"Volume171,omitempty" json:"Volume171,omitempty" yaml:"Volume171,omitempty"`

	Volume172 float64 `xml:"Volume172,omitempty" json:"Volume172,omitempty" yaml:"Volume172,omitempty"`

	Volume173 float64 `xml:"Volume173,omitempty" json:"Volume173,omitempty" yaml:"Volume173,omitempty"`

	Volume174 float64 `xml:"Volume174,omitempty" json:"Volume174,omitempty" yaml:"Volume174,omitempty"`

	Volume175 float64 `xml:"Volume175,omitempty" json:"Volume175,omitempty" yaml:"Volume175,omitempty"`

	Volume176 float64 `xml:"Volume176,omitempty" json:"Volume176,omitempty" yaml:"Volume176,omitempty"`

	Volume177 float64 `xml:"Volume177,omitempty" json:"Volume177,omitempty" yaml:"Volume177,omitempty"`

	Volume178 float64 `xml:"Volume178,omitempty" json:"Volume178,omitempty" yaml:"Volume178,omitempty"`

	Volume179 float64 `xml:"Volume179,omitempty" json:"Volume179,omitempty" yaml:"Volume179,omitempty"`

	Volume180 float64 `xml:"Volume180,omitempty" json:"Volume180,omitempty" yaml:"Volume180,omitempty"`

	Volume181 float64 `xml:"Volume181,omitempty" json:"Volume181,omitempty" yaml:"Volume181,omitempty"`

	Volume182 float64 `xml:"Volume182,omitempty" json:"Volume182,omitempty" yaml:"Volume182,omitempty"`

	Volume183 float64 `xml:"Volume183,omitempty" json:"Volume183,omitempty" yaml:"Volume183,omitempty"`

	Volume184 float64 `xml:"Volume184,omitempty" json:"Volume184,omitempty" yaml:"Volume184,omitempty"`

	Volume185 float64 `xml:"Volume185,omitempty" json:"Volume185,omitempty" yaml:"Volume185,omitempty"`

	Volume186 float64 `xml:"Volume186,omitempty" json:"Volume186,omitempty" yaml:"Volume186,omitempty"`

	Volume187 float64 `xml:"Volume187,omitempty" json:"Volume187,omitempty" yaml:"Volume187,omitempty"`

	Volume188 float64 `xml:"Volume188,omitempty" json:"Volume188,omitempty" yaml:"Volume188,omitempty"`

	Volume189 float64 `xml:"Volume189,omitempty" json:"Volume189,omitempty" yaml:"Volume189,omitempty"`

	Volume190 float64 `xml:"Volume190,omitempty" json:"Volume190,omitempty" yaml:"Volume190,omitempty"`

	Volume191 float64 `xml:"Volume191,omitempty" json:"Volume191,omitempty" yaml:"Volume191,omitempty"`

	Volume192 float64 `xml:"Volume192,omitempty" json:"Volume192,omitempty" yaml:"Volume192,omitempty"`

	Volume193 float64 `xml:"Volume193,omitempty" json:"Volume193,omitempty" yaml:"Volume193,omitempty"`

	Volume194 float64 `xml:"Volume194,omitempty" json:"Volume194,omitempty" yaml:"Volume194,omitempty"`

	Volume195 float64 `xml:"Volume195,omitempty" json:"Volume195,omitempty" yaml:"Volume195,omitempty"`

	Volume196 float64 `xml:"Volume196,omitempty" json:"Volume196,omitempty" yaml:"Volume196,omitempty"`

	Volume197 float64 `xml:"Volume197,omitempty" json:"Volume197,omitempty" yaml:"Volume197,omitempty"`

	Volume198 float64 `xml:"Volume198,omitempty" json:"Volume198,omitempty" yaml:"Volume198,omitempty"`

	Volume199 float64 `xml:"Volume199,omitempty" json:"Volume199,omitempty" yaml:"Volume199,omitempty"`

	Volume200 float64 `xml:"Volume200,omitempty" json:"Volume200,omitempty" yaml:"Volume200,omitempty"`
}

type GetRealTimeCommissionsRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`
}

type GetRealTimeCommissionsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetRealTimeCommissionsResult"`

	*ApiResponse

	Commissions *ArrayOfCommissionResponse `xml:"Commissions,omitempty" json:"Commissions,omitempty" yaml:"Commissions,omitempty"`
}

type ArrayOfCommissionResponse struct {
	CommissionResponse []*CommissionResponse `xml:"CommissionResponse,omitempty" json:"CommissionResponse,omitempty" yaml:"CommissionResponse,omitempty"`
}

type CommissionResponse struct {
	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	PeriodType int32 `xml:"PeriodType,omitempty" json:"PeriodType,omitempty" yaml:"PeriodType,omitempty"`

	PeriodID int32 `xml:"PeriodID,omitempty" json:"PeriodID,omitempty" yaml:"PeriodID,omitempty"`

	PeriodDescription string `xml:"PeriodDescription,omitempty" json:"PeriodDescription,omitempty" yaml:"PeriodDescription,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty" yaml:"CurrencyCode,omitempty"`

	CommissionTotal float64 `xml:"CommissionTotal,omitempty" json:"CommissionTotal,omitempty" yaml:"CommissionTotal,omitempty"`

	Bonuses *ArrayOfCommissionBonusResponse `xml:"Bonuses,omitempty" json:"Bonuses,omitempty" yaml:"Bonuses,omitempty"`
}

type ArrayOfCommissionBonusResponse struct {
	CommissionBonusResponse []*CommissionBonusResponse `xml:"CommissionBonusResponse,omitempty" json:"CommissionBonusResponse,omitempty" yaml:"CommissionBonusResponse,omitempty"`
}

type CommissionBonusResponse struct {
	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	Amount float64 `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`

	BonusID int32 `xml:"BonusID,omitempty" json:"BonusID,omitempty" yaml:"BonusID,omitempty"`
}

type GetRealTimeCommissionDetailRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	PeriodType int32 `xml:"PeriodType,omitempty" json:"PeriodType,omitempty" yaml:"PeriodType,omitempty"`

	PeriodID int32 `xml:"PeriodID,omitempty" json:"PeriodID,omitempty" yaml:"PeriodID,omitempty"`

	BonusID int32 `xml:"BonusID,omitempty" json:"BonusID,omitempty" yaml:"BonusID,omitempty"`
}

type GetRealTimeCommissionDetailResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetRealTimeCommissionDetailResult"`

	*ApiResponse

	CommissionDetails *ArrayOfCommissionDetailResponse `xml:"CommissionDetails,omitempty" json:"CommissionDetails,omitempty" yaml:"CommissionDetails,omitempty"`
}

type ArrayOfCommissionDetailResponse struct {
	CommissionDetailResponse []*CommissionDetailResponse `xml:"CommissionDetailResponse,omitempty" json:"CommissionDetailResponse,omitempty" yaml:"CommissionDetailResponse,omitempty"`
}

type CommissionDetailResponse struct {
	FromCustomerID int32 `xml:"FromCustomerID,omitempty" json:"FromCustomerID,omitempty" yaml:"FromCustomerID,omitempty"`

	FromCustomerName string `xml:"FromCustomerName,omitempty" json:"FromCustomerName,omitempty" yaml:"FromCustomerName,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	Level int32 `xml:"Level,omitempty" json:"Level,omitempty" yaml:"Level,omitempty"`

	PaidLevel int32 `xml:"PaidLevel,omitempty" json:"PaidLevel,omitempty" yaml:"PaidLevel,omitempty"`

	SourceAmount float64 `xml:"SourceAmount,omitempty" json:"SourceAmount,omitempty" yaml:"SourceAmount,omitempty"`

	Percentage float64 `xml:"Percentage,omitempty" json:"Percentage,omitempty" yaml:"Percentage,omitempty"`

	CommissionAmount float64 `xml:"CommissionAmount,omitempty" json:"CommissionAmount,omitempty" yaml:"CommissionAmount,omitempty"`
}

type GetRankQualificationsRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	RankID int32 `xml:"RankID,omitempty" json:"RankID,omitempty" yaml:"RankID,omitempty"`

	PeriodType int32 `xml:"PeriodType,omitempty" json:"PeriodType,omitempty" yaml:"PeriodType,omitempty"`

	PeriodID int32 `xml:"PeriodID,omitempty" json:"PeriodID,omitempty" yaml:"PeriodID,omitempty"`

	CultureCode string `xml:"CultureCode,omitempty" json:"CultureCode,omitempty" yaml:"CultureCode,omitempty"`

	RankGroupID int32 `xml:"RankGroupID,omitempty" json:"RankGroupID,omitempty" yaml:"RankGroupID,omitempty"`
}

type GetRankQualificationsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetRankQualificationsResult"`

	*ApiResponse

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	RankID int32 `xml:"RankID,omitempty" json:"RankID,omitempty" yaml:"RankID,omitempty"`

	RankDescription string `xml:"RankDescription,omitempty" json:"RankDescription,omitempty" yaml:"RankDescription,omitempty"`

	Qualifies bool `xml:"Qualifies,omitempty" json:"Qualifies,omitempty" yaml:"Qualifies,omitempty"`

	QualifiesOverride bool `xml:"QualifiesOverride,omitempty" json:"QualifiesOverride,omitempty" yaml:"QualifiesOverride,omitempty"`

	PayeeQualificationLegs *ArrayOfArrayOfQualificationResponse `xml:"PayeeQualificationLegs,omitempty" json:"PayeeQualificationLegs,omitempty" yaml:"PayeeQualificationLegs,omitempty"`

	BackRankID int32 `xml:"BackRankID,omitempty" json:"BackRankID,omitempty" yaml:"BackRankID,omitempty"`

	BackRankDescription string `xml:"BackRankDescription,omitempty" json:"BackRankDescription,omitempty" yaml:"BackRankDescription,omitempty"`

	NextRankID int32 `xml:"NextRankID,omitempty" json:"NextRankID,omitempty" yaml:"NextRankID,omitempty"`

	NextRankDescription string `xml:"NextRankDescription,omitempty" json:"NextRankDescription,omitempty" yaml:"NextRankDescription,omitempty"`

	Score float64 `xml:"Score,omitempty" json:"Score,omitempty" yaml:"Score,omitempty"`
}

type ArrayOfArrayOfQualificationResponse struct {
	ArrayOfQualificationResponse []*ArrayOfQualificationResponse `xml:"ArrayOfQualificationResponse,omitempty" json:"ArrayOfQualificationResponse,omitempty" yaml:"ArrayOfQualificationResponse,omitempty"`
}

type ArrayOfQualificationResponse struct {
	QualificationResponse []*QualificationResponse `xml:"QualificationResponse,omitempty" json:"QualificationResponse,omitempty" yaml:"QualificationResponse,omitempty"`
}

type QualificationResponse struct {
	QualificationDescription string `xml:"QualificationDescription,omitempty" json:"QualificationDescription,omitempty" yaml:"QualificationDescription,omitempty"`

	Required string `xml:"Required,omitempty" json:"Required,omitempty" yaml:"Required,omitempty"`

	Actual string `xml:"Actual,omitempty" json:"Actual,omitempty" yaml:"Actual,omitempty"`

	Qualifies bool `xml:"Qualifies,omitempty" json:"Qualifies,omitempty" yaml:"Qualifies,omitempty"`

	QualifiesOverride bool `xml:"QualifiesOverride,omitempty" json:"QualifiesOverride,omitempty" yaml:"QualifiesOverride,omitempty"`

	SupportingTable struct {
	} `xml:"SupportingTable,omitempty" json:"SupportingTable,omitempty" yaml:"SupportingTable,omitempty"`

	Completed float64 `xml:"Completed,omitempty" json:"Completed,omitempty" yaml:"Completed,omitempty"`

	Weight float64 `xml:"Weight,omitempty" json:"Weight,omitempty" yaml:"Weight,omitempty"`

	Score float64 `xml:"Score,omitempty" json:"Score,omitempty" yaml:"Score,omitempty"`
}

type GetQualificationOverridesRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	OverrideID int32 `xml:"OverrideID,omitempty" json:"OverrideID,omitempty" yaml:"OverrideID,omitempty"`

	BatchSize int32 `xml:"BatchSize,omitempty" json:"BatchSize,omitempty" yaml:"BatchSize,omitempty"`
}

type GetQualitificationOverridesResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetQualificationOverridesResult"`

	*ApiResponse

	QualificationOverrides *ArrayOfGetQualificationOverrideResponse `xml:"QualificationOverrides,omitempty" json:"QualificationOverrides,omitempty" yaml:"QualificationOverrides,omitempty"`
}

type ArrayOfGetQualificationOverrideResponse struct {
	GetQualificationOverrideResponse []*GetQualificationOverrideResponse `xml:"GetQualificationOverrideResponse,omitempty" json:"GetQualificationOverrideResponse,omitempty" yaml:"GetQualificationOverrideResponse,omitempty"`
}

type GetQualificationOverrideResponse struct {
	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	OverrideID int32 `xml:"OverrideID,omitempty" json:"OverrideID,omitempty" yaml:"OverrideID,omitempty"`

	Qualifies bool `xml:"Qualifies,omitempty" json:"Qualifies,omitempty" yaml:"Qualifies,omitempty"`

	PeriodType int32 `xml:"PeriodType,omitempty" json:"PeriodType,omitempty" yaml:"PeriodType,omitempty"`

	StartPeriodID int32 `xml:"StartPeriodID,omitempty" json:"StartPeriodID,omitempty" yaml:"StartPeriodID,omitempty"`

	EndPeriodID int32 `xml:"EndPeriodID,omitempty" json:"EndPeriodID,omitempty" yaml:"EndPeriodID,omitempty"`

	ModifiedDate time.Time `xml:"ModifiedDate,omitempty" json:"ModifiedDate,omitempty" yaml:"ModifiedDate,omitempty"`

	ModifiedBy string `xml:"ModifiedBy,omitempty" json:"ModifiedBy,omitempty" yaml:"ModifiedBy,omitempty"`

	Amount float64 `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`
}

type SetQualificationOverrideRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	OverrideID int32 `xml:"OverrideID,omitempty" json:"OverrideID,omitempty" yaml:"OverrideID,omitempty"`

	Qualifies bool `xml:"Qualifies,omitempty" json:"Qualifies,omitempty" yaml:"Qualifies,omitempty"`

	PeriodType int32 `xml:"PeriodType,omitempty" json:"PeriodType,omitempty" yaml:"PeriodType,omitempty"`

	StartPeriodID int32 `xml:"StartPeriodID,omitempty" json:"StartPeriodID,omitempty" yaml:"StartPeriodID,omitempty"`

	EndPeriodID int32 `xml:"EndPeriodID,omitempty" json:"EndPeriodID,omitempty" yaml:"EndPeriodID,omitempty"`

	Amount float64 `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`
}

type SetQualificationOverrideResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetQualificationOverrideResult"`

	*ApiResponse
}

type DeleteQualificationOverrideRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	OverrideID int32 `xml:"OverrideID,omitempty" json:"OverrideID,omitempty" yaml:"OverrideID,omitempty"`
}

type DeleteQualificationOverrideResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteQualificationOverrideResult"`

	*ApiResponse
}

type AdjustInventoryRequest struct {
	*ApiRequest

	WarehouseID int32 `xml:"WarehouseID,omitempty" json:"WarehouseID,omitempty" yaml:"WarehouseID,omitempty"`

	ItemCode string `xml:"ItemCode,omitempty" json:"ItemCode,omitempty" yaml:"ItemCode,omitempty"`

	Quantity int32 `xml:"Quantity,omitempty" json:"Quantity,omitempty" yaml:"Quantity,omitempty"`

	Notes string `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`
}

type AdjustInventoryResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ AdjustInventoryResult"`

	*ApiResponse
}

type GetCustomerSiteRequest struct {
	*ApiRequest

	WebAlias string `xml:"WebAlias,omitempty" json:"WebAlias,omitempty" yaml:"WebAlias,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`
}

type GetCustomerSiteResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomerSiteResult"`

	*ApiResponse

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	WebAlias string `xml:"WebAlias,omitempty" json:"WebAlias,omitempty" yaml:"WebAlias,omitempty"`

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`

	Company string `xml:"Company,omitempty" json:"Company,omitempty" yaml:"Company,omitempty"`

	Address1 string `xml:"Address1,omitempty" json:"Address1,omitempty" yaml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty" json:"Address2,omitempty" yaml:"Address2,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty" json:"Zip,omitempty" yaml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`

	Email string `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`

	Phone string `xml:"Phone,omitempty" json:"Phone,omitempty" yaml:"Phone,omitempty"`

	Phone2 string `xml:"Phone2,omitempty" json:"Phone2,omitempty" yaml:"Phone2,omitempty"`

	Fax string `xml:"Fax,omitempty" json:"Fax,omitempty" yaml:"Fax,omitempty"`

	Notes1 string `xml:"Notes1,omitempty" json:"Notes1,omitempty" yaml:"Notes1,omitempty"`

	Notes2 string `xml:"Notes2,omitempty" json:"Notes2,omitempty" yaml:"Notes2,omitempty"`

	Notes3 string `xml:"Notes3,omitempty" json:"Notes3,omitempty" yaml:"Notes3,omitempty"`

	Notes4 string `xml:"Notes4,omitempty" json:"Notes4,omitempty" yaml:"Notes4,omitempty"`

	Url1 string `xml:"Url1,omitempty" json:"Url1,omitempty" yaml:"Url1,omitempty"`

	Url2 string `xml:"Url2,omitempty" json:"Url2,omitempty" yaml:"Url2,omitempty"`

	Url3 string `xml:"Url3,omitempty" json:"Url3,omitempty" yaml:"Url3,omitempty"`

	Url4 string `xml:"Url4,omitempty" json:"Url4,omitempty" yaml:"Url4,omitempty"`

	Url5 string `xml:"Url5,omitempty" json:"Url5,omitempty" yaml:"Url5,omitempty"`

	Url6 string `xml:"Url6,omitempty" json:"Url6,omitempty" yaml:"Url6,omitempty"`

	Url7 string `xml:"Url7,omitempty" json:"Url7,omitempty" yaml:"Url7,omitempty"`

	Url8 string `xml:"Url8,omitempty" json:"Url8,omitempty" yaml:"Url8,omitempty"`

	Url9 string `xml:"Url9,omitempty" json:"Url9,omitempty" yaml:"Url9,omitempty"`

	Url10 string `xml:"Url10,omitempty" json:"Url10,omitempty" yaml:"Url10,omitempty"`

	Url1Description string `xml:"Url1Description,omitempty" json:"Url1Description,omitempty" yaml:"Url1Description,omitempty"`

	Url2Description string `xml:"Url2Description,omitempty" json:"Url2Description,omitempty" yaml:"Url2Description,omitempty"`

	Url3Description string `xml:"Url3Description,omitempty" json:"Url3Description,omitempty" yaml:"Url3Description,omitempty"`

	Url4Description string `xml:"Url4Description,omitempty" json:"Url4Description,omitempty" yaml:"Url4Description,omitempty"`

	Url5Description string `xml:"Url5Description,omitempty" json:"Url5Description,omitempty" yaml:"Url5Description,omitempty"`

	Url6Description string `xml:"Url6Description,omitempty" json:"Url6Description,omitempty" yaml:"Url6Description,omitempty"`

	Url7Description string `xml:"Url7Description,omitempty" json:"Url7Description,omitempty" yaml:"Url7Description,omitempty"`

	Url8Description string `xml:"Url8Description,omitempty" json:"Url8Description,omitempty" yaml:"Url8Description,omitempty"`

	Url9Description string `xml:"Url9Description,omitempty" json:"Url9Description,omitempty" yaml:"Url9Description,omitempty"`

	Url10Description string `xml:"Url10Description,omitempty" json:"Url10Description,omitempty" yaml:"Url10Description,omitempty"`

	Image1 string `xml:"Image1,omitempty" json:"Image1,omitempty" yaml:"Image1,omitempty"`

	Image2 string `xml:"Image2,omitempty" json:"Image2,omitempty" yaml:"Image2,omitempty"`

	ImageUrl1 string `xml:"ImageUrl1,omitempty" json:"ImageUrl1,omitempty" yaml:"ImageUrl1,omitempty"`

	ImageUrl2 string `xml:"ImageUrl2,omitempty" json:"ImageUrl2,omitempty" yaml:"ImageUrl2,omitempty"`
}

type GetCustomerExtendedRequest struct {
	*ApiRequest

	ExtendedGroupID int32 `xml:"ExtendedGroupID,omitempty" json:"ExtendedGroupID,omitempty" yaml:"ExtendedGroupID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	CustomerExtendedID int32 `xml:"CustomerExtendedID,omitempty" json:"CustomerExtendedID,omitempty" yaml:"CustomerExtendedID,omitempty"`

	Field1 string `xml:"Field1,omitempty" json:"Field1,omitempty" yaml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty" json:"Field2,omitempty" yaml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty" json:"Field3,omitempty" yaml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty" json:"Field4,omitempty" yaml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty" json:"Field5,omitempty" yaml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty" json:"Field6,omitempty" yaml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty" json:"Field7,omitempty" yaml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty" json:"Field8,omitempty" yaml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty" json:"Field9,omitempty" yaml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty" json:"Field10,omitempty" yaml:"Field10,omitempty"`

	Field11 string `xml:"Field11,omitempty" json:"Field11,omitempty" yaml:"Field11,omitempty"`

	Field12 string `xml:"Field12,omitempty" json:"Field12,omitempty" yaml:"Field12,omitempty"`

	Field13 string `xml:"Field13,omitempty" json:"Field13,omitempty" yaml:"Field13,omitempty"`

	Field14 string `xml:"Field14,omitempty" json:"Field14,omitempty" yaml:"Field14,omitempty"`

	Field15 string `xml:"Field15,omitempty" json:"Field15,omitempty" yaml:"Field15,omitempty"`

	Field16 string `xml:"Field16,omitempty" json:"Field16,omitempty" yaml:"Field16,omitempty"`

	Field17 string `xml:"Field17,omitempty" json:"Field17,omitempty" yaml:"Field17,omitempty"`

	Field18 string `xml:"Field18,omitempty" json:"Field18,omitempty" yaml:"Field18,omitempty"`

	Field19 string `xml:"Field19,omitempty" json:"Field19,omitempty" yaml:"Field19,omitempty"`

	Field20 string `xml:"Field20,omitempty" json:"Field20,omitempty" yaml:"Field20,omitempty"`

	Field21 string `xml:"Field21,omitempty" json:"Field21,omitempty" yaml:"Field21,omitempty"`

	Field22 string `xml:"Field22,omitempty" json:"Field22,omitempty" yaml:"Field22,omitempty"`

	Field23 string `xml:"Field23,omitempty" json:"Field23,omitempty" yaml:"Field23,omitempty"`

	Field24 string `xml:"Field24,omitempty" json:"Field24,omitempty" yaml:"Field24,omitempty"`

	Field25 string `xml:"Field25,omitempty" json:"Field25,omitempty" yaml:"Field25,omitempty"`

	Field26 string `xml:"Field26,omitempty" json:"Field26,omitempty" yaml:"Field26,omitempty"`

	Field27 string `xml:"Field27,omitempty" json:"Field27,omitempty" yaml:"Field27,omitempty"`

	Field28 string `xml:"Field28,omitempty" json:"Field28,omitempty" yaml:"Field28,omitempty"`

	Field29 string `xml:"Field29,omitempty" json:"Field29,omitempty" yaml:"Field29,omitempty"`

	Field30 string `xml:"Field30,omitempty" json:"Field30,omitempty" yaml:"Field30,omitempty"`

	GreaterThanCustomerExtendedID int32 `xml:"GreaterThanCustomerExtendedID,omitempty" json:"GreaterThanCustomerExtendedID,omitempty" yaml:"GreaterThanCustomerExtendedID,omitempty"`

	GreaterThanModifiedDate time.Time `xml:"GreaterThanModifiedDate,omitempty" json:"GreaterThanModifiedDate,omitempty" yaml:"GreaterThanModifiedDate,omitempty"`

	BatchSize int32 `xml:"BatchSize,omitempty" json:"BatchSize,omitempty" yaml:"BatchSize,omitempty"`
}

type GetCustomerExtendedResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomerExtendedResult"`

	*ApiResponse

	Items *ArrayOfCustomerExtendedResponse `xml:"Items,omitempty" json:"Items,omitempty" yaml:"Items,omitempty"`
}

type ArrayOfCustomerExtendedResponse struct {
	CustomerExtendedResponse []*CustomerExtendedResponse `xml:"CustomerExtendedResponse,omitempty" json:"CustomerExtendedResponse,omitempty" yaml:"CustomerExtendedResponse,omitempty"`
}

type CustomerExtendedResponse struct {
	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	ExtendedGroupID int32 `xml:"ExtendedGroupID,omitempty" json:"ExtendedGroupID,omitempty" yaml:"ExtendedGroupID,omitempty"`

	CustomerExtendedID int32 `xml:"CustomerExtendedID,omitempty" json:"CustomerExtendedID,omitempty" yaml:"CustomerExtendedID,omitempty"`

	Field1 string `xml:"Field1,omitempty" json:"Field1,omitempty" yaml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty" json:"Field2,omitempty" yaml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty" json:"Field3,omitempty" yaml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty" json:"Field4,omitempty" yaml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty" json:"Field5,omitempty" yaml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty" json:"Field6,omitempty" yaml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty" json:"Field7,omitempty" yaml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty" json:"Field8,omitempty" yaml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty" json:"Field9,omitempty" yaml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty" json:"Field10,omitempty" yaml:"Field10,omitempty"`

	Field11 string `xml:"Field11,omitempty" json:"Field11,omitempty" yaml:"Field11,omitempty"`

	Field12 string `xml:"Field12,omitempty" json:"Field12,omitempty" yaml:"Field12,omitempty"`

	Field13 string `xml:"Field13,omitempty" json:"Field13,omitempty" yaml:"Field13,omitempty"`

	Field14 string `xml:"Field14,omitempty" json:"Field14,omitempty" yaml:"Field14,omitempty"`

	Field15 string `xml:"Field15,omitempty" json:"Field15,omitempty" yaml:"Field15,omitempty"`

	Field16 string `xml:"Field16,omitempty" json:"Field16,omitempty" yaml:"Field16,omitempty"`

	Field17 string `xml:"Field17,omitempty" json:"Field17,omitempty" yaml:"Field17,omitempty"`

	Field18 string `xml:"Field18,omitempty" json:"Field18,omitempty" yaml:"Field18,omitempty"`

	Field19 string `xml:"Field19,omitempty" json:"Field19,omitempty" yaml:"Field19,omitempty"`

	Field20 string `xml:"Field20,omitempty" json:"Field20,omitempty" yaml:"Field20,omitempty"`

	Field21 string `xml:"Field21,omitempty" json:"Field21,omitempty" yaml:"Field21,omitempty"`

	Field22 string `xml:"Field22,omitempty" json:"Field22,omitempty" yaml:"Field22,omitempty"`

	Field23 string `xml:"Field23,omitempty" json:"Field23,omitempty" yaml:"Field23,omitempty"`

	Field24 string `xml:"Field24,omitempty" json:"Field24,omitempty" yaml:"Field24,omitempty"`

	Field25 string `xml:"Field25,omitempty" json:"Field25,omitempty" yaml:"Field25,omitempty"`

	Field26 string `xml:"Field26,omitempty" json:"Field26,omitempty" yaml:"Field26,omitempty"`

	Field27 string `xml:"Field27,omitempty" json:"Field27,omitempty" yaml:"Field27,omitempty"`

	Field28 string `xml:"Field28,omitempty" json:"Field28,omitempty" yaml:"Field28,omitempty"`

	Field29 string `xml:"Field29,omitempty" json:"Field29,omitempty" yaml:"Field29,omitempty"`

	Field30 string `xml:"Field30,omitempty" json:"Field30,omitempty" yaml:"Field30,omitempty"`

	ModifiedDate time.Time `xml:"ModifiedDate,omitempty" json:"ModifiedDate,omitempty" yaml:"ModifiedDate,omitempty"`
}

type GetCustomerBillingRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`
}

type GetCustomerBillingResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomerBillingResult"`

	*ApiResponse

	PrimaryCreditCard *CreditCardAccountResponse `xml:"PrimaryCreditCard,omitempty" json:"PrimaryCreditCard,omitempty" yaml:"PrimaryCreditCard,omitempty"`

	SecondaryCreditCard *CreditCardAccountResponse `xml:"SecondaryCreditCard,omitempty" json:"SecondaryCreditCard,omitempty" yaml:"SecondaryCreditCard,omitempty"`

	BankAccount *BankAccountResponse `xml:"BankAccount,omitempty" json:"BankAccount,omitempty" yaml:"BankAccount,omitempty"`

	PrimaryWalletAccount *WalletAccountResponse `xml:"PrimaryWalletAccount,omitempty" json:"PrimaryWalletAccount,omitempty" yaml:"PrimaryWalletAccount,omitempty"`

	SecondaryWallletAccount *WalletAccountResponse `xml:"SecondaryWallletAccount,omitempty" json:"SecondaryWallletAccount,omitempty" yaml:"SecondaryWallletAccount,omitempty"`
}

type CreditCardAccountResponse struct {
	CreditCardNumberDisplay string `xml:"CreditCardNumberDisplay,omitempty" json:"CreditCardNumberDisplay,omitempty" yaml:"CreditCardNumberDisplay,omitempty"`

	ExpirationMonth int32 `xml:"ExpirationMonth,omitempty" json:"ExpirationMonth,omitempty" yaml:"ExpirationMonth,omitempty"`

	ExpirationYear int32 `xml:"ExpirationYear,omitempty" json:"ExpirationYear,omitempty" yaml:"ExpirationYear,omitempty"`

	CreditCardType int32 `xml:"CreditCardType,omitempty" json:"CreditCardType,omitempty" yaml:"CreditCardType,omitempty"`

	CreditCardTypeDescription string `xml:"CreditCardTypeDescription,omitempty" json:"CreditCardTypeDescription,omitempty" yaml:"CreditCardTypeDescription,omitempty"`

	BillingName string `xml:"BillingName,omitempty" json:"BillingName,omitempty" yaml:"BillingName,omitempty"`

	BillingAddress string `xml:"BillingAddress,omitempty" json:"BillingAddress,omitempty" yaml:"BillingAddress,omitempty"`

	BillingCity string `xml:"BillingCity,omitempty" json:"BillingCity,omitempty" yaml:"BillingCity,omitempty"`

	BillingState string `xml:"BillingState,omitempty" json:"BillingState,omitempty" yaml:"BillingState,omitempty"`

	BillingZip string `xml:"BillingZip,omitempty" json:"BillingZip,omitempty" yaml:"BillingZip,omitempty"`

	BillingCountry string `xml:"BillingCountry,omitempty" json:"BillingCountry,omitempty" yaml:"BillingCountry,omitempty"`

	BillingAddress2 string `xml:"BillingAddress2,omitempty" json:"BillingAddress2,omitempty" yaml:"BillingAddress2,omitempty"`
}

type BankAccountResponse struct {
	BankAccountNumberDisplay string `xml:"BankAccountNumberDisplay,omitempty" json:"BankAccountNumberDisplay,omitempty" yaml:"BankAccountNumberDisplay,omitempty"`

	BankRoutingNumber string `xml:"BankRoutingNumber,omitempty" json:"BankRoutingNumber,omitempty" yaml:"BankRoutingNumber,omitempty"`

	BankName string `xml:"BankName,omitempty" json:"BankName,omitempty" yaml:"BankName,omitempty"`

	BankAccountType *BankAccountType `xml:"BankAccountType,omitempty" json:"BankAccountType,omitempty" yaml:"BankAccountType,omitempty"`

	NameOnAccount string `xml:"NameOnAccount,omitempty" json:"NameOnAccount,omitempty" yaml:"NameOnAccount,omitempty"`

	BillingAddress string `xml:"BillingAddress,omitempty" json:"BillingAddress,omitempty" yaml:"BillingAddress,omitempty"`

	BillingCity string `xml:"BillingCity,omitempty" json:"BillingCity,omitempty" yaml:"BillingCity,omitempty"`

	BillingState string `xml:"BillingState,omitempty" json:"BillingState,omitempty" yaml:"BillingState,omitempty"`

	BillingZip string `xml:"BillingZip,omitempty" json:"BillingZip,omitempty" yaml:"BillingZip,omitempty"`

	BillingCountry string `xml:"BillingCountry,omitempty" json:"BillingCountry,omitempty" yaml:"BillingCountry,omitempty"`

	BillingAddress2 string `xml:"BillingAddress2,omitempty" json:"BillingAddress2,omitempty" yaml:"BillingAddress2,omitempty"`
}

type WalletAccountResponse struct {
	WalletType int32 `xml:"WalletType,omitempty" json:"WalletType,omitempty" yaml:"WalletType,omitempty"`

	WalletAccountDisplay string `xml:"WalletAccountDisplay,omitempty" json:"WalletAccountDisplay,omitempty" yaml:"WalletAccountDisplay,omitempty"`
}

type GetPaymentsRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`
}

type GetPaymentsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetPaymentsResult"`

	*ApiResponse

	Payments *ArrayOfPaymentResponse `xml:"Payments,omitempty" json:"Payments,omitempty" yaml:"Payments,omitempty"`
}

type ArrayOfPaymentResponse struct {
	PaymentResponse []*PaymentResponse `xml:"PaymentResponse,omitempty" json:"PaymentResponse,omitempty" yaml:"PaymentResponse,omitempty"`
}

type PaymentResponse struct {
	PaymentID int32 `xml:"PaymentID,omitempty" json:"PaymentID,omitempty" yaml:"PaymentID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	PaymentType *PaymentType `xml:"PaymentType,omitempty" json:"PaymentType,omitempty" yaml:"PaymentType,omitempty"`

	PaymentDate time.Time `xml:"PaymentDate,omitempty" json:"PaymentDate,omitempty" yaml:"PaymentDate,omitempty"`

	Amount float64 `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty" yaml:"CurrencyCode,omitempty"`

	BillingName string `xml:"BillingName,omitempty" json:"BillingName,omitempty" yaml:"BillingName,omitempty"`

	BillingAddress1 string `xml:"BillingAddress1,omitempty" json:"BillingAddress1,omitempty" yaml:"BillingAddress1,omitempty"`

	BillingAddress2 string `xml:"BillingAddress2,omitempty" json:"BillingAddress2,omitempty" yaml:"BillingAddress2,omitempty"`

	BillingCity string `xml:"BillingCity,omitempty" json:"BillingCity,omitempty" yaml:"BillingCity,omitempty"`

	BillingState string `xml:"BillingState,omitempty" json:"BillingState,omitempty" yaml:"BillingState,omitempty"`

	BillingZip string `xml:"BillingZip,omitempty" json:"BillingZip,omitempty" yaml:"BillingZip,omitempty"`

	BillingCountry string `xml:"BillingCountry,omitempty" json:"BillingCountry,omitempty" yaml:"BillingCountry,omitempty"`

	BankName string `xml:"BankName,omitempty" json:"BankName,omitempty" yaml:"BankName,omitempty"`

	Memo string `xml:"Memo,omitempty" json:"Memo,omitempty" yaml:"Memo,omitempty"`

	CreditCardNumberDisplay string `xml:"CreditCardNumberDisplay,omitempty" json:"CreditCardNumberDisplay,omitempty" yaml:"CreditCardNumberDisplay,omitempty"`

	AuthorizationCode string `xml:"AuthorizationCode,omitempty" json:"AuthorizationCode,omitempty" yaml:"AuthorizationCode,omitempty"`

	CreditCardType int32 `xml:"CreditCardType,omitempty" json:"CreditCardType,omitempty" yaml:"CreditCardType,omitempty"`

	CreditCardTypeDescription string `xml:"CreditCardTypeDescription,omitempty" json:"CreditCardTypeDescription,omitempty" yaml:"CreditCardTypeDescription,omitempty"`
}

type FundPaymentCardRequest struct {
	*BaseCreatePayoutRequest

	PaymentCardTypeID int32 `xml:"PaymentCardTypeID,omitempty" json:"PaymentCardTypeID,omitempty" yaml:"PaymentCardTypeID,omitempty"`

	BillIDList *ArrayOfInt `xml:"BillIDList,omitempty" json:"BillIDList,omitempty" yaml:"BillIDList,omitempty"`
}

type FundPaymentCardResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ FundPaymentCardResult"`

	*CreatePayoutResponse
}

type CreatePaymentPointAccountRequest struct {
	*BaseCreatePaymentRequest

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	PointAccountID int32 `xml:"PointAccountID,omitempty" json:"PointAccountID,omitempty" yaml:"PointAccountID,omitempty"`

	PaymentDate time.Time `xml:"PaymentDate,omitempty" json:"PaymentDate,omitempty" yaml:"PaymentDate,omitempty"`

	Amount float64 `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`

	Memo string `xml:"Memo,omitempty" json:"Memo,omitempty" yaml:"Memo,omitempty"`
}

type CreatePaymentPointAccountResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreatePaymentPointAccountResult"`

	*BaseCreatePaymentResponse

	PaymentID int32 `xml:"PaymentID,omitempty" json:"PaymentID,omitempty" yaml:"PaymentID,omitempty"`
}

type CreatePaymentCheckRequest struct {
	*BaseCreatePaymentRequest

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	PaymentDate time.Time `xml:"PaymentDate,omitempty" json:"PaymentDate,omitempty" yaml:"PaymentDate,omitempty"`

	Amount float64 `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`

	Memo string `xml:"Memo,omitempty" json:"Memo,omitempty" yaml:"Memo,omitempty"`

	CheckNumber string `xml:"CheckNumber,omitempty" json:"CheckNumber,omitempty" yaml:"CheckNumber,omitempty"`

	CheckAccountNumber string `xml:"CheckAccountNumber,omitempty" json:"CheckAccountNumber,omitempty" yaml:"CheckAccountNumber,omitempty"`

	CheckRoutingNumber string `xml:"CheckRoutingNumber,omitempty" json:"CheckRoutingNumber,omitempty" yaml:"CheckRoutingNumber,omitempty"`

	CheckDate time.Time `xml:"CheckDate,omitempty" json:"CheckDate,omitempty" yaml:"CheckDate,omitempty"`
}

type CreatePaymentCheckResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreatePaymentCheckResult"`

	*BaseCreatePaymentResponse

	PaymentID int32 `xml:"PaymentID,omitempty" json:"PaymentID,omitempty" yaml:"PaymentID,omitempty"`
}

type GetCustomReportRequest struct {
	*ApiRequest

	ReportID int32 `xml:"ReportID,omitempty" json:"ReportID,omitempty" yaml:"ReportID,omitempty"`

	Parameters *ArrayOfParameterRequest `xml:"Parameters,omitempty" json:"Parameters,omitempty" yaml:"Parameters,omitempty"`
}

type ArrayOfParameterRequest struct {
	ParameterRequest []*ParameterRequest `xml:"ParameterRequest,omitempty" json:"ParameterRequest,omitempty" yaml:"ParameterRequest,omitempty"`
}

type ParameterRequest struct {
	ParameterName string `xml:"ParameterName,omitempty" json:"ParameterName,omitempty" yaml:"ParameterName,omitempty"`

	Value struct {
	} `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
}

type GetCustomReportResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomReportResult"`

	*ApiResponse

	ReportData struct {
		Schema *Schema `xml:"schema,omitempty" json:",omitempty" yaml:",omitempty"`
	} `xml:"ReportData,omitempty" json:"ReportData,omitempty" yaml:"ReportData,omitempty"`
}

type GetReportRequest struct {
	*ApiRequest

	ReportID int32 `xml:"ReportID,omitempty" json:"ReportID,omitempty" yaml:"ReportID,omitempty"`

	Parameters *ArrayOfReportParameterRequest `xml:"Parameters,omitempty" json:"Parameters,omitempty" yaml:"Parameters,omitempty"`
}

type ArrayOfReportParameterRequest struct {
	ReportParameterRequest []*ReportParameterRequest `xml:"ReportParameterRequest,omitempty" json:"ReportParameterRequest,omitempty" yaml:"ReportParameterRequest,omitempty"`
}

type ReportParameterRequest struct {
	ParameterName string `xml:"ParameterName,omitempty" json:"ParameterName,omitempty" yaml:"ParameterName,omitempty"`

	Value string `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
}

type GetReportResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetReportResult"`

	*ApiResponse

	ReportData struct {
		Schema *Schema `xml:"schema,omitempty" json:",omitempty" yaml:",omitempty"`
	} `xml:"ReportData,omitempty" json:"ReportData,omitempty" yaml:"ReportData,omitempty"`
}

type ChargeWalletAccountResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ChargeWalletAccountResult"`

	*CreatePaymentResponse

	Amount float64 `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`

	AuthorizationCode string `xml:"AuthorizationCode,omitempty" json:"AuthorizationCode,omitempty" yaml:"AuthorizationCode,omitempty"`
}

type GetAccountDirectDepositRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`
}

type GetAccountDirectDepositResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetAccountDirectDepositResult"`

	*ApiResponse

	NameOnAccount string `xml:"NameOnAccount,omitempty" json:"NameOnAccount,omitempty" yaml:"NameOnAccount,omitempty"`

	BankAccountNumberDisplay string `xml:"BankAccountNumberDisplay,omitempty" json:"BankAccountNumberDisplay,omitempty" yaml:"BankAccountNumberDisplay,omitempty"`

	BankRoutingNumber string `xml:"BankRoutingNumber,omitempty" json:"BankRoutingNumber,omitempty" yaml:"BankRoutingNumber,omitempty"`

	DepositAccountType *DepositAccountType `xml:"DepositAccountType,omitempty" json:"DepositAccountType,omitempty" yaml:"DepositAccountType,omitempty"`

	BankName string `xml:"BankName,omitempty" json:"BankName,omitempty" yaml:"BankName,omitempty"`

	BankAddress string `xml:"BankAddress,omitempty" json:"BankAddress,omitempty" yaml:"BankAddress,omitempty"`

	BankCity string `xml:"BankCity,omitempty" json:"BankCity,omitempty" yaml:"BankCity,omitempty"`

	BankState string `xml:"BankState,omitempty" json:"BankState,omitempty" yaml:"BankState,omitempty"`

	BankZip string `xml:"BankZip,omitempty" json:"BankZip,omitempty" yaml:"BankZip,omitempty"`

	BankCountry string `xml:"BankCountry,omitempty" json:"BankCountry,omitempty" yaml:"BankCountry,omitempty"`

	BankAccountType *BankAccountType `xml:"BankAccountType,omitempty" json:"BankAccountType,omitempty" yaml:"BankAccountType,omitempty"`

	Iban string `xml:"Iban,omitempty" json:"Iban,omitempty" yaml:"Iban,omitempty"`

	SwiftCode string `xml:"SwiftCode,omitempty" json:"SwiftCode,omitempty" yaml:"SwiftCode,omitempty"`
}

type OptInEmailResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ OptInEmailResult"`

	*ApiResponse

	RecordsAffected int32 `xml:"RecordsAffected,omitempty" json:"RecordsAffected,omitempty" yaml:"RecordsAffected,omitempty"`
}

type OptInSmsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ OptInSmsResult"`

	*ApiResponse

	RecordsAffected int32 `xml:"RecordsAffected,omitempty" json:"RecordsAffected,omitempty" yaml:"RecordsAffected,omitempty"`
}

type SetCustomerSiteImageRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetCusotmerSiteImageRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	ImageName string `xml:"ImageName,omitempty" json:"ImageName,omitempty" yaml:"ImageName,omitempty"`

	ImageData []byte `xml:"ImageData,omitempty" json:"ImageData,omitempty" yaml:"ImageData,omitempty"`

	CustomerSiteImageType *CustomerSiteImageType `xml:"CustomerSiteImageType,omitempty" json:"CustomerSiteImageType,omitempty" yaml:"CustomerSiteImageType,omitempty"`
}

type SetCustomerSiteImageResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetCustomerSiteImageResult"`

	*ApiResponse
}

type LoginCustomerRequest struct {
	*ApiRequest

	LoginName string `xml:"LoginName,omitempty" json:"LoginName,omitempty" yaml:"LoginName,omitempty"`

	Password string `xml:"Password,omitempty" json:"Password,omitempty" yaml:"Password,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`
}

type LoginCustomerResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ LoginCustomerResult"`

	*ApiResponse

	SessionID string `xml:"SessionID,omitempty" json:"SessionID,omitempty" yaml:"SessionID,omitempty"`
}

type GetLoginSessionRequest struct {
	*ApiRequest

	SessionID string `xml:"SessionID,omitempty" json:"SessionID,omitempty" yaml:"SessionID,omitempty"`
}

type GetLoginSessionResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetLoginSessionResult"`

	*ApiResponse

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`
}

type AuthenticateCustomerRequest struct {
	*ApiRequest

	LoginName string `xml:"LoginName,omitempty" json:"LoginName,omitempty" yaml:"LoginName,omitempty"`

	Password string `xml:"Password,omitempty" json:"Password,omitempty" yaml:"Password,omitempty"`
}

type AuthenticateCustomerResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ AuthenticateCustomerResult"`

	*ApiResponse

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`
}

type AuthenticateUserRequest struct {
	*ApiRequest

	LoginName string `xml:"LoginName,omitempty" json:"LoginName,omitempty" yaml:"LoginName,omitempty"`

	Password string `xml:"Password,omitempty" json:"Password,omitempty" yaml:"Password,omitempty"`
}

type AuthenticateUserResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ AuthenticateUserResult"`

	*ApiResponse

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`
}

type GetUserPermissionsRequest struct {
	*ApiRequest

	LoginName string `xml:"LoginName,omitempty" json:"LoginName,omitempty" yaml:"LoginName,omitempty"`
}

type GetUserPermissionsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetUserPermissionsResult"`

	*ApiResponse

	RestrictToCustomerTypes *ArrayOfInt `xml:"RestrictToCustomerTypes,omitempty" json:"RestrictToCustomerTypes,omitempty" yaml:"RestrictToCustomerTypes,omitempty"`

	RestrictToCustomerStatuses *ArrayOfInt `xml:"RestrictToCustomerStatuses,omitempty" json:"RestrictToCustomerStatuses,omitempty" yaml:"RestrictToCustomerStatuses,omitempty"`

	RestrictToWarehouses *ArrayOfInt `xml:"RestrictToWarehouses,omitempty" json:"RestrictToWarehouses,omitempty" yaml:"RestrictToWarehouses,omitempty"`

	RestrictToCountries *ArrayOfString `xml:"RestrictToCountries,omitempty" json:"RestrictToCountries,omitempty" yaml:"RestrictToCountries,omitempty"`

	RestrictToCurrencies *ArrayOfString `xml:"RestrictToCurrencies,omitempty" json:"RestrictToCurrencies,omitempty" yaml:"RestrictToCurrencies,omitempty"`

	ViewDeletedCustomers bool `xml:"ViewDeletedCustomers,omitempty" json:"ViewDeletedCustomers,omitempty" yaml:"ViewDeletedCustomers,omitempty"`

	AllowRemoteCheckPrint bool `xml:"AllowRemoteCheckPrint,omitempty" json:"AllowRemoteCheckPrint,omitempty" yaml:"AllowRemoteCheckPrint,omitempty"`

	AllowOverrideItemPrice bool `xml:"AllowOverrideItemPrice,omitempty" json:"AllowOverrideItemPrice,omitempty" yaml:"AllowOverrideItemPrice,omitempty"`

	AllowStatementPrint bool `xml:"AllowStatementPrint,omitempty" json:"AllowStatementPrint,omitempty" yaml:"AllowStatementPrint,omitempty"`

	DefaultWarehouseID int32 `xml:"DefaultWarehouseID,omitempty" json:"DefaultWarehouseID,omitempty" yaml:"DefaultWarehouseID,omitempty"`

	LanguageID int32 `xml:"LanguageID,omitempty" json:"LanguageID,omitempty" yaml:"LanguageID,omitempty"`

	CultureCode string `xml:"CultureCode,omitempty" json:"CultureCode,omitempty" yaml:"CultureCode,omitempty"`
}

type ChangeOrderStatusRequest struct {
	*ApiRequest

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	OrderStatus *OrderStatusType `xml:"OrderStatus,omitempty" json:"OrderStatus,omitempty" yaml:"OrderStatus,omitempty"`
}

type ChangeOrderStatusResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ChangeOrderStatusResult"`

	*ApiResponse
}

type ChangeAutoOrderStatusRequest struct {
	*ApiRequest

	AutoOrderID int32 `xml:"AutoOrderID,omitempty" json:"AutoOrderID,omitempty" yaml:"AutoOrderID,omitempty"`

	AutoOrderStatus *AutoOrderStatusType `xml:"AutoOrderStatus,omitempty" json:"AutoOrderStatus,omitempty" yaml:"AutoOrderStatus,omitempty"`
}

type ChangeAutoOrderStatusResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ChangeAutoOrderStatusResult"`

	*ApiResponse
}

type GetShipMethodsRequest struct {
	*ApiRequest

	WarehouseID int32 `xml:"WarehouseID,omitempty" json:"WarehouseID,omitempty" yaml:"WarehouseID,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty" yaml:"CurrencyCode,omitempty"`

	OrderSubTotal float64 `xml:"OrderSubTotal,omitempty" json:"OrderSubTotal,omitempty" yaml:"OrderSubTotal,omitempty"`

	OrderWieght float64 `xml:"OrderWieght,omitempty" json:"OrderWieght,omitempty" yaml:"OrderWieght,omitempty"`
}

type GetShipMethodsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetShipMethodsResult"`

	*ApiResponse

	ShipMethods *ArrayOfShipMethodResponse `xml:"ShipMethods,omitempty" json:"ShipMethods,omitempty" yaml:"ShipMethods,omitempty"`
}

type GetOrdersRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	OrderIDs *ArrayOfInt `xml:"OrderIDs,omitempty" json:"OrderIDs,omitempty" yaml:"OrderIDs,omitempty"`

	OrderDateStart time.Time `xml:"OrderDateStart,omitempty" json:"OrderDateStart,omitempty" yaml:"OrderDateStart,omitempty"`

	OrderDateEnd time.Time `xml:"OrderDateEnd,omitempty" json:"OrderDateEnd,omitempty" yaml:"OrderDateEnd,omitempty"`

	OrderStatus *OrderStatusType `xml:"OrderStatus,omitempty" json:"OrderStatus,omitempty" yaml:"OrderStatus,omitempty"`

	WarehouseID int32 `xml:"WarehouseID,omitempty" json:"WarehouseID,omitempty" yaml:"WarehouseID,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty" yaml:"CurrencyCode,omitempty"`

	ReturnCustomer bool `xml:"ReturnCustomer,omitempty" json:"ReturnCustomer,omitempty" yaml:"ReturnCustomer,omitempty"`

	ReturnKitDetails bool `xml:"ReturnKitDetails,omitempty" json:"ReturnKitDetails,omitempty" yaml:"ReturnKitDetails,omitempty"`

	GreaterThanOrderID int32 `xml:"GreaterThanOrderID,omitempty" json:"GreaterThanOrderID,omitempty" yaml:"GreaterThanOrderID,omitempty"`

	GreaterThanModifiedDate time.Time `xml:"GreaterThanModifiedDate,omitempty" json:"GreaterThanModifiedDate,omitempty" yaml:"GreaterThanModifiedDate,omitempty"`

	BatchSize int32 `xml:"BatchSize,omitempty" json:"BatchSize,omitempty" yaml:"BatchSize,omitempty"`

	Other11 string `xml:"Other11,omitempty" json:"Other11,omitempty" yaml:"Other11,omitempty"`

	Other12 string `xml:"Other12,omitempty" json:"Other12,omitempty" yaml:"Other12,omitempty"`

	Other13 string `xml:"Other13,omitempty" json:"Other13,omitempty" yaml:"Other13,omitempty"`

	Other14 string `xml:"Other14,omitempty" json:"Other14,omitempty" yaml:"Other14,omitempty"`

	Other15 string `xml:"Other15,omitempty" json:"Other15,omitempty" yaml:"Other15,omitempty"`

	Other16 string `xml:"Other16,omitempty" json:"Other16,omitempty" yaml:"Other16,omitempty"`

	Other17 string `xml:"Other17,omitempty" json:"Other17,omitempty" yaml:"Other17,omitempty"`

	Other18 string `xml:"Other18,omitempty" json:"Other18,omitempty" yaml:"Other18,omitempty"`

	Other19 string `xml:"Other19,omitempty" json:"Other19,omitempty" yaml:"Other19,omitempty"`

	Other20 string `xml:"Other20,omitempty" json:"Other20,omitempty" yaml:"Other20,omitempty"`

	PartyID int32 `xml:"PartyID,omitempty" json:"PartyID,omitempty" yaml:"PartyID,omitempty"`
}

type OrderResponse struct {
	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	OrderStatus *OrderStatusType `xml:"OrderStatus,omitempty" json:"OrderStatus,omitempty" yaml:"OrderStatus,omitempty"`

	OrderDate time.Time `xml:"OrderDate,omitempty" json:"OrderDate,omitempty" yaml:"OrderDate,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty" yaml:"CurrencyCode,omitempty"`

	WarehouseID int32 `xml:"WarehouseID,omitempty" json:"WarehouseID,omitempty" yaml:"WarehouseID,omitempty"`

	ShipMethodID int32 `xml:"ShipMethodID,omitempty" json:"ShipMethodID,omitempty" yaml:"ShipMethodID,omitempty"`

	PriceType int32 `xml:"PriceType,omitempty" json:"PriceType,omitempty" yaml:"PriceType,omitempty"`

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`

	Company string `xml:"Company,omitempty" json:"Company,omitempty" yaml:"Company,omitempty"`

	Address1 string `xml:"Address1,omitempty" json:"Address1,omitempty" yaml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty" json:"Address2,omitempty" yaml:"Address2,omitempty"`

	Address3 string `xml:"Address3,omitempty" json:"Address3,omitempty" yaml:"Address3,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty" json:"Zip,omitempty" yaml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`

	County string `xml:"County,omitempty" json:"County,omitempty" yaml:"County,omitempty"`

	Email string `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`

	Phone string `xml:"Phone,omitempty" json:"Phone,omitempty" yaml:"Phone,omitempty"`

	Notes string `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`

	Total float64 `xml:"Total,omitempty" json:"Total,omitempty" yaml:"Total,omitempty"`

	SubTotal float64 `xml:"SubTotal,omitempty" json:"SubTotal,omitempty" yaml:"SubTotal,omitempty"`

	TaxTotal float64 `xml:"TaxTotal,omitempty" json:"TaxTotal,omitempty" yaml:"TaxTotal,omitempty"`

	ShippingTotal float64 `xml:"ShippingTotal,omitempty" json:"ShippingTotal,omitempty" yaml:"ShippingTotal,omitempty"`

	DiscountTotal float64 `xml:"DiscountTotal,omitempty" json:"DiscountTotal,omitempty" yaml:"DiscountTotal,omitempty"`

	DiscountPercent float64 `xml:"DiscountPercent,omitempty" json:"DiscountPercent,omitempty" yaml:"DiscountPercent,omitempty"`

	WeightTotal float64 `xml:"WeightTotal,omitempty" json:"WeightTotal,omitempty" yaml:"WeightTotal,omitempty"`

	BusinessVolumeTotal float64 `xml:"BusinessVolumeTotal,omitempty" json:"BusinessVolumeTotal,omitempty" yaml:"BusinessVolumeTotal,omitempty"`

	CommissionableVolumeTotal float64 `xml:"CommissionableVolumeTotal,omitempty" json:"CommissionableVolumeTotal,omitempty" yaml:"CommissionableVolumeTotal,omitempty"`

	TrackingNumber1 string `xml:"TrackingNumber1,omitempty" json:"TrackingNumber1,omitempty" yaml:"TrackingNumber1,omitempty"`

	TrackingNumber2 string `xml:"TrackingNumber2,omitempty" json:"TrackingNumber2,omitempty" yaml:"TrackingNumber2,omitempty"`

	TrackingNumber3 string `xml:"TrackingNumber3,omitempty" json:"TrackingNumber3,omitempty" yaml:"TrackingNumber3,omitempty"`

	TrackingNumber4 string `xml:"TrackingNumber4,omitempty" json:"TrackingNumber4,omitempty" yaml:"TrackingNumber4,omitempty"`

	TrackingNumber5 string `xml:"TrackingNumber5,omitempty" json:"TrackingNumber5,omitempty" yaml:"TrackingNumber5,omitempty"`

	Other1Total float64 `xml:"Other1Total,omitempty" json:"Other1Total,omitempty" yaml:"Other1Total,omitempty"`

	Other2Total float64 `xml:"Other2Total,omitempty" json:"Other2Total,omitempty" yaml:"Other2Total,omitempty"`

	Other3Total float64 `xml:"Other3Total,omitempty" json:"Other3Total,omitempty" yaml:"Other3Total,omitempty"`

	Other4Total float64 `xml:"Other4Total,omitempty" json:"Other4Total,omitempty" yaml:"Other4Total,omitempty"`

	Other5Total float64 `xml:"Other5Total,omitempty" json:"Other5Total,omitempty" yaml:"Other5Total,omitempty"`

	Other6Total float64 `xml:"Other6Total,omitempty" json:"Other6Total,omitempty" yaml:"Other6Total,omitempty"`

	Other7Total float64 `xml:"Other7Total,omitempty" json:"Other7Total,omitempty" yaml:"Other7Total,omitempty"`

	Other8Total float64 `xml:"Other8Total,omitempty" json:"Other8Total,omitempty" yaml:"Other8Total,omitempty"`

	Other9Total float64 `xml:"Other9Total,omitempty" json:"Other9Total,omitempty" yaml:"Other9Total,omitempty"`

	Other10Total float64 `xml:"Other10Total,omitempty" json:"Other10Total,omitempty" yaml:"Other10Total,omitempty"`

	ShippingTax float64 `xml:"ShippingTax,omitempty" json:"ShippingTax,omitempty" yaml:"ShippingTax,omitempty"`

	OrderTax float64 `xml:"OrderTax,omitempty" json:"OrderTax,omitempty" yaml:"OrderTax,omitempty"`

	FedTaxTotal float64 `xml:"FedTaxTotal,omitempty" json:"FedTaxTotal,omitempty" yaml:"FedTaxTotal,omitempty"`

	StateTaxTotal float64 `xml:"StateTaxTotal,omitempty" json:"StateTaxTotal,omitempty" yaml:"StateTaxTotal,omitempty"`

	FedShippingTax float64 `xml:"FedShippingTax,omitempty" json:"FedShippingTax,omitempty" yaml:"FedShippingTax,omitempty"`

	StateShippingTax float64 `xml:"StateShippingTax,omitempty" json:"StateShippingTax,omitempty" yaml:"StateShippingTax,omitempty"`

	CityShippingTax float64 `xml:"CityShippingTax,omitempty" json:"CityShippingTax,omitempty" yaml:"CityShippingTax,omitempty"`

	CityLocalShippingTax float64 `xml:"CityLocalShippingTax,omitempty" json:"CityLocalShippingTax,omitempty" yaml:"CityLocalShippingTax,omitempty"`

	CountyShippingTax float64 `xml:"CountyShippingTax,omitempty" json:"CountyShippingTax,omitempty" yaml:"CountyShippingTax,omitempty"`

	CountyLocalShippingTax float64 `xml:"CountyLocalShippingTax,omitempty" json:"CountyLocalShippingTax,omitempty" yaml:"CountyLocalShippingTax,omitempty"`

	Other11 string `xml:"Other11,omitempty" json:"Other11,omitempty" yaml:"Other11,omitempty"`

	Other12 string `xml:"Other12,omitempty" json:"Other12,omitempty" yaml:"Other12,omitempty"`

	Other13 string `xml:"Other13,omitempty" json:"Other13,omitempty" yaml:"Other13,omitempty"`

	Other14 string `xml:"Other14,omitempty" json:"Other14,omitempty" yaml:"Other14,omitempty"`

	Other15 string `xml:"Other15,omitempty" json:"Other15,omitempty" yaml:"Other15,omitempty"`

	Other16 string `xml:"Other16,omitempty" json:"Other16,omitempty" yaml:"Other16,omitempty"`

	Other17 string `xml:"Other17,omitempty" json:"Other17,omitempty" yaml:"Other17,omitempty"`

	Other18 string `xml:"Other18,omitempty" json:"Other18,omitempty" yaml:"Other18,omitempty"`

	Other19 string `xml:"Other19,omitempty" json:"Other19,omitempty" yaml:"Other19,omitempty"`

	Other20 string `xml:"Other20,omitempty" json:"Other20,omitempty" yaml:"Other20,omitempty"`

	ModifiedDate time.Time `xml:"ModifiedDate,omitempty" json:"ModifiedDate,omitempty" yaml:"ModifiedDate,omitempty"`

	OrderType *OrderType `xml:"OrderType,omitempty" json:"OrderType,omitempty" yaml:"OrderType,omitempty"`

	ShippedDate time.Time `xml:"ShippedDate,omitempty" json:"ShippedDate,omitempty" yaml:"ShippedDate,omitempty"`

	CreatedDate time.Time `xml:"CreatedDate,omitempty" json:"CreatedDate,omitempty" yaml:"CreatedDate,omitempty"`

	CreatedBy string `xml:"CreatedBy,omitempty" json:"CreatedBy,omitempty" yaml:"CreatedBy,omitempty"`

	ModifiedBy string `xml:"ModifiedBy,omitempty" json:"ModifiedBy,omitempty" yaml:"ModifiedBy,omitempty"`

	TaxFedRate float64 `xml:"TaxFedRate,omitempty" json:"TaxFedRate,omitempty" yaml:"TaxFedRate,omitempty"`

	TaxStateRate float64 `xml:"TaxStateRate,omitempty" json:"TaxStateRate,omitempty" yaml:"TaxStateRate,omitempty"`

	TaxCityRate float64 `xml:"TaxCityRate,omitempty" json:"TaxCityRate,omitempty" yaml:"TaxCityRate,omitempty"`

	TaxCityLocalRate float64 `xml:"TaxCityLocalRate,omitempty" json:"TaxCityLocalRate,omitempty" yaml:"TaxCityLocalRate,omitempty"`

	TaxCountyRate float64 `xml:"TaxCountyRate,omitempty" json:"TaxCountyRate,omitempty" yaml:"TaxCountyRate,omitempty"`

	TaxCountyLocalRate float64 `xml:"TaxCountyLocalRate,omitempty" json:"TaxCountyLocalRate,omitempty" yaml:"TaxCountyLocalRate,omitempty"`

	TaxManualRate float64 `xml:"TaxManualRate,omitempty" json:"TaxManualRate,omitempty" yaml:"TaxManualRate,omitempty"`

	TaxCity string `xml:"TaxCity,omitempty" json:"TaxCity,omitempty" yaml:"TaxCity,omitempty"`

	TaxCounty string `xml:"TaxCounty,omitempty" json:"TaxCounty,omitempty" yaml:"TaxCounty,omitempty"`

	TaxState string `xml:"TaxState,omitempty" json:"TaxState,omitempty" yaml:"TaxState,omitempty"`

	TaxZip string `xml:"TaxZip,omitempty" json:"TaxZip,omitempty" yaml:"TaxZip,omitempty"`

	TaxCountry string `xml:"TaxCountry,omitempty" json:"TaxCountry,omitempty" yaml:"TaxCountry,omitempty"`

	TaxIsExempt bool `xml:"TaxIsExempt,omitempty" json:"TaxIsExempt,omitempty" yaml:"TaxIsExempt,omitempty"`

	TaxIsOverRide bool `xml:"TaxIsOverRide,omitempty" json:"TaxIsOverRide,omitempty" yaml:"TaxIsOverRide,omitempty"`

	Details *ArrayOfOrderDetailResponse `xml:"Details,omitempty" json:"Details,omitempty" yaml:"Details,omitempty"`

	Payments *ArrayOfPaymentResponse `xml:"Payments,omitempty" json:"Payments,omitempty" yaml:"Payments,omitempty"`

	ExpectedPayments *ArrayOfExpectedPaymentResponse `xml:"ExpectedPayments,omitempty" json:"ExpectedPayments,omitempty" yaml:"ExpectedPayments,omitempty"`

	Customer *CustomerResponse `xml:"Customer,omitempty" json:"Customer,omitempty" yaml:"Customer,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty" json:"MiddleName,omitempty" yaml:"MiddleName,omitempty"`

	NameSuffix string `xml:"NameSuffix,omitempty" json:"NameSuffix,omitempty" yaml:"NameSuffix,omitempty"`

	AutoOrderID int32 `xml:"AutoOrderID,omitempty" json:"AutoOrderID,omitempty" yaml:"AutoOrderID,omitempty"`

	PartyID int32 `xml:"PartyID,omitempty" json:"PartyID,omitempty" yaml:"PartyID,omitempty"`

	Reference1 string `xml:"Reference1,omitempty" json:"Reference1,omitempty" yaml:"Reference1,omitempty"`

	IsRMA bool `xml:"IsRMA,omitempty" json:"IsRMA,omitempty" yaml:"IsRMA,omitempty"`
}

type ArrayOfExpectedPaymentResponse struct {
	ExpectedPaymentResponse []*ExpectedPaymentResponse `xml:"ExpectedPaymentResponse,omitempty" json:"ExpectedPaymentResponse,omitempty" yaml:"ExpectedPaymentResponse,omitempty"`
}

type ExpectedPaymentResponse struct {
	ExpectedPaymentID int32 `xml:"ExpectedPaymentID,omitempty" json:"ExpectedPaymentID,omitempty" yaml:"ExpectedPaymentID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	PaymentType *PaymentType `xml:"PaymentType,omitempty" json:"PaymentType,omitempty" yaml:"PaymentType,omitempty"`

	PaymentDate time.Time `xml:"PaymentDate,omitempty" json:"PaymentDate,omitempty" yaml:"PaymentDate,omitempty"`

	Amount float64 `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty" yaml:"CurrencyCode,omitempty"`

	BillingName string `xml:"BillingName,omitempty" json:"BillingName,omitempty" yaml:"BillingName,omitempty"`

	AuthorizationCode string `xml:"AuthorizationCode,omitempty" json:"AuthorizationCode,omitempty" yaml:"AuthorizationCode,omitempty"`
}

type ArrayOfOrderResponse struct {
	OrderResponse []*OrderResponse `xml:"OrderResponse,omitempty" json:"OrderResponse,omitempty" yaml:"OrderResponse,omitempty"`
}

type GetOrdersResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetOrdersResult"`

	*ApiResponse

	Orders *ArrayOfOrderResponse `xml:"Orders,omitempty" json:"Orders,omitempty" yaml:"Orders,omitempty"`

	RecordCount int32 `xml:"RecordCount,omitempty" json:"RecordCount,omitempty" yaml:"RecordCount,omitempty"`
}

type GetOrderTotalsRequest struct {
	*ApiRequest

	StartDate time.Time `xml:"StartDate,omitempty" json:"StartDate,omitempty" yaml:"StartDate,omitempty"`

	EndDate time.Time `xml:"EndDate,omitempty" json:"EndDate,omitempty" yaml:"EndDate,omitempty"`
}

type OrderTotalByCurrency struct {
	CurrencyCode string `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty" yaml:"CurrencyCode,omitempty"`

	Amount float64 `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`

	Count int32 `xml:"Count,omitempty" json:"Count,omitempty" yaml:"Count,omitempty"`
}

type ArrayOfOrderTotalByCurrency struct {
	OrderTotalByCurrency []*OrderTotalByCurrency `xml:"OrderTotalByCurrency,omitempty" json:"OrderTotalByCurrency,omitempty" yaml:"OrderTotalByCurrency,omitempty"`
}

type GetOrderTotalsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetOrderTotalsResult"`

	*ApiResponse

	StartDate time.Time `xml:"StartDate,omitempty" json:"StartDate,omitempty" yaml:"StartDate,omitempty"`

	EndDate time.Time `xml:"EndDate,omitempty" json:"EndDate,omitempty" yaml:"EndDate,omitempty"`

	AcceptedByCurrency *ArrayOfOrderTotalByCurrency `xml:"AcceptedByCurrency,omitempty" json:"AcceptedByCurrency,omitempty" yaml:"AcceptedByCurrency,omitempty"`
}

type GetAutoOrdersRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	AutoOrderID int32 `xml:"AutoOrderID,omitempty" json:"AutoOrderID,omitempty" yaml:"AutoOrderID,omitempty"`

	AutoOrderStatus *AutoOrderStatusType `xml:"AutoOrderStatus,omitempty" json:"AutoOrderStatus,omitempty" yaml:"AutoOrderStatus,omitempty"`
}

type AutoOrderResponse struct {
	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	AutoOrderID int32 `xml:"AutoOrderID,omitempty" json:"AutoOrderID,omitempty" yaml:"AutoOrderID,omitempty"`

	AutoOrderStatus *AutoOrderStatusType `xml:"AutoOrderStatus,omitempty" json:"AutoOrderStatus,omitempty" yaml:"AutoOrderStatus,omitempty"`

	Frequency *FrequencyType `xml:"Frequency,omitempty" json:"Frequency,omitempty" yaml:"Frequency,omitempty"`

	StartDate time.Time `xml:"StartDate,omitempty" json:"StartDate,omitempty" yaml:"StartDate,omitempty"`

	StopDate time.Time `xml:"StopDate,omitempty" json:"StopDate,omitempty" yaml:"StopDate,omitempty"`

	LastRunDate time.Time `xml:"LastRunDate,omitempty" json:"LastRunDate,omitempty" yaml:"LastRunDate,omitempty"`

	NextRunDate time.Time `xml:"NextRunDate,omitempty" json:"NextRunDate,omitempty" yaml:"NextRunDate,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty" yaml:"CurrencyCode,omitempty"`

	WarehouseID int32 `xml:"WarehouseID,omitempty" json:"WarehouseID,omitempty" yaml:"WarehouseID,omitempty"`

	ShipMethodID int32 `xml:"ShipMethodID,omitempty" json:"ShipMethodID,omitempty" yaml:"ShipMethodID,omitempty"`

	PaymentType *AutoOrderPaymentType `xml:"PaymentType,omitempty" json:"PaymentType,omitempty" yaml:"PaymentType,omitempty"`

	ProcessType *AutoOrderProcessType `xml:"ProcessType,omitempty" json:"ProcessType,omitempty" yaml:"ProcessType,omitempty"`

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`

	Company string `xml:"Company,omitempty" json:"Company,omitempty" yaml:"Company,omitempty"`

	Address1 string `xml:"Address1,omitempty" json:"Address1,omitempty" yaml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty" json:"Address2,omitempty" yaml:"Address2,omitempty"`

	Address3 string `xml:"Address3,omitempty" json:"Address3,omitempty" yaml:"Address3,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty" json:"Zip,omitempty" yaml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`

	County string `xml:"County,omitempty" json:"County,omitempty" yaml:"County,omitempty"`

	Email string `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`

	Phone string `xml:"Phone,omitempty" json:"Phone,omitempty" yaml:"Phone,omitempty"`

	Notes string `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`

	Total float64 `xml:"Total,omitempty" json:"Total,omitempty" yaml:"Total,omitempty"`

	SubTotal float64 `xml:"SubTotal,omitempty" json:"SubTotal,omitempty" yaml:"SubTotal,omitempty"`

	TaxTotal float64 `xml:"TaxTotal,omitempty" json:"TaxTotal,omitempty" yaml:"TaxTotal,omitempty"`

	ShippingTotal float64 `xml:"ShippingTotal,omitempty" json:"ShippingTotal,omitempty" yaml:"ShippingTotal,omitempty"`

	DiscountTotal float64 `xml:"DiscountTotal,omitempty" json:"DiscountTotal,omitempty" yaml:"DiscountTotal,omitempty"`

	BusinessVolumeTotal float64 `xml:"BusinessVolumeTotal,omitempty" json:"BusinessVolumeTotal,omitempty" yaml:"BusinessVolumeTotal,omitempty"`

	CommissionableVolumeTotal float64 `xml:"CommissionableVolumeTotal,omitempty" json:"CommissionableVolumeTotal,omitempty" yaml:"CommissionableVolumeTotal,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	Other11 string `xml:"Other11,omitempty" json:"Other11,omitempty" yaml:"Other11,omitempty"`

	Other12 string `xml:"Other12,omitempty" json:"Other12,omitempty" yaml:"Other12,omitempty"`

	Other13 string `xml:"Other13,omitempty" json:"Other13,omitempty" yaml:"Other13,omitempty"`

	Other14 string `xml:"Other14,omitempty" json:"Other14,omitempty" yaml:"Other14,omitempty"`

	Other15 string `xml:"Other15,omitempty" json:"Other15,omitempty" yaml:"Other15,omitempty"`

	Other16 string `xml:"Other16,omitempty" json:"Other16,omitempty" yaml:"Other16,omitempty"`

	Other17 string `xml:"Other17,omitempty" json:"Other17,omitempty" yaml:"Other17,omitempty"`

	Other18 string `xml:"Other18,omitempty" json:"Other18,omitempty" yaml:"Other18,omitempty"`

	Other19 string `xml:"Other19,omitempty" json:"Other19,omitempty" yaml:"Other19,omitempty"`

	Other20 string `xml:"Other20,omitempty" json:"Other20,omitempty" yaml:"Other20,omitempty"`

	Details *ArrayOfAutoOrderDetailResponse `xml:"Details,omitempty" json:"Details,omitempty" yaml:"Details,omitempty"`

	ModifiedDate time.Time `xml:"ModifiedDate,omitempty" json:"ModifiedDate,omitempty" yaml:"ModifiedDate,omitempty"`

	ModifiedBy string `xml:"ModifiedBy,omitempty" json:"ModifiedBy,omitempty" yaml:"ModifiedBy,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty" json:"MiddleName,omitempty" yaml:"MiddleName,omitempty"`

	NameSuffix string `xml:"NameSuffix,omitempty" json:"NameSuffix,omitempty" yaml:"NameSuffix,omitempty"`

	SpecificDayInterval int32 `xml:"SpecificDayInterval,omitempty" json:"SpecificDayInterval,omitempty" yaml:"SpecificDayInterval,omitempty"`
}

type ArrayOfAutoOrderDetailResponse struct {
	AutoOrderDetailResponse []*AutoOrderDetailResponse `xml:"AutoOrderDetailResponse,omitempty" json:"AutoOrderDetailResponse,omitempty" yaml:"AutoOrderDetailResponse,omitempty"`
}

type AutoOrderDetailResponse struct {
	ItemCode string `xml:"ItemCode,omitempty" json:"ItemCode,omitempty" yaml:"ItemCode,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	Quantity float64 `xml:"Quantity,omitempty" json:"Quantity,omitempty" yaml:"Quantity,omitempty"`

	PriceEach float64 `xml:"PriceEach,omitempty" json:"PriceEach,omitempty" yaml:"PriceEach,omitempty"`

	PriceTotal float64 `xml:"PriceTotal,omitempty" json:"PriceTotal,omitempty" yaml:"PriceTotal,omitempty"`

	BusinessVolumeEach float64 `xml:"BusinessVolumeEach,omitempty" json:"BusinessVolumeEach,omitempty" yaml:"BusinessVolumeEach,omitempty"`

	BusinesVolume float64 `xml:"BusinesVolume,omitempty" json:"BusinesVolume,omitempty" yaml:"BusinesVolume,omitempty"`

	CommissionableVolumeEach float64 `xml:"CommissionableVolumeEach,omitempty" json:"CommissionableVolumeEach,omitempty" yaml:"CommissionableVolumeEach,omitempty"`

	CommissionableVolume float64 `xml:"CommissionableVolume,omitempty" json:"CommissionableVolume,omitempty" yaml:"CommissionableVolume,omitempty"`

	ParentItemCode string `xml:"ParentItemCode,omitempty" json:"ParentItemCode,omitempty" yaml:"ParentItemCode,omitempty"`

	PriceEachOverride float64 `xml:"PriceEachOverride,omitempty" json:"PriceEachOverride,omitempty" yaml:"PriceEachOverride,omitempty"`

	TaxableEachOverride float64 `xml:"TaxableEachOverride,omitempty" json:"TaxableEachOverride,omitempty" yaml:"TaxableEachOverride,omitempty"`

	ShippingPriceEachOverride float64 `xml:"ShippingPriceEachOverride,omitempty" json:"ShippingPriceEachOverride,omitempty" yaml:"ShippingPriceEachOverride,omitempty"`

	BusinessVolumeEachOverride float64 `xml:"BusinessVolumeEachOverride,omitempty" json:"BusinessVolumeEachOverride,omitempty" yaml:"BusinessVolumeEachOverride,omitempty"`

	CommissionableVolumeEachOverride float64 `xml:"CommissionableVolumeEachOverride,omitempty" json:"CommissionableVolumeEachOverride,omitempty" yaml:"CommissionableVolumeEachOverride,omitempty"`

	Reference1 string `xml:"Reference1,omitempty" json:"Reference1,omitempty" yaml:"Reference1,omitempty"`

	ProcessWhileDate time.Time `xml:"ProcessWhileDate,omitempty" json:"ProcessWhileDate,omitempty" yaml:"ProcessWhileDate,omitempty"`

	SkipUntilDate time.Time `xml:"SkipUntilDate,omitempty" json:"SkipUntilDate,omitempty" yaml:"SkipUntilDate,omitempty"`
}

type ArrayOfAutoOrderResponse struct {
	AutoOrderResponse []*AutoOrderResponse `xml:"AutoOrderResponse,omitempty" json:"AutoOrderResponse,omitempty" yaml:"AutoOrderResponse,omitempty"`
}

type GetAutoOrdersResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetAutoOrdersResult"`

	*ApiResponse

	AutoOrders *ArrayOfAutoOrderResponse `xml:"AutoOrders,omitempty" json:"AutoOrders,omitempty" yaml:"AutoOrders,omitempty"`
}

type ChangeOrderStatusBatchRequest struct {
	*ApiRequest

	OrderStatus *OrderStatusType `xml:"OrderStatus,omitempty" json:"OrderStatus,omitempty" yaml:"OrderStatus,omitempty"`

	Details *ArrayOfOrderBatchDetailRequest `xml:"Details,omitempty" json:"Details,omitempty" yaml:"Details,omitempty"`
}

type ArrayOfOrderBatchDetailRequest struct {
	OrderBatchDetailRequest []*OrderBatchDetailRequest `xml:"OrderBatchDetailRequest,omitempty" json:"OrderBatchDetailRequest,omitempty" yaml:"OrderBatchDetailRequest,omitempty"`
}

type OrderBatchDetailRequest struct {
	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`

	TrackingNumber1 string `xml:"TrackingNumber1,omitempty" json:"TrackingNumber1,omitempty" yaml:"TrackingNumber1,omitempty"`

	TrackingNumber2 string `xml:"TrackingNumber2,omitempty" json:"TrackingNumber2,omitempty" yaml:"TrackingNumber2,omitempty"`

	TrackingNumber3 string `xml:"TrackingNumber3,omitempty" json:"TrackingNumber3,omitempty" yaml:"TrackingNumber3,omitempty"`

	TrackingNumber4 string `xml:"TrackingNumber4,omitempty" json:"TrackingNumber4,omitempty" yaml:"TrackingNumber4,omitempty"`

	TrackingNumber5 string `xml:"TrackingNumber5,omitempty" json:"TrackingNumber5,omitempty" yaml:"TrackingNumber5,omitempty"`

	ShippedDate time.Time `xml:"ShippedDate,omitempty" json:"ShippedDate,omitempty" yaml:"ShippedDate,omitempty"`
}

type ChangeOrderStatusBatchResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ChangeOrderStatusBatchResult"`

	*ApiResponse
}

type MergeCustomerRequest struct {
	*ApiRequest

	ToCustomerID int32 `xml:"ToCustomerID,omitempty" json:"ToCustomerID,omitempty" yaml:"ToCustomerID,omitempty"`

	FromCustomerID int32 `xml:"FromCustomerID,omitempty" json:"FromCustomerID,omitempty" yaml:"FromCustomerID,omitempty"`
}

type MergeCustomerResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ MergeCustomerResult"`

	*ApiResponse
}

type PlaceEnrollerNodeRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	ToEnrollerID int32 `xml:"ToEnrollerID,omitempty" json:"ToEnrollerID,omitempty" yaml:"ToEnrollerID,omitempty"`

	Reason string `xml:"Reason,omitempty" json:"Reason,omitempty" yaml:"Reason,omitempty"`
}

type PlaceEnrollerNodeResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ PlaceEnrollerNodeResult"`

	*ApiResponse
}

type PlaceStackNodeRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	ToParentID int32 `xml:"ToParentID,omitempty" json:"ToParentID,omitempty" yaml:"ToParentID,omitempty"`

	Reason string `xml:"Reason,omitempty" json:"Reason,omitempty" yaml:"Reason,omitempty"`
}

type PlaceStackNodeResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ PlaceStackNodeResult"`

	*ApiResponse
}

type PlaceUniLevelNodeRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	ToSponsorID int32 `xml:"ToSponsorID,omitempty" json:"ToSponsorID,omitempty" yaml:"ToSponsorID,omitempty"`

	Reason string `xml:"Reason,omitempty" json:"Reason,omitempty" yaml:"Reason,omitempty"`

	OptionalPlacement int32 `xml:"OptionalPlacement,omitempty" json:"OptionalPlacement,omitempty" yaml:"OptionalPlacement,omitempty"`

	OptionalFindAvailable bool `xml:"OptionalFindAvailable,omitempty" json:"OptionalFindAvailable,omitempty" yaml:"OptionalFindAvailable,omitempty"`

	OptionalUnilevelBuildTypeID int32 `xml:"OptionalUnilevelBuildTypeID,omitempty" json:"OptionalUnilevelBuildTypeID,omitempty" yaml:"OptionalUnilevelBuildTypeID,omitempty"`
}

type PlaceUniLevelNodeResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ PlaceUniLevelNodeResult"`

	*ApiResponse
}

type PlaceBinaryNodeRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	ToParentID int32 `xml:"ToParentID,omitempty" json:"ToParentID,omitempty" yaml:"ToParentID,omitempty"`

	PlacementType *BinaryPlacementType `xml:"PlacementType,omitempty" json:"PlacementType,omitempty" yaml:"PlacementType,omitempty"`

	Reason string `xml:"Reason,omitempty" json:"Reason,omitempty" yaml:"Reason,omitempty"`
}

type PlaceBinaryNodeResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ PlaceBinaryNodeResult"`

	*ApiResponse
}

type GetBinaryPreferenceRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`
}

type GetBinaryPreferenceResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetBinaryPreferenceResult"`

	*ApiResponse

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	PlacementType *BinaryPlacementType `xml:"PlacementType,omitempty" json:"PlacementType,omitempty" yaml:"PlacementType,omitempty"`
}

type SetBinaryPreferenceRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	PlacementType *BinaryPlacementType `xml:"PlacementType,omitempty" json:"PlacementType,omitempty" yaml:"PlacementType,omitempty"`
}

type SetBinaryPreferenceResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetBinaryPreferenceResult"`

	*ApiResponse
}

type PlaceMatrixNodeRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	ToParentCustomerID int32 `xml:"ToParentCustomerID,omitempty" json:"ToParentCustomerID,omitempty" yaml:"ToParentCustomerID,omitempty"`

	ToParentMatrixID int32 `xml:"ToParentMatrixID,omitempty" json:"ToParentMatrixID,omitempty" yaml:"ToParentMatrixID,omitempty"`

	Reason string `xml:"Reason,omitempty" json:"Reason,omitempty" yaml:"Reason,omitempty"`

	Placement int32 `xml:"Placement,omitempty" json:"Placement,omitempty" yaml:"Placement,omitempty"`
}

type PlaceMatrixNodeResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ PlaceMatrixNodeResult"`

	*ApiResponse
}

type GetCountryRegionsRequest struct {
	*ApiRequest

	CountryCode string `xml:"CountryCode,omitempty" json:"CountryCode,omitempty" yaml:"CountryCode,omitempty"`
}

type RegionResponse struct {
	RegionCode string `xml:"RegionCode,omitempty" json:"RegionCode,omitempty" yaml:"RegionCode,omitempty"`

	RegionName string `xml:"RegionName,omitempty" json:"RegionName,omitempty" yaml:"RegionName,omitempty"`
}

type CountryResponse struct {
	CountryCode string `xml:"CountryCode,omitempty" json:"CountryCode,omitempty" yaml:"CountryCode,omitempty"`

	CountryName string `xml:"CountryName,omitempty" json:"CountryName,omitempty" yaml:"CountryName,omitempty"`
}

type ArrayOfCountryResponse struct {
	CountryResponse []*CountryResponse `xml:"CountryResponse,omitempty" json:"CountryResponse,omitempty" yaml:"CountryResponse,omitempty"`
}

type ArrayOfRegionResponse struct {
	RegionResponse []*RegionResponse `xml:"RegionResponse,omitempty" json:"RegionResponse,omitempty" yaml:"RegionResponse,omitempty"`
}

type GetCountryRegionsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCountryRegionsResult"`

	*ApiResponse

	Countries *ArrayOfCountryResponse `xml:"Countries,omitempty" json:"Countries,omitempty" yaml:"Countries,omitempty"`

	SelectedCountry string `xml:"SelectedCountry,omitempty" json:"SelectedCountry,omitempty" yaml:"SelectedCountry,omitempty"`

	Regions *ArrayOfRegionResponse `xml:"Regions,omitempty" json:"Regions,omitempty" yaml:"Regions,omitempty"`
}

type GetDownlineRequest struct {
	*ApiRequest

	TreeType *TreeType `xml:"TreeType,omitempty" json:"TreeType,omitempty" yaml:"TreeType,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	PeriodType int32 `xml:"PeriodType,omitempty" json:"PeriodType,omitempty" yaml:"PeriodType,omitempty"`

	PeriodID int32 `xml:"PeriodID,omitempty" json:"PeriodID,omitempty" yaml:"PeriodID,omitempty"`

	MaxLevelDepth int32 `xml:"MaxLevelDepth,omitempty" json:"MaxLevelDepth,omitempty" yaml:"MaxLevelDepth,omitempty"`

	CustomerTypes *ArrayOfInt `xml:"CustomerTypes,omitempty" json:"CustomerTypes,omitempty" yaml:"CustomerTypes,omitempty"`

	Ranks *ArrayOfInt `xml:"Ranks,omitempty" json:"Ranks,omitempty" yaml:"Ranks,omitempty"`

	PayRanks *ArrayOfInt `xml:"PayRanks,omitempty" json:"PayRanks,omitempty" yaml:"PayRanks,omitempty"`

	VolumeFilters *ArrayOfVolumeFilter `xml:"VolumeFilters,omitempty" json:"VolumeFilters,omitempty" yaml:"VolumeFilters,omitempty"`

	CustomerStatusTypes *ArrayOfInt `xml:"CustomerStatusTypes,omitempty" json:"CustomerStatusTypes,omitempty" yaml:"CustomerStatusTypes,omitempty"`

	BatchSize int32 `xml:"BatchSize,omitempty" json:"BatchSize,omitempty" yaml:"BatchSize,omitempty"`

	SortByLevel bool `xml:"SortByLevel,omitempty" json:"SortByLevel,omitempty" yaml:"SortByLevel,omitempty"`

	BatchOffset int32 `xml:"BatchOffset,omitempty" json:"BatchOffset,omitempty" yaml:"BatchOffset,omitempty"`
}

type ArrayOfVolumeFilter struct {
	VolumeFilter []*VolumeFilter `xml:"VolumeFilter,omitempty" json:"VolumeFilter,omitempty" yaml:"VolumeFilter,omitempty"`
}

type VolumeFilter struct {
	VolumeID int32 `xml:"VolumeID,omitempty" json:"VolumeID,omitempty" yaml:"VolumeID,omitempty"`

	Compare *NumericCompareType `xml:"Compare,omitempty" json:"Compare,omitempty" yaml:"Compare,omitempty"`

	Value float64 `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
}

type NodeResponse struct {
	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	NodeID int32 `xml:"NodeID,omitempty" json:"NodeID,omitempty" yaml:"NodeID,omitempty"`

	ParentID int32 `xml:"ParentID,omitempty" json:"ParentID,omitempty" yaml:"ParentID,omitempty"`

	Level int32 `xml:"Level,omitempty" json:"Level,omitempty" yaml:"Level,omitempty"`

	Position int32 `xml:"Position,omitempty" json:"Position,omitempty" yaml:"Position,omitempty"`

	CustomerType int32 `xml:"CustomerType,omitempty" json:"CustomerType,omitempty" yaml:"CustomerType,omitempty"`

	CustomerStatus int32 `xml:"CustomerStatus,omitempty" json:"CustomerStatus,omitempty" yaml:"CustomerStatus,omitempty"`

	RankID int32 `xml:"RankID,omitempty" json:"RankID,omitempty" yaml:"RankID,omitempty"`

	PayRankID int32 `xml:"PayRankID,omitempty" json:"PayRankID,omitempty" yaml:"PayRankID,omitempty"`

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`

	Company string `xml:"Company,omitempty" json:"Company,omitempty" yaml:"Company,omitempty"`

	Email string `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`

	Phone string `xml:"Phone,omitempty" json:"Phone,omitempty" yaml:"Phone,omitempty"`

	CreatedDate time.Time `xml:"CreatedDate,omitempty" json:"CreatedDate,omitempty" yaml:"CreatedDate,omitempty"`

	Volume1 float64 `xml:"Volume1,omitempty" json:"Volume1,omitempty" yaml:"Volume1,omitempty"`

	Volume2 float64 `xml:"Volume2,omitempty" json:"Volume2,omitempty" yaml:"Volume2,omitempty"`

	Volume3 float64 `xml:"Volume3,omitempty" json:"Volume3,omitempty" yaml:"Volume3,omitempty"`

	Volume4 float64 `xml:"Volume4,omitempty" json:"Volume4,omitempty" yaml:"Volume4,omitempty"`

	Volume5 float64 `xml:"Volume5,omitempty" json:"Volume5,omitempty" yaml:"Volume5,omitempty"`

	Volume6 float64 `xml:"Volume6,omitempty" json:"Volume6,omitempty" yaml:"Volume6,omitempty"`

	Volume7 float64 `xml:"Volume7,omitempty" json:"Volume7,omitempty" yaml:"Volume7,omitempty"`

	Volume8 float64 `xml:"Volume8,omitempty" json:"Volume8,omitempty" yaml:"Volume8,omitempty"`

	Volume9 float64 `xml:"Volume9,omitempty" json:"Volume9,omitempty" yaml:"Volume9,omitempty"`

	Volume10 float64 `xml:"Volume10,omitempty" json:"Volume10,omitempty" yaml:"Volume10,omitempty"`

	Volume11 float64 `xml:"Volume11,omitempty" json:"Volume11,omitempty" yaml:"Volume11,omitempty"`

	Volume12 float64 `xml:"Volume12,omitempty" json:"Volume12,omitempty" yaml:"Volume12,omitempty"`

	Volume13 float64 `xml:"Volume13,omitempty" json:"Volume13,omitempty" yaml:"Volume13,omitempty"`

	Volume14 float64 `xml:"Volume14,omitempty" json:"Volume14,omitempty" yaml:"Volume14,omitempty"`

	Volume15 float64 `xml:"Volume15,omitempty" json:"Volume15,omitempty" yaml:"Volume15,omitempty"`

	Volume16 float64 `xml:"Volume16,omitempty" json:"Volume16,omitempty" yaml:"Volume16,omitempty"`

	Volume17 float64 `xml:"Volume17,omitempty" json:"Volume17,omitempty" yaml:"Volume17,omitempty"`

	Volume18 float64 `xml:"Volume18,omitempty" json:"Volume18,omitempty" yaml:"Volume18,omitempty"`

	Volume19 float64 `xml:"Volume19,omitempty" json:"Volume19,omitempty" yaml:"Volume19,omitempty"`

	Volume20 float64 `xml:"Volume20,omitempty" json:"Volume20,omitempty" yaml:"Volume20,omitempty"`

	Volume21 float64 `xml:"Volume21,omitempty" json:"Volume21,omitempty" yaml:"Volume21,omitempty"`

	Volume22 float64 `xml:"Volume22,omitempty" json:"Volume22,omitempty" yaml:"Volume22,omitempty"`

	Volume23 float64 `xml:"Volume23,omitempty" json:"Volume23,omitempty" yaml:"Volume23,omitempty"`

	Volume24 float64 `xml:"Volume24,omitempty" json:"Volume24,omitempty" yaml:"Volume24,omitempty"`

	Volume25 float64 `xml:"Volume25,omitempty" json:"Volume25,omitempty" yaml:"Volume25,omitempty"`

	Volume26 float64 `xml:"Volume26,omitempty" json:"Volume26,omitempty" yaml:"Volume26,omitempty"`

	Volume27 float64 `xml:"Volume27,omitempty" json:"Volume27,omitempty" yaml:"Volume27,omitempty"`

	Volume28 float64 `xml:"Volume28,omitempty" json:"Volume28,omitempty" yaml:"Volume28,omitempty"`

	Volume29 float64 `xml:"Volume29,omitempty" json:"Volume29,omitempty" yaml:"Volume29,omitempty"`

	Volume30 float64 `xml:"Volume30,omitempty" json:"Volume30,omitempty" yaml:"Volume30,omitempty"`

	Volume31 float64 `xml:"Volume31,omitempty" json:"Volume31,omitempty" yaml:"Volume31,omitempty"`

	Volume32 float64 `xml:"Volume32,omitempty" json:"Volume32,omitempty" yaml:"Volume32,omitempty"`

	Volume33 float64 `xml:"Volume33,omitempty" json:"Volume33,omitempty" yaml:"Volume33,omitempty"`

	Volume34 float64 `xml:"Volume34,omitempty" json:"Volume34,omitempty" yaml:"Volume34,omitempty"`

	Volume35 float64 `xml:"Volume35,omitempty" json:"Volume35,omitempty" yaml:"Volume35,omitempty"`

	Volume36 float64 `xml:"Volume36,omitempty" json:"Volume36,omitempty" yaml:"Volume36,omitempty"`

	Volume37 float64 `xml:"Volume37,omitempty" json:"Volume37,omitempty" yaml:"Volume37,omitempty"`

	Volume38 float64 `xml:"Volume38,omitempty" json:"Volume38,omitempty" yaml:"Volume38,omitempty"`

	Volume39 float64 `xml:"Volume39,omitempty" json:"Volume39,omitempty" yaml:"Volume39,omitempty"`

	Volume40 float64 `xml:"Volume40,omitempty" json:"Volume40,omitempty" yaml:"Volume40,omitempty"`

	Volume41 float64 `xml:"Volume41,omitempty" json:"Volume41,omitempty" yaml:"Volume41,omitempty"`

	Volume42 float64 `xml:"Volume42,omitempty" json:"Volume42,omitempty" yaml:"Volume42,omitempty"`

	Volume43 float64 `xml:"Volume43,omitempty" json:"Volume43,omitempty" yaml:"Volume43,omitempty"`

	Volume44 float64 `xml:"Volume44,omitempty" json:"Volume44,omitempty" yaml:"Volume44,omitempty"`

	Volume45 float64 `xml:"Volume45,omitempty" json:"Volume45,omitempty" yaml:"Volume45,omitempty"`

	Volume46 float64 `xml:"Volume46,omitempty" json:"Volume46,omitempty" yaml:"Volume46,omitempty"`

	Volume47 float64 `xml:"Volume47,omitempty" json:"Volume47,omitempty" yaml:"Volume47,omitempty"`

	Volume48 float64 `xml:"Volume48,omitempty" json:"Volume48,omitempty" yaml:"Volume48,omitempty"`

	Volume49 float64 `xml:"Volume49,omitempty" json:"Volume49,omitempty" yaml:"Volume49,omitempty"`

	Volume50 float64 `xml:"Volume50,omitempty" json:"Volume50,omitempty" yaml:"Volume50,omitempty"`

	Volume51 float64 `xml:"Volume51,omitempty" json:"Volume51,omitempty" yaml:"Volume51,omitempty"`

	Volume52 float64 `xml:"Volume52,omitempty" json:"Volume52,omitempty" yaml:"Volume52,omitempty"`

	Volume53 float64 `xml:"Volume53,omitempty" json:"Volume53,omitempty" yaml:"Volume53,omitempty"`

	Volume54 float64 `xml:"Volume54,omitempty" json:"Volume54,omitempty" yaml:"Volume54,omitempty"`

	Volume55 float64 `xml:"Volume55,omitempty" json:"Volume55,omitempty" yaml:"Volume55,omitempty"`

	Volume56 float64 `xml:"Volume56,omitempty" json:"Volume56,omitempty" yaml:"Volume56,omitempty"`

	Volume57 float64 `xml:"Volume57,omitempty" json:"Volume57,omitempty" yaml:"Volume57,omitempty"`

	Volume58 float64 `xml:"Volume58,omitempty" json:"Volume58,omitempty" yaml:"Volume58,omitempty"`

	Volume59 float64 `xml:"Volume59,omitempty" json:"Volume59,omitempty" yaml:"Volume59,omitempty"`

	Volume60 float64 `xml:"Volume60,omitempty" json:"Volume60,omitempty" yaml:"Volume60,omitempty"`

	Volume61 float64 `xml:"Volume61,omitempty" json:"Volume61,omitempty" yaml:"Volume61,omitempty"`

	Volume62 float64 `xml:"Volume62,omitempty" json:"Volume62,omitempty" yaml:"Volume62,omitempty"`

	Volume63 float64 `xml:"Volume63,omitempty" json:"Volume63,omitempty" yaml:"Volume63,omitempty"`

	Volume64 float64 `xml:"Volume64,omitempty" json:"Volume64,omitempty" yaml:"Volume64,omitempty"`

	Volume65 float64 `xml:"Volume65,omitempty" json:"Volume65,omitempty" yaml:"Volume65,omitempty"`

	Volume66 float64 `xml:"Volume66,omitempty" json:"Volume66,omitempty" yaml:"Volume66,omitempty"`

	Volume67 float64 `xml:"Volume67,omitempty" json:"Volume67,omitempty" yaml:"Volume67,omitempty"`

	Volume68 float64 `xml:"Volume68,omitempty" json:"Volume68,omitempty" yaml:"Volume68,omitempty"`

	Volume69 float64 `xml:"Volume69,omitempty" json:"Volume69,omitempty" yaml:"Volume69,omitempty"`

	Volume70 float64 `xml:"Volume70,omitempty" json:"Volume70,omitempty" yaml:"Volume70,omitempty"`

	Volume71 float64 `xml:"Volume71,omitempty" json:"Volume71,omitempty" yaml:"Volume71,omitempty"`

	Volume72 float64 `xml:"Volume72,omitempty" json:"Volume72,omitempty" yaml:"Volume72,omitempty"`

	Volume73 float64 `xml:"Volume73,omitempty" json:"Volume73,omitempty" yaml:"Volume73,omitempty"`

	Volume74 float64 `xml:"Volume74,omitempty" json:"Volume74,omitempty" yaml:"Volume74,omitempty"`

	Volume75 float64 `xml:"Volume75,omitempty" json:"Volume75,omitempty" yaml:"Volume75,omitempty"`

	Volume76 float64 `xml:"Volume76,omitempty" json:"Volume76,omitempty" yaml:"Volume76,omitempty"`

	Volume77 float64 `xml:"Volume77,omitempty" json:"Volume77,omitempty" yaml:"Volume77,omitempty"`

	Volume78 float64 `xml:"Volume78,omitempty" json:"Volume78,omitempty" yaml:"Volume78,omitempty"`

	Volume79 float64 `xml:"Volume79,omitempty" json:"Volume79,omitempty" yaml:"Volume79,omitempty"`

	Volume80 float64 `xml:"Volume80,omitempty" json:"Volume80,omitempty" yaml:"Volume80,omitempty"`

	Volume81 float64 `xml:"Volume81,omitempty" json:"Volume81,omitempty" yaml:"Volume81,omitempty"`

	Volume82 float64 `xml:"Volume82,omitempty" json:"Volume82,omitempty" yaml:"Volume82,omitempty"`

	Volume83 float64 `xml:"Volume83,omitempty" json:"Volume83,omitempty" yaml:"Volume83,omitempty"`

	Volume84 float64 `xml:"Volume84,omitempty" json:"Volume84,omitempty" yaml:"Volume84,omitempty"`

	Volume85 float64 `xml:"Volume85,omitempty" json:"Volume85,omitempty" yaml:"Volume85,omitempty"`

	Volume86 float64 `xml:"Volume86,omitempty" json:"Volume86,omitempty" yaml:"Volume86,omitempty"`

	Volume87 float64 `xml:"Volume87,omitempty" json:"Volume87,omitempty" yaml:"Volume87,omitempty"`

	Volume88 float64 `xml:"Volume88,omitempty" json:"Volume88,omitempty" yaml:"Volume88,omitempty"`

	Volume89 float64 `xml:"Volume89,omitempty" json:"Volume89,omitempty" yaml:"Volume89,omitempty"`

	Volume90 float64 `xml:"Volume90,omitempty" json:"Volume90,omitempty" yaml:"Volume90,omitempty"`

	Volume91 float64 `xml:"Volume91,omitempty" json:"Volume91,omitempty" yaml:"Volume91,omitempty"`

	Volume92 float64 `xml:"Volume92,omitempty" json:"Volume92,omitempty" yaml:"Volume92,omitempty"`

	Volume93 float64 `xml:"Volume93,omitempty" json:"Volume93,omitempty" yaml:"Volume93,omitempty"`

	Volume94 float64 `xml:"Volume94,omitempty" json:"Volume94,omitempty" yaml:"Volume94,omitempty"`

	Volume95 float64 `xml:"Volume95,omitempty" json:"Volume95,omitempty" yaml:"Volume95,omitempty"`

	Volume96 float64 `xml:"Volume96,omitempty" json:"Volume96,omitempty" yaml:"Volume96,omitempty"`

	Volume97 float64 `xml:"Volume97,omitempty" json:"Volume97,omitempty" yaml:"Volume97,omitempty"`

	Volume98 float64 `xml:"Volume98,omitempty" json:"Volume98,omitempty" yaml:"Volume98,omitempty"`

	Volume99 float64 `xml:"Volume99,omitempty" json:"Volume99,omitempty" yaml:"Volume99,omitempty"`

	Volume100 float64 `xml:"Volume100,omitempty" json:"Volume100,omitempty" yaml:"Volume100,omitempty"`

	Volume101 float64 `xml:"Volume101,omitempty" json:"Volume101,omitempty" yaml:"Volume101,omitempty"`

	Volume102 float64 `xml:"Volume102,omitempty" json:"Volume102,omitempty" yaml:"Volume102,omitempty"`

	Volume103 float64 `xml:"Volume103,omitempty" json:"Volume103,omitempty" yaml:"Volume103,omitempty"`

	Volume104 float64 `xml:"Volume104,omitempty" json:"Volume104,omitempty" yaml:"Volume104,omitempty"`

	Volume105 float64 `xml:"Volume105,omitempty" json:"Volume105,omitempty" yaml:"Volume105,omitempty"`

	Volume106 float64 `xml:"Volume106,omitempty" json:"Volume106,omitempty" yaml:"Volume106,omitempty"`

	Volume107 float64 `xml:"Volume107,omitempty" json:"Volume107,omitempty" yaml:"Volume107,omitempty"`

	Volume108 float64 `xml:"Volume108,omitempty" json:"Volume108,omitempty" yaml:"Volume108,omitempty"`

	Volume109 float64 `xml:"Volume109,omitempty" json:"Volume109,omitempty" yaml:"Volume109,omitempty"`

	Volume110 float64 `xml:"Volume110,omitempty" json:"Volume110,omitempty" yaml:"Volume110,omitempty"`

	Volume111 float64 `xml:"Volume111,omitempty" json:"Volume111,omitempty" yaml:"Volume111,omitempty"`

	Volume112 float64 `xml:"Volume112,omitempty" json:"Volume112,omitempty" yaml:"Volume112,omitempty"`

	Volume113 float64 `xml:"Volume113,omitempty" json:"Volume113,omitempty" yaml:"Volume113,omitempty"`

	Volume114 float64 `xml:"Volume114,omitempty" json:"Volume114,omitempty" yaml:"Volume114,omitempty"`

	Volume115 float64 `xml:"Volume115,omitempty" json:"Volume115,omitempty" yaml:"Volume115,omitempty"`

	Volume116 float64 `xml:"Volume116,omitempty" json:"Volume116,omitempty" yaml:"Volume116,omitempty"`

	Volume117 float64 `xml:"Volume117,omitempty" json:"Volume117,omitempty" yaml:"Volume117,omitempty"`

	Volume118 float64 `xml:"Volume118,omitempty" json:"Volume118,omitempty" yaml:"Volume118,omitempty"`

	Volume119 float64 `xml:"Volume119,omitempty" json:"Volume119,omitempty" yaml:"Volume119,omitempty"`

	Volume120 float64 `xml:"Volume120,omitempty" json:"Volume120,omitempty" yaml:"Volume120,omitempty"`

	Volume121 float64 `xml:"Volume121,omitempty" json:"Volume121,omitempty" yaml:"Volume121,omitempty"`

	Volume122 float64 `xml:"Volume122,omitempty" json:"Volume122,omitempty" yaml:"Volume122,omitempty"`

	Volume123 float64 `xml:"Volume123,omitempty" json:"Volume123,omitempty" yaml:"Volume123,omitempty"`

	Volume124 float64 `xml:"Volume124,omitempty" json:"Volume124,omitempty" yaml:"Volume124,omitempty"`

	Volume125 float64 `xml:"Volume125,omitempty" json:"Volume125,omitempty" yaml:"Volume125,omitempty"`

	Volume126 float64 `xml:"Volume126,omitempty" json:"Volume126,omitempty" yaml:"Volume126,omitempty"`

	Volume127 float64 `xml:"Volume127,omitempty" json:"Volume127,omitempty" yaml:"Volume127,omitempty"`

	Volume128 float64 `xml:"Volume128,omitempty" json:"Volume128,omitempty" yaml:"Volume128,omitempty"`

	Volume129 float64 `xml:"Volume129,omitempty" json:"Volume129,omitempty" yaml:"Volume129,omitempty"`

	Volume130 float64 `xml:"Volume130,omitempty" json:"Volume130,omitempty" yaml:"Volume130,omitempty"`

	Volume131 float64 `xml:"Volume131,omitempty" json:"Volume131,omitempty" yaml:"Volume131,omitempty"`

	Volume132 float64 `xml:"Volume132,omitempty" json:"Volume132,omitempty" yaml:"Volume132,omitempty"`

	Volume133 float64 `xml:"Volume133,omitempty" json:"Volume133,omitempty" yaml:"Volume133,omitempty"`

	Volume134 float64 `xml:"Volume134,omitempty" json:"Volume134,omitempty" yaml:"Volume134,omitempty"`

	Volume135 float64 `xml:"Volume135,omitempty" json:"Volume135,omitempty" yaml:"Volume135,omitempty"`

	Volume136 float64 `xml:"Volume136,omitempty" json:"Volume136,omitempty" yaml:"Volume136,omitempty"`

	Volume137 float64 `xml:"Volume137,omitempty" json:"Volume137,omitempty" yaml:"Volume137,omitempty"`

	Volume138 float64 `xml:"Volume138,omitempty" json:"Volume138,omitempty" yaml:"Volume138,omitempty"`

	Volume139 float64 `xml:"Volume139,omitempty" json:"Volume139,omitempty" yaml:"Volume139,omitempty"`

	Volume140 float64 `xml:"Volume140,omitempty" json:"Volume140,omitempty" yaml:"Volume140,omitempty"`

	Volume141 float64 `xml:"Volume141,omitempty" json:"Volume141,omitempty" yaml:"Volume141,omitempty"`

	Volume142 float64 `xml:"Volume142,omitempty" json:"Volume142,omitempty" yaml:"Volume142,omitempty"`

	Volume143 float64 `xml:"Volume143,omitempty" json:"Volume143,omitempty" yaml:"Volume143,omitempty"`

	Volume144 float64 `xml:"Volume144,omitempty" json:"Volume144,omitempty" yaml:"Volume144,omitempty"`

	Volume145 float64 `xml:"Volume145,omitempty" json:"Volume145,omitempty" yaml:"Volume145,omitempty"`

	Volume146 float64 `xml:"Volume146,omitempty" json:"Volume146,omitempty" yaml:"Volume146,omitempty"`

	Volume147 float64 `xml:"Volume147,omitempty" json:"Volume147,omitempty" yaml:"Volume147,omitempty"`

	Volume148 float64 `xml:"Volume148,omitempty" json:"Volume148,omitempty" yaml:"Volume148,omitempty"`

	Volume149 float64 `xml:"Volume149,omitempty" json:"Volume149,omitempty" yaml:"Volume149,omitempty"`

	Volume150 float64 `xml:"Volume150,omitempty" json:"Volume150,omitempty" yaml:"Volume150,omitempty"`

	Volume151 float64 `xml:"Volume151,omitempty" json:"Volume151,omitempty" yaml:"Volume151,omitempty"`

	Volume152 float64 `xml:"Volume152,omitempty" json:"Volume152,omitempty" yaml:"Volume152,omitempty"`

	Volume153 float64 `xml:"Volume153,omitempty" json:"Volume153,omitempty" yaml:"Volume153,omitempty"`

	Volume154 float64 `xml:"Volume154,omitempty" json:"Volume154,omitempty" yaml:"Volume154,omitempty"`

	Volume155 float64 `xml:"Volume155,omitempty" json:"Volume155,omitempty" yaml:"Volume155,omitempty"`

	Volume156 float64 `xml:"Volume156,omitempty" json:"Volume156,omitempty" yaml:"Volume156,omitempty"`

	Volume157 float64 `xml:"Volume157,omitempty" json:"Volume157,omitempty" yaml:"Volume157,omitempty"`

	Volume158 float64 `xml:"Volume158,omitempty" json:"Volume158,omitempty" yaml:"Volume158,omitempty"`

	Volume159 float64 `xml:"Volume159,omitempty" json:"Volume159,omitempty" yaml:"Volume159,omitempty"`

	Volume160 float64 `xml:"Volume160,omitempty" json:"Volume160,omitempty" yaml:"Volume160,omitempty"`

	Volume161 float64 `xml:"Volume161,omitempty" json:"Volume161,omitempty" yaml:"Volume161,omitempty"`

	Volume162 float64 `xml:"Volume162,omitempty" json:"Volume162,omitempty" yaml:"Volume162,omitempty"`

	Volume163 float64 `xml:"Volume163,omitempty" json:"Volume163,omitempty" yaml:"Volume163,omitempty"`

	Volume164 float64 `xml:"Volume164,omitempty" json:"Volume164,omitempty" yaml:"Volume164,omitempty"`

	Volume165 float64 `xml:"Volume165,omitempty" json:"Volume165,omitempty" yaml:"Volume165,omitempty"`

	Volume166 float64 `xml:"Volume166,omitempty" json:"Volume166,omitempty" yaml:"Volume166,omitempty"`

	Volume167 float64 `xml:"Volume167,omitempty" json:"Volume167,omitempty" yaml:"Volume167,omitempty"`

	Volume168 float64 `xml:"Volume168,omitempty" json:"Volume168,omitempty" yaml:"Volume168,omitempty"`

	Volume169 float64 `xml:"Volume169,omitempty" json:"Volume169,omitempty" yaml:"Volume169,omitempty"`

	Volume170 float64 `xml:"Volume170,omitempty" json:"Volume170,omitempty" yaml:"Volume170,omitempty"`

	Volume171 float64 `xml:"Volume171,omitempty" json:"Volume171,omitempty" yaml:"Volume171,omitempty"`

	Volume172 float64 `xml:"Volume172,omitempty" json:"Volume172,omitempty" yaml:"Volume172,omitempty"`

	Volume173 float64 `xml:"Volume173,omitempty" json:"Volume173,omitempty" yaml:"Volume173,omitempty"`

	Volume174 float64 `xml:"Volume174,omitempty" json:"Volume174,omitempty" yaml:"Volume174,omitempty"`

	Volume175 float64 `xml:"Volume175,omitempty" json:"Volume175,omitempty" yaml:"Volume175,omitempty"`

	Volume176 float64 `xml:"Volume176,omitempty" json:"Volume176,omitempty" yaml:"Volume176,omitempty"`

	Volume177 float64 `xml:"Volume177,omitempty" json:"Volume177,omitempty" yaml:"Volume177,omitempty"`

	Volume178 float64 `xml:"Volume178,omitempty" json:"Volume178,omitempty" yaml:"Volume178,omitempty"`

	Volume179 float64 `xml:"Volume179,omitempty" json:"Volume179,omitempty" yaml:"Volume179,omitempty"`

	Volume180 float64 `xml:"Volume180,omitempty" json:"Volume180,omitempty" yaml:"Volume180,omitempty"`

	Volume181 float64 `xml:"Volume181,omitempty" json:"Volume181,omitempty" yaml:"Volume181,omitempty"`

	Volume182 float64 `xml:"Volume182,omitempty" json:"Volume182,omitempty" yaml:"Volume182,omitempty"`

	Volume183 float64 `xml:"Volume183,omitempty" json:"Volume183,omitempty" yaml:"Volume183,omitempty"`

	Volume184 float64 `xml:"Volume184,omitempty" json:"Volume184,omitempty" yaml:"Volume184,omitempty"`

	Volume185 float64 `xml:"Volume185,omitempty" json:"Volume185,omitempty" yaml:"Volume185,omitempty"`

	Volume186 float64 `xml:"Volume186,omitempty" json:"Volume186,omitempty" yaml:"Volume186,omitempty"`

	Volume187 float64 `xml:"Volume187,omitempty" json:"Volume187,omitempty" yaml:"Volume187,omitempty"`

	Volume188 float64 `xml:"Volume188,omitempty" json:"Volume188,omitempty" yaml:"Volume188,omitempty"`

	Volume189 float64 `xml:"Volume189,omitempty" json:"Volume189,omitempty" yaml:"Volume189,omitempty"`

	Volume190 float64 `xml:"Volume190,omitempty" json:"Volume190,omitempty" yaml:"Volume190,omitempty"`

	Volume191 float64 `xml:"Volume191,omitempty" json:"Volume191,omitempty" yaml:"Volume191,omitempty"`

	Volume192 float64 `xml:"Volume192,omitempty" json:"Volume192,omitempty" yaml:"Volume192,omitempty"`

	Volume193 float64 `xml:"Volume193,omitempty" json:"Volume193,omitempty" yaml:"Volume193,omitempty"`

	Volume194 float64 `xml:"Volume194,omitempty" json:"Volume194,omitempty" yaml:"Volume194,omitempty"`

	Volume195 float64 `xml:"Volume195,omitempty" json:"Volume195,omitempty" yaml:"Volume195,omitempty"`

	Volume196 float64 `xml:"Volume196,omitempty" json:"Volume196,omitempty" yaml:"Volume196,omitempty"`

	Volume197 float64 `xml:"Volume197,omitempty" json:"Volume197,omitempty" yaml:"Volume197,omitempty"`

	Volume198 float64 `xml:"Volume198,omitempty" json:"Volume198,omitempty" yaml:"Volume198,omitempty"`

	Volume199 float64 `xml:"Volume199,omitempty" json:"Volume199,omitempty" yaml:"Volume199,omitempty"`

	Volume200 float64 `xml:"Volume200,omitempty" json:"Volume200,omitempty" yaml:"Volume200,omitempty"`
}

type ArrayOfNodeResponse struct {
	NodeResponse []*NodeResponse `xml:"NodeResponse,omitempty" json:"NodeResponse,omitempty" yaml:"NodeResponse,omitempty"`
}

type GetDownlineResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetDownlineResult"`

	*ApiResponse

	Nodes *ArrayOfNodeResponse `xml:"Nodes,omitempty" json:"Nodes,omitempty" yaml:"Nodes,omitempty"`

	RecordCount int32 `xml:"RecordCount,omitempty" json:"RecordCount,omitempty" yaml:"RecordCount,omitempty"`
}

type GetUplineRequest struct {
	*ApiRequest

	TreeType *TreeType `xml:"TreeType,omitempty" json:"TreeType,omitempty" yaml:"TreeType,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	PeriodType int32 `xml:"PeriodType,omitempty" json:"PeriodType,omitempty" yaml:"PeriodType,omitempty"`

	PeriodID int32 `xml:"PeriodID,omitempty" json:"PeriodID,omitempty" yaml:"PeriodID,omitempty"`
}

type GetUplineResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetUplineResult"`

	*ApiResponse

	Nodes *ArrayOfNodeResponse `xml:"Nodes,omitempty" json:"Nodes,omitempty" yaml:"Nodes,omitempty"`
}

type DequeueCustomerEventsRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DequeuCustomerEventsRequest"`

	*ApiRequest
}

type CustomerEventResponse struct {
	EventID int32 `xml:"EventID,omitempty" json:"EventID,omitempty" yaml:"EventID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	EventDescription string `xml:"EventDescription,omitempty" json:"EventDescription,omitempty" yaml:"EventDescription,omitempty"`

	Fields *ArrayOfCustomerEventField `xml:"Fields,omitempty" json:"Fields,omitempty" yaml:"Fields,omitempty"`

	EventDate time.Time `xml:"EventDate,omitempty" json:"EventDate,omitempty" yaml:"EventDate,omitempty"`
}

type ArrayOfCustomerEventField struct {
	CustomerEventField []*CustomerEventField `xml:"CustomerEventField,omitempty" json:"CustomerEventField,omitempty" yaml:"CustomerEventField,omitempty"`
}

type CustomerEventField struct {
	Name string `xml:"Name,omitempty" json:"Name,omitempty" yaml:"Name,omitempty"`

	Value int32 `xml:"Value,omitempty" json:"Value,omitempty" yaml:"Value,omitempty"`
}

type ArrayOfCustomerEventResponse struct {
	CustomerEventResponse []*CustomerEventResponse `xml:"CustomerEventResponse,omitempty" json:"CustomerEventResponse,omitempty" yaml:"CustomerEventResponse,omitempty"`
}

type DequeueCustomerEventsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DequeueCustomerEventsResult"`

	*ApiResponse

	CustomerEvents *ArrayOfCustomerEventResponse `xml:"CustomerEvents,omitempty" json:"CustomerEvents,omitempty" yaml:"CustomerEvents,omitempty"`
}

type CreatePointTransactionRequest struct {
	*ApiRequest

	PointAccountID int32 `xml:"PointAccountID,omitempty" json:"PointAccountID,omitempty" yaml:"PointAccountID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	Amount float64 `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`

	Reference string `xml:"Reference,omitempty" json:"Reference,omitempty" yaml:"Reference,omitempty"`

	TransactionType *PointTransactionType `xml:"TransactionType,omitempty" json:"TransactionType,omitempty" yaml:"TransactionType,omitempty"`
}

type CreatePointTransactionResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreatePointTransactionResult"`

	*ApiResponse
}

type GetPointAccountRequest struct {
	*ApiRequest

	PointAccountID int32 `xml:"PointAccountID,omitempty" json:"PointAccountID,omitempty" yaml:"PointAccountID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`
}

type GetPointAccountResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetPointAccountResult"`

	*ApiResponse

	Balance float64 `xml:"Balance,omitempty" json:"Balance,omitempty" yaml:"Balance,omitempty"`
}

type GetSubscriptionRequest struct {
	*ApiRequest

	SubscriptionID int32 `xml:"SubscriptionID,omitempty" json:"SubscriptionID,omitempty" yaml:"SubscriptionID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`
}

type GetSubscriptionResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetSubscriptionResult"`

	*ApiResponse

	Status *SubscriptionStatus `xml:"Status,omitempty" json:"Status,omitempty" yaml:"Status,omitempty"`

	StartDate time.Time `xml:"StartDate,omitempty" json:"StartDate,omitempty" yaml:"StartDate,omitempty"`

	ExpireDate time.Time `xml:"ExpireDate,omitempty" json:"ExpireDate,omitempty" yaml:"ExpireDate,omitempty"`
}

type ValidateRequest struct {
	*ApiRequest
}

type IsEnrollerChildValidateRequest struct {
	*ValidateRequest

	ParentID int32 `xml:"ParentID,omitempty" json:"ParentID,omitempty" yaml:"ParentID,omitempty"`

	ChildID int32 `xml:"ChildID,omitempty" json:"ChildID,omitempty" yaml:"ChildID,omitempty"`
}

type IsTaxIDAvailableValidateRequest struct {
	*ValidateRequest

	TaxID string `xml:"TaxID,omitempty" json:"TaxID,omitempty" yaml:"TaxID,omitempty"`

	TaxTypeID int32 `xml:"TaxTypeID,omitempty" json:"TaxTypeID,omitempty" yaml:"TaxTypeID,omitempty"`

	ExcludeCustomerID int32 `xml:"ExcludeCustomerID,omitempty" json:"ExcludeCustomerID,omitempty" yaml:"ExcludeCustomerID,omitempty"`
}

type IsMatrixChildValidateRequest struct {
	*ValidateRequest

	ParentID int32 `xml:"ParentID,omitempty" json:"ParentID,omitempty" yaml:"ParentID,omitempty"`

	ChildID int32 `xml:"ChildID,omitempty" json:"ChildID,omitempty" yaml:"ChildID,omitempty"`
}

type IsUniLevelChildValidateRequest struct {
	*ValidateRequest

	ParentID int32 `xml:"ParentID,omitempty" json:"ParentID,omitempty" yaml:"ParentID,omitempty"`

	ChildID int32 `xml:"ChildID,omitempty" json:"ChildID,omitempty" yaml:"ChildID,omitempty"`
}

type IsLoginNameAvailableValidateRequest struct {
	*ValidateRequest

	LoginName string `xml:"LoginName,omitempty" json:"LoginName,omitempty" yaml:"LoginName,omitempty"`
}

type ValidateResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ValidateResult"`

	*ApiResponse

	IsValid bool `xml:"IsValid,omitempty" json:"IsValid,omitempty" yaml:"IsValid,omitempty"`
}

type VerifyAddressRequest struct {
	*ApiRequest

	Address string `xml:"Address,omitempty" json:"Address,omitempty" yaml:"Address,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty" json:"Zip,omitempty" yaml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`
}

type VerifyAddressResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ VerifyAddressResult"`

	*ApiResponse

	Address string `xml:"Address,omitempty" json:"Address,omitempty" yaml:"Address,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`

	County string `xml:"County,omitempty" json:"County,omitempty" yaml:"County,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty" json:"Zip,omitempty" yaml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`
}

type OptOutEmailRequest struct {
	*ApiRequest

	Email string `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`
}

type OptOutEmailResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ OptOutEmailResult"`

	*ApiResponse

	RecordsAffected int32 `xml:"RecordsAffected,omitempty" json:"RecordsAffected,omitempty" yaml:"RecordsAffected,omitempty"`
}

type OptOutSmsRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	PhoneNumber string `xml:"PhoneNumber,omitempty" json:"PhoneNumber,omitempty" yaml:"PhoneNumber,omitempty"`
}

type OptOutSmsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ OptOutSmsResult"`

	*ApiResponse

	RecordsAffected int32 `xml:"RecordsAffected,omitempty" json:"RecordsAffected,omitempty" yaml:"RecordsAffected,omitempty"`
}

type GetShoppingCartRequest struct {
	*ApiRequest

	ShoppingID string `xml:"ShoppingID,omitempty" json:"ShoppingID,omitempty" yaml:"ShoppingID,omitempty"`
}

type GetShoppingCartResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetShoppingCartResult"`

	*ApiResponse

	ExistingOrderID int32 `xml:"ExistingOrderID,omitempty" json:"ExistingOrderID,omitempty" yaml:"ExistingOrderID,omitempty"`

	ExistingAutoOrderID int32 `xml:"ExistingAutoOrderID,omitempty" json:"ExistingAutoOrderID,omitempty" yaml:"ExistingAutoOrderID,omitempty"`

	Details *ArrayOfOrderDetailResponse `xml:"Details,omitempty" json:"Details,omitempty" yaml:"Details,omitempty"`
}

type GetWarehousesRequest struct {
	*ApiRequest
}

type ArrayOfWarehouseResponse struct {
	WarehouseResponse []*WarehouseResponse `xml:"WarehouseResponse,omitempty" json:"WarehouseResponse,omitempty" yaml:"WarehouseResponse,omitempty"`
}

type GetWarehousesResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetWarehousesResult"`

	*ApiResponse

	Warehouses *ArrayOfWarehouseResponse `xml:"Warehouses,omitempty" json:"Warehouses,omitempty" yaml:"Warehouses,omitempty"`
}

type GetSessionRequest struct {
	*ApiRequest

	SessionID string `xml:"SessionID,omitempty" json:"SessionID,omitempty" yaml:"SessionID,omitempty"`
}

type GetSessionResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetSessionResult"`

	*ApiResponse

	SessionData string `xml:"SessionData,omitempty" json:"SessionData,omitempty" yaml:"SessionData,omitempty"`
}

type SetSessionRequest struct {
	*ApiRequest

	SessionID string `xml:"SessionID,omitempty" json:"SessionID,omitempty" yaml:"SessionID,omitempty"`

	SessionData string `xml:"SessionData,omitempty" json:"SessionData,omitempty" yaml:"SessionData,omitempty"`
}

type SetSessionResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetSessionResult"`

	*ApiResponse
}

type GetItemsRequest struct {
	*ApiRequest

	CurrencyCode string `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty" yaml:"CurrencyCode,omitempty"`

	PriceType int32 `xml:"PriceType,omitempty" json:"PriceType,omitempty" yaml:"PriceType,omitempty"`

	WarehouseID int32 `xml:"WarehouseID,omitempty" json:"WarehouseID,omitempty" yaml:"WarehouseID,omitempty"`

	ItemCodes *ArrayOfString `xml:"ItemCodes,omitempty" json:"ItemCodes,omitempty" yaml:"ItemCodes,omitempty"`

	WebID int32 `xml:"WebID,omitempty" json:"WebID,omitempty" yaml:"WebID,omitempty"`

	WebCategoryID int32 `xml:"WebCategoryID,omitempty" json:"WebCategoryID,omitempty" yaml:"WebCategoryID,omitempty"`

	ReturnLongDetail bool `xml:"ReturnLongDetail,omitempty" json:"ReturnLongDetail,omitempty" yaml:"ReturnLongDetail,omitempty"`

	RestrictToWarehouse bool `xml:"RestrictToWarehouse,omitempty" json:"RestrictToWarehouse,omitempty" yaml:"RestrictToWarehouse,omitempty"`

	LanguageID int32 `xml:"LanguageID,omitempty" json:"LanguageID,omitempty" yaml:"LanguageID,omitempty"`

	ExcludeHideFromSearch bool `xml:"ExcludeHideFromSearch,omitempty" json:"ExcludeHideFromSearch,omitempty" yaml:"ExcludeHideFromSearch,omitempty"`
}

type ItemResponse struct {
	ItemCode string `xml:"ItemCode,omitempty" json:"ItemCode,omitempty" yaml:"ItemCode,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	Price float64 `xml:"Price,omitempty" json:"Price,omitempty" yaml:"Price,omitempty"`

	CommissionableVolume float64 `xml:"CommissionableVolume,omitempty" json:"CommissionableVolume,omitempty" yaml:"CommissionableVolume,omitempty"`

	BusinessVolume float64 `xml:"BusinessVolume,omitempty" json:"BusinessVolume,omitempty" yaml:"BusinessVolume,omitempty"`

	Other1Price float64 `xml:"Other1Price,omitempty" json:"Other1Price,omitempty" yaml:"Other1Price,omitempty"`

	Other2Price float64 `xml:"Other2Price,omitempty" json:"Other2Price,omitempty" yaml:"Other2Price,omitempty"`

	Other3Price float64 `xml:"Other3Price,omitempty" json:"Other3Price,omitempty" yaml:"Other3Price,omitempty"`

	Other4Price float64 `xml:"Other4Price,omitempty" json:"Other4Price,omitempty" yaml:"Other4Price,omitempty"`

	Other5Price float64 `xml:"Other5Price,omitempty" json:"Other5Price,omitempty" yaml:"Other5Price,omitempty"`

	Other6Price float64 `xml:"Other6Price,omitempty" json:"Other6Price,omitempty" yaml:"Other6Price,omitempty"`

	Other7Price float64 `xml:"Other7Price,omitempty" json:"Other7Price,omitempty" yaml:"Other7Price,omitempty"`

	Other8Price float64 `xml:"Other8Price,omitempty" json:"Other8Price,omitempty" yaml:"Other8Price,omitempty"`

	Other9Price float64 `xml:"Other9Price,omitempty" json:"Other9Price,omitempty" yaml:"Other9Price,omitempty"`

	Other10Price float64 `xml:"Other10Price,omitempty" json:"Other10Price,omitempty" yaml:"Other10Price,omitempty"`

	Category string `xml:"Category,omitempty" json:"Category,omitempty" yaml:"Category,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty" json:"CategoryID,omitempty" yaml:"CategoryID,omitempty"`

	TinyPicture string `xml:"TinyPicture,omitempty" json:"TinyPicture,omitempty" yaml:"TinyPicture,omitempty"`

	SmallPicture string `xml:"SmallPicture,omitempty" json:"SmallPicture,omitempty" yaml:"SmallPicture,omitempty"`

	LargePicture string `xml:"LargePicture,omitempty" json:"LargePicture,omitempty" yaml:"LargePicture,omitempty"`

	ShortDetail string `xml:"ShortDetail,omitempty" json:"ShortDetail,omitempty" yaml:"ShortDetail,omitempty"`

	ShortDetail2 string `xml:"ShortDetail2,omitempty" json:"ShortDetail2,omitempty" yaml:"ShortDetail2,omitempty"`

	ShortDetail3 string `xml:"ShortDetail3,omitempty" json:"ShortDetail3,omitempty" yaml:"ShortDetail3,omitempty"`

	ShortDetail4 string `xml:"ShortDetail4,omitempty" json:"ShortDetail4,omitempty" yaml:"ShortDetail4,omitempty"`

	LongDetail string `xml:"LongDetail,omitempty" json:"LongDetail,omitempty" yaml:"LongDetail,omitempty"`

	LongDetail2 string `xml:"LongDetail2,omitempty" json:"LongDetail2,omitempty" yaml:"LongDetail2,omitempty"`

	LongDetail3 string `xml:"LongDetail3,omitempty" json:"LongDetail3,omitempty" yaml:"LongDetail3,omitempty"`

	LongDetail4 string `xml:"LongDetail4,omitempty" json:"LongDetail4,omitempty" yaml:"LongDetail4,omitempty"`

	InventoryStatus *InventoryStatusType `xml:"InventoryStatus,omitempty" json:"InventoryStatus,omitempty" yaml:"InventoryStatus,omitempty"`

	StockLevel int32 `xml:"StockLevel,omitempty" json:"StockLevel,omitempty" yaml:"StockLevel,omitempty"`

	AvailableStockLevel int32 `xml:"AvailableStockLevel,omitempty" json:"AvailableStockLevel,omitempty" yaml:"AvailableStockLevel,omitempty"`

	MaxAllowedOnOrder int32 `xml:"MaxAllowedOnOrder,omitempty" json:"MaxAllowedOnOrder,omitempty" yaml:"MaxAllowedOnOrder,omitempty"`

	Field1 string `xml:"Field1,omitempty" json:"Field1,omitempty" yaml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty" json:"Field2,omitempty" yaml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty" json:"Field3,omitempty" yaml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty" json:"Field4,omitempty" yaml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty" json:"Field5,omitempty" yaml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty" json:"Field6,omitempty" yaml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty" json:"Field7,omitempty" yaml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty" json:"Field8,omitempty" yaml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty" json:"Field9,omitempty" yaml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty" json:"Field10,omitempty" yaml:"Field10,omitempty"`

	OtherCheck1 bool `xml:"OtherCheck1,omitempty" json:"OtherCheck1,omitempty" yaml:"OtherCheck1,omitempty"`

	OtherCheck2 bool `xml:"OtherCheck2,omitempty" json:"OtherCheck2,omitempty" yaml:"OtherCheck2,omitempty"`

	OtherCheck3 bool `xml:"OtherCheck3,omitempty" json:"OtherCheck3,omitempty" yaml:"OtherCheck3,omitempty"`

	OtherCheck4 bool `xml:"OtherCheck4,omitempty" json:"OtherCheck4,omitempty" yaml:"OtherCheck4,omitempty"`

	OtherCheck5 bool `xml:"OtherCheck5,omitempty" json:"OtherCheck5,omitempty" yaml:"OtherCheck5,omitempty"`

	IsVirtual bool `xml:"IsVirtual,omitempty" json:"IsVirtual,omitempty" yaml:"IsVirtual,omitempty"`

	AllowOnAutoOrder bool `xml:"AllowOnAutoOrder,omitempty" json:"AllowOnAutoOrder,omitempty" yaml:"AllowOnAutoOrder,omitempty"`

	IsGroupMaster bool `xml:"IsGroupMaster,omitempty" json:"IsGroupMaster,omitempty" yaml:"IsGroupMaster,omitempty"`

	GroupDescription string `xml:"GroupDescription,omitempty" json:"GroupDescription,omitempty" yaml:"GroupDescription,omitempty"`

	GroupMembersDescription string `xml:"GroupMembersDescription,omitempty" json:"GroupMembersDescription,omitempty" yaml:"GroupMembersDescription,omitempty"`

	GroupMembers *ArrayOfItemMemberResponse `xml:"GroupMembers,omitempty" json:"GroupMembers,omitempty" yaml:"GroupMembers,omitempty"`

	IsDynamicKitMaster bool `xml:"IsDynamicKitMaster,omitempty" json:"IsDynamicKitMaster,omitempty" yaml:"IsDynamicKitMaster,omitempty"`

	HideFromSearch bool `xml:"HideFromSearch,omitempty" json:"HideFromSearch,omitempty" yaml:"HideFromSearch,omitempty"`

	KitMembers *ArrayOfKitMemberResponse `xml:"KitMembers,omitempty" json:"KitMembers,omitempty" yaml:"KitMembers,omitempty"`

	TaxablePrice float64 `xml:"TaxablePrice,omitempty" json:"TaxablePrice,omitempty" yaml:"TaxablePrice,omitempty"`

	ShippingPrice float64 `xml:"ShippingPrice,omitempty" json:"ShippingPrice,omitempty" yaml:"ShippingPrice,omitempty"`
}

type ArrayOfItemMemberResponse struct {
	ItemMemberResponse []*ItemMemberResponse `xml:"ItemMemberResponse,omitempty" json:"ItemMemberResponse,omitempty" yaml:"ItemMemberResponse,omitempty"`
}

type ItemMemberResponse struct {
	ItemCode string `xml:"ItemCode,omitempty" json:"ItemCode,omitempty" yaml:"ItemCode,omitempty"`

	MemberDescription string `xml:"MemberDescription,omitempty" json:"MemberDescription,omitempty" yaml:"MemberDescription,omitempty"`

	ItemDescription string `xml:"ItemDescription,omitempty" json:"ItemDescription,omitempty" yaml:"ItemDescription,omitempty"`

	InventoryStatus *InventoryStatusType `xml:"InventoryStatus,omitempty" json:"InventoryStatus,omitempty" yaml:"InventoryStatus,omitempty"`

	StockLevel int32 `xml:"StockLevel,omitempty" json:"StockLevel,omitempty" yaml:"StockLevel,omitempty"`

	AvailableStockLevel int32 `xml:"AvailableStockLevel,omitempty" json:"AvailableStockLevel,omitempty" yaml:"AvailableStockLevel,omitempty"`
}

type ArrayOfKitMemberResponse struct {
	KitMemberResponse []*KitMemberResponse `xml:"KitMemberResponse,omitempty" json:"KitMemberResponse,omitempty" yaml:"KitMemberResponse,omitempty"`
}

type KitMemberResponse struct {
	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	KitMemberItems *ArrayOfKitMemberItemResponse `xml:"KitMemberItems,omitempty" json:"KitMemberItems,omitempty" yaml:"KitMemberItems,omitempty"`
}

type ArrayOfKitMemberItemResponse struct {
	KitMemberItemResponse []*KitMemberItemResponse `xml:"KitMemberItemResponse,omitempty" json:"KitMemberItemResponse,omitempty" yaml:"KitMemberItemResponse,omitempty"`
}

type KitMemberItemResponse struct {
	ItemCode string `xml:"ItemCode,omitempty" json:"ItemCode,omitempty" yaml:"ItemCode,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	InventoryStatus *InventoryStatusType `xml:"InventoryStatus,omitempty" json:"InventoryStatus,omitempty" yaml:"InventoryStatus,omitempty"`
}

type ArrayOfItemResponse struct {
	ItemResponse []*ItemResponse `xml:"ItemResponse,omitempty" json:"ItemResponse,omitempty" yaml:"ItemResponse,omitempty"`
}

type GetItemsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetItemsResult"`

	*ApiResponse

	Items *ArrayOfItemResponse `xml:"Items,omitempty" json:"Items,omitempty" yaml:"Items,omitempty"`
}

type GetLanguagesRequest struct {
	*ApiRequest
}

type LanguageResponse struct {
	LanguageID int32 `xml:"LanguageID,omitempty" json:"LanguageID,omitempty" yaml:"LanguageID,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	CultureCode string `xml:"CultureCode,omitempty" json:"CultureCode,omitempty" yaml:"CultureCode,omitempty"`
}

type ArrayOfLanguageResponse struct {
	LanguageResponse []*LanguageResponse `xml:"LanguageResponse,omitempty" json:"LanguageResponse,omitempty" yaml:"LanguageResponse,omitempty"`
}

type GetLanguagesResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCompanyLanguagesResult"`

	*ApiResponse

	CompanyLanguages *ArrayOfLanguageResponse `xml:"CompanyLanguages,omitempty" json:"CompanyLanguages,omitempty" yaml:"CompanyLanguages,omitempty"`
}

type CreateWebCategoryRequest struct {
	*ApiRequest

	WebID int32 `xml:"WebID,omitempty" json:"WebID,omitempty" yaml:"WebID,omitempty"`

	ParentID int32 `xml:"ParentID,omitempty" json:"ParentID,omitempty" yaml:"ParentID,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
}

type CreateWebCategoryResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateWebCategoryResult"`

	*ApiResponse

	CategoryID int32 `xml:"CategoryID,omitempty" json:"CategoryID,omitempty" yaml:"CategoryID,omitempty"`

	WebID int32 `xml:"WebID,omitempty" json:"WebID,omitempty" yaml:"WebID,omitempty"`
}

type UpdateWebCategoryRequest struct {
	*ApiRequest

	WebID int32 `xml:"WebID,omitempty" json:"WebID,omitempty" yaml:"WebID,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty" json:"CategoryID,omitempty" yaml:"CategoryID,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
}

type UpdateWebCategoryResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateWebCategoryResult"`

	*ApiResponse

	CategoryID int32 `xml:"CategoryID,omitempty" json:"CategoryID,omitempty" yaml:"CategoryID,omitempty"`

	WebID int32 `xml:"WebID,omitempty" json:"WebID,omitempty" yaml:"WebID,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`
}

type DeleteWebCategoryRequest struct {
	*ApiRequest

	WebID int32 `xml:"WebID,omitempty" json:"WebID,omitempty" yaml:"WebID,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty" json:"CategoryID,omitempty" yaml:"CategoryID,omitempty"`
}

type DeleteWebCategoryResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteWebCategoryResult"`

	*ApiResponse

	CategoryID int32 `xml:"CategoryID,omitempty" json:"CategoryID,omitempty" yaml:"CategoryID,omitempty"`

	WebID int32 `xml:"WebID,omitempty" json:"WebID,omitempty" yaml:"WebID,omitempty"`
}

type AddProductsToCategoryRequest struct {
	*ApiRequest

	WebID int32 `xml:"WebID,omitempty" json:"WebID,omitempty" yaml:"WebID,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty" json:"CategoryID,omitempty" yaml:"CategoryID,omitempty"`

	ItemCodes *ArrayOfString `xml:"ItemCodes,omitempty" json:"ItemCodes,omitempty" yaml:"ItemCodes,omitempty"`
}

type AddProductsToCategoryResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ AddProductsToCategoryResult"`

	*ApiResponse

	CategoryID int32 `xml:"CategoryID,omitempty" json:"CategoryID,omitempty" yaml:"CategoryID,omitempty"`

	WebID int32 `xml:"WebID,omitempty" json:"WebID,omitempty" yaml:"WebID,omitempty"`
}

type DeleteProductFromCategoryRequest struct {
	*ApiRequest

	WebID int32 `xml:"WebID,omitempty" json:"WebID,omitempty" yaml:"WebID,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty" json:"CategoryID,omitempty" yaml:"CategoryID,omitempty"`

	ItemCode string `xml:"ItemCode,omitempty" json:"ItemCode,omitempty" yaml:"ItemCode,omitempty"`
}

type DeleteProductFromCategoryResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteProductFromCategoryResult"`

	*ApiResponse

	CategoryID int32 `xml:"CategoryID,omitempty" json:"CategoryID,omitempty" yaml:"CategoryID,omitempty"`

	WebID int32 `xml:"WebID,omitempty" json:"WebID,omitempty" yaml:"WebID,omitempty"`

	ItemCode string `xml:"ItemCode,omitempty" json:"ItemCode,omitempty" yaml:"ItemCode,omitempty"`
}

type GetCompanyNewsRequest struct {
	*ApiRequest

	StartDate time.Time `xml:"StartDate,omitempty" json:"StartDate,omitempty" yaml:"StartDate,omitempty"`

	EndDate time.Time `xml:"EndDate,omitempty" json:"EndDate,omitempty" yaml:"EndDate,omitempty"`

	DepartmentType int32 `xml:"DepartmentType,omitempty" json:"DepartmentType,omitempty" yaml:"DepartmentType,omitempty"`
}

type CompanyNewsResponse struct {
	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	NewsID int32 `xml:"NewsID,omitempty" json:"NewsID,omitempty" yaml:"NewsID,omitempty"`

	CreatedDate time.Time `xml:"CreatedDate,omitempty" json:"CreatedDate,omitempty" yaml:"CreatedDate,omitempty"`

	WebSettings *NewsWebSettings `xml:"WebSettings,omitempty" json:"WebSettings,omitempty" yaml:"WebSettings,omitempty"`

	CompanySettings *NewsCompanySettings `xml:"CompanySettings,omitempty" json:"CompanySettings,omitempty" yaml:"CompanySettings,omitempty"`
}

type ArrayOfCompanyNewsResponse struct {
	CompanyNewsResponse []*CompanyNewsResponse `xml:"CompanyNewsResponse,omitempty" json:"CompanyNewsResponse,omitempty" yaml:"CompanyNewsResponse,omitempty"`
}

type GetCompanyNewsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCompanyNewsResult"`

	*ApiResponse

	CompanyNews *ArrayOfCompanyNewsResponse `xml:"CompanyNews,omitempty" json:"CompanyNews,omitempty" yaml:"CompanyNews,omitempty"`
}

type GetCompanyNewsItemRequest struct {
	*ApiRequest

	NewsID int32 `xml:"NewsID,omitempty" json:"NewsID,omitempty" yaml:"NewsID,omitempty"`
}

type DepartmentInfo struct {
	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	DepartmentType int32 `xml:"DepartmentType,omitempty" json:"DepartmentType,omitempty" yaml:"DepartmentType,omitempty"`
}

type ArrayOfDepartmentInfo struct {
	DepartmentInfo []*DepartmentInfo `xml:"DepartmentInfo,omitempty" json:"DepartmentInfo,omitempty" yaml:"DepartmentInfo,omitempty"`
}

type GetCompanyNewsItemResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCompanyNewsItemResult"`

	*ApiResponse

	Description string `xml:"Description,omitempty" json:"Description,omitempty" yaml:"Description,omitempty"`

	NewsID int32 `xml:"NewsID,omitempty" json:"NewsID,omitempty" yaml:"NewsID,omitempty"`

	CreatedDate time.Time `xml:"CreatedDate,omitempty" json:"CreatedDate,omitempty" yaml:"CreatedDate,omitempty"`

	WebSettings *NewsWebSettings `xml:"WebSettings,omitempty" json:"WebSettings,omitempty" yaml:"WebSettings,omitempty"`

	CompanySettings *NewsCompanySettings `xml:"CompanySettings,omitempty" json:"CompanySettings,omitempty" yaml:"CompanySettings,omitempty"`

	Content string `xml:"Content,omitempty" json:"Content,omitempty" yaml:"Content,omitempty"`

	Departments *ArrayOfDepartmentInfo `xml:"Departments,omitempty" json:"Departments,omitempty" yaml:"Departments,omitempty"`
}

type GetRandomMessageRequest struct {
	*ApiRequest

	LanguageID int32 `xml:"LanguageID,omitempty" json:"LanguageID,omitempty" yaml:"LanguageID,omitempty"`
}

type GetRandomMessageResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetRandomMessageResult"`

	*ApiResponse

	RandomMessageID int32 `xml:"RandomMessageID,omitempty" json:"RandomMessageID,omitempty" yaml:"RandomMessageID,omitempty"`

	Content string `xml:"Content,omitempty" json:"Content,omitempty" yaml:"Content,omitempty"`
}

type FireResponderRequest struct {
	*ApiRequest

	ResponderID int32 `xml:"ResponderID,omitempty" json:"ResponderID,omitempty" yaml:"ResponderID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty" json:"OrderID,omitempty" yaml:"OrderID,omitempty"`
}

type FireResponderResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ FireResponderResult"`

	*ApiResponse
}

type SendSmsRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	Message string `xml:"Message,omitempty" json:"Message,omitempty" yaml:"Message,omitempty"`

	Phone string `xml:"Phone,omitempty" json:"Phone,omitempty" yaml:"Phone,omitempty"`
}

type SendSmsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SendSmsResult"`

	*ApiResponse
}

type CreateVendorBillRequest struct {
	*ApiRequest

	VendorBillStatusTypeID int32 `xml:"VendorBillStatusTypeID,omitempty" json:"VendorBillStatusTypeID,omitempty" yaml:"VendorBillStatusTypeID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty" json:"CurrencyCode,omitempty" yaml:"CurrencyCode,omitempty"`

	DueDate time.Time `xml:"DueDate,omitempty" json:"DueDate,omitempty" yaml:"DueDate,omitempty"`

	Amount float64 `xml:"Amount,omitempty" json:"Amount,omitempty" yaml:"Amount,omitempty"`

	Reference string `xml:"Reference,omitempty" json:"Reference,omitempty" yaml:"Reference,omitempty"`

	Notes string `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`
}

type CreateVendorBillResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateVendorBillResult"`

	*ApiResponse

	VendorBillID int32 `xml:"VendorBillID,omitempty" json:"VendorBillID,omitempty" yaml:"VendorBillID,omitempty"`
}

type CreateCustomerContactRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty" json:"MiddleName,omitempty" yaml:"MiddleName,omitempty"`

	Company string `xml:"Company,omitempty" json:"Company,omitempty" yaml:"Company,omitempty"`

	BusinessPhone string `xml:"BusinessPhone,omitempty" json:"BusinessPhone,omitempty" yaml:"BusinessPhone,omitempty"`

	HomePhone string `xml:"HomePhone,omitempty" json:"HomePhone,omitempty" yaml:"HomePhone,omitempty"`

	Mobile string `xml:"Mobile,omitempty" json:"Mobile,omitempty" yaml:"Mobile,omitempty"`

	Fax string `xml:"Fax,omitempty" json:"Fax,omitempty" yaml:"Fax,omitempty"`

	Email string `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`

	Address1 string `xml:"Address1,omitempty" json:"Address1,omitempty" yaml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty" json:"Address2,omitempty" yaml:"Address2,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty" json:"Zip,omitempty" yaml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`

	BirthDate time.Time `xml:"BirthDate,omitempty" json:"BirthDate,omitempty" yaml:"BirthDate,omitempty"`

	Notes string `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`

	LinkedIn string `xml:"LinkedIn,omitempty" json:"LinkedIn,omitempty" yaml:"LinkedIn,omitempty"`

	Facebook string `xml:"Facebook,omitempty" json:"Facebook,omitempty" yaml:"Facebook,omitempty"`

	Blog string `xml:"Blog,omitempty" json:"Blog,omitempty" yaml:"Blog,omitempty"`

	MySpace string `xml:"MySpace,omitempty" json:"MySpace,omitempty" yaml:"MySpace,omitempty"`

	GooglePlus string `xml:"GooglePlus,omitempty" json:"GooglePlus,omitempty" yaml:"GooglePlus,omitempty"`

	Twitter string `xml:"Twitter,omitempty" json:"Twitter,omitempty" yaml:"Twitter,omitempty"`
}

type CreateCustomerContactResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateCustomerContactResult"`

	*ApiResponse

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	CustomerContactID int32 `xml:"CustomerContactID,omitempty" json:"CustomerContactID,omitempty" yaml:"CustomerContactID,omitempty"`
}

type UpdateCustomerContactRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	CustomerContactID int32 `xml:"CustomerContactID,omitempty" json:"CustomerContactID,omitempty" yaml:"CustomerContactID,omitempty"`

	FirstName string `xml:"FirstName,omitempty" json:"FirstName,omitempty" yaml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty" json:"LastName,omitempty" yaml:"LastName,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty" json:"MiddleName,omitempty" yaml:"MiddleName,omitempty"`

	Company string `xml:"Company,omitempty" json:"Company,omitempty" yaml:"Company,omitempty"`

	BusinessPhone string `xml:"BusinessPhone,omitempty" json:"BusinessPhone,omitempty" yaml:"BusinessPhone,omitempty"`

	HomePhone string `xml:"HomePhone,omitempty" json:"HomePhone,omitempty" yaml:"HomePhone,omitempty"`

	Mobile string `xml:"Mobile,omitempty" json:"Mobile,omitempty" yaml:"Mobile,omitempty"`

	Fax string `xml:"Fax,omitempty" json:"Fax,omitempty" yaml:"Fax,omitempty"`

	Email string `xml:"Email,omitempty" json:"Email,omitempty" yaml:"Email,omitempty"`

	Address1 string `xml:"Address1,omitempty" json:"Address1,omitempty" yaml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty" json:"Address2,omitempty" yaml:"Address2,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty" json:"Zip,omitempty" yaml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`

	BirthDate time.Time `xml:"BirthDate,omitempty" json:"BirthDate,omitempty" yaml:"BirthDate,omitempty"`

	Notes string `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`

	LinkedIn string `xml:"LinkedIn,omitempty" json:"LinkedIn,omitempty" yaml:"LinkedIn,omitempty"`

	Facebook string `xml:"Facebook,omitempty" json:"Facebook,omitempty" yaml:"Facebook,omitempty"`

	Blog string `xml:"Blog,omitempty" json:"Blog,omitempty" yaml:"Blog,omitempty"`

	MySpace string `xml:"MySpace,omitempty" json:"MySpace,omitempty" yaml:"MySpace,omitempty"`

	GooglePlus string `xml:"GooglePlus,omitempty" json:"GooglePlus,omitempty" yaml:"GooglePlus,omitempty"`

	Twitter string `xml:"Twitter,omitempty" json:"Twitter,omitempty" yaml:"Twitter,omitempty"`
}

type UpdateCustomerContactResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateCustomerContactResult"`

	*ApiResponse

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	CustomerContactID int32 `xml:"CustomerContactID,omitempty" json:"CustomerContactID,omitempty" yaml:"CustomerContactID,omitempty"`
}

type DeleteCustomerContactRequest struct {
	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	CustomerContactID int32 `xml:"CustomerContactID,omitempty" json:"CustomerContactID,omitempty" yaml:"CustomerContactID,omitempty"`
}

type DeleteCustomerContactResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteCustomerContactResult"`

	*ApiResponse

	CustomerID int32 `xml:"CustomerID,omitempty" json:"CustomerID,omitempty" yaml:"CustomerID,omitempty"`

	CustomerContactID int32 `xml:"CustomerContactID,omitempty" json:"CustomerContactID,omitempty" yaml:"CustomerContactID,omitempty"`
}

type CreateCalendarItemRequest struct {
	*ApiRequest

	UserID int32 `xml:"UserID,omitempty" json:"UserID,omitempty" yaml:"UserID,omitempty"`

	CalendarID int32 `xml:"CalendarID,omitempty" json:"CalendarID,omitempty" yaml:"CalendarID,omitempty"`

	CalendarItemType *CalendarItemType `xml:"CalendarItemType,omitempty" json:"CalendarItemType,omitempty" yaml:"CalendarItemType,omitempty"`

	CalendarItemStatusType *CalendarItemStatusType `xml:"CalendarItemStatusType,omitempty" json:"CalendarItemStatusType,omitempty" yaml:"CalendarItemStatusType,omitempty"`

	CalendarItemPriorityType *CalendarItemPriorityType `xml:"CalendarItemPriorityType,omitempty" json:"CalendarItemPriorityType,omitempty" yaml:"CalendarItemPriorityType,omitempty"`

	Subject string `xml:"Subject,omitempty" json:"Subject,omitempty" yaml:"Subject,omitempty"`

	Location string `xml:"Location,omitempty" json:"Location,omitempty" yaml:"Location,omitempty"`

	Notes string `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`

	StartDate time.Time `xml:"StartDate,omitempty" json:"StartDate,omitempty" yaml:"StartDate,omitempty"`

	EndDate time.Time `xml:"EndDate,omitempty" json:"EndDate,omitempty" yaml:"EndDate,omitempty"`

	TimeZone int32 `xml:"TimeZone,omitempty" json:"TimeZone,omitempty" yaml:"TimeZone,omitempty"`

	Address1 string `xml:"Address1,omitempty" json:"Address1,omitempty" yaml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty" json:"Address2,omitempty" yaml:"Address2,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`

	PostalCode string `xml:"PostalCode,omitempty" json:"PostalCode,omitempty" yaml:"PostalCode,omitempty"`

	ContactInfo string `xml:"ContactInfo,omitempty" json:"ContactInfo,omitempty" yaml:"ContactInfo,omitempty"`

	ContactPhone string `xml:"ContactPhone,omitempty" json:"ContactPhone,omitempty" yaml:"ContactPhone,omitempty"`

	ContactPhoneType *ContactPhoneType `xml:"ContactPhoneType,omitempty" json:"ContactPhoneType,omitempty" yaml:"ContactPhoneType,omitempty"`

	ContactEmail string `xml:"ContactEmail,omitempty" json:"ContactEmail,omitempty" yaml:"ContactEmail,omitempty"`

	EventHost string `xml:"EventHost,omitempty" json:"EventHost,omitempty" yaml:"EventHost,omitempty"`

	SpecialGuests string `xml:"SpecialGuests,omitempty" json:"SpecialGuests,omitempty" yaml:"SpecialGuests,omitempty"`

	EventFlyer string `xml:"EventFlyer,omitempty" json:"EventFlyer,omitempty" yaml:"EventFlyer,omitempty"`

	EventCostInfo string `xml:"EventCostInfo,omitempty" json:"EventCostInfo,omitempty" yaml:"EventCostInfo,omitempty"`

	EventConferenceCallOrWebinar string `xml:"EventConferenceCallOrWebinar,omitempty" json:"EventConferenceCallOrWebinar,omitempty" yaml:"EventConferenceCallOrWebinar,omitempty"`

	EventRegistrationInfo string `xml:"EventRegistrationInfo,omitempty" json:"EventRegistrationInfo,omitempty" yaml:"EventRegistrationInfo,omitempty"`

	EventTags string `xml:"EventTags,omitempty" json:"EventTags,omitempty" yaml:"EventTags,omitempty"`

	IsShared bool `xml:"IsShared,omitempty" json:"IsShared,omitempty" yaml:"IsShared,omitempty"`
}

type CreateCalendarItemResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateCalendarItemResult"`

	*ApiResponse

	CalendarID int32 `xml:"CalendarID,omitempty" json:"CalendarID,omitempty" yaml:"CalendarID,omitempty"`

	CalendarItemID int32 `xml:"CalendarItemID,omitempty" json:"CalendarItemID,omitempty" yaml:"CalendarItemID,omitempty"`
}

type UpdateCalendarItemRequest struct {
	*ApiRequest

	UserID int32 `xml:"UserID,omitempty" json:"UserID,omitempty" yaml:"UserID,omitempty"`

	CalendarID int32 `xml:"CalendarID,omitempty" json:"CalendarID,omitempty" yaml:"CalendarID,omitempty"`

	CalendarItemID int32 `xml:"CalendarItemID,omitempty" json:"CalendarItemID,omitempty" yaml:"CalendarItemID,omitempty"`

	CalendarItemType *CalendarItemType `xml:"CalendarItemType,omitempty" json:"CalendarItemType,omitempty" yaml:"CalendarItemType,omitempty"`

	CalendarItemStatusType *CalendarItemStatusType `xml:"CalendarItemStatusType,omitempty" json:"CalendarItemStatusType,omitempty" yaml:"CalendarItemStatusType,omitempty"`

	CalendarItemPriorityType *CalendarItemPriorityType `xml:"CalendarItemPriorityType,omitempty" json:"CalendarItemPriorityType,omitempty" yaml:"CalendarItemPriorityType,omitempty"`

	Subject string `xml:"Subject,omitempty" json:"Subject,omitempty" yaml:"Subject,omitempty"`

	Location string `xml:"Location,omitempty" json:"Location,omitempty" yaml:"Location,omitempty"`

	Notes string `xml:"Notes,omitempty" json:"Notes,omitempty" yaml:"Notes,omitempty"`

	StartDate time.Time `xml:"StartDate,omitempty" json:"StartDate,omitempty" yaml:"StartDate,omitempty"`

	EndDate time.Time `xml:"EndDate,omitempty" json:"EndDate,omitempty" yaml:"EndDate,omitempty"`

	TimeZone int32 `xml:"TimeZone,omitempty" json:"TimeZone,omitempty" yaml:"TimeZone,omitempty"`

	Address1 string `xml:"Address1,omitempty" json:"Address1,omitempty" yaml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty" json:"Address2,omitempty" yaml:"Address2,omitempty"`

	City string `xml:"City,omitempty" json:"City,omitempty" yaml:"City,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty" yaml:"State,omitempty"`

	Country string `xml:"Country,omitempty" json:"Country,omitempty" yaml:"Country,omitempty"`

	PostalCode string `xml:"PostalCode,omitempty" json:"PostalCode,omitempty" yaml:"PostalCode,omitempty"`

	ContactInfo string `xml:"ContactInfo,omitempty" json:"ContactInfo,omitempty" yaml:"ContactInfo,omitempty"`

	ContactPhone string `xml:"ContactPhone,omitempty" json:"ContactPhone,omitempty" yaml:"ContactPhone,omitempty"`

	ContactPhoneType *ContactPhoneType `xml:"ContactPhoneType,omitempty" json:"ContactPhoneType,omitempty" yaml:"ContactPhoneType,omitempty"`

	ContactEmail string `xml:"ContactEmail,omitempty" json:"ContactEmail,omitempty" yaml:"ContactEmail,omitempty"`

	EventHost string `xml:"EventHost,omitempty" json:"EventHost,omitempty" yaml:"EventHost,omitempty"`

	SpecialGuests string `xml:"SpecialGuests,omitempty" json:"SpecialGuests,omitempty" yaml:"SpecialGuests,omitempty"`

	EventFlyer string `xml:"EventFlyer,omitempty" json:"EventFlyer,omitempty" yaml:"EventFlyer,omitempty"`

	EventCostInfo string `xml:"EventCostInfo,omitempty" json:"EventCostInfo,omitempty" yaml:"EventCostInfo,omitempty"`

	EventConferenceCallOrWebinar string `xml:"EventConferenceCallOrWebinar,omitempty" json:"EventConferenceCallOrWebinar,omitempty" yaml:"EventConferenceCallOrWebinar,omitempty"`

	EventRegistrationInfo string `xml:"EventRegistrationInfo,omitempty" json:"EventRegistrationInfo,omitempty" yaml:"EventRegistrationInfo,omitempty"`

	EventTags string `xml:"EventTags,omitempty" json:"EventTags,omitempty" yaml:"EventTags,omitempty"`

	IsShared bool `xml:"IsShared,omitempty" json:"IsShared,omitempty" yaml:"IsShared,omitempty"`
}

type UpdateCalendarItemResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateCalendarItemResult"`

	*ApiResponse

	CalendarID int32 `xml:"CalendarID,omitempty" json:"CalendarID,omitempty" yaml:"CalendarID,omitempty"`

	CalendarItemID int32 `xml:"CalendarItemID,omitempty" json:"CalendarItemID,omitempty" yaml:"CalendarItemID,omitempty"`
}

type DeleteCalendarItemRequest struct {
	*ApiRequest

	UserID int32 `xml:"UserID,omitempty" json:"UserID,omitempty" yaml:"UserID,omitempty"`

	CalendarID int32 `xml:"CalendarID,omitempty" json:"CalendarID,omitempty" yaml:"CalendarID,omitempty"`

	CalendarItemID int32 `xml:"CalendarItemID,omitempty" json:"CalendarItemID,omitempty" yaml:"CalendarItemID,omitempty"`
}

type DeleteCalendarItemResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteCalendarItemResult"`

	*ApiResponse

	CalendarID int32 `xml:"CalendarID,omitempty" json:"CalendarID,omitempty" yaml:"CalendarID,omitempty"`

	CalendarItemID int32 `xml:"CalendarItemID,omitempty" json:"CalendarItemID,omitempty" yaml:"CalendarItemID,omitempty"`
}

type ExigoApiSoap struct {
	client *SOAPClient
}

func NewExigoApiSoap(url string, tls bool, auth *BasicAuth) *ExigoApiSoap {
	if url == "" {
		url = "http://api.exigo.com/3.0/ExigoApi.asmx"
	}
	client := NewSOAPClient(url, tls, auth)

	return &ExigoApiSoap{
		client: client,
	}
}

func (service *ExigoApiSoap) SetHeader(header interface{}) {
	service.client.SetHeader(header)
}

/*
   Creates an email.
*/
func (service *ExigoApiSoap) CreateEmail(request *CreateEmailRequest) (*CreateEmailResponse, error) {
	response := new(CreateEmailResponse)
	err := service.client.Call("http://api.exigo.com/CreateEmail", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Moves an email to a different folder.
*/
func (service *ExigoApiSoap) MoveEmail(request *MoveEmailRequest) (*MoveEmailResponse, error) {
	response := new(MoveEmailResponse)
	err := service.client.Call("http://api.exigo.com/MoveEmail", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Updates email's status.
*/
func (service *ExigoApiSoap) UpdateEmailStatus(request *UpdateEmailStatusRequest) (*UpdateEmailStatusResponse, error) {
	response := new(UpdateEmailStatusResponse)
	err := service.client.Call("http://api.exigo.com/UpdateEmailStatus", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Gets an email attachment.
*/
func (service *ExigoApiSoap) GetEmailAttachment(request *GetEmailAttachmentRequest) (*GetEmailAttachmentResponse, error) {
	response := new(GetEmailAttachmentResponse)
	err := service.client.Call("http://api.exigo.com/GetEmailAttachment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Deletes an email forever.
*/
func (service *ExigoApiSoap) DeleteEmail(request *DeleteEmailRequest) (*DeleteEmailResponse, error) {
	response := new(DeleteEmailResponse)
	err := service.client.Call("http://api.exigo.com/DeleteEmail", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Creates an email template.
*/
func (service *ExigoApiSoap) CreateEmailTemplate(request *CreateEmailTemplateRequest) (*CreateEmailTemplateResponse, error) {
	response := new(CreateEmailTemplateResponse)
	err := service.client.Call("http://api.exigo.com/CreateEmailTemplate", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Updates an email template.
*/
func (service *ExigoApiSoap) UpdateEmailTemplate(request *UpdateEmailTemplateRequest) (*UpdateEmailTemplateResponse, error) {
	response := new(UpdateEmailTemplateResponse)
	err := service.client.Call("http://api.exigo.com/UpdateEmailTemplate", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Deletes an email template.
*/
func (service *ExigoApiSoap) DeleteEmailTemplate(request *DeleteEmailTemplateRequest) (*DeleteEmailTemplateResponse, error) {
	response := new(DeleteEmailTemplateResponse)
	err := service.client.Call("http://api.exigo.com/DeleteEmailTemplate", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Ensures all the basic mail folders exist.
*/
func (service *ExigoApiSoap) EnsureMailFolders(request *EnsureMailFoldersRequest) (*EnsureMailFoldersResponse, error) {
	response := new(EnsureMailFoldersResponse)
	err := service.client.Call("http://api.exigo.com/EnsureMailFolders", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Creates a mail folder.
*/
func (service *ExigoApiSoap) CreateMailFolder(request *CreateMailFolderRequest) (*CreateMailFolderResponse, error) {
	response := new(CreateMailFolderResponse)
	err := service.client.Call("http://api.exigo.com/CreateMailFolder", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Updates a mail folder.
*/
func (service *ExigoApiSoap) UpdateMailFolder(request *UpdateMailFolderRequest) (*UpdateMailFolderResponse, error) {
	response := new(UpdateMailFolderResponse)
	err := service.client.Call("http://api.exigo.com/UpdateMailFolder", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Deletes a mail folder.
*/
func (service *ExigoApiSoap) DeleteMailFolder(request *DeleteMailFolderRequest) (*DeleteMailFolderResponse, error) {
	response := new(DeleteMailFolderResponse)
	err := service.client.Call("http://api.exigo.com/DeleteMailFolder", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Deletes all emails from a folder.
*/
func (service *ExigoApiSoap) EmptyMailFolder(request *EmptyMailFolderRequest) (*EmptyMailFolderResponse, error) {
	response := new(EmptyMailFolderResponse)
	err := service.client.Call("http://api.exigo.com/EmptyMailFolder", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Returns a list of social network(s) for the customer.
*/
func (service *ExigoApiSoap) GetCustomerSocialNetworks(request *GetCustomerSocialNetworksRequest) (*GetCustomerSocialNetworksResponse, error) {
	response := new(GetCustomerSocialNetworksResponse)
	err := service.client.Call("http://api.exigo.com/GetCustomerSocialNetworks", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Returns a list of social network(s) for the customer lead.
*/
func (service *ExigoApiSoap) GetCustomerLeadSocialNetworks(request *GetCustomerLeadSocialNetworksRequest) (*GetCustomerLeadSocialNetworksResponse, error) {
	response := new(GetCustomerLeadSocialNetworksResponse)
	err := service.client.Call("http://api.exigo.com/GetCustomerLeadSocialNetworks", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Deletes all existing social network(s) for the customer, then inserts the provided list of social network(s) for the customer.
*/
func (service *ExigoApiSoap) SetCustomerSocialNetworks(request *SetCustomerSocialNetworksRequest) (*SetCustomerSocialNetworksResponse, error) {
	response := new(SetCustomerSocialNetworksResponse)
	err := service.client.Call("http://api.exigo.com/SetCustomerSocialNetworks", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Deletes all existing social network(s) for the customer lead, then inserts the provided list of social network(s) for the customer lead.
*/
func (service *ExigoApiSoap) SetCustomerLeadSocialNetworks(request *SetCustomerLeadSocialNetworksRequest) (*SetCustomerLeadSocialNetworksResponse, error) {
	response := new(SetCustomerLeadSocialNetworksResponse)
	err := service.client.Call("http://api.exigo.com/SetCustomerLeadSocialNetworks", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Create new wall item in CustomerWall table.
*/
func (service *ExigoApiSoap) CreateCustomerWallItem(request *CreateCustomerWallItemRequest) (*CreateCustomerWallItemResponse, error) {
	response := new(CreateCustomerWallItemResponse)
	err := service.client.Call("http://api.exigo.com/CreateCustomerWallItem", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Delete wall item(s) from CustomerWall table.
*/
func (service *ExigoApiSoap) DeleteCustomerWallItem(request *DeleteCustomerWallItemRequest) (*DeleteCustomerWallItemResponse, error) {
	response := new(DeleteCustomerWallItemResponse)
	err := service.client.Call("http://api.exigo.com/DeleteCustomerWallItem", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Get wall item(s) from CustomerWall table.
*/
func (service *ExigoApiSoap) GetCustomerWall(request *GetCustomerWallRequest) (*GetCustomerWallResponse, error) {
	response := new(GetCustomerWallResponse)
	err := service.client.Call("http://api.exigo.com/GetCustomerWall", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Updates an Item (and optional warehouse/price info)
*/
func (service *ExigoApiSoap) UpdateItem(request *UpdateItemRequest) (*UpdateItemResponse, error) {
	response := new(UpdateItemResponse)
	err := service.client.Call("http://api.exigo.com/UpdateItem", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Adds an Item (and optional warehouse/price info)
*/
func (service *ExigoApiSoap) CreateItem(request *CreateItemRequest) (*CreateItemResponse, error) {
	response := new(CreateItemResponse)
	err := service.client.Call("http://api.exigo.com/CreateItem", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Sets pricing info for an item)
*/
func (service *ExigoApiSoap) SetItemPrice(request *SetItemPriceRequest) (*SetItemPriceResponse, error) {
	response := new(SetItemPriceResponse)
	err := service.client.Call("http://api.exigo.com/SetItemPrice", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Sets warehouse info for an item)
*/
func (service *ExigoApiSoap) SetItemWarehouse(request *SetItemWarehouseRequest) (*SetItemWarehouseResponse, error) {
	response := new(SetItemWarehouseResponse)
	err := service.client.Call("http://api.exigo.com/SetItemWarehouse", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Sets country/region  info for an item)
*/
func (service *ExigoApiSoap) SetItemCountryRegion(request *SetItemCountryRegionRequest) (*SetItemCountryRegionResponse, error) {
	response := new(SetItemCountryRegionResponse)
	err := service.client.Call("http://api.exigo.com/SetItemCountryRegion", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Returns item, country and region properties along with taxings for an item.
*/
func (service *ExigoApiSoap) GetItemCountryRegion(request *GetItemCountryRegionRequest) (*GetItemCountryRegionResponse, error) {
	response := new(GetItemCountryRegionResponse)
	err := service.client.Call("http://api.exigo.com/GetItemCountryRegion", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Set Images for an Item
*/
func (service *ExigoApiSoap) SetItemImage(request *SetItemImageRequest) (*SetItemImageResponse, error) {
	response := new(SetItemImageResponse)
	err := service.client.Call("http://api.exigo.com/SetItemImage", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Set Images
*/
func (service *ExigoApiSoap) SetImageFile(request *SetImageFileRequest) (*SetImageFileResponse, error) {
	response := new(SetImageFileResponse)
	err := service.client.Call("http://api.exigo.com/SetImageFile", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Creates a file for the customer in their default directory
*/
func (service *ExigoApiSoap) CreateCustomerFile(request *CreateCustomerFileRequest) (*CreateCustomerFileResponse, error) {
	response := new(CreateCustomerFileResponse)
	err := service.client.Call("http://api.exigo.com/CreateCustomerFile", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Creates a customer history record.
*/
func (service *ExigoApiSoap) CreateCustomerInquiry(request *CreateCustomerInquiryRequest) (*CreateCustomerInquiryResponse, error) {
	response := new(CreateCustomerInquiryResponse)
	err := service.client.Call("http://api.exigo.com/CreateCustomerInquiry", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Returns the sum of orders, payments, and adjustments per currency.
*/
func (service *ExigoApiSoap) GetCustomerBalances(request *GetCustomerBalancesRequest) (*GetCustomerBalancesResponse, error) {
	response := new(GetCustomerBalancesResponse)
	err := service.client.Call("http://api.exigo.com/GetCustomerBalances", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Create an adjustment for a customer's account in a currency.
*/
func (service *ExigoApiSoap) CreateCustomerBalanceAdjustment(request *CreateCustomerBalanceAdjustmentRequest) (*CreateCustomerBalanceAdjustmentResponse, error) {
	response := new(CreateCustomerBalanceAdjustmentResponse)
	err := service.client.Call("http://api.exigo.com/CreateCustomerBalanceAdjustment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Creates a new Party */
func (service *ExigoApiSoap) CreateParty(request *CreatePartyRequest) (*CreatePartyResponse, error) {
	response := new(CreatePartyResponse)
	err := service.client.Call("http://api.exigo.com/CreateParty", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Updates selected Party instance */
func (service *ExigoApiSoap) UpdateParty(request *UpdatePartyRequest) (*UpdatePartyResponse, error) {
	response := new(UpdatePartyResponse)
	err := service.client.Call("http://api.exigo.com/UpdateParty", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Gets Parties list */
func (service *ExigoApiSoap) GetParties(request *GetPartiesRequest) (*GetPartiesResponse, error) {
	response := new(GetPartiesResponse)
	err := service.client.Call("http://api.exigo.com/GetParties", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Gets a list of guests */
func (service *ExigoApiSoap) GetGuests(request *GetGuestsRequest) (*GetGuestsResponse, error) {
	response := new(GetGuestsResponse)
	err := service.client.Call("http://api.exigo.com/GetGuests", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Gets a list of guests belonging to a party */
func (service *ExigoApiSoap) GetPartyGuests(request *GetPartyGuestsRequest) (*GetPartyGuestsResponse, error) {
	response := new(GetPartyGuestsResponse)
	err := service.client.Call("http://api.exigo.com/GetPartyGuests", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Returns a list of social network(s) for the guest. */
func (service *ExigoApiSoap) GetGuestSocialNetworks(request *GetGuestSocialNetworksRequest) (*GetGuestSocialNetworksResponse, error) {
	response := new(GetGuestSocialNetworksResponse)
	err := service.client.Call("http://api.exigo.com/GetGuestSocialNetworks", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Deletes all existing social network(s) for the guest, then inserts the provided list of social network(s) for the guest. */
func (service *ExigoApiSoap) SetGuestSocialNetworks(request *SetGuestSocialNetworksRequest) (*SetGuestSocialNetworksResponse, error) {
	response := new(SetGuestSocialNetworksResponse)
	err := service.client.Call("http://api.exigo.com/SetGuestSocialNetworks", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Creates a new guest. Can optionally be put in a party. */
func (service *ExigoApiSoap) CreateGuest(request *CreateGuestRequest) (*CreateGuestResponse, error) {
	response := new(CreateGuestResponse)
	err := service.client.Call("http://api.exigo.com/CreateGuest", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Updates one or more fields on an existing guest. */
func (service *ExigoApiSoap) UpdateGuest(request *UpdateGuestRequest) (*UpdateGuestResponse, error) {
	response := new(UpdateGuestResponse)
	err := service.client.Call("http://api.exigo.com/UpdateGuest", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Adds one or more guests to a party. */
func (service *ExigoApiSoap) AddPartyGuests(request *AddPartyGuestsRequest) (*AddPartyGuestsResponse, error) {
	response := new(AddPartyGuestsResponse)
	err := service.client.Call("http://api.exigo.com/AddPartyGuests", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Removes one or more guests from a party. */
func (service *ExigoApiSoap) RemovePartyGuests(request *RemovePartyGuestsRequest) (*RemovePartyGuestsResponse, error) {
	response := new(RemovePartyGuestsResponse)
	err := service.client.Call("http://api.exigo.com/RemovePartyGuests", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Create new Extended DB schema. */
func (service *ExigoApiSoap) CreateExtendedDbSchema(request *CreateExtendedDbSchemeRequest) (*CreateExtendedDbSchemaResponse, error) {
	response := new(CreateExtendedDbSchemaResponse)
	err := service.client.Call("http://api.exigo.com/CreateExtendedDbSchema", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Get Extended DB schema. */
func (service *ExigoApiSoap) GetExtendedDbSchema(request *GetSchemaRequest) (*GetSchemaResponse, error) {
	response := new(GetSchemaResponse)
	err := service.client.Call("http://api.exigo.com/GetExtendedDbSchema", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Delete Extended DB schema. */
func (service *ExigoApiSoap) DeleteExtendedDbSchema(request *DeleteSchemaRequest) (*DeleteSchemaResponse, error) {
	response := new(DeleteSchemaResponse)
	err := service.client.Call("http://api.exigo.com/DeleteExtendedDbSchema", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Create Extended DB Entity. */
func (service *ExigoApiSoap) CreateExtendedDbEntity(request *CreateEntityRequest) (*CreateEntityResponse, error) {
	response := new(CreateEntityResponse)
	err := service.client.Call("http://api.exigo.com/CreateExtendedDbEntity", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Get Extended DB Entity. */
func (service *ExigoApiSoap) GetExtendedDbEntity(request *GetEntityRequest) (*GetEntityResponse, error) {
	response := new(GetEntityResponse)
	err := service.client.Call("http://api.exigo.com/GetExtendedDbEntity", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Update Extended DB Entity. */
func (service *ExigoApiSoap) UpdateExtendedDbEntity(request *UpdateEntityRequest) (*UpdateEntityResponse, error) {
	response := new(UpdateEntityResponse)
	err := service.client.Call("http://api.exigo.com/UpdateExtendedDbEntity", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Delete Extended DB Entity. */
func (service *ExigoApiSoap) DeleteExtendedDbEntity(request *DeleteEntityRequest) (*DeleteEntityResponse, error) {
	response := new(DeleteEntityResponse)
	err := service.client.Call("http://api.exigo.com/DeleteExtendedDbEntity", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Start a Sandbox */
func (service *ExigoApiSoap) StartSandbox(request *StartSandboxRequest) (*StartSandboxResponse, error) {
	response := new(StartSandboxResponse)
	err := service.client.Call("http://api.exigo.com/StartSandbox", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Stop a Sandbox */
func (service *ExigoApiSoap) StopSandbox(request *StopSandboxRequest) (*StopSandboxResponse, error) {
	response := new(StopSandboxResponse)
	err := service.client.Call("http://api.exigo.com/StopSandbox", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Refresh a Sandbox */
func (service *ExigoApiSoap) RefreshSandbox(request *RefreshSandboxRequest) (*RefreshSandboxResponse, error) {
	response := new(RefreshSandboxResponse)
	err := service.client.Call("http://api.exigo.com/RefreshSandbox", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Get List of Sandboxes */
func (service *ExigoApiSoap) GetSandbox(request *GetSandboxRequest) (*GetSandboxResponse, error) {
	response := new(GetSandboxResponse)
	err := service.client.Call("http://api.exigo.com/GetSandbox", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Returns one or more customers that match the filter critera passed in.
*/
func (service *ExigoApiSoap) GetCustomers(request *GetCustomersRequest) (*GetCustomersResponse, error) {
	response := new(GetCustomersResponse)
	err := service.client.Call("http://api.exigo.com/GetCustomers", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Returns a list of customer files for a given customer
*/
func (service *ExigoApiSoap) GetFileContents(request *GetFileContentsRequest) (*GetFileContentsResponse, error) {
	response := new(GetFileContentsResponse)
	err := service.client.Call("http://api.exigo.com/GetFileContents", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Returns a customer file for the given criteria
*/
func (service *ExigoApiSoap) GetFiles(request *GetFilesRequest) (*GetFilesResponse, error) {
	response := new(GetFilesResponse)
	err := service.client.Call("http://api.exigo.com/GetFiles", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Creates a new Warehouse
*/
func (service *ExigoApiSoap) CreateWarehouse(request *CreateWarehouseRequest) (*CreateWarehouseResponse, error) {
	response := new(CreateWarehouseResponse)
	err := service.client.Call("http://api.exigo.com/CreateWarehouse", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Sets the children for a static kit
*/
func (service *ExigoApiSoap) SetItemKitMembers(request *SetItemKitMembersRequest) (*SetItemKitMembersResponse, error) {
	response := new(SetItemKitMembersResponse)
	err := service.client.Call("http://api.exigo.com/SetItemKitMembers", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Sends a specified Email using SendGrid
*/
func (service *ExigoApiSoap) SendEmail(request *SendEmailRequest) (*SendEmailResponse, error) {
	response := new(SendEmailResponse)
	err := service.client.Call("http://api.exigo.com/SendEmail", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Returns the notes field for one or more customers that match the filter critera passed in.
*/
func (service *ExigoApiSoap) GetCustomerNotes(request *GetCustomerNotesRequest) (*GetCustomerNotesResponse, error) {
	response := new(GetCustomerNotesResponse)
	err := service.client.Call("http://api.exigo.com/GetCustomerNotes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Appends a specified value to the specified Customer's Notes record
*/
func (service *ExigoApiSoap) AppendCustomerNotes(request *AppendCustomerNotesRequest) (*AppendCustomerNotesResponse, error) {
	response := new(AppendCustomerNotesResponse)
	err := service.client.Call("http://api.exigo.com/AppendCustomerNotes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Returns volume records for a current period in a periodtype.
*/
func (service *ExigoApiSoap) GetVolumes(request *GetVolumesRequest) (*GetVolumesResponse, error) {
	response := new(GetVolumesResponse)
	err := service.client.Call("http://api.exigo.com/GetVolumes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Returns realtime commissions available for open periods.
*/
func (service *ExigoApiSoap) GetRealTimeCommissions(request *GetRealTimeCommissionsRequest) (*GetRealTimeCommissionsResponse, error) {
	response := new(GetRealTimeCommissionsResponse)
	err := service.client.Call("http://api.exigo.com/GetRealTimeCommissions", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Returns realtime commission detail for a customer/bonus.
*/
func (service *ExigoApiSoap) GetRealTimeCommissionDetail(request *GetRealTimeCommissionDetailRequest) (*GetRealTimeCommissionDetailResponse, error) {
	response := new(GetRealTimeCommissionDetailResponse)
	err := service.client.Call("http://api.exigo.com/GetRealTimeCommissionDetail", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Returns rank qualification report card.
*/
func (service *ExigoApiSoap) GetRankQualifications(request *GetRankQualificationsRequest) (*GetRankQualificationsResponse, error) {
	response := new(GetRankQualificationsResponse)
	err := service.client.Call("http://api.exigo.com/GetRankQualifications", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Returns Customer Overrides
*/
func (service *ExigoApiSoap) GetQualificationOverrides(request *GetQualificationOverridesRequest) (*GetQualitificationOverridesResponse, error) {
	response := new(GetQualitificationOverridesResponse)
	err := service.client.Call("http://api.exigo.com/GetQualificationOverrides", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Create or Edit a Customer Override */
func (service *ExigoApiSoap) SetQualificationOverride(request *SetQualificationOverrideRequest) (*SetQualificationOverrideResponse, error) {
	response := new(SetQualificationOverrideResponse)
	err := service.client.Call("http://api.exigo.com/SetQualificationOverride", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Delete a Customer Override */
func (service *ExigoApiSoap) DeleteQualificationOverride(request *DeleteQualificationOverrideRequest) (*DeleteQualificationOverrideResponse, error) {
	response := new(DeleteQualificationOverrideResponse)
	err := service.client.Call("http://api.exigo.com/DeleteQualificationOverride", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Posts an adjustment to item inventory.
*/
func (service *ExigoApiSoap) AdjustInventory(request *AdjustInventoryRequest) (*AdjustInventoryResponse, error) {
	response := new(AdjustInventoryResponse)
	err := service.client.Call("http://api.exigo.com/AdjustInventory", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Returns the public web site info setup for a customer. Pass in CustomerID or WebAlias.
*/
func (service *ExigoApiSoap) GetCustomerSite(request *GetCustomerSiteRequest) (*GetCustomerSiteResponse, error) {
	response := new(GetCustomerSiteResponse)
	err := service.client.Call("http://api.exigo.com/GetCustomerSite", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Returns extended properties setup for a customer.
*/
func (service *ExigoApiSoap) GetCustomerExtended(request *GetCustomerExtendedRequest) (*GetCustomerExtendedResponse, error) {
	response := new(GetCustomerExtendedResponse)
	err := service.client.Call("http://api.exigo.com/GetCustomerExtended", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Creates a new customer extended record.
*/
func (service *ExigoApiSoap) CreateCustomerExtended(request *CreateCustomerExtendedRequest) (*CreateCustomerExtendedResponse, error) {
	response := new(CreateCustomerExtendedResponse)
	err := service.client.Call("http://api.exigo.com/CreateCustomerExtended", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Updates one customer extended record.
*/
func (service *ExigoApiSoap) UpdateCustomerExtended(request *UpdateCustomerExtendedRequest) (*UpdateCustomerExtendedResponse, error) {
	response := new(UpdateCustomerExtendedResponse)
	err := service.client.Call("http://api.exigo.com/UpdateCustomerExtended", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Removes one customer extended record.
*/
func (service *ExigoApiSoap) DeleteCustomerExtended(request *DeleteCustomerExtendedRequest) (*DeleteCustomerExtendedResponse, error) {
	response := new(DeleteCustomerExtendedResponse)
	err := service.client.Call("http://api.exigo.com/DeleteCustomerExtended", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Retruns billing accounts on file for a customer. Only returns last four digits of a credit card.
*/
func (service *ExigoApiSoap) GetCustomerBilling(request *GetCustomerBillingRequest) (*GetCustomerBillingResponse, error) {
	response := new(GetCustomerBillingResponse)
	err := service.client.Call("http://api.exigo.com/GetCustomerBilling", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Calculates pricing, tax, shipping and volume info for a potential order. This is for calculation only and does not store a permanent record.
Client provides address, itemcodes, and quantity. Server calculates and returns pricing, tax, shipping and volume info.
*/
func (service *ExigoApiSoap) CalculateOrder(request *CalculateOrderRequest) (*CalculateOrderResponse, error) {
	response := new(CalculateOrderResponse)
	err := service.client.Call("http://api.exigo.com/CalculateOrder", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Create a new Recurring Order template for an existing customer or a new customer as part of a transaction.
Client provides schedule, item codes, quantity, and shipping info. Server calculates and commits data, then returns pricing, tax, shipping and volume info.
You can also update/overwrite an existing autoorder by using the OverwriteExistingAutoOrder and ExistingAutoOrderID properties.
*/
func (service *ExigoApiSoap) CreateAutoOrder(request *CreateAutoOrderRequest) (*CreateAutoOrderResponse, error) {
	response := new(CreateAutoOrderResponse)
	err := service.client.Call("http://api.exigo.com/CreateAutoOrder", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Creates a new Customer. Can optionally be put in one or more trees.
*/
func (service *ExigoApiSoap) CreateCustomer(request *CreateCustomerRequest) (*CreateCustomerResponse, error) {
	response := new(CreateCustomerResponse)
	err := service.client.Call("http://api.exigo.com/CreateCustomer", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Gets payment record(s) for a customer or order. */
func (service *ExigoApiSoap) GetPayments(request *GetPaymentsRequest) (*GetPaymentsResponse, error) {
	response := new(GetPaymentsResponse)
	err := service.client.Call("http://api.exigo.com/GetPayments", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Creates a new payment type using cash, money order etc. */
func (service *ExigoApiSoap) CreatePayment(request *CreatePaymentRequest) (*CreatePaymentResponse, error) {
	response := new(CreatePaymentResponse)
	err := service.client.Call("http://api.exigo.com/CreatePayment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Creates a new payout for one or more bills and updates status */
func (service *ExigoApiSoap) CreatePayout(request *CreatePayoutRequest) (*CreatePayoutResponse, error) {
	response := new(CreatePayoutResponse)
	err := service.client.Call("http://api.exigo.com/CreatePayout", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Creates a new payout and funds a payment card for one or more bills and updates status */
func (service *ExigoApiSoap) FundPaymentCard(request *FundPaymentCardRequest) (*FundPaymentCardResponse, error) {
	response := new(FundPaymentCardResponse)
	err := service.client.Call("http://api.exigo.com/FundPaymentCard", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Creates a new bill */
func (service *ExigoApiSoap) CreateBill(request *CreateBillRequest) (*CreateBillResponse, error) {
	response := new(CreateBillResponse)
	err := service.client.Call("http://api.exigo.com/CreateBill", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Creates a new payment for wallet account. */
func (service *ExigoApiSoap) CreatePaymentWallet(request *CreatePaymentWalletRequest) (*CreatePaymentResponse, error) {
	response := new(CreatePaymentResponse)
	err := service.client.Call("http://api.exigo.com/CreatePaymentWallet", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Creates a new payment for point account. */
func (service *ExigoApiSoap) CreatePaymentPointAccount(request *CreatePaymentPointAccountRequest) (*CreatePaymentPointAccountResponse, error) {
	response := new(CreatePaymentPointAccountResponse)
	err := service.client.Call("http://api.exigo.com/CreatePaymentPointAccount", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Creates a new payment using check information. */
func (service *ExigoApiSoap) CreatePaymentCheck(request *CreatePaymentCheckRequest) (*CreatePaymentCheckResponse, error) {
	response := new(CreatePaymentCheckResponse)
	err := service.client.Call("http://api.exigo.com/CreatePaymentCheck", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Creates a new expected payment type using cash, money order etc. */
func (service *ExigoApiSoap) CreateExpectedPayment(request *CreateExpectedPaymentRequest) (*CreateExpectedPaymentResponse, error) {
	response := new(CreateExpectedPaymentResponse)
	err := service.client.Call("http://api.exigo.com/CreateExpectedPayment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Returns a custom report in dataset format. */
func (service *ExigoApiSoap) GetCustomReport(request *GetCustomReportRequest) (*GetCustomReportResponse, error) {
	response := new(GetCustomReportResponse)
	err := service.client.Call("http://api.exigo.com/GetCustomReport", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Returns a custom report in dataset format. */
func (service *ExigoApiSoap) GetReport(request *GetReportRequest) (*GetReportResponse, error) {
	response := new(GetReportResponse)
	err := service.client.Call("http://api.exigo.com/GetReport", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Creates a new expected payment of type Bank Bire. */
func (service *ExigoApiSoap) CreateExpectedBankWire(request *CreateExpectedBankWireRequest) (*CreateExpectedBankWireResponse, error) {
	response := new(CreateExpectedBankWireResponse)
	err := service.client.Call("http://api.exigo.com/CreateExpectedBankWire", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Creates a specific credit card payment for an existing order, or a new order in a transaction.
Note: This does not charge a card, and you must have an authorization code to submit this.
*/
func (service *ExigoApiSoap) CreatePaymentCreditCard(request *CreatePaymentCreditCardRequest) (*CreatePaymentCreditCardResponse, error) {
	response := new(CreatePaymentCreditCardResponse)
	err := service.client.Call("http://api.exigo.com/CreatePaymentCreditCard", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Deprecated. Use ChargeCreditCardToken instead. */
func (service *ExigoApiSoap) ChargeCreditCard(request *ChargeCreditCardRequest) (*ChargeCreditCardResponse, error) {
	response := new(ChargeCreditCardResponse)
	err := service.client.Call("http://api.exigo.com/ChargeCreditCard", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Deprecated. Use ChargeCreditCardTokenOnFile instead.
*/
func (service *ExigoApiSoap) ChargeCreditCardOnFile(request *ChargeCreditCardOnFileRequest) (*ChargeCreditCardResponse, error) {
	response := new(ChargeCreditCardResponse)
	err := service.client.Call("http://api.exigo.com/ChargeCreditCardOnFile", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Attempts to charge a card token with the amount found on an exiting order, or a new order in a transaction.
If the charge is successful, the order is changed to accepted.
*/
func (service *ExigoApiSoap) ChargeCreditCardToken(request *ChargeCreditCardTokenRequest) (*ChargeCreditCardResponse, error) {
	response := new(ChargeCreditCardResponse)
	err := service.client.Call("http://api.exigo.com/ChargeCreditCardToken", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Attempts to charge a card token with the amount found on the specified orders.
*/
func (service *ExigoApiSoap) ChargeGroupOrderCreditCardToken(request *ChargeGroupOrderCreditCardTokenRequest) (*ChargeGroupOrderCreditCardTokenResponse, error) {
	response := new(ChargeGroupOrderCreditCardTokenResponse)
	err := service.client.Call("http://api.exigo.com/ChargeGroupOrderCreditCardToken", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Attempts to charge a card token on file with the amount found on an exiting order, or a new order in a transaction.
If the charge is successful, the order is changed to accepted.
*/
func (service *ExigoApiSoap) ChargeCreditCardTokenOnFile(request *ChargeCreditCardTokenOnFileRequest) (*ChargeCreditCardResponse, error) {
	response := new(ChargeCreditCardResponse)
	err := service.client.Call("http://api.exigo.com/ChargeCreditCardTokenOnFile", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Attempts to authorize only a credit card token. A follow up ChargePriorAuthorization will need to be issued to actually capture and settle the authorization.
*/
func (service *ExigoApiSoap) AuthorizeOnlyCreditCardToken(request *AuthorizeOnlyCreditCardTokenRequest) (*AuthorizeOnlyCreditCardResponse, error) {
	response := new(AuthorizeOnlyCreditCardResponse)
	err := service.client.Call("http://api.exigo.com/AuthorizeOnlyCreditCardToken", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Attempts to authorize only a credit card token on file. A follow up ChargePriorAuthorization will need to be issued to actually capture and settle the authorization.
*/
func (service *ExigoApiSoap) AuthorizeOnlyCreditCardTokenOnFile(request *AuthorizeOnlyCreditCardTokenOnFileRequest) (*AuthorizeOnlyCreditCardResponse, error) {
	response := new(AuthorizeOnlyCreditCardResponse)
	err := service.client.Call("http://api.exigo.com/AuthorizeOnlyCreditCardTokenOnFile", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Attempts to charge a prior authorization with the amount found on an existing order, or a new order in a transaction.
If the charge is successful, the order is changed to accepted.
*/
func (service *ExigoApiSoap) ChargePriorAuthorization(request *ChargePriorAuthorizationRequest) (*ChargeCreditCardResponse, error) {
	response := new(ChargeCreditCardResponse)
	err := service.client.Call("http://api.exigo.com/ChargePriorAuthorization", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Refunds a prior charge. To be used with a return order.
*/
func (service *ExigoApiSoap) RefundPriorCreditCardCharge(request *RefundPriorCreditCardChargeRequest) (*RefundPriorCreditCardChargeResponse, error) {
	response := new(RefundPriorCreditCardChargeResponse)
	err := service.client.Call("http://api.exigo.com/RefundPriorCreditCardCharge", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Refunds a prior wallet charge. To be used with a return order.
*/
func (service *ExigoApiSoap) RefundPriorWalletCharge(request *RefundPriorWalletChargeRequest) (*RefundPriorWalletChargeResponse, error) {
	response := new(RefundPriorWalletChargeResponse)
	err := service.client.Call("http://api.exigo.com/RefundPriorWalletCharge", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Attempts to debit a bank account with the amount found on an exiting order, or a new order in a transaction.
If the debit request is successful, the order is changed to accepted.
*/
func (service *ExigoApiSoap) DebitBankAccount(request *DebitBankAccountRequest) (*DebitBankAccountResponse, error) {
	response := new(DebitBankAccountResponse)
	err := service.client.Call("http://api.exigo.com/DebitBankAccount", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Attempts to debit a bank account on file with the amount found on an exiting order, or a new order in a transaction.
If the debit request is successful, the order is changed to accepted.
*/
func (service *ExigoApiSoap) DebitBankAccountOnFile(request *DebitBankAccountOnFileRequest) (*DebitBankAccountResponse, error) {
	response := new(DebitBankAccountResponse)
	err := service.client.Call("http://api.exigo.com/DebitBankAccountOnFile", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Attempts to debit a wallet account with the amount found on an exiting order, or a new order in a transaction.
If the debit request is successful, the order is changed to accepted.
*/
func (service *ExigoApiSoap) ChargeWalletAccount(request *ChargeWalletAccountRequest) (*ChargeWalletAccountResponse, error) {
	response := new(ChargeWalletAccountResponse)
	err := service.client.Call("http://api.exigo.com/ChargeWalletAccount", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Attempts to debit a wallet account with the amount found on an exiting order, or a new order in a transaction.
If the debit request is successful, the order is changed to accepted.
*/
func (service *ExigoApiSoap) ChargeWalletAccountOnFile(request *ChargeWalletAccountOnFileRequest) (*ChargeWalletAccountResponse, error) {
	response := new(ChargeWalletAccountResponse)
	err := service.client.Call("http://api.exigo.com/ChargeWalletAccountOnFile", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Setup primary or secondary credit card on file for an existing customer, or new customer in a transaction. */
func (service *ExigoApiSoap) SetAccountCreditCard(request *SetAccountCreditCardRequest) (*SetAccountResponse, error) {
	response := new(SetAccountResponse)
	err := service.client.Call("http://api.exigo.com/SetAccountCreditCard", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Setup primary or secondary credit card token on file for an existing customer, or new customer in a transaction. */
func (service *ExigoApiSoap) SetAccountCreditCardToken(request *SetAccountCreditCardTokenRequest) (*SetAccountResponse, error) {
	response := new(SetAccountResponse)
	err := service.client.Call("http://api.exigo.com/SetAccountCreditCardToken", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Setup primary or secondary wallet account on file for an existing customer, or new customer in a transaction. */
func (service *ExigoApiSoap) SetAccountWallet(request *SetAccountWalletRequest) (*SetAccountResponse, error) {
	response := new(SetAccountResponse)
	err := service.client.Call("http://api.exigo.com/SetAccountWallet", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Setup bank account on file an existing customer, or new customer in a transaction. You can also empty the account settings by simply leaving everything but CustomerID empty. */
func (service *ExigoApiSoap) SetAccountChecking(request *SetAccountCheckingRequest) (*SetAccountResponse, error) {
	response := new(SetAccountResponse)
	err := service.client.Call("http://api.exigo.com/SetAccountChecking", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Setup direct deposit info for an existing customer, or new customer in a transaction. You can also empty the account settings by simply leaving everything but CustomerID empty. */
func (service *ExigoApiSoap) SetAccountDirectDeposit(request *SetAccountDirectDepositRequest) (*SetAccountResponse, error) {
	response := new(SetAccountResponse)
	err := service.client.Call("http://api.exigo.com/SetAccountDirectDeposit", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Returns direct deposit info for an existing customer. */
func (service *ExigoApiSoap) GetAccountDirectDeposit(request *GetAccountDirectDepositRequest) (*GetAccountDirectDepositResponse, error) {
	response := new(GetAccountDirectDepositResponse)
	err := service.client.Call("http://api.exigo.com/GetAccountDirectDeposit", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Setup opt in settings for Email by Email address. */
func (service *ExigoApiSoap) OptInEmail(request *OptInEmailRequest) (*OptInEmailResponse, error) {
	response := new(OptInEmailResponse)
	err := service.client.Call("http://api.exigo.com/OptInEmail", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Setup opt in settings for Sms messaging by phone number. */
func (service *ExigoApiSoap) OptInSms(request *OptInSmsRequest) (*OptInSmsResponse, error) {
	response := new(OptInSmsResponse)
	err := service.client.Call("http://api.exigo.com/OptInSms", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Setup public web site info for an existing customer, or new customer in a transaction. */
func (service *ExigoApiSoap) SetCustomerSite(request *SetCustomerSiteRequest) (*SetCustomerSiteResponse, error) {
	response := new(SetCustomerSiteResponse)
	err := service.client.Call("http://api.exigo.com/SetCustomerSite", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Setup public web site image for an existing customer. */
func (service *ExigoApiSoap) SetCustomerSiteImage(request *SetCustomerSiteImageRequest) (*SetCustomerSiteImageResponse, error) {
	response := new(SetCustomerSiteImageResponse)
	err := service.client.Call("http://api.exigo.com/SetCustomerSiteImage", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Creates a new Order for an existing customer. Can create an order for a new customer if part of a transaction.
Server will calculate all pricing, tax, shipping and volume info unless overridden in the request.
*/
func (service *ExigoApiSoap) CreateOrder(request *CreateOrderRequest) (*CreateOrderResponse, error) {
	response := new(CreateOrderResponse)
	err := service.client.Call("http://api.exigo.com/CreateOrder", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Imports a new Order for an existing customer.
Use this for import routines or other instances where you wish to supply all calculation data.
It is advised to use CreateOrder unless you know the exact breakdown of all pricing, tax, shipping and volume info.
*/
func (service *ExigoApiSoap) CreateOrderImport(request *CreateOrderImportRequest) (*CreateOrderImportResponse, error) {
	response := new(CreateOrderImportResponse)
	err := service.client.Call("http://api.exigo.com/CreateOrderImport", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Commits multiple requests in one step. All business rules are checked before commit and it is a commit-all or fail-all.
*/
func (service *ExigoApiSoap) ProcessTransaction(request *TransactionalRequest) (*TransactionalResponse, error) {
	response := new(TransactionalResponse)
	err := service.client.Call("http://api.exigo.com/ProcessTransaction", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Deprecated. Use AuthenticateCustomer instead.
*/
func (service *ExigoApiSoap) LoginCustomer(request *LoginCustomerRequest) (*LoginCustomerResponse, error) {
	response := new(LoginCustomerResponse)
	err := service.client.Call("http://api.exigo.com/LoginCustomer", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Deprecated. Use AuthenticateCustomer instead.
*/
func (service *ExigoApiSoap) GetLoginSession(request *GetLoginSessionRequest) (*GetLoginSessionResponse, error) {
	response := new(GetLoginSessionResponse)
	err := service.client.Call("http://api.exigo.com/GetLoginSession", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Authenticates a customer for login using UserName and Password combination.
*/
func (service *ExigoApiSoap) AuthenticateCustomer(request *AuthenticateCustomerRequest) (*AuthenticateCustomerResponse, error) {
	response := new(AuthenticateCustomerResponse)
	err := service.client.Call("http://api.exigo.com/AuthenticateCustomer", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Authenticates a corporate user for login using UserName and Password combination.
*/
func (service *ExigoApiSoap) AuthenticateUser(request *AuthenticateUserRequest) (*AuthenticateUserResponse, error) {
	response := new(AuthenticateUserResponse)
	err := service.client.Call("http://api.exigo.com/AuthenticateUser", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Returns permissions associated with the corporate user account.
*/
func (service *ExigoApiSoap) GetUserPermissions(request *GetUserPermissionsRequest) (*GetUserPermissionsResponse, error) {
	response := new(GetUserPermissionsResponse)
	err := service.client.Call("http://api.exigo.com/GetUserPermissions", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Changes an existing order from one status to another.
*/
func (service *ExigoApiSoap) ChangeOrderStatus(request *ChangeOrderStatusRequest) (*ChangeOrderStatusResponse, error) {
	response := new(ChangeOrderStatusResponse)
	err := service.client.Call("http://api.exigo.com/ChangeOrderStatus", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Changes the status of an AutoOrder. This is typically used to delete or suspend an auto order.
*/
func (service *ExigoApiSoap) ChangeAutoOrderStatus(request *ChangeAutoOrderStatusRequest) (*ChangeAutoOrderStatusResponse, error) {
	response := new(ChangeAutoOrderStatusResponse)
	err := service.client.Call("http://api.exigo.com/ChangeAutoOrderStatus", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
This has been deprecated. Use CalculateOrder with ReturnShipMethods=true instead.
*/
func (service *ExigoApiSoap) GetShipMethods(request *GetShipMethodsRequest) (*GetShipMethodsResponse, error) {
	response := new(GetShipMethodsResponse)
	err := service.client.Call("http://api.exigo.com/GetShipMethods", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Returns one or more orders. You can submit several optional filter fields to control the results.
*/
func (service *ExigoApiSoap) GetOrders(request *GetOrdersRequest) (*GetOrdersResponse, error) {
	response := new(GetOrdersResponse)
	err := service.client.Call("http://api.exigo.com/GetOrders", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Returns the totals of accepted orders within a date range
*/
func (service *ExigoApiSoap) GetOrderTotals(request *GetOrderTotalsRequest) (*GetOrderTotalsResponse, error) {
	response := new(GetOrderTotalsResponse)
	err := service.client.Call("http://api.exigo.com/GetOrderTotals", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Returns any AutoOrders setup and active for a given customer.
*/
func (service *ExigoApiSoap) GetAutoOrders(request *GetAutoOrdersRequest) (*GetAutoOrdersResponse, error) {
	response := new(GetAutoOrdersResponse)
	err := service.client.Call("http://api.exigo.com/GetAutoOrders", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Updates one or more fields on an existing customer. Can participate in a transaction.
*/
func (service *ExigoApiSoap) UpdateCustomer(request *UpdateCustomerRequest) (*UpdateCustomerResponse, error) {
	response := new(UpdateCustomerResponse)
	err := service.client.Call("http://api.exigo.com/UpdateCustomer", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Updates one or more fields on an existing order. Can participate in a transaction.
*/
func (service *ExigoApiSoap) UpdateOrder(request *UpdateOrderRequest) (*UpdateOrderResponse, error) {
	response := new(UpdateOrderResponse)
	err := service.client.Call("http://api.exigo.com/UpdateOrder", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Create a new Order Detail for an existing order without recalculating. Can participate in a transaction.
*/
func (service *ExigoApiSoap) CreateOrderDetail(request *CreateOrderDetailRequest) (*CreateOrderDetailResponse, error) {
	response := new(CreateOrderDetailResponse)
	err := service.client.Call("http://api.exigo.com/CreateOrderDetail", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Update an existing Order Detail for an existing order without recalculating. Can participate in a transaction.
*/
func (service *ExigoApiSoap) UpdateOrderDetail(request *UpdateOrderDetailRequest) (*UpdateOrderDetailResponse, error) {
	response := new(UpdateOrderDetailResponse)
	err := service.client.Call("http://api.exigo.com/UpdateOrderDetail", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Delete an existing Order Detail for an existing order without recalculating. Can participate in a transaction.
*/
func (service *ExigoApiSoap) DeleteOrderDetail(request *DeleteOrderDetailRequest) (*DeleteOrderDetailResponse, error) {
	response := new(DeleteOrderDetailResponse)
	err := service.client.Call("http://api.exigo.com/DeleteOrderDetail", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Changes existing orders from one status to another in a batch. Can optionally update tracking numbers as well.
*/
func (service *ExigoApiSoap) ChangeOrderStatusBatch(request *ChangeOrderStatusBatchRequest) (*ChangeOrderStatusBatchResponse, error) {
	response := new(ChangeOrderStatusBatchResponse)
	err := service.client.Call("http://api.exigo.com/ChangeOrderStatusBatch", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Merges Order, Payments and AutoOrder data from two Customers
*/
func (service *ExigoApiSoap) MergeCustomer(request *MergeCustomerRequest) (*MergeCustomerResponse, error) {
	response := new(MergeCustomerResponse)
	err := service.client.Call("http://api.exigo.com/MergeCustomer", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Inserts/Moves a customer within the Enroller Tree.
*/
func (service *ExigoApiSoap) PlaceEnrollerNode(request *PlaceEnrollerNodeRequest) (*PlaceEnrollerNodeResponse, error) {
	response := new(PlaceEnrollerNodeResponse)
	err := service.client.Call("http://api.exigo.com/PlaceEnrollerNode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Inserts/Moves a customer within the Stack Tree.
*/
func (service *ExigoApiSoap) PlaceStackNode(request *PlaceStackNodeRequest) (*PlaceStackNodeResponse, error) {
	response := new(PlaceStackNodeResponse)
	err := service.client.Call("http://api.exigo.com/PlaceStackNode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Inserts/Moves a customer within the UniLevel (Sponsor) Tree.
*/
func (service *ExigoApiSoap) PlaceUniLevelNode(request *PlaceUniLevelNodeRequest) (*PlaceUniLevelNodeResponse, error) {
	response := new(PlaceUniLevelNodeResponse)
	err := service.client.Call("http://api.exigo.com/PlaceUniLevelNode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Inserts/Moves a customer within the Binary Tree.
*/
func (service *ExigoApiSoap) PlaceBinaryNode(request *PlaceBinaryNodeRequest) (*PlaceBinaryNodeResponse, error) {
	response := new(PlaceBinaryNodeResponse)
	err := service.client.Call("http://api.exigo.com/PlaceBinaryNode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Returns the currently configured placement preference for new placements under a customer.
*/
func (service *ExigoApiSoap) GetBinaryPreference(request *GetBinaryPreferenceRequest) (*GetBinaryPreferenceResponse, error) {
	response := new(GetBinaryPreferenceResponse)
	err := service.client.Call("http://api.exigo.com/GetBinaryPreference", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Sets the placement preference for new placements under a customer.
*/
func (service *ExigoApiSoap) SetBinaryPreference(request *SetBinaryPreferenceRequest) (*SetBinaryPreferenceResponse, error) {
	response := new(SetBinaryPreferenceResponse)
	err := service.client.Call("http://api.exigo.com/SetBinaryPreference", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Inserts/Moves a customer within the Matrix Tree.
*/
func (service *ExigoApiSoap) PlaceMatrixNode(request *PlaceMatrixNodeRequest) (*PlaceMatrixNodeResponse, error) {
	response := new(PlaceMatrixNodeResponse)
	err := service.client.Call("http://api.exigo.com/PlaceMatrixNode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
Returns countries setup for company as well as the regions for a single country requested.
*/
func (service *ExigoApiSoap) GetCountryRegions(request *GetCountryRegionsRequest) (*GetCountryRegionsResponse, error) {
	response := new(GetCountryRegionsResponse)
	err := service.client.Call("http://api.exigo.com/GetCountryRegions", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Gets nodes in a downline.
*/
func (service *ExigoApiSoap) GetDownline(request *GetDownlineRequest) (*GetDownlineResponse, error) {
	response := new(GetDownlineResponse)
	err := service.client.Call("http://api.exigo.com/GetDownline", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Gets nodes in an upline.
*/
func (service *ExigoApiSoap) GetUpline(request *GetUplineRequest) (*GetUplineResponse, error) {
	response := new(GetUplineResponse)
	err := service.client.Call("http://api.exigo.com/GetUpline", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   This will 'pop' the oldest 100 customer events off the event queue.
*/
func (service *ExigoApiSoap) DequeueCustomerEvents(request *DequeueCustomerEventsRequest) (*DequeueCustomerEventsResponse, error) {
	response := new(DequeueCustomerEventsResponse)
	err := service.client.Call("http://api.exigo.com/DequeueCustomerEvents", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Create a transaction that adjusts or redeems points from a customer's point account.
*/
func (service *ExigoApiSoap) CreatePointTransaction(request *CreatePointTransactionRequest) (*CreatePointTransactionResponse, error) {
	response := new(CreatePointTransactionResponse)
	err := service.client.Call("http://api.exigo.com/CreatePointTransaction", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Get the balance of a point account for a customer.
*/
func (service *ExigoApiSoap) GetPointAccount(request *GetPointAccountRequest) (*GetPointAccountResponse, error) {
	response := new(GetPointAccountResponse)
	err := service.client.Call("http://api.exigo.com/GetPointAccount", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Get the subscription account for a customer.
*/
func (service *ExigoApiSoap) GetSubscription(request *GetSubscriptionRequest) (*GetSubscriptionResponse, error) {
	response := new(GetSubscriptionResponse)
	err := service.client.Call("http://api.exigo.com/GetSubscription", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Validates Business Rules.
*/
func (service *ExigoApiSoap) Validate(request *ValidateRequest) (*ValidateResponse, error) {
	response := new(ValidateResponse)
	err := service.client.Call("http://api.exigo.com/Validate", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Verifies and cleans up an address.
*/
func (service *ExigoApiSoap) VerifyAddress(request *VerifyAddressRequest) (*VerifyAddressResponse, error) {
	response := new(VerifyAddressResponse)
	err := service.client.Call("http://api.exigo.com/VerifyAddress", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Opts out all customers and customer contacts with specified email from receiving broadcast emails.
*/
func (service *ExigoApiSoap) OptOutEmail(request *OptOutEmailRequest) (*OptOutEmailResponse, error) {
	response := new(OptOutEmailResponse)
	err := service.client.Call("http://api.exigo.com/OptOutEmail", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Opts out all customers and customer contacts with specified phone number from receiving broadcast SMS messages.
*/
func (service *ExigoApiSoap) OptOutSms(request *OptOutSmsRequest) (*OptOutSmsResponse, error) {
	response := new(OptOutSmsResponse)
	err := service.client.Call("http://api.exigo.com/OptOutSms", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Gets a existing Shopping Cart Session.
*/
func (service *ExigoApiSoap) GetShoppingCart(request *GetShoppingCartRequest) (*GetShoppingCartResponse, error) {
	response := new(GetShoppingCartResponse)
	err := service.client.Call("http://api.exigo.com/GetShoppingCart", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Gets a list of warehouses setup in the system.
*/
func (service *ExigoApiSoap) GetWarehouses(request *GetWarehousesRequest) (*GetWarehousesResponse, error) {
	response := new(GetWarehousesResponse)
	err := service.client.Call("http://api.exigo.com/GetWarehouses", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Gets a generic session string for a unique sessionID.
*/
func (service *ExigoApiSoap) GetSession(request *GetSessionRequest) (*GetSessionResponse, error) {
	response := new(GetSessionResponse)
	err := service.client.Call("http://api.exigo.com/GetSession", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Sets a generic session string for a unique sessionID.
*/
func (service *ExigoApiSoap) SetSession(request *SetSessionRequest) (*SetSessionResponse, error) {
	response := new(SetSessionResponse)
	err := service.client.Call("http://api.exigo.com/SetSession", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Returns price, volume, and description for one or more item codes.
*/
func (service *ExigoApiSoap) GetItems(request *GetItemsRequest) (*GetItemsResponse, error) {
	response := new(GetItemsResponse)
	err := service.client.Call("http://api.exigo.com/GetItems", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Gets languages list available to the company.
*/
func (service *ExigoApiSoap) GetCompanyLanguages(request *GetLanguagesRequest) (*GetLanguagesResponse, error) {
	response := new(GetLanguagesResponse)
	err := service.client.Call("http://api.exigo.com/GetCompanyLanguages", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Creates a new product web category.
*/
func (service *ExigoApiSoap) CreateWebCategory(request *CreateWebCategoryRequest) (*CreateWebCategoryResponse, error) {
	response := new(CreateWebCategoryResponse)
	err := service.client.Call("http://api.exigo.com/CreateWebCategory", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Updates web category information.
*/
func (service *ExigoApiSoap) UpdateWebCategory(request *UpdateWebCategoryRequest) (*UpdateWebCategoryResponse, error) {
	response := new(UpdateWebCategoryResponse)
	err := service.client.Call("http://api.exigo.com/UpdateWebCategory", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Deletes a web category.
*/
func (service *ExigoApiSoap) DeleteWebCategory(request *DeleteWebCategoryRequest) (*DeleteWebCategoryResponse, error) {
	response := new(DeleteWebCategoryResponse)
	err := service.client.Call("http://api.exigo.com/DeleteWebCategory", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Adds items/products to a web category.
*/
func (service *ExigoApiSoap) AddProductsToCategory(request *AddProductsToCategoryRequest) (*AddProductsToCategoryResponse, error) {
	response := new(AddProductsToCategoryResponse)
	err := service.client.Call("http://api.exigo.com/AddProductsToCategory", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Deletes a product from a web category.
*/
func (service *ExigoApiSoap) DeleteProductFromCategory(request *DeleteProductFromCategoryRequest) (*DeleteProductFromCategoryResponse, error) {
	response := new(DeleteProductFromCategoryResponse)
	err := service.client.Call("http://api.exigo.com/DeleteProductFromCategory", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Gets a list of company news descriptions.
*/
func (service *ExigoApiSoap) GetCompanyNews(request *GetCompanyNewsRequest) (*GetCompanyNewsResponse, error) {
	response := new(GetCompanyNewsResponse)
	err := service.client.Call("http://api.exigo.com/GetCompanyNews", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Gets company news details.
*/
func (service *ExigoApiSoap) GetCompanyNewsItem(request *GetCompanyNewsItemRequest) (*GetCompanyNewsItemResponse, error) {
	response := new(GetCompanyNewsItemResponse)
	err := service.client.Call("http://api.exigo.com/GetCompanyNewsItem", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Gets a random message (defined in the Exigo Admin).
*/
func (service *ExigoApiSoap) GetRandomMessage(request *GetRandomMessageRequest) (*GetRandomMessageResponse, error) {
	response := new(GetRandomMessageResponse)
	err := service.client.Call("http://api.exigo.com/GetRandomMessage", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Fires an email responder to customer or order email address on file.
*/
func (service *ExigoApiSoap) FireResponder(request *FireResponderRequest) (*FireResponderResponse, error) {
	response := new(FireResponderResponse)
	err := service.client.Call("http://api.exigo.com/FireResponder", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Sends an SMS message.
*/
func (service *ExigoApiSoap) SendSms(request *SendSmsRequest) (*SendSmsResponse, error) {
	response := new(SendSmsResponse)
	err := service.client.Call("http://api.exigo.com/SendSms", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Creates a new VendorBill for a customer on file.
*/
func (service *ExigoApiSoap) CreateVendorBill(request *CreateVendorBillRequest) (*CreateVendorBillResponse, error) {
	response := new(CreateVendorBillResponse)
	err := service.client.Call("http://api.exigo.com/CreateVendorBill", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Creates a new customer contact.
*/
func (service *ExigoApiSoap) CreateCustomerContact(request *CreateCustomerContactRequest) (*CreateCustomerContactResponse, error) {
	response := new(CreateCustomerContactResponse)
	err := service.client.Call("http://api.exigo.com/CreateCustomerContact", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Updates customer contact information.
*/
func (service *ExigoApiSoap) UpdateCustomerContact(request *UpdateCustomerContactRequest) (*UpdateCustomerContactResponse, error) {
	response := new(UpdateCustomerContactResponse)
	err := service.client.Call("http://api.exigo.com/UpdateCustomerContact", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Deletes a customer contact.
*/
func (service *ExigoApiSoap) DeleteCustomerContact(request *DeleteCustomerContactRequest) (*DeleteCustomerContactResponse, error) {
	response := new(DeleteCustomerContactResponse)
	err := service.client.Call("http://api.exigo.com/DeleteCustomerContact", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Deletes a customer lead.
*/
func (service *ExigoApiSoap) DeleteCustomerLead(request *DeleteCustomerLeadRequest) (*DeleteCustomerLeadResponse, error) {
	response := new(DeleteCustomerLeadResponse)
	err := service.client.Call("http://api.exigo.com/DeleteCustomerLead", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Creates a new customer lead.
*/
func (service *ExigoApiSoap) CreateCustomerLead(request *CreateCustomerLeadRequest) (*CreateCustomerLeadResponse, error) {
	response := new(CreateCustomerLeadResponse)
	err := service.client.Call("http://api.exigo.com/CreateCustomerLead", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Gets a list of customer leads.
*/
func (service *ExigoApiSoap) GetCustomerLeads(request *GetCustomerLeadsRequest) (*GetCustomerLeadsResponse, error) {
	response := new(GetCustomerLeadsResponse)
	err := service.client.Call("http://api.exigo.com/GetCustomerLeads", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Updates a customer leads.
*/
func (service *ExigoApiSoap) UpdateCustomerLead(request *UpdateCustomerLeadRequest) (*UpdateCustomerLeadResponse, error) {
	response := new(UpdateCustomerLeadResponse)
	err := service.client.Call("http://api.exigo.com/UpdateCustomerLead", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Creates a new calendar item.
*/
func (service *ExigoApiSoap) CreateCalendarItem(request *CreateCalendarItemRequest) (*CreateCalendarItemResponse, error) {
	response := new(CreateCalendarItemResponse)
	err := service.client.Call("http://api.exigo.com/CreateCalendarItem", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Updates calendar item information.
*/
func (service *ExigoApiSoap) UpdateCalendarItem(request *UpdateCalendarItemRequest) (*UpdateCalendarItemResponse, error) {
	response := new(UpdateCalendarItemResponse)
	err := service.client.Call("http://api.exigo.com/UpdateCalendarItem", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
   Deletes a calendar item.
*/
func (service *ExigoApiSoap) DeleteCalendarItem(request *DeleteCalendarItemRequest) (*DeleteCalendarItemResponse, error) {
	response := new(DeleteCalendarItemResponse)
	err := service.client.Call("http://api.exigo.com/DeleteCalendarItem", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  *SOAPHeader
	Body    SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Header interface{}
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url    string
	tls    bool
	auth   *BasicAuth
	header interface{}
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func NewSOAPClient(url string, tls bool, auth *BasicAuth) *SOAPClient {
	return &SOAPClient{
		url:  url,
		tls:  tls,
		auth: auth,
	}
}

func (s *SOAPClient) SetHeader(header interface{}) {
	s.header = header
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{}

	if s.header != nil {
		envelope.Header = &SOAPHeader{Header: s.header}
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	//log.Println(buffer.String())

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	if soapAction != "" {
		req.Header.Add("SOAPAction", soapAction)
	}

	req.Header.Set("User-Agent", "gowsdl/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.tls,
		},
		Dial: dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	//log.Println(string(rawbody))
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		log.Println(string(rawbody))
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		log.Println(string(rawbody))
		return fault
	}

	return nil
}
