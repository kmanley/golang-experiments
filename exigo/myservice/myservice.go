package myservice

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"log"
	"math/rand"
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
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateEmailRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	MailFolderType *MailForderType `xml:"MailFolderType,omitempty"`

	Priority *MailPriority `xml:"Priority,omitempty"`

	MailStatusType *MailStatusType `xml:"MailStatusType,omitempty"`

	Subject string `xml:"Subject,omitempty"`

	MailFrom string `xml:"MailFrom,omitempty"`

	MailTo string `xml:"MailTo,omitempty"`

	ReplyTo string `xml:"ReplyTo,omitempty"`

	MailCC string `xml:"MailCC,omitempty"`

	MailBCC string `xml:"MailBCC,omitempty"`

	Content string `xml:"Content,omitempty"`

	SmtpServer string `xml:"SmtpServer,omitempty"`

	Attachments *ArrayOfEmailAttachment `xml:"Attachments,omitempty"`

	ForwardedAttachments *ArrayOfForwardedAttachment `xml:"ForwardedAttachments,omitempty"`
}

type ApiRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ApiRequest"`
}

type DeleteOrderDetailRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteOrderDetailRequest"`

	*ApiRequest

	OrderID int32 `xml:"OrderID,omitempty"`

	OrderLine int32 `xml:"OrderLine,omitempty"`
}

type UpdateOrderDetailRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateOrderDetailRequest"`

	*ApiRequest

	OrderID int32 `xml:"OrderID,omitempty"`

	ItemID int32 `xml:"ItemID,omitempty"`

	ItemCode string `xml:"ItemCode,omitempty"`

	OrderLine int32 `xml:"OrderLine,omitempty"`

	Description string `xml:"Description,omitempty"`

	Qty float64 `xml:"Qty,omitempty"`

	PriceEach float64 `xml:"PriceEach,omitempty"`

	PriceExt float64 `xml:"PriceExt,omitempty"`

	BVEach float64 `xml:"BVEach,omitempty"`

	BusinessVolume float64 `xml:"BusinessVolume,omitempty"`

	CVEach float64 `xml:"CVEach,omitempty"`

	CommissionableVolume float64 `xml:"CommissionableVolume,omitempty"`

	ShippingPriceEach float64 `xml:"ShippingPriceEach,omitempty"`

	ChargeShippingOn float64 `xml:"ChargeShippingOn,omitempty"`

	IsTaxedInRegion bool `xml:"IsTaxedInRegion,omitempty"`

	IsTaxedInRegionFed bool `xml:"IsTaxedInRegionFed,omitempty"`

	IsTaxedInRegionState bool `xml:"IsTaxedInRegionState,omitempty"`

	TaxablePriceEach float64 `xml:"TaxablePriceEach,omitempty"`

	Taxable float64 `xml:"Taxable,omitempty"`

	CombinedTax float64 `xml:"CombinedTax,omitempty"`

	FedTax float64 `xml:"FedTax,omitempty"`

	StateTax float64 `xml:"StateTax,omitempty"`

	CityTax float64 `xml:"CityTax,omitempty"`

	CityLocalTax float64 `xml:"CityLocalTax,omitempty"`

	CountyTax float64 `xml:"CountyTax,omitempty"`

	CountyLocalTax float64 `xml:"CountyLocalTax,omitempty"`

	ManualTax float64 `xml:"ManualTax,omitempty"`

	IsBackOrder bool `xml:"IsBackOrder,omitempty"`

	WeightEach float64 `xml:"WeightEach,omitempty"`

	Other1Each float64 `xml:"Other1Each,omitempty"`

	Other1 float64 `xml:"Other1,omitempty"`

	Other2Each float64 `xml:"Other2Each,omitempty"`

	Other2 float64 `xml:"Other2,omitempty"`

	Other3Each float64 `xml:"Other3Each,omitempty"`

	Other3 float64 `xml:"Other3,omitempty"`

	Other4Each float64 `xml:"Other4Each,omitempty"`

	Other4 float64 `xml:"Other4,omitempty"`

	Other5Each float64 `xml:"Other5Each,omitempty"`

	Other5 float64 `xml:"Other5,omitempty"`

	Other6Each float64 `xml:"Other6Each,omitempty"`

	Other6 float64 `xml:"Other6,omitempty"`

	Other7Each float64 `xml:"Other7Each,omitempty"`

	Other7 float64 `xml:"Other7,omitempty"`

	Other8Each float64 `xml:"Other8Each,omitempty"`

	Other8 float64 `xml:"Other8,omitempty"`

	Other9Each float64 `xml:"Other9Each,omitempty"`

	Other9 float64 `xml:"Other9,omitempty"`

	Other10Each float64 `xml:"Other10Each,omitempty"`

	Other10 float64 `xml:"Other10,omitempty"`

	DiscountExt float64 `xml:"DiscountExt,omitempty"`

	OriginalTaxableEach float64 `xml:"OriginalTaxableEach,omitempty"`

	OriginalBVEach float64 `xml:"OriginalBVEach,omitempty"`

	OriginalCVEach float64 `xml:"OriginalCVEach,omitempty"`

	StateTaxable float64 `xml:"StateTaxable,omitempty"`

	IsStateTaxOverride bool `xml:"IsStateTaxOverride,omitempty"`

	DynamicKitItemID int32 `xml:"DynamicKitItemID,omitempty"`

	HandlingFee float64 `xml:"HandlingFee,omitempty"`

	Reference1 string `xml:"Reference1,omitempty"`
}

type CreateOrderDetailRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateOrderDetailRequest"`

	*ApiRequest

	OrderID int32 `xml:"OrderID,omitempty"`

	OrderLine int32 `xml:"OrderLine,omitempty"`

	ItemCode string `xml:"ItemCode,omitempty"`

	Description string `xml:"Description,omitempty"`

	Qty float64 `xml:"Qty,omitempty"`

	PriceEach float64 `xml:"PriceEach,omitempty"`

	PriceExt float64 `xml:"PriceExt,omitempty"`

	BVEach float64 `xml:"BVEach,omitempty"`

	BusinessVolume float64 `xml:"BusinessVolume,omitempty"`

	CVEach float64 `xml:"CVEach,omitempty"`

	CommissionableVolume float64 `xml:"CommissionableVolume,omitempty"`

	ShippingPriceEach float64 `xml:"ShippingPriceEach,omitempty"`

	ChargeShippingOn float64 `xml:"ChargeShippingOn,omitempty"`

	IsTaxedInRegion bool `xml:"IsTaxedInRegion,omitempty"`

	IsTaxedInRegionFed bool `xml:"IsTaxedInRegionFed,omitempty"`

	IsTaxedInRegionState bool `xml:"IsTaxedInRegionState,omitempty"`

	TaxablePriceEach float64 `xml:"TaxablePriceEach,omitempty"`

	Taxable float64 `xml:"Taxable,omitempty"`

	CombinedTax float64 `xml:"CombinedTax,omitempty"`

	FedTax float64 `xml:"FedTax,omitempty"`

	StateTax float64 `xml:"StateTax,omitempty"`

	CityTax float64 `xml:"CityTax,omitempty"`

	CityLocalTax float64 `xml:"CityLocalTax,omitempty"`

	CountyTax float64 `xml:"CountyTax,omitempty"`

	CountyLocalTax float64 `xml:"CountyLocalTax,omitempty"`

	ManualTax float64 `xml:"ManualTax,omitempty"`

	IsBackOrder bool `xml:"IsBackOrder,omitempty"`

	WeightEach float64 `xml:"WeightEach,omitempty"`

	Other1Each float64 `xml:"Other1Each,omitempty"`

	Other1 float64 `xml:"Other1,omitempty"`

	Other2Each float64 `xml:"Other2Each,omitempty"`

	Other2 float64 `xml:"Other2,omitempty"`

	Other3Each float64 `xml:"Other3Each,omitempty"`

	Other3 float64 `xml:"Other3,omitempty"`

	Other4Each float64 `xml:"Other4Each,omitempty"`

	Other4 float64 `xml:"Other4,omitempty"`

	Other5Each float64 `xml:"Other5Each,omitempty"`

	Other5 float64 `xml:"Other5,omitempty"`

	Other6Each float64 `xml:"Other6Each,omitempty"`

	Other6 float64 `xml:"Other6,omitempty"`

	Other7Each float64 `xml:"Other7Each,omitempty"`

	Other7 float64 `xml:"Other7,omitempty"`

	Other8Each float64 `xml:"Other8Each,omitempty"`

	Other8 float64 `xml:"Other8,omitempty"`

	Other9Each float64 `xml:"Other9Each,omitempty"`

	Other9 float64 `xml:"Other9,omitempty"`

	Other10Each float64 `xml:"Other10Each,omitempty"`

	Other10 float64 `xml:"Other10,omitempty"`

	DiscountExt float64 `xml:"DiscountExt,omitempty"`

	OriginalTaxableEach float64 `xml:"OriginalTaxableEach,omitempty"`

	OriginalBVEach float64 `xml:"OriginalBVEach,omitempty"`

	OriginalCVEach float64 `xml:"OriginalCVEach,omitempty"`

	StateTaxable float64 `xml:"StateTaxable,omitempty"`

	IsStateTaxOverride bool `xml:"IsStateTaxOverride,omitempty"`

	DynamicKitItemID int32 `xml:"DynamicKitItemID,omitempty"`

	HandlingFee float64 `xml:"HandlingFee,omitempty"`

	Reference1 string `xml:"Reference1,omitempty"`
}

type BaseAuthorizeOnlyCreditCardTokenRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ BaseAuthorizeOnlyCreditCardTokenRequest"`

	*ApiRequest
}

type AuthorizeOnlyCreditCardTokenOnFileRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ AuthorizeOnlyCreditCardTokenOnFileRequest"`

	*BaseAuthorizeOnlyCreditCardTokenRequest

	CreditCardAccountType *AccountCreditCardType `xml:"CreditCardAccountType,omitempty"`

	WarehouseID int32 `xml:"WarehouseID,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	Amount float64 `xml:"Amount,omitempty"`
}

type AuthorizeOnlyCreditCardTokenRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ AuthorizeOnlyCreditCardTokenRequest"`

	*BaseAuthorizeOnlyCreditCardTokenRequest

	CreditCardToken string `xml:"CreditCardToken,omitempty"`

	Amount float64 `xml:"Amount,omitempty"`

	BillingName string `xml:"BillingName,omitempty"`

	BillingAddress string `xml:"BillingAddress,omitempty"`

	BillingCity string `xml:"BillingCity,omitempty"`

	BillingState string `xml:"BillingState,omitempty"`

	BillingZip string `xml:"BillingZip,omitempty"`

	BillingCountry string `xml:"BillingCountry,omitempty"`

	WarehouseID int32 `xml:"WarehouseID,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	CvcCode string `xml:"CvcCode,omitempty"`
}

type SendEmailRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SendEmailRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	MailFrom string `xml:"MailFrom,omitempty"`

	MailTo string `xml:"MailTo,omitempty"`

	Subject string `xml:"Subject,omitempty"`

	Body string `xml:"Body,omitempty"`
}

type SetItemKitMembersRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetItemKitMembersRequest"`

	*ApiRequest

	ParentItemCode string `xml:"ParentItemCode,omitempty"`

	ItemKitMembers *ArrayOfKitMember `xml:"ItemKitMembers,omitempty"`
}

type ArrayOfKitMember struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfKitMember"`

	KitMember []*KitMember `xml:"KitMember,omitempty"`
}

type KitMember struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ KitMember"`

	ItemID int32 `xml:"ItemID,omitempty"`

	ItemCode string `xml:"ItemCode,omitempty"`

	Quantity int32 `xml:"Quantity,omitempty"`
}

type GetFileContentsRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetFileContentsRequest"`

	*ApiRequest

	FileName string `xml:"FileName,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty"`
}

type GetFilesRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetFilesRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`
}

type ChargeGroupOrderCreditCardTokenRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ChargeGroupOrderCreditCardTokenRequest"`

	*ApiRequest

	_orders *ArrayOfGroupOrder `xml:"_orders,omitempty"`

	CreditCardToken string `xml:"CreditCardToken,omitempty"`

	BillingName string `xml:"BillingName,omitempty"`

	BillingAddress string `xml:"BillingAddress,omitempty"`

	BillingAddress2 string `xml:"BillingAddress2,omitempty"`

	BillingCity string `xml:"BillingCity,omitempty"`

	BillingState string `xml:"BillingState,omitempty"`

	BillingZip string `xml:"BillingZip,omitempty"`

	BillingCountry string `xml:"BillingCountry,omitempty"`

	CvcCode string `xml:"CvcCode,omitempty"`

	IssueNumber string `xml:"IssueNumber,omitempty"`

	CreditCardType int32 `xml:"CreditCardType,omitempty"`

	MasterOrderID int32 `xml:"MasterOrderID,omitempty"`

	Orders *ArrayOfGroupOrder `xml:"Orders,omitempty"`

	MerchantWarehouseIDOverride int32 `xml:"MerchantWarehouseIDOverride,omitempty"`

	ClientIPAddress string `xml:"ClientIPAddress,omitempty"`

	OtherData1 string `xml:"OtherData1,omitempty"`

	OtherData2 string `xml:"OtherData2,omitempty"`

	OtherData3 string `xml:"OtherData3,omitempty"`

	OtherData4 string `xml:"OtherData4,omitempty"`

	OtherData5 string `xml:"OtherData5,omitempty"`

	OtherData6 string `xml:"OtherData6,omitempty"`

	OtherData7 string `xml:"OtherData7,omitempty"`

	OtherData8 string `xml:"OtherData8,omitempty"`

	OtherData9 string `xml:"OtherData9,omitempty"`

	OtherData10 string `xml:"OtherData10,omitempty"`

	PaymentMemo string `xml:"PaymentMemo,omitempty"`
}

type ArrayOfGroupOrder struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfGroupOrder"`

	GroupOrder []*GroupOrder `xml:"GroupOrder,omitempty"`
}

type GroupOrder struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GroupOrder"`

	OrderID int32 `xml:"OrderID,omitempty"`

	MaxAmount float64 `xml:"MaxAmount,omitempty"`
}

type DeleteCustomerExtendedRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteCustomerExtendedRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	ExtendedGroupID int32 `xml:"ExtendedGroupID,omitempty"`

	CustomerExtendedID int32 `xml:"CustomerExtendedID,omitempty"`
}

type UpdateCustomerExtendedRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateCustomerExtendedRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	ExtendedGroupID int32 `xml:"ExtendedGroupID,omitempty"`

	CustomerExtendedID int32 `xml:"CustomerExtendedID,omitempty"`

	Field1 string `xml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty"`

	Field11 string `xml:"Field11,omitempty"`

	Field12 string `xml:"Field12,omitempty"`

	Field13 string `xml:"Field13,omitempty"`

	Field14 string `xml:"Field14,omitempty"`

	Field15 string `xml:"Field15,omitempty"`

	Field16 string `xml:"Field16,omitempty"`

	Field17 string `xml:"Field17,omitempty"`

	Field18 string `xml:"Field18,omitempty"`

	Field19 string `xml:"Field19,omitempty"`

	Field20 string `xml:"Field20,omitempty"`

	Field21 string `xml:"Field21,omitempty"`

	Field22 string `xml:"Field22,omitempty"`

	Field23 string `xml:"Field23,omitempty"`

	Field24 string `xml:"Field24,omitempty"`

	Field25 string `xml:"Field25,omitempty"`

	Field26 string `xml:"Field26,omitempty"`

	Field27 string `xml:"Field27,omitempty"`

	Field28 string `xml:"Field28,omitempty"`

	Field29 string `xml:"Field29,omitempty"`

	Field30 string `xml:"Field30,omitempty"`
}

type CreateCustomerExtendedRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateCustomerExtendedRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	ExtendedGroupID int32 `xml:"ExtendedGroupID,omitempty"`

	Field1 string `xml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty"`

	Field11 string `xml:"Field11,omitempty"`

	Field12 string `xml:"Field12,omitempty"`

	Field13 string `xml:"Field13,omitempty"`

	Field14 string `xml:"Field14,omitempty"`

	Field15 string `xml:"Field15,omitempty"`

	Field16 string `xml:"Field16,omitempty"`

	Field17 string `xml:"Field17,omitempty"`

	Field18 string `xml:"Field18,omitempty"`

	Field19 string `xml:"Field19,omitempty"`

	Field20 string `xml:"Field20,omitempty"`

	Field21 string `xml:"Field21,omitempty"`

	Field22 string `xml:"Field22,omitempty"`

	Field23 string `xml:"Field23,omitempty"`

	Field24 string `xml:"Field24,omitempty"`

	Field25 string `xml:"Field25,omitempty"`

	Field26 string `xml:"Field26,omitempty"`

	Field27 string `xml:"Field27,omitempty"`

	Field28 string `xml:"Field28,omitempty"`

	Field29 string `xml:"Field29,omitempty"`

	Field30 string `xml:"Field30,omitempty"`
}

type UpdatePartyRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdatePartyRequest"`

	*ApiRequest

	PartyID int32 `xml:"PartyID,omitempty"`

	PartyType int32 `xml:"PartyType,omitempty"`

	PartyStatusType int32 `xml:"PartyStatusType,omitempty"`

	HostID int32 `xml:"HostID,omitempty"`

	DistributorID int32 `xml:"DistributorID,omitempty"`

	StartDate time.Time `xml:"StartDate,omitempty"`

	CloseDate time.Time `xml:"CloseDate,omitempty"`

	Description string `xml:"Description,omitempty"`

	EventStart time.Time `xml:"EventStart,omitempty"`

	EventEnd time.Time `xml:"EventEnd,omitempty"`

	LanguageID int32 `xml:"LanguageID,omitempty"`

	Information string `xml:"Information,omitempty"`

	Address *PartyAddress `xml:"Address,omitempty"`

	BookingPartyID int32 `xml:"BookingPartyID,omitempty"`

	Field1 string `xml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty"`
}

type PartyAddress struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ PartyAddress"`

	Address1 string `xml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty"`

	City string `xml:"City,omitempty"`

	State string `xml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty"`
}

type CreatePartyRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreatePartyRequest"`

	*ApiRequest

	PartyType int32 `xml:"PartyType,omitempty"`

	PartyStatusType int32 `xml:"PartyStatusType,omitempty"`

	HostID int32 `xml:"HostID,omitempty"`

	DistributorID int32 `xml:"DistributorID,omitempty"`

	StartDate time.Time `xml:"StartDate,omitempty"`

	CloseDate time.Time `xml:"CloseDate,omitempty"`

	Description string `xml:"Description,omitempty"`

	EventStart time.Time `xml:"EventStart,omitempty"`

	EventEnd time.Time `xml:"EventEnd,omitempty"`

	LanguageID int32 `xml:"LanguageID,omitempty"`

	Information string `xml:"Information,omitempty"`

	Address *PartyAddress `xml:"Address,omitempty"`

	BookingPartyID int32 `xml:"BookingPartyID,omitempty"`

	Field1 string `xml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty"`
}

type CreateBillRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateBillRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	IsOtherIncome bool `xml:"IsOtherIncome,omitempty"`

	DueDate time.Time `xml:"DueDate,omitempty"`

	Amount float64 `xml:"Amount,omitempty"`

	Reference string `xml:"Reference,omitempty"`

	Notes string `xml:"Notes,omitempty"`

	BillStatusTypeID int32 `xml:"BillStatusTypeID,omitempty"`

	PayableTypeIDOverride int32 `xml:"PayableTypeIDOverride,omitempty"`
}

type BaseCreatePayoutRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ BaseCreatePayoutRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	BankAccountID int32 `xml:"BankAccountID,omitempty"`

	Reference string `xml:"Reference,omitempty"`

	TransactionNote string `xml:"TransactionNote,omitempty"`

	PaymentDate time.Time `xml:"PaymentDate,omitempty"`
}

type CreatePayoutRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreatePayoutRequest"`

	*BaseCreatePayoutRequest

	BillIDs_ToPay *ArrayOfInt `xml:"BillIDs_ToPay,omitempty"`

	VendorPaymentTypeID int32 `xml:"VendorPaymentTypeID,omitempty"`
}

type ArrayOfInt struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfInt"`

	Int []int32 `xml:"int,omitempty"`
}

type CreateCustomerInquiryRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateCustomerInquiryRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	Detail string `xml:"Detail,omitempty"`

	Description string `xml:"Description,omitempty"`

	AssignToUser string `xml:"AssignToUser,omitempty"`

	CustomerInquiryStatusID int32 `xml:"CustomerInquiryStatusID,omitempty"`

	CustomerInquiryCategoryID int32 `xml:"CustomerInquiryCategoryID,omitempty"`
}

type CreateCustomerFileRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateCustomerFileRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	FileName string `xml:"FileName,omitempty"`

	FileData []byte `xml:"FileData,omitempty"`

	OverwriteExistingFile bool `xml:"OverwriteExistingFile,omitempty"`
}

type SetItemCountryRegionRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetItemCountryRegionRequest"`

	*ApiRequest

	ItemCode string `xml:"ItemCode,omitempty"`

	CountryCode string `xml:"CountryCode,omitempty"`

	RegionCode string `xml:"RegionCode,omitempty"`

	Taxed bool `xml:"Taxed,omitempty"`

	TaxedFed bool `xml:"TaxedFed,omitempty"`

	TaxedState bool `xml:"TaxedState,omitempty"`

	UseTaxOverride bool `xml:"UseTaxOverride,omitempty"`

	TaxOverridePct float64 `xml:"TaxOverridePct,omitempty"`
}

type SetItemWarehouseRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetItemWarehouseRequest"`

	*ApiRequest

	AllowedUserWarehouses *ArrayOfInt `xml:"AllowedUserWarehouses,omitempty"`

	AllowedWarehouseManagementTypes *ArrayOfInt `xml:"AllowedWarehouseManagementTypes,omitempty"`

	ItemCode string `xml:"ItemCode,omitempty"`

	WarehouseID int32 `xml:"WarehouseID,omitempty"`

	IsAvailable bool `xml:"IsAvailable,omitempty"`

	ItemManageTypeID int32 `xml:"ItemManageTypeID,omitempty"`
}

type SetItemPriceRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetItemPriceRequest"`

	*ApiRequest

	ItemCode string `xml:"ItemCode,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	PriceType int32 `xml:"PriceType,omitempty"`

	Price float64 `xml:"Price,omitempty"`

	BusinessVolume float64 `xml:"BusinessVolume,omitempty"`

	CommissionableVolume float64 `xml:"CommissionableVolume,omitempty"`

	TaxablePrice float64 `xml:"TaxablePrice,omitempty"`

	ShippingPrice float64 `xml:"ShippingPrice,omitempty"`

	Other1Price float64 `xml:"Other1Price,omitempty"`

	Other2Price float64 `xml:"Other2Price,omitempty"`

	Other3Price float64 `xml:"Other3Price,omitempty"`

	Other4Price float64 `xml:"Other4Price,omitempty"`

	Other5Price float64 `xml:"Other5Price,omitempty"`

	Other6Price float64 `xml:"Other6Price,omitempty"`

	Other7Price float64 `xml:"Other7Price,omitempty"`

	Other8Price float64 `xml:"Other8Price,omitempty"`

	Other9Price float64 `xml:"Other9Price,omitempty"`

	Other10Price float64 `xml:"Other10Price,omitempty"`
}

type CreateItemRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateItemRequest"`

	*ApiRequest

	ItemCode string `xml:"ItemCode,omitempty"`

	Description string `xml:"Description,omitempty"`

	Weight float64 `xml:"Weight,omitempty"`

	Notes string `xml:"Notes,omitempty"`

	AvailableInAllCountryRegions bool `xml:"AvailableInAllCountryRegions,omitempty"`

	TaxedInAllCountryRegions bool `xml:"TaxedInAllCountryRegions,omitempty"`

	AvailableInAllWarehouses bool `xml:"AvailableInAllWarehouses,omitempty"`

	IsVirtual bool `xml:"IsVirtual,omitempty"`

	ItemTypeID int32 `xml:"ItemTypeID,omitempty"`

	OtherCheck1 bool `xml:"OtherCheck1,omitempty"`

	OtherCheck2 bool `xml:"OtherCheck2,omitempty"`

	OtherCheck3 bool `xml:"OtherCheck3,omitempty"`

	OtherCheck4 bool `xml:"OtherCheck4,omitempty"`

	OtherCheck5 bool `xml:"OtherCheck5,omitempty"`

	Field1 string `xml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty"`

	HideFromSearch bool `xml:"HideFromSearch,omitempty"`
}

type CreateCustomerWallItemRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateCustomerWallItemRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	Text string `xml:"Text,omitempty"`

	Field1 string `xml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty"`
}

type DeleteCustomerWallItemRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteCustomerWallItemRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	WallItemID int32 `xml:"WallItemID,omitempty"`

	OlderThanEntryDate time.Time `xml:"OlderThanEntryDate,omitempty"`
}

type GetCustomerWallRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomerWallRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	WallItemID int32 `xml:"WallItemID,omitempty"`

	OlderThanEntryDate time.Time `xml:"OlderThanEntryDate,omitempty"`

	BatchSize int32 `xml:"BatchSize,omitempty"`
}

type CreateCustomerLeadRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateCustomerLeadRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	FirstName string `xml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty"`

	Company string `xml:"Company,omitempty"`

	Phone string `xml:"Phone,omitempty"`

	Phone2 string `xml:"Phone2,omitempty"`

	MobilePhone string `xml:"MobilePhone,omitempty"`

	Fax string `xml:"Fax,omitempty"`

	Email string `xml:"Email,omitempty"`

	Address1 string `xml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty"`

	City string `xml:"City,omitempty"`

	State string `xml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty"`

	BirthDate time.Time `xml:"BirthDate,omitempty"`

	Notes string `xml:"Notes,omitempty"`
}

type UpdateCustomerLeadRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateCustomerLeadRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	CustomerLeadID int32 `xml:"CustomerLeadID,omitempty"`

	FirstName string `xml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty"`

	Company string `xml:"Company,omitempty"`

	Phone string `xml:"Phone,omitempty"`

	Phone2 string `xml:"Phone2,omitempty"`

	MobilePhone string `xml:"MobilePhone,omitempty"`

	Fax string `xml:"Fax,omitempty"`

	Email string `xml:"Email,omitempty"`

	Address1 string `xml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty"`

	City string `xml:"City,omitempty"`

	State string `xml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty"`

	BirthDate time.Time `xml:"BirthDate,omitempty"`

	Notes string `xml:"Notes,omitempty"`
}

type GetCustomerLeadsRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomerLeadsRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	CustomerLeadID int32 `xml:"CustomerLeadID,omitempty"`

	GreaterThanCustomerLeadID int32 `xml:"GreaterThanCustomerLeadID,omitempty"`

	BatchSize int32 `xml:"BatchSize,omitempty"`
}

type SetCustomerLeadSocialNetworksRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetCustomerLeadSocialNetworksRequest"`

	*ApiRequest

	CustomerLeadID int32 `xml:"CustomerLeadID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	CustomerLeadSocialNetworks *ArrayOfCustomerLeadSocialNetworkRequest `xml:"CustomerLeadSocialNetworks,omitempty"`
}

type ArrayOfCustomerLeadSocialNetworkRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfCustomerLeadSocialNetworkRequest"`

	CustomerLeadSocialNetworkRequest []*CustomerLeadSocialNetworkRequest `xml:"CustomerLeadSocialNetworkRequest,omitempty"`
}

type CustomerLeadSocialNetworkRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CustomerLeadSocialNetworkRequest"`

	SocialNetworkID int32 `xml:"SocialNetworkID,omitempty"`

	Url string `xml:"Url,omitempty"`
}

type GetCustomerLeadSocialNetworksRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomerLeadSocialNetworksRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	CustomerLeadID int32 `xml:"CustomerLeadID,omitempty"`
}

type GetCustomerSocialNetworksRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomerSocialNetworksRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`
}

type SetCustomerSocialNetworksRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetCustomerSocialNetworksRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	CustomerSocialNetworks *ArrayOfCustomerSocialNetworkRequest `xml:"CustomerSocialNetworks,omitempty"`
}

type ArrayOfCustomerSocialNetworkRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfCustomerSocialNetworkRequest"`

	CustomerSocialNetworkRequest []*CustomerSocialNetworkRequest `xml:"CustomerSocialNetworkRequest,omitempty"`
}

type CustomerSocialNetworkRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CustomerSocialNetworkRequest"`

	SocialNetworkID int32 `xml:"SocialNetworkID,omitempty"`

	Url string `xml:"Url,omitempty"`
}

type SetCustomerSiteRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetCustomerSiteRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	WebAlias string `xml:"WebAlias,omitempty"`

	FirstName string `xml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty"`

	Company string `xml:"Company,omitempty"`

	Address1 string `xml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty"`

	City string `xml:"City,omitempty"`

	State string `xml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty"`

	Email string `xml:"Email,omitempty"`

	Phone string `xml:"Phone,omitempty"`

	Phone2 string `xml:"Phone2,omitempty"`

	Fax string `xml:"Fax,omitempty"`

	Notes1 string `xml:"Notes1,omitempty"`

	Notes2 string `xml:"Notes2,omitempty"`

	Notes3 string `xml:"Notes3,omitempty"`

	Notes4 string `xml:"Notes4,omitempty"`

	Url1 string `xml:"Url1,omitempty"`

	Url2 string `xml:"Url2,omitempty"`

	Url3 string `xml:"Url3,omitempty"`

	Url4 string `xml:"Url4,omitempty"`

	Url5 string `xml:"Url5,omitempty"`

	Url6 string `xml:"Url6,omitempty"`

	Url7 string `xml:"Url7,omitempty"`

	Url8 string `xml:"Url8,omitempty"`

	Url9 string `xml:"Url9,omitempty"`

	Url10 string `xml:"Url10,omitempty"`

	Url1Description string `xml:"Url1Description,omitempty"`

	Url2Description string `xml:"Url2Description,omitempty"`

	Url3Description string `xml:"Url3Description,omitempty"`

	Url4Description string `xml:"Url4Description,omitempty"`

	Url5Description string `xml:"Url5Description,omitempty"`

	Url6Description string `xml:"Url6Description,omitempty"`

	Url7Description string `xml:"Url7Description,omitempty"`

	Url8Description string `xml:"Url8Description,omitempty"`

	Url9Description string `xml:"Url9Description,omitempty"`

	Url10Description string `xml:"Url10Description,omitempty"`
}

type OptInSmsRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ OptInSmsRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	PhoneNumber string `xml:"PhoneNumber,omitempty"`
}

type OptInEmailRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ OptInEmailRequest"`

	*ApiRequest

	Email string `xml:"Email,omitempty"`
}

type UpdateOrderRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateOrderRequest"`

	*ApiRequest

	OrderID int32 `xml:"OrderID,omitempty"`

	OrderDate time.Time `xml:"OrderDate,omitempty"`

	DeclineCount int32 `xml:"DeclineCount,omitempty"`

	OrderTy int32 `xml:"OrderTy,omitempty"`

	OrderStatus int32 `xml:"OrderStatus,omitempty"`

	PriceTy int32 `xml:"PriceTy,omitempty"`

	Total float64 `xml:"Total,omitempty"`

	SubTotal float64 `xml:"SubTotal,omitempty"`

	Shipping float64 `xml:"Shipping,omitempty"`

	OrderTax float64 `xml:"OrderTax,omitempty"`

	ShippingTax float64 `xml:"ShippingTax,omitempty"`

	FedShippingTax float64 `xml:"FedShippingTax,omitempty"`

	StateShippingTax float64 `xml:"StateShippingTax,omitempty"`

	CityShippingTax float64 `xml:"CityShippingTax,omitempty"`

	CityLocalShippingTax float64 `xml:"CityLocalShippingTax,omitempty"`

	CountyShippingTax float64 `xml:"CountyShippingTax,omitempty"`

	CountyLocalShippingTax float64 `xml:"CountyLocalShippingTax,omitempty"`

	ManualTaxRate float64 `xml:"ManualTaxRate,omitempty"`

	TotalTax float64 `xml:"TotalTax,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	PaymentMethod int32 `xml:"PaymentMethod,omitempty"`

	WarehouseID int32 `xml:"WarehouseID,omitempty"`

	BatchID int32 `xml:"BatchID,omitempty"`

	PreviousBalance float64 `xml:"PreviousBalance,omitempty"`

	OverrideShipping bool `xml:"OverrideShipping,omitempty"`

	OverrideTax bool `xml:"OverrideTax,omitempty"`

	BusinessVolume float64 `xml:"BusinessVolume,omitempty"`

	CommissionableVolume float64 `xml:"CommissionableVolume,omitempty"`

	Other1 float64 `xml:"Other1,omitempty"`

	Other2 float64 `xml:"Other2,omitempty"`

	Other3 float64 `xml:"Other3,omitempty"`

	Other4 float64 `xml:"Other4,omitempty"`

	Other5 float64 `xml:"Other5,omitempty"`

	Discount float64 `xml:"Discount,omitempty"`

	DiscountPercent float64 `xml:"DiscountPercent,omitempty"`

	Weight float64 `xml:"Weight,omitempty"`

	Sourcety int32 `xml:"Sourcety,omitempty"`

	Notes string `xml:"Notes,omitempty"`

	Usr string `xml:"Usr,omitempty"`

	FirstName string `xml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty"`

	Company string `xml:"Company,omitempty"`

	Address1 string `xml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty"`

	City string `xml:"City,omitempty"`

	State string `xml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty"`

	Email string `xml:"Email,omitempty"`

	Phone string `xml:"Phone,omitempty"`

	SuppressPackSlipPrice bool `xml:"SuppressPackSlipPrice,omitempty"`

	ShipMethodID int32 `xml:"ShipMethodID,omitempty"`

	AutoOrderID int32 `xml:"AutoOrderID,omitempty"`

	CreatedBy string `xml:"CreatedBy,omitempty"`

	ReturnOrderID int32 `xml:"ReturnOrderID,omitempty"`

	OrderRankID int32 `xml:"OrderRankID,omitempty"`

	OrderPayRankID int32 `xml:"OrderPayRankID,omitempty"`

	AddressIsVerified bool `xml:"AddressIsVerified,omitempty"`

	County string `xml:"County,omitempty"`

	IsRMA bool `xml:"IsRMA,omitempty"`

	BackOrderFromID int32 `xml:"BackOrderFromID,omitempty"`

	CreditsEarned float64 `xml:"CreditsEarned,omitempty"`

	TotalFedTax float64 `xml:"TotalFedTax,omitempty"`

	TotalStateTax float64 `xml:"TotalStateTax,omitempty"`

	ManualShippingTax float64 `xml:"ManualShippingTax,omitempty"`

	ReplacementOrderID int32 `xml:"ReplacementOrderID,omitempty"`

	LockedDate time.Time `xml:"LockedDate,omitempty"`

	CommissionedDate time.Time `xml:"CommissionedDate,omitempty"`

	Flag1 bool `xml:"Flag1,omitempty"`

	Flag2 bool `xml:"Flag2,omitempty"`

	Flag3 bool `xml:"Flag3,omitempty"`

	Other6 float64 `xml:"Other6,omitempty"`

	Other7 float64 `xml:"Other7,omitempty"`

	Other8 float64 `xml:"Other8,omitempty"`

	Other9 float64 `xml:"Other9,omitempty"`

	Other10 float64 `xml:"Other10,omitempty"`

	OriginalWarehouseID int32 `xml:"OriginalWarehouseID,omitempty"`

	PickupName string `xml:"PickupName,omitempty"`

	TransferToID int32 `xml:"TransferToID,omitempty"`

	IsCommissionable bool `xml:"IsCommissionable,omitempty"`

	FulfilledBy string `xml:"FulfilledBy,omitempty"`

	CreditApplied float64 `xml:"CreditApplied,omitempty"`

	ShippedDate time.Time `xml:"ShippedDate,omitempty"`

	TaxLockDate time.Time `xml:"TaxLockDate,omitempty"`

	TotalTaxable float64 `xml:"TotalTaxable,omitempty"`

	ReturnCategoryID int32 `xml:"ReturnCategoryID,omitempty"`

	ReplacementCategoryID int32 `xml:"ReplacementCategoryID,omitempty"`

	CalculatedShipping float64 `xml:"CalculatedShipping,omitempty"`

	HandlingFee float64 `xml:"HandlingFee,omitempty"`

	OrderProcessTy int32 `xml:"OrderProcessTy,omitempty"`

	ActualCarrier int32 `xml:"ActualCarrier,omitempty"`

	ParentOrderID int32 `xml:"ParentOrderID,omitempty"`

	CustomerTy int32 `xml:"CustomerTy,omitempty"`

	Reference string `xml:"Reference,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty"`

	NameSuffix string `xml:"NameSuffix,omitempty"`

	Address3 string `xml:"Address3,omitempty"`

	PartyID int32 `xml:"PartyID,omitempty"`

	TrackingNumber1 string `xml:"TrackingNumber1,omitempty"`

	TrackingNumber2 string `xml:"TrackingNumber2,omitempty"`

	TrackingNumber3 string `xml:"TrackingNumber3,omitempty"`

	TrackingNumber4 string `xml:"TrackingNumber4,omitempty"`

	TrackingNumber5 string `xml:"TrackingNumber5,omitempty"`

	WebCarrierID *OrderShipCarrier `xml:"WebCarrierID,omitempty"`

	WebCarrierID2 *OrderShipCarrier `xml:"WebCarrierID2,omitempty"`

	WebCarrierID3 *OrderShipCarrier `xml:"WebCarrierID3,omitempty"`

	WebCarrierID4 *OrderShipCarrier `xml:"WebCarrierID4,omitempty"`

	WebCarrierID5 *OrderShipCarrier `xml:"WebCarrierID5,omitempty"`

	Other11 string `xml:"Other11,omitempty"`

	Other12 string `xml:"Other12,omitempty"`

	Other13 string `xml:"Other13,omitempty"`

	Other14 string `xml:"Other14,omitempty"`

	Other15 string `xml:"Other15,omitempty"`

	Other16 string `xml:"Other16,omitempty"`

	Other17 string `xml:"Other17,omitempty"`

	Other18 string `xml:"Other18,omitempty"`

	Other19 string `xml:"Other19,omitempty"`

	Other20 string `xml:"Other20,omitempty"`
}

type UpdateCustomerRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateCustomerRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	FirstName string `xml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty"`

	Company string `xml:"Company,omitempty"`

	CustomerType int32 `xml:"CustomerType,omitempty"`

	CustomerStatus int32 `xml:"CustomerStatus,omitempty"`

	Email string `xml:"Email,omitempty"`

	Phone string `xml:"Phone,omitempty"`

	Phone2 string `xml:"Phone2,omitempty"`

	MobilePhone string `xml:"MobilePhone,omitempty"`

	Fax string `xml:"Fax,omitempty"`

	MainAddress1 string `xml:"MainAddress1,omitempty"`

	MainAddress2 string `xml:"MainAddress2,omitempty"`

	MainAddress3 string `xml:"MainAddress3,omitempty"`

	MainCity string `xml:"MainCity,omitempty"`

	MainState string `xml:"MainState,omitempty"`

	MainZip string `xml:"MainZip,omitempty"`

	MainCountry string `xml:"MainCountry,omitempty"`

	MainCounty string `xml:"MainCounty,omitempty"`

	MailAddress1 string `xml:"MailAddress1,omitempty"`

	MailAddress2 string `xml:"MailAddress2,omitempty"`

	MailAddress3 string `xml:"MailAddress3,omitempty"`

	MailCity string `xml:"MailCity,omitempty"`

	MailState string `xml:"MailState,omitempty"`

	MailZip string `xml:"MailZip,omitempty"`

	MailCountry string `xml:"MailCountry,omitempty"`

	MailCounty string `xml:"MailCounty,omitempty"`

	OtherAddress1 string `xml:"OtherAddress1,omitempty"`

	OtherAddress2 string `xml:"OtherAddress2,omitempty"`

	OtherAddress3 string `xml:"OtherAddress3,omitempty"`

	OtherCity string `xml:"OtherCity,omitempty"`

	OtherState string `xml:"OtherState,omitempty"`

	OtherZip string `xml:"OtherZip,omitempty"`

	OtherCountry string `xml:"OtherCountry,omitempty"`

	OtherCounty string `xml:"OtherCounty,omitempty"`

	CanLogin bool `xml:"CanLogin,omitempty"`

	LoginName string `xml:"LoginName,omitempty"`

	LoginPassword string `xml:"LoginPassword,omitempty"`

	TaxID string `xml:"TaxID,omitempty"`

	SalesTaxID string `xml:"SalesTaxID,omitempty"`

	SalesTaxExemptExpireDate time.Time `xml:"SalesTaxExemptExpireDate,omitempty"`

	IsSalesTaxExempt bool `xml:"IsSalesTaxExempt,omitempty"`

	BirthDate time.Time `xml:"BirthDate,omitempty"`

	Field1 string `xml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty"`

	Field11 string `xml:"Field11,omitempty"`

	Field12 string `xml:"Field12,omitempty"`

	Field13 string `xml:"Field13,omitempty"`

	Field14 string `xml:"Field14,omitempty"`

	Field15 string `xml:"Field15,omitempty"`

	SubscribeToBroadcasts bool `xml:"SubscribeToBroadcasts,omitempty"`

	SubscribeFromIPAddress string `xml:"SubscribeFromIPAddress,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	PayableToName string `xml:"PayableToName,omitempty"`

	PayableType *PayableType `xml:"PayableType,omitempty"`

	DefaultWarehouseID int32 `xml:"DefaultWarehouseID,omitempty"`

	CheckThreshold float64 `xml:"CheckThreshold,omitempty"`

	CreatedDate time.Time `xml:"CreatedDate,omitempty"`

	TaxIDType *TaxIDType `xml:"TaxIDType,omitempty"`

	LanguageID int32 `xml:"LanguageID,omitempty"`

	Gender *Gender `xml:"Gender,omitempty"`

	VatRegistration string `xml:"VatRegistration,omitempty"`

	Date1 time.Time `xml:"Date1,omitempty"`

	Date2 time.Time `xml:"Date2,omitempty"`

	Date3 time.Time `xml:"Date3,omitempty"`

	Date4 time.Time `xml:"Date4,omitempty"`

	Date5 time.Time `xml:"Date5,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty"`

	NameSuffix string `xml:"NameSuffix,omitempty"`

	BinaryPlacementPreference int32 `xml:"BinaryPlacementPreference,omitempty"`

	UseBinaryHoldingTank bool `xml:"UseBinaryHoldingTank,omitempty"`

	MainAddressVerified bool `xml:"MainAddressVerified,omitempty"`

	MailAddressVerified bool `xml:"MailAddressVerified,omitempty"`

	OtherAddressVerified bool `xml:"OtherAddressVerified,omitempty"`
}

type BaseCalculateOrderRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ BaseCalculateOrderRequest"`

	*ApiRequest
}

type CalculateOrderRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CalculateOrderRequest"`

	*BaseCalculateOrderRequest

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	WarehouseID int32 `xml:"WarehouseID,omitempty"`

	ShipMethodID int32 `xml:"ShipMethodID,omitempty"`

	PriceType int32 `xml:"PriceType,omitempty"`

	Address1 string `xml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty"`

	Address3 string `xml:"Address3,omitempty"`

	City string `xml:"City,omitempty"`

	State string `xml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty"`

	County string `xml:"County,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	OrderType *OrderType `xml:"OrderType,omitempty"`

	TaxRateOverride float64 `xml:"TaxRateOverride,omitempty"`

	ShippingAmountOverride float64 `xml:"ShippingAmountOverride,omitempty"`

	ReturnOrderID int32 `xml:"ReturnOrderID,omitempty"`

	Other11 string `xml:"Other11,omitempty"`

	Other12 string `xml:"Other12,omitempty"`

	Other13 string `xml:"Other13,omitempty"`

	Other14 string `xml:"Other14,omitempty"`

	Other15 string `xml:"Other15,omitempty"`

	Other16 string `xml:"Other16,omitempty"`

	Other17 string `xml:"Other17,omitempty"`

	Other18 string `xml:"Other18,omitempty"`

	Other19 string `xml:"Other19,omitempty"`

	Other20 string `xml:"Other20,omitempty"`

	Details *ArrayOfOrderDetailRequest `xml:"Details,omitempty"`

	ReturnShipMethods bool `xml:"ReturnShipMethods,omitempty"`

	PartyID int32 `xml:"PartyID,omitempty"`
}

type ArrayOfOrderDetailRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfOrderDetailRequest"`

	OrderDetailRequest []*OrderDetailRequest `xml:"OrderDetailRequest,omitempty"`
}

type OrderDetailRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ OrderDetailRequest"`

	ItemCode string `xml:"ItemCode,omitempty"`

	Quantity float64 `xml:"Quantity,omitempty"`

	ParentItemCode string `xml:"ParentItemCode,omitempty"`

	PriceEachOverride float64 `xml:"PriceEachOverride,omitempty"`

	TaxableEachOverride float64 `xml:"TaxableEachOverride,omitempty"`

	ShippingPriceEachOverride float64 `xml:"ShippingPriceEachOverride,omitempty"`

	BusinessVolumeEachOverride float64 `xml:"BusinessVolumeEachOverride,omitempty"`

	CommissionableVolumeEachOverride float64 `xml:"CommissionableVolumeEachOverride,omitempty"`

	Other1EachOverride float64 `xml:"Other1EachOverride,omitempty"`

	Other2EachOverride float64 `xml:"Other2EachOverride,omitempty"`

	Other3EachOverride float64 `xml:"Other3EachOverride,omitempty"`

	Other4EachOverride float64 `xml:"Other4EachOverride,omitempty"`

	Other5EachOverride float64 `xml:"Other5EachOverride,omitempty"`

	Other6EachOverride float64 `xml:"Other6EachOverride,omitempty"`

	Other7EachOverride float64 `xml:"Other7EachOverride,omitempty"`

	Other8EachOverride float64 `xml:"Other8EachOverride,omitempty"`

	Other9EachOverride float64 `xml:"Other9EachOverride,omitempty"`

	Other10EachOverride float64 `xml:"Other10EachOverride,omitempty"`

	DescriptionOverride string `xml:"DescriptionOverride,omitempty"`

	Reference1 string `xml:"Reference1,omitempty"`

	AdvancedAutoOptions *AdvancedAutoOptionsRequest `xml:"AdvancedAutoOptions,omitempty"`
}

type AdvancedAutoOptionsRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ AdvancedAutoOptionsRequest"`

	ProcessWhileDate time.Time `xml:"ProcessWhileDate,omitempty"`

	SkipUntilDate time.Time `xml:"SkipUntilDate,omitempty"`
}

type CreateAutoOrderRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateAutoOrderRequest"`

	*BaseCalculateOrderRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	Frequency *FrequencyType `xml:"Frequency,omitempty"`

	StartDate time.Time `xml:"StartDate,omitempty"`

	StopDate time.Time `xml:"StopDate,omitempty"`

	SpecificDayInterval int32 `xml:"SpecificDayInterval,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	WarehouseID int32 `xml:"WarehouseID,omitempty"`

	ShipMethodID int32 `xml:"ShipMethodID,omitempty"`

	PriceType int32 `xml:"PriceType,omitempty"`

	PaymentType *AutoOrderPaymentType `xml:"PaymentType,omitempty"`

	ProcessType *AutoOrderProcessType `xml:"ProcessType,omitempty"`

	FirstName string `xml:"FirstName,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty"`

	LastName string `xml:"LastName,omitempty"`

	NameSuffix string `xml:"NameSuffix,omitempty"`

	Company string `xml:"Company,omitempty"`

	Address1 string `xml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty"`

	Address3 string `xml:"Address3,omitempty"`

	City string `xml:"City,omitempty"`

	State string `xml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty"`

	County string `xml:"County,omitempty"`

	Email string `xml:"Email,omitempty"`

	Phone string `xml:"Phone,omitempty"`

	Notes string `xml:"Notes,omitempty"`

	Other11 string `xml:"Other11,omitempty"`

	Other12 string `xml:"Other12,omitempty"`

	Other13 string `xml:"Other13,omitempty"`

	Other14 string `xml:"Other14,omitempty"`

	Other15 string `xml:"Other15,omitempty"`

	Other16 string `xml:"Other16,omitempty"`

	Other17 string `xml:"Other17,omitempty"`

	Other18 string `xml:"Other18,omitempty"`

	Other19 string `xml:"Other19,omitempty"`

	Other20 string `xml:"Other20,omitempty"`

	Description string `xml:"Description,omitempty"`

	OverwriteExistingAutoOrder bool `xml:"OverwriteExistingAutoOrder,omitempty"`

	ExistingAutoOrderID int32 `xml:"ExistingAutoOrderID,omitempty"`

	Details *ArrayOfOrderDetailRequest `xml:"Details,omitempty"`
}

type CreateOrderRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateOrderRequest"`

	*BaseCalculateOrderRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	OrderStatus *OrderStatusType `xml:"OrderStatus,omitempty"`

	OrderDate time.Time `xml:"OrderDate,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	WarehouseID int32 `xml:"WarehouseID,omitempty"`

	ShipMethodID int32 `xml:"ShipMethodID,omitempty"`

	PriceType int32 `xml:"PriceType,omitempty"`

	FirstName string `xml:"FirstName,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty"`

	LastName string `xml:"LastName,omitempty"`

	NameSuffix string `xml:"NameSuffix,omitempty"`

	Company string `xml:"Company,omitempty"`

	Address1 string `xml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty"`

	Address3 string `xml:"Address3,omitempty"`

	City string `xml:"City,omitempty"`

	State string `xml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty"`

	County string `xml:"County,omitempty"`

	Email string `xml:"Email,omitempty"`

	Phone string `xml:"Phone,omitempty"`

	Notes string `xml:"Notes,omitempty"`

	Other11 string `xml:"Other11,omitempty"`

	Other12 string `xml:"Other12,omitempty"`

	Other13 string `xml:"Other13,omitempty"`

	Other14 string `xml:"Other14,omitempty"`

	Other15 string `xml:"Other15,omitempty"`

	Other16 string `xml:"Other16,omitempty"`

	Other17 string `xml:"Other17,omitempty"`

	Other18 string `xml:"Other18,omitempty"`

	Other19 string `xml:"Other19,omitempty"`

	Other20 string `xml:"Other20,omitempty"`

	OrderType *OrderType `xml:"OrderType,omitempty"`

	TaxRateOverride float64 `xml:"TaxRateOverride,omitempty"`

	ShippingAmountOverride float64 `xml:"ShippingAmountOverride,omitempty"`

	UseManualOrderID bool `xml:"UseManualOrderID,omitempty"`

	ManualOrderID int32 `xml:"ManualOrderID,omitempty"`

	TransferVolumeToID int32 `xml:"TransferVolumeToID,omitempty"`

	ReturnOrderID int32 `xml:"ReturnOrderID,omitempty"`

	OverwriteExistingOrder bool `xml:"OverwriteExistingOrder,omitempty"`

	ExistingOrderID int32 `xml:"ExistingOrderID,omitempty"`

	PartyID int32 `xml:"PartyID,omitempty"`

	Details *ArrayOfOrderDetailRequest `xml:"Details,omitempty"`

	SuppressPackSlipPrice bool `xml:"SuppressPackSlipPrice,omitempty"`
}

type SetAccountDirectDepositRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetAccountDirectDepositRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	NameOnAccount string `xml:"NameOnAccount,omitempty"`

	BankAccountNumber string `xml:"BankAccountNumber,omitempty"`

	BankRoutingNumber string `xml:"BankRoutingNumber,omitempty"`

	DepositAccountType *DepositAccountType `xml:"DepositAccountType,omitempty"`

	BankName string `xml:"BankName,omitempty"`

	BankAddress string `xml:"BankAddress,omitempty"`

	BankCity string `xml:"BankCity,omitempty"`

	BankState string `xml:"BankState,omitempty"`

	BankZip string `xml:"BankZip,omitempty"`

	BankCountry string `xml:"BankCountry,omitempty"`

	Iban string `xml:"Iban,omitempty"`

	SwiftCode string `xml:"SwiftCode,omitempty"`

	BankAccountType *BankAccountType `xml:"BankAccountType,omitempty"`
}

type SetAccountWalletRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetAccountWalletRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	WalletAccountType *AccountWalletType `xml:"WalletAccountType,omitempty"`

	WalletType int32 `xml:"WalletType,omitempty"`

	WalletAccount string `xml:"WalletAccount,omitempty"`
}

type SetAccountCreditCardTokenRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetAccountCreditCardTokenRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	CreditCardAccountType *AccountCreditCardType `xml:"CreditCardAccountType,omitempty"`

	CreditCardToken string `xml:"CreditCardToken,omitempty"`

	ExpirationMonth int32 `xml:"ExpirationMonth,omitempty"`

	ExpirationYear int32 `xml:"ExpirationYear,omitempty"`

	CreditCardType int32 `xml:"CreditCardType,omitempty"`

	BillingName string `xml:"BillingName,omitempty"`

	UseMainAddress bool `xml:"UseMainAddress,omitempty"`

	BillingAddress string `xml:"BillingAddress,omitempty"`

	BillingAddress2 string `xml:"BillingAddress2,omitempty"`

	BillingCity string `xml:"BillingCity,omitempty"`

	BillingState string `xml:"BillingState,omitempty"`

	BillingZip string `xml:"BillingZip,omitempty"`

	BillingCountry string `xml:"BillingCountry,omitempty"`

	HideFromWeb bool `xml:"HideFromWeb,omitempty"`
}

type SetAccountCreditCardRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetAccountCreditCardRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	CreditCardAccountType *AccountCreditCardType `xml:"CreditCardAccountType,omitempty"`

	CreditCardNumber string `xml:"CreditCardNumber,omitempty"`

	ExpirationMonth int32 `xml:"ExpirationMonth,omitempty"`

	ExpirationYear int32 `xml:"ExpirationYear,omitempty"`

	CvcCode string `xml:"CvcCode,omitempty"`

	IssueCode string `xml:"IssueCode,omitempty"`

	CreditCardType int32 `xml:"CreditCardType,omitempty"`

	BillingName string `xml:"BillingName,omitempty"`

	UseMainAddress bool `xml:"UseMainAddress,omitempty"`

	BillingAddress string `xml:"BillingAddress,omitempty"`

	BillingCity string `xml:"BillingCity,omitempty"`

	BillingState string `xml:"BillingState,omitempty"`

	BillingZip string `xml:"BillingZip,omitempty"`

	BillingCountry string `xml:"BillingCountry,omitempty"`

	HideFromWeb bool `xml:"HideFromWeb,omitempty"`
}

type SetAccountCheckingRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetAccountCheckingRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	BankAccountNumber string `xml:"BankAccountNumber,omitempty"`

	BankRoutingNumber string `xml:"BankRoutingNumber,omitempty"`

	BankName string `xml:"BankName,omitempty"`

	BankAccountType *BankAccountType `xml:"BankAccountType,omitempty"`

	NameOnAccount string `xml:"NameOnAccount,omitempty"`

	UseMainAddress bool `xml:"UseMainAddress,omitempty"`

	BillingAddress string `xml:"BillingAddress,omitempty"`

	BillingCity string `xml:"BillingCity,omitempty"`

	BillingState string `xml:"BillingState,omitempty"`

	BillingZip string `xml:"BillingZip,omitempty"`

	BillingCountry string `xml:"BillingCountry,omitempty"`

	DriversLicenseNumber string `xml:"DriversLicenseNumber,omitempty"`
}

type CreateOrderImportRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateOrderImportRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	FirstName string `xml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty"`

	Company string `xml:"Company,omitempty"`

	Phone string `xml:"Phone,omitempty"`

	Email string `xml:"Email,omitempty"`

	ShipMethodID int32 `xml:"ShipMethodID,omitempty"`

	Address1 string `xml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty"`

	City string `xml:"City,omitempty"`

	State string `xml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty"`

	County string `xml:"County,omitempty"`

	Notes string `xml:"Notes,omitempty"`

	WarehouseID int32 `xml:"WarehouseID,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	ShippingStateTax float64 `xml:"ShippingStateTax,omitempty"`

	ShippingFedTax float64 `xml:"ShippingFedTax,omitempty"`

	ShippingCountyLocalTax float64 `xml:"ShippingCountyLocalTax,omitempty"`

	ShippingCountyTax float64 `xml:"ShippingCountyTax,omitempty"`

	ShippingCityLocalTax float64 `xml:"ShippingCityLocalTax,omitempty"`

	ShippingCityTax float64 `xml:"ShippingCityTax,omitempty"`

	Shipping float64 `xml:"Shipping,omitempty"`

	PriceType int32 `xml:"PriceType,omitempty"`

	OrderStatus *OrderStatusType `xml:"OrderStatus,omitempty"`

	OrderDate time.Time `xml:"OrderDate,omitempty"`

	Other11 string `xml:"Other11,omitempty"`

	Other12 string `xml:"Other12,omitempty"`

	Other13 string `xml:"Other13,omitempty"`

	Other14 string `xml:"Other14,omitempty"`

	Other15 string `xml:"Other15,omitempty"`

	Other16 string `xml:"Other16,omitempty"`

	Other17 string `xml:"Other17,omitempty"`

	Other18 string `xml:"Other18,omitempty"`

	Other19 string `xml:"Other19,omitempty"`

	Other20 string `xml:"Other20,omitempty"`

	OrderType *OrderType `xml:"OrderType,omitempty"`

	UseManualOrderID bool `xml:"UseManualOrderID,omitempty"`

	ManualOrderID int32 `xml:"ManualOrderID,omitempty"`

	ReturnOrderID int32 `xml:"ReturnOrderID,omitempty"`

	OrderDetails *ArrayOfOrderImportDetail `xml:"OrderDetails,omitempty"`

	PartyID int32 `xml:"PartyID,omitempty"`
}

type ArrayOfOrderImportDetail struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfOrderImportDetail"`

	OrderImportDetail []*OrderImportDetail `xml:"OrderImportDetail,omitempty"`
}

type OrderImportDetail struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ OrderImportDetail"`

	ParentItemCode string `xml:"ParentItemCode,omitempty"`

	ItemCode string `xml:"ItemCode,omitempty"`

	Description string `xml:"Description,omitempty"`

	Qty float64 `xml:"Qty,omitempty"`

	WeightEach float64 `xml:"WeightEach,omitempty"`

	CountyLocalTax float64 `xml:"CountyLocalTax,omitempty"`

	CountyTax float64 `xml:"CountyTax,omitempty"`

	CityTax float64 `xml:"CityTax,omitempty"`

	StateTax float64 `xml:"StateTax,omitempty"`

	FedTax float64 `xml:"FedTax,omitempty"`

	TaxablePriceEach float64 `xml:"TaxablePriceEach,omitempty"`

	CVEach float64 `xml:"CVEach,omitempty"`

	BVEach float64 `xml:"BVEach,omitempty"`

	PriceEach float64 `xml:"PriceEach,omitempty"`

	Other10Each float64 `xml:"Other10Each,omitempty"`

	Other9Each float64 `xml:"Other9Each,omitempty"`

	Other8Each float64 `xml:"Other8Each,omitempty"`

	Other7Each float64 `xml:"Other7Each,omitempty"`

	Other6Each float64 `xml:"Other6Each,omitempty"`

	Other5Each float64 `xml:"Other5Each,omitempty"`

	Other4Each float64 `xml:"Other4Each,omitempty"`

	Other3Each float64 `xml:"Other3Each,omitempty"`

	Other2Each float64 `xml:"Other2Each,omitempty"`

	Other1Each float64 `xml:"Other1Each,omitempty"`
}

type CreateCustomerRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateCustomerRequest"`

	*ApiRequest

	FirstName string `xml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty"`

	Company string `xml:"Company,omitempty"`

	CustomerType int32 `xml:"CustomerType,omitempty"`

	CustomerStatus int32 `xml:"CustomerStatus,omitempty"`

	Email string `xml:"Email,omitempty"`

	Phone string `xml:"Phone,omitempty"`

	Phone2 string `xml:"Phone2,omitempty"`

	MobilePhone string `xml:"MobilePhone,omitempty"`

	Fax string `xml:"Fax,omitempty"`

	Notes string `xml:"Notes,omitempty"`

	MainAddress1 string `xml:"MainAddress1,omitempty"`

	MainAddress2 string `xml:"MainAddress2,omitempty"`

	MainAddress3 string `xml:"MainAddress3,omitempty"`

	MainCity string `xml:"MainCity,omitempty"`

	MainState string `xml:"MainState,omitempty"`

	MainZip string `xml:"MainZip,omitempty"`

	MainCountry string `xml:"MainCountry,omitempty"`

	MainCounty string `xml:"MainCounty,omitempty"`

	MailAddress1 string `xml:"MailAddress1,omitempty"`

	MailAddress2 string `xml:"MailAddress2,omitempty"`

	MailAddress3 string `xml:"MailAddress3,omitempty"`

	MailCity string `xml:"MailCity,omitempty"`

	MailState string `xml:"MailState,omitempty"`

	MailZip string `xml:"MailZip,omitempty"`

	MailCountry string `xml:"MailCountry,omitempty"`

	MailCounty string `xml:"MailCounty,omitempty"`

	OtherAddress1 string `xml:"OtherAddress1,omitempty"`

	OtherAddress2 string `xml:"OtherAddress2,omitempty"`

	OtherAddress3 string `xml:"OtherAddress3,omitempty"`

	OtherCity string `xml:"OtherCity,omitempty"`

	OtherState string `xml:"OtherState,omitempty"`

	OtherZip string `xml:"OtherZip,omitempty"`

	OtherCountry string `xml:"OtherCountry,omitempty"`

	OtherCounty string `xml:"OtherCounty,omitempty"`

	CanLogin bool `xml:"CanLogin,omitempty"`

	LoginName string `xml:"LoginName,omitempty"`

	LoginPassword string `xml:"LoginPassword,omitempty"`

	InsertEnrollerTree bool `xml:"InsertEnrollerTree,omitempty"`

	EnrollerID int32 `xml:"EnrollerID,omitempty"`

	InsertUnilevelTree bool `xml:"InsertUnilevelTree,omitempty"`

	SponsorID int32 `xml:"SponsorID,omitempty"`

	UseManualCustomerID bool `xml:"UseManualCustomerID,omitempty"`

	ManualCustomerID int32 `xml:"ManualCustomerID,omitempty"`

	TaxID string `xml:"TaxID,omitempty"`

	SalesTaxID string `xml:"SalesTaxID,omitempty"`

	SalesTaxExemptExpireDate time.Time `xml:"SalesTaxExemptExpireDate,omitempty"`

	IsSalesTaxExempt bool `xml:"IsSalesTaxExempt,omitempty"`

	BirthDate time.Time `xml:"BirthDate,omitempty"`

	Field1 string `xml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty"`

	Field11 string `xml:"Field11,omitempty"`

	Field12 string `xml:"Field12,omitempty"`

	Field13 string `xml:"Field13,omitempty"`

	Field14 string `xml:"Field14,omitempty"`

	Field15 string `xml:"Field15,omitempty"`

	SubscribeToBroadcasts bool `xml:"SubscribeToBroadcasts,omitempty"`

	SubscribeFromIPAddress string `xml:"SubscribeFromIPAddress,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	PayableToName string `xml:"PayableToName,omitempty"`

	EntryDate time.Time `xml:"EntryDate,omitempty"`

	DefaultWarehouseID int32 `xml:"DefaultWarehouseID,omitempty"`

	PayableType *PayableType `xml:"PayableType,omitempty"`

	CheckThreshold float64 `xml:"CheckThreshold,omitempty"`

	TaxIDType *TaxIDType `xml:"TaxIDType,omitempty"`

	LanguageID int32 `xml:"LanguageID,omitempty"`

	Gender *Gender `xml:"Gender,omitempty"`

	VatRegistration string `xml:"VatRegistration,omitempty"`

	Date1 time.Time `xml:"Date1,omitempty"`

	Date2 time.Time `xml:"Date2,omitempty"`

	Date3 time.Time `xml:"Date3,omitempty"`

	Date4 time.Time `xml:"Date4,omitempty"`

	Date5 time.Time `xml:"Date5,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty"`

	NameSuffix string `xml:"NameSuffix,omitempty"`

	BinaryPlacementPreference int32 `xml:"BinaryPlacementPreference,omitempty"`

	UseBinaryHoldingTank bool `xml:"UseBinaryHoldingTank,omitempty"`

	MainAddressVerified bool `xml:"MainAddressVerified,omitempty"`

	MailAddressVerified bool `xml:"MailAddressVerified,omitempty"`

	OtherAddressVerified bool `xml:"OtherAddressVerified,omitempty"`
}

type TransactionalRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ TransactionalRequest"`

	*ApiRequest

	TransactionRequests *ArrayOfApiRequest `xml:"TransactionRequests,omitempty"`
}

type ArrayOfApiRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfApiRequest"`

	ApiRequest []*ApiRequest `xml:"ApiRequest,omitempty"`
}

type BaseCreateExpectedPaymentRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ BaseCreateExpectedPaymentRequest"`

	*ApiRequest
}

type CreateExpectedBankWireRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateExpectedBankWireRequest"`

	*BaseCreateExpectedPaymentRequest

	OrderID int32 `xml:"OrderID,omitempty"`

	PaymentDate time.Time `xml:"PaymentDate,omitempty"`

	Amount float64 `xml:"Amount,omitempty"`

	BankName string `xml:"BankName,omitempty"`

	NameOnAccount string `xml:"NameOnAccount,omitempty"`

	Memo string `xml:"Memo,omitempty"`

	AuthorizationCode string `xml:"AuthorizationCode,omitempty"`
}

type CreateExpectedCODRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateExpectedCODRequest"`

	*BaseCreateExpectedPaymentRequest

	OrderID int32 `xml:"OrderID,omitempty"`

	PaymentDate time.Time `xml:"PaymentDate,omitempty"`

	Amount float64 `xml:"Amount,omitempty"`
}

type CreateExpectedPaymentRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateExpectedPaymentRequest"`

	*BaseCreateExpectedPaymentRequest

	OrderID int32 `xml:"OrderID,omitempty"`

	PaymentDate time.Time `xml:"PaymentDate,omitempty"`

	Amount float64 `xml:"Amount,omitempty"`

	PaymentType *PaymentType `xml:"PaymentType,omitempty"`

	AuthorizationCode string `xml:"AuthorizationCode,omitempty"`
}

type BaseCreatePaymentRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ BaseCreatePaymentRequest"`

	*ApiRequest
}

type ChargePriorAuthorizationRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ChargePriorAuthorizationRequest"`

	*BaseCreatePaymentRequest

	MerchantTransactionKey string `xml:"MerchantTransactionKey,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty"`

	MaxAmount float64 `xml:"MaxAmount,omitempty"`
}

type BaseChargeWalletAccountRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ BaseChargeWalletAccountRequest"`

	*BaseCreatePaymentRequest
}

type ChargeWalletAccountRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ChargeWalletAccountRequest"`

	*BaseChargeWalletAccountRequest

	WalletAccountNumber string `xml:"WalletAccountNumber,omitempty"`

	WalletTy int32 `xml:"WalletTy,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty"`

	Memo string `xml:"Memo,omitempty"`

	MaxAmount float64 `xml:"MaxAmount,omitempty"`
}

type ChargeWalletAccountOnFileRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ChargeWalletAccountOnFileRequest"`

	*BaseCreatePaymentRequest

	WalletAccountType *AccountWalletType `xml:"WalletAccountType,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty"`

	Memo string `xml:"Memo,omitempty"`

	MaxAmount float64 `xml:"MaxAmount,omitempty"`
}

type BaseDebitBankAccountRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ BaseDebitBankAccountRequest"`

	*BaseCreatePaymentRequest
}

type DebitBankAccountRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DebitBankAccountRequest"`

	*BaseDebitBankAccountRequest

	BankAccountNumber string `xml:"BankAccountNumber,omitempty"`

	BankRoutingNumber string `xml:"BankRoutingNumber,omitempty"`

	BankName string `xml:"BankName,omitempty"`

	BankAccountType *BankAccountType `xml:"BankAccountType,omitempty"`

	CheckNumber string `xml:"CheckNumber,omitempty"`

	NameOnAccount string `xml:"NameOnAccount,omitempty"`

	BillingAddress string `xml:"BillingAddress,omitempty"`

	BillingAddress2 string `xml:"BillingAddress2,omitempty"`

	BillingCity string `xml:"BillingCity,omitempty"`

	BillingState string `xml:"BillingState,omitempty"`

	BillingZip string `xml:"BillingZip,omitempty"`

	BillingCountry string `xml:"BillingCountry,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty"`

	MaxAmount float64 `xml:"MaxAmount,omitempty"`
}

type DebitBankAccountOnFileRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DebitBankAccountOnFileRequest"`

	*BaseDebitBankAccountRequest

	OrderID int32 `xml:"OrderID,omitempty"`

	MaxAmount float64 `xml:"MaxAmount,omitempty"`
}

type RefundPriorWalletChargeRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ RefundPriorWalletChargeRequest"`

	*BaseCreatePaymentRequest

	ReturnPaymentID int32 `xml:"ReturnPaymentID,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty"`

	MaxAmount float64 `xml:"MaxAmount,omitempty"`
}

type CreatePaymentRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreatePaymentRequest"`

	*BaseCreatePaymentRequest

	OrderID int32 `xml:"OrderID,omitempty"`

	PaymentDate time.Time `xml:"PaymentDate,omitempty"`

	Amount float64 `xml:"Amount,omitempty"`

	PaymentType *PaymentType `xml:"PaymentType,omitempty"`

	AuthorizationCode string `xml:"AuthorizationCode,omitempty"`

	Memo string `xml:"Memo,omitempty"`

	BillingName string `xml:"BillingName,omitempty"`

	BillingAddress string `xml:"BillingAddress,omitempty"`

	BillingAddress2 string `xml:"BillingAddress2,omitempty"`

	BillingCity string `xml:"BillingCity,omitempty"`

	BillingState string `xml:"BillingState,omitempty"`

	BillingZip string `xml:"BillingZip,omitempty"`

	BillingCountry string `xml:"BillingCountry,omitempty"`
}

type BaseCreatePaymentCreditCardRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ BaseCreatePaymentCreditCardRequest"`

	*BaseCreatePaymentRequest
}

type RefundPriorCreditCardChargeRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ RefundPriorCreditCardChargeRequest"`

	*BaseCreatePaymentCreditCardRequest

	ReturnPaymentID int32 `xml:"ReturnPaymentID,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty"`

	MaxAmount float64 `xml:"MaxAmount,omitempty"`
}

type CreatePaymentCreditCardRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreatePaymentCreditCardRequest"`

	*BaseCreatePaymentCreditCardRequest

	OrderID int32 `xml:"OrderID,omitempty"`

	PaymentDate time.Time `xml:"PaymentDate,omitempty"`

	Amount float64 `xml:"Amount,omitempty"`

	CreditCardNumber string `xml:"CreditCardNumber,omitempty"`

	ExpirationMonth int32 `xml:"ExpirationMonth,omitempty"`

	ExpirationYear int32 `xml:"ExpirationYear,omitempty"`

	BillingName string `xml:"BillingName,omitempty"`

	BillingAddress string `xml:"BillingAddress,omitempty"`

	BillingAddress2 string `xml:"BillingAddress2,omitempty"`

	BillingCity string `xml:"BillingCity,omitempty"`

	BillingState string `xml:"BillingState,omitempty"`

	BillingZip string `xml:"BillingZip,omitempty"`

	BillingCountry string `xml:"BillingCountry,omitempty"`

	CreditCardType int32 `xml:"CreditCardType,omitempty"`

	AuthorizationCode string `xml:"AuthorizationCode,omitempty"`

	Memo string `xml:"Memo,omitempty"`
}

type BaseChargeCreditCardRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ BaseChargeCreditCardRequest"`

	*BaseCreatePaymentCreditCardRequest
}

type ChargeCreditCardTokenOnFileRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ChargeCreditCardTokenOnFileRequest"`

	*BaseChargeCreditCardRequest

	CreditCardAccountType *AccountCreditCardType `xml:"CreditCardAccountType,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty"`

	CvcCode string `xml:"CvcCode,omitempty"`

	MaxAmount float64 `xml:"MaxAmount,omitempty"`

	MerchantWarehouseIDOverride int32 `xml:"MerchantWarehouseIDOverride,omitempty"`

	ClientIPAddress string `xml:"ClientIPAddress,omitempty"`

	OtherData1 string `xml:"OtherData1,omitempty"`

	OtherData2 string `xml:"OtherData2,omitempty"`

	OtherData3 string `xml:"OtherData3,omitempty"`

	OtherData4 string `xml:"OtherData4,omitempty"`

	OtherData5 string `xml:"OtherData5,omitempty"`

	OtherData6 string `xml:"OtherData6,omitempty"`

	OtherData7 string `xml:"OtherData7,omitempty"`

	OtherData8 string `xml:"OtherData8,omitempty"`

	OtherData9 string `xml:"OtherData9,omitempty"`

	OtherData10 string `xml:"OtherData10,omitempty"`
}

type ChargeCreditCardTokenRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ChargeCreditCardTokenRequest"`

	*BaseChargeCreditCardRequest

	CreditCardToken string `xml:"CreditCardToken,omitempty"`

	BillingName string `xml:"BillingName,omitempty"`

	BillingAddress string `xml:"BillingAddress,omitempty"`

	BillingAddress2 string `xml:"BillingAddress2,omitempty"`

	BillingCity string `xml:"BillingCity,omitempty"`

	BillingState string `xml:"BillingState,omitempty"`

	BillingZip string `xml:"BillingZip,omitempty"`

	BillingCountry string `xml:"BillingCountry,omitempty"`

	CvcCode string `xml:"CvcCode,omitempty"`

	IssueNumber string `xml:"IssueNumber,omitempty"`

	CreditCardType int32 `xml:"CreditCardType,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty"`

	MaxAmount float64 `xml:"MaxAmount,omitempty"`

	MerchantWarehouseIDOverride int32 `xml:"MerchantWarehouseIDOverride,omitempty"`

	ClientIPAddress string `xml:"ClientIPAddress,omitempty"`

	OtherData1 string `xml:"OtherData1,omitempty"`

	OtherData2 string `xml:"OtherData2,omitempty"`

	OtherData3 string `xml:"OtherData3,omitempty"`

	OtherData4 string `xml:"OtherData4,omitempty"`

	OtherData5 string `xml:"OtherData5,omitempty"`

	OtherData6 string `xml:"OtherData6,omitempty"`

	OtherData7 string `xml:"OtherData7,omitempty"`

	OtherData8 string `xml:"OtherData8,omitempty"`

	OtherData9 string `xml:"OtherData9,omitempty"`

	OtherData10 string `xml:"OtherData10,omitempty"`
}

type ChargeCreditCardOnFileRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ChargeCreditCardOnFileRequest"`

	*BaseChargeCreditCardRequest

	CreditCardAccountType *AccountCreditCardType `xml:"CreditCardAccountType,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty"`

	MaxAmount float64 `xml:"MaxAmount,omitempty"`
}

type ChargeCreditCardRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ChargeCreditCardRequest"`

	*BaseChargeCreditCardRequest

	CreditCardNumber string `xml:"CreditCardNumber,omitempty"`

	ExpirationMonth int32 `xml:"ExpirationMonth,omitempty"`

	ExpirationYear int32 `xml:"ExpirationYear,omitempty"`

	BillingName string `xml:"BillingName,omitempty"`

	BillingAddress string `xml:"BillingAddress,omitempty"`

	BillingAddress2 string `xml:"BillingAddress2,omitempty"`

	BillingCity string `xml:"BillingCity,omitempty"`

	BillingState string `xml:"BillingState,omitempty"`

	BillingZip string `xml:"BillingZip,omitempty"`

	BillingCountry string `xml:"BillingCountry,omitempty"`

	CvcCode string `xml:"CvcCode,omitempty"`

	IssueNumber string `xml:"IssueNumber,omitempty"`

	CreditCardType int32 `xml:"CreditCardType,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty"`

	MaxAmount float64 `xml:"MaxAmount,omitempty"`
}

type CreatePaymentWalletRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreatePaymentWalletRequest"`

	*BaseCreatePaymentRequest

	OrderID int32 `xml:"OrderID,omitempty"`

	PaymentDate time.Time `xml:"PaymentDate,omitempty"`

	Amount float64 `xml:"Amount,omitempty"`

	WalletType int32 `xml:"WalletType,omitempty"`

	WalletAccount string `xml:"WalletAccount,omitempty"`

	AuthorizationCode string `xml:"AuthorizationCode,omitempty"`

	Memo string `xml:"Memo,omitempty"`

	BillingName string `xml:"BillingName,omitempty"`
}

type DeleteCustomerLeadRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteCustomerLeadRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	CustomerLeadID int32 `xml:"CustomerLeadID,omitempty"`
}

type ArrayOfEmailAttachment struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfEmailAttachment"`

	EmailAttachment []*EmailAttachment `xml:"EmailAttachment,omitempty"`
}

type EmailAttachment struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ EmailAttachment"`

	BinaryData []byte `xml:"BinaryData,omitempty"`

	FileName string `xml:"FileName,omitempty"`

	ContentLength int32 `xml:"ContentLength,omitempty"`
}

type ArrayOfForwardedAttachment struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfForwardedAttachment"`

	ForwardedAttachment []*ForwardedAttachment `xml:"ForwardedAttachment,omitempty"`
}

type ForwardedAttachment struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ForwardedAttachment"`

	MailID int32 `xml:"MailID,omitempty"`

	AttachmentID int32 `xml:"AttachmentID,omitempty"`
}

type CreateEmailResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateEmailResponse"`

	*ApiResponse
}

type ApiResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ApiResponse"`

	Result *ApiResult //`xml:"Result,omitempty"`
}

type ApiResult struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ApiResult"`

	Status *ResultStatus `xml:"Status,omitempty"`

	Errors *ArrayOfString `xml:"Errors,omitempty"`

	TransactionKey string `xml:"TransactionKey,omitempty"`
}

type ArrayOfString struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfString"`

	String []string `xml:"string,omitempty"`
}

type UpdateOrderDetailResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateOrderDetailResponse"`

	*ApiResponse
}

type CreateOrderDetailResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateOrderDetailResponse"`

	*ApiResponse

	OrderID int32 `xml:"OrderID,omitempty"`

	OrderLine int32 `xml:"OrderLine,omitempty"`
}

type SendEmailResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SendEmailResponse"`

	*ApiResponse
}

type SetItemKitMembersResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetItemKitMembersResponse"`

	*ApiResponse
}

type GetFileContentsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetFileContentsResponse"`

	*ApiResponse

	File []byte `xml:"File,omitempty"`
}

type GetFilesResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetFilesResponse"`

	*ApiResponse

	CustomerFileList *ArrayOfCustomerFilesResponse `xml:"CustomerFileList,omitempty"`
}

type ArrayOfCustomerFilesResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfCustomerFilesResponse"`

	CustomerFilesResponse []*CustomerFilesResponse `xml:"CustomerFilesResponse,omitempty"`
}

type CustomerFilesResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CustomerFilesResponse"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	FileName string `xml:"FileName,omitempty"`
}

type ChargeGroupOrderCreditCardTokenResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ChargeGroupOrderCreditCardTokenResponse"`

	*ApiResponse

	_paymentIDs *ArrayOfPaymentsResponse `xml:"_paymentIDs,omitempty"`

	Amount float64 `xml:"Amount,omitempty"`

	AuthorizationCode string `xml:"AuthorizationCode,omitempty"`

	Payments *ArrayOfPaymentsResponse `xml:"Payments,omitempty"`
}

type ArrayOfPaymentsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfPaymentsResponse"`

	PaymentsResponse []*PaymentsResponse `xml:"PaymentsResponse,omitempty"`
}

type PaymentsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ PaymentsResponse"`

	_PaymentID int32 `xml:"_PaymentID,omitempty"`

	_OrderID int32 `xml:"_OrderID,omitempty"`
}

type DeleteCustomerExtendedResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteCustomerExtendedResponse"`

	*ApiResponse

	CustomerExtendedID int32 `xml:"CustomerExtendedID,omitempty"`
}

type UpdateCustomerExtendedResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateCustomerExtendedResponse"`

	*ApiResponse
}

type CreateCustomerExtendedResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateCustomerExtendedResponse"`

	*ApiResponse

	CustomerExtendedID int32 `xml:"CustomerExtendedID,omitempty"`
}

type CreateCustomerLeadResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateCustomerLeadResponse"`

	*ApiResponse

	CustomerLeadID int32 `xml:"CustomerLeadID,omitempty"`
}

type CreatePartyResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreatePartyResponse"`

	*ApiResponse

	PartyID int32 `xml:"PartyID,omitempty"`
}

type CreateBillResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateBillResponse"`

	*ApiResponse

	BillID int32 `xml:"BillID,omitempty"`
}

type CreatePayoutResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreatePayoutResponse"`

	*ApiResponse

	PayoutID int32 `xml:"PayoutID,omitempty"`

	TotalDollarAmount float64 `xml:"TotalDollarAmount,omitempty"`
}

type CreateCustomerInquiryResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateCustomerInquiryResponse"`

	*ApiResponse

	NewCustomerHistoryID int32 `xml:"NewCustomerHistoryID,omitempty"`
}

type CreateCustomerFileResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateCustomerFileResponse"`

	*ApiResponse

	FolderID int32 `xml:"FolderID,omitempty"`

	FileID int32 `xml:"FileID,omitempty"`
}

type SetItemCountryRegionResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetItemCountryRegionResponse"`

	*ApiResponse
}

type GetItemCountryRegionResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetItemCountryRegionResponse"`

	*ApiResponse

	ItemCountryRegions *ArrayOfItemCountryRegionResponse `xml:"ItemCountryRegions,omitempty"`
}

type ArrayOfItemCountryRegionResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfItemCountryRegionResponse"`

	ItemCountryRegionResponse []*ItemCountryRegionResponse `xml:"ItemCountryRegionResponse,omitempty"`
}

type ItemCountryRegionResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ItemCountryRegionResponse"`

	ItemCode string `xml:"ItemCode,omitempty"`

	CountryCode string `xml:"CountryCode,omitempty"`

	RegionCode string `xml:"RegionCode,omitempty"`

	Taxed bool `xml:"Taxed,omitempty"`

	TaxedFed bool `xml:"TaxedFed,omitempty"`

	TaxedState bool `xml:"TaxedState,omitempty"`

	UseTaxOverride bool `xml:"UseTaxOverride,omitempty"`

	TaxOverridePct float64 `xml:"TaxOverridePct,omitempty"`
}

type CreateItemResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateItemResponse"`

	*ApiResponse

	ItemCode string `xml:"ItemCode,omitempty"`
}

type GetCustomerWallResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomerWallResponse"`

	*ApiResponse

	CustomerWallItems *ArrayOfCustomerWallItemResponse `xml:"CustomerWallItems,omitempty"`
}

type ArrayOfCustomerWallItemResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfCustomerWallItemResponse"`

	CustomerWallItemResponse []*CustomerWallItemResponse `xml:"CustomerWallItemResponse,omitempty"`
}

type CustomerWallItemResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CustomerWallItemResponse"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	WallItemID int32 `xml:"WallItemID,omitempty"`

	Text string `xml:"Text,omitempty"`

	EntryDate time.Time `xml:"EntryDate,omitempty"`

	Field1 string `xml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty"`
}

type DeleteCustomerWallItemResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteCustomerWallItemResponse"`

	*ApiResponse

	CountOfDeletedRows int32 `xml:"CountOfDeletedRows,omitempty"`
}

type CreateCustomerWallItemResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateCustomerWallItemResponse"`

	*ApiResponse

	WallItemID int32 `xml:"WallItemID,omitempty"`
}

type GetCustomerLeadsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomerLeadsResponse"`

	*ApiResponse

	CustomerLeads *ArrayOfCustomerLeadsResponse `xml:"CustomerLeads,omitempty"`
}

type ArrayOfCustomerLeadsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfCustomerLeadsResponse"`

	CustomerLeadsResponse []*CustomerLeadsResponse `xml:"CustomerLeadsResponse,omitempty"`
}

type CustomerLeadsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CustomerLeadsResponse"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	CustomerLeadID int32 `xml:"CustomerLeadID,omitempty"`

	FirstName string `xml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty"`

	Company string `xml:"Company,omitempty"`

	Phone string `xml:"Phone,omitempty"`

	Phone2 string `xml:"Phone2,omitempty"`

	MobilePhone string `xml:"MobilePhone,omitempty"`

	Fax string `xml:"Fax,omitempty"`

	Email string `xml:"Email,omitempty"`

	Address1 string `xml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty"`

	City string `xml:"City,omitempty"`

	State string `xml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty"`

	BirthDate time.Time `xml:"BirthDate,omitempty"`

	Notes string `xml:"Notes,omitempty"`
}

type DeleteCustomerLeadResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteCustomerLeadResponse"`

	*ApiResponse
}

type DeleteOrderDetailResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteOrderDetailResponse"`

	*ApiResponse
}

type UpdatePartyResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdatePartyResponse"`

	*ApiResponse
}

type GetCustomerLeadSocialNetworksResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomerLeadSocialNetworksResponse"`

	*ApiResponse

	CustomerLeadSocialNetwork *ArrayOfCustomerLeadSocialNetworksResponse `xml:"CustomerLeadSocialNetwork,omitempty"`
}

type ArrayOfCustomerLeadSocialNetworksResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfCustomerLeadSocialNetworksResponse"`

	CustomerLeadSocialNetworksResponse []*CustomerLeadSocialNetworksResponse `xml:"CustomerLeadSocialNetworksResponse,omitempty"`
}

type CustomerLeadSocialNetworksResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CustomerLeadSocialNetworksResponse"`

	SocialNetworkID int32 `xml:"SocialNetworkID,omitempty"`

	SocialNetworkDescription string `xml:"SocialNetworkDescription,omitempty"`

	Url string `xml:"Url,omitempty"`
}

type SetCustomerSocialNetworksResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetCustomerSocialNetworksResponse"`

	*ApiResponse
}

type GetCustomerSocialNetworksResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomerSocialNetworksResponse"`

	*ApiResponse

	CustomerSocialNetwork *ArrayOfCustomerSocialNetworksResponse `xml:"CustomerSocialNetwork,omitempty"`
}

type ArrayOfCustomerSocialNetworksResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfCustomerSocialNetworksResponse"`

	CustomerSocialNetworksResponse []*CustomerSocialNetworksResponse `xml:"CustomerSocialNetworksResponse,omitempty"`
}

type CustomerSocialNetworksResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CustomerSocialNetworksResponse"`

	SocialNetworkID int32 `xml:"SocialNetworkID,omitempty"`

	SocialNetworkDescription string `xml:"SocialNetworkDescription,omitempty"`

	Url string `xml:"Url,omitempty"`
}

type SetCustomerSiteResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetCustomerSiteResponse"`

	*ApiResponse
}

type UpdateCustomerResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateCustomerResponse"`

	*ApiResponse
}

type UpdateOrderResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateOrderResponse"`

	*ApiResponse
}

type BaseCalculateOrderResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ BaseCalculateOrderResponse"`

	*ApiResponse
}

type CalculateOrderResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CalculateOrderResponse"`

	*BaseCalculateOrderResponse

	Total float64 `xml:"Total,omitempty"`

	SubTotal float64 `xml:"SubTotal,omitempty"`

	TaxTotal float64 `xml:"TaxTotal,omitempty"`

	ShippingTotal float64 `xml:"ShippingTotal,omitempty"`

	DiscountTotal float64 `xml:"DiscountTotal,omitempty"`

	DiscountPercent float64 `xml:"DiscountPercent,omitempty"`

	WeightTotal float64 `xml:"WeightTotal,omitempty"`

	BusinessVolumeTotal float64 `xml:"BusinessVolumeTotal,omitempty"`

	CommissionableVolumeTotal float64 `xml:"CommissionableVolumeTotal,omitempty"`

	Other1Total float64 `xml:"Other1Total,omitempty"`

	Other2Total float64 `xml:"Other2Total,omitempty"`

	Other3Total float64 `xml:"Other3Total,omitempty"`

	Other4Total float64 `xml:"Other4Total,omitempty"`

	Other5Total float64 `xml:"Other5Total,omitempty"`

	Other6Total float64 `xml:"Other6Total,omitempty"`

	Other7Total float64 `xml:"Other7Total,omitempty"`

	Other8Total float64 `xml:"Other8Total,omitempty"`

	Other9Total float64 `xml:"Other9Total,omitempty"`

	Other10Total float64 `xml:"Other10Total,omitempty"`

	Other11 string `xml:"Other11,omitempty"`

	Other12 string `xml:"Other12,omitempty"`

	Other13 string `xml:"Other13,omitempty"`

	Other14 string `xml:"Other14,omitempty"`

	Other15 string `xml:"Other15,omitempty"`

	Other16 string `xml:"Other16,omitempty"`

	Other17 string `xml:"Other17,omitempty"`

	Other18 string `xml:"Other18,omitempty"`

	Other19 string `xml:"Other19,omitempty"`

	Other20 string `xml:"Other20,omitempty"`

	ShippingTax float64 `xml:"ShippingTax,omitempty"`

	OrderTax float64 `xml:"OrderTax,omitempty"`

	FedTaxTotal float64 `xml:"FedTaxTotal,omitempty"`

	StateTaxTotal float64 `xml:"StateTaxTotal,omitempty"`

	Details *ArrayOfOrderDetailResponse `xml:"Details,omitempty"`

	ShipMethods *ArrayOfShipMethodResponse `xml:"ShipMethods,omitempty"`

	Warnings *ArrayOfString `xml:"Warnings,omitempty"`
}

type ArrayOfOrderDetailResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfOrderDetailResponse"`

	OrderDetailResponse []*OrderDetailResponse `xml:"OrderDetailResponse,omitempty"`
}

type OrderDetailResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ OrderDetailResponse"`

	ItemCode string `xml:"ItemCode,omitempty"`

	Description string `xml:"Description,omitempty"`

	Quantity float64 `xml:"Quantity,omitempty"`

	PriceEach float64 `xml:"PriceEach,omitempty"`

	PriceTotal float64 `xml:"PriceTotal,omitempty"`

	Tax float64 `xml:"Tax,omitempty"`

	WeightEach float64 `xml:"WeightEach,omitempty"`

	Weight float64 `xml:"Weight,omitempty"`

	BusinessVolumeEach float64 `xml:"BusinessVolumeEach,omitempty"`

	BusinesVolume float64 `xml:"BusinesVolume,omitempty"`

	CommissionableVolumeEach float64 `xml:"CommissionableVolumeEach,omitempty"`

	CommissionableVolume float64 `xml:"CommissionableVolume,omitempty"`

	Other1Each float64 `xml:"Other1Each,omitempty"`

	Other1 float64 `xml:"Other1,omitempty"`

	Other2Each float64 `xml:"Other2Each,omitempty"`

	Other2 float64 `xml:"Other2,omitempty"`

	Other3Each float64 `xml:"Other3Each,omitempty"`

	Other3 float64 `xml:"Other3,omitempty"`

	Other4Each float64 `xml:"Other4Each,omitempty"`

	Other4 float64 `xml:"Other4,omitempty"`

	Other5Each float64 `xml:"Other5Each,omitempty"`

	Other5 float64 `xml:"Other5,omitempty"`

	Other6Each float64 `xml:"Other6Each,omitempty"`

	Other6 float64 `xml:"Other6,omitempty"`

	Other7Each float64 `xml:"Other7Each,omitempty"`

	Other7 float64 `xml:"Other7,omitempty"`

	Other8Each float64 `xml:"Other8Each,omitempty"`

	Other8 float64 `xml:"Other8,omitempty"`

	Other9Each float64 `xml:"Other9Each,omitempty"`

	Other9 float64 `xml:"Other9,omitempty"`

	Other10Each float64 `xml:"Other10Each,omitempty"`

	Other10 float64 `xml:"Other10,omitempty"`

	ParentItemCode string `xml:"ParentItemCode,omitempty"`

	Taxable float64 `xml:"Taxable,omitempty"`

	FedTax float64 `xml:"FedTax,omitempty"`

	StateTax float64 `xml:"StateTax,omitempty"`

	CityTax float64 `xml:"CityTax,omitempty"`

	CityLocalTax float64 `xml:"CityLocalTax,omitempty"`

	CountyTax float64 `xml:"CountyTax,omitempty"`

	CountyLocalTax float64 `xml:"CountyLocalTax,omitempty"`

	ManualTax float64 `xml:"ManualTax,omitempty"`

	IsStateTaxOverride bool `xml:"IsStateTaxOverride,omitempty"`

	OrderLine int32 `xml:"OrderLine,omitempty"`

	Reference1 string `xml:"Reference1,omitempty"`

	ShippingPriceEach float64 `xml:"ShippingPriceEach,omitempty"`
}

type ArrayOfShipMethodResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfShipMethodResponse"`

	ShipMethodResponse []*ShipMethodResponse `xml:"ShipMethodResponse,omitempty"`
}

type ShipMethodResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ShipMethodResponse"`

	ShipMethodID int32 `xml:"ShipMethodID,omitempty"`

	Description string `xml:"Description,omitempty"`

	ShippingAmount float64 `xml:"ShippingAmount,omitempty"`
}

type CreateAutoOrderResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateAutoOrderResponse"`

	*BaseCalculateOrderResponse

	AutoOrderID int32 `xml:"AutoOrderID,omitempty"`

	Description string `xml:"Description,omitempty"`

	Total float64 `xml:"Total,omitempty"`

	SubTotal float64 `xml:"SubTotal,omitempty"`

	TaxTotal float64 `xml:"TaxTotal,omitempty"`

	ShippingTotal float64 `xml:"ShippingTotal,omitempty"`

	DiscountTotal float64 `xml:"DiscountTotal,omitempty"`

	WeightTotal float64 `xml:"WeightTotal,omitempty"`

	BusinessVolumeTotal float64 `xml:"BusinessVolumeTotal,omitempty"`

	CommissionableVolumeTotal float64 `xml:"CommissionableVolumeTotal,omitempty"`

	Other1Total float64 `xml:"Other1Total,omitempty"`

	Other2Total float64 `xml:"Other2Total,omitempty"`

	Other3Total float64 `xml:"Other3Total,omitempty"`

	Other4Total float64 `xml:"Other4Total,omitempty"`

	Other5Total float64 `xml:"Other5Total,omitempty"`

	Other6Total float64 `xml:"Other6Total,omitempty"`

	Other7Total float64 `xml:"Other7Total,omitempty"`

	Other8Total float64 `xml:"Other8Total,omitempty"`

	Other9Total float64 `xml:"Other9Total,omitempty"`

	Other10Total float64 `xml:"Other10Total,omitempty"`

	ShippingTax float64 `xml:"ShippingTax,omitempty"`

	OrderTax float64 `xml:"OrderTax,omitempty"`

	Details *ArrayOfOrderDetailResponse `xml:"Details,omitempty"`
}

type CreateOrderResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateOrderResponse"`

	*BaseCalculateOrderResponse

	OrderID int32 `xml:"OrderID,omitempty"`

	Total float64 `xml:"Total,omitempty"`

	SubTotal float64 `xml:"SubTotal,omitempty"`

	TaxTotal float64 `xml:"TaxTotal,omitempty"`

	ShippingTotal float64 `xml:"ShippingTotal,omitempty"`

	DiscountTotal float64 `xml:"DiscountTotal,omitempty"`

	WeightTotal float64 `xml:"WeightTotal,omitempty"`

	BusinessVolumeTotal float64 `xml:"BusinessVolumeTotal,omitempty"`

	CommissionableVolumeTotal float64 `xml:"CommissionableVolumeTotal,omitempty"`

	Other1Total float64 `xml:"Other1Total,omitempty"`

	Other2Total float64 `xml:"Other2Total,omitempty"`

	Other3Total float64 `xml:"Other3Total,omitempty"`

	Other4Total float64 `xml:"Other4Total,omitempty"`

	Other5Total float64 `xml:"Other5Total,omitempty"`

	Other6Total float64 `xml:"Other6Total,omitempty"`

	Other7Total float64 `xml:"Other7Total,omitempty"`

	Other8Total float64 `xml:"Other8Total,omitempty"`

	Other9Total float64 `xml:"Other9Total,omitempty"`

	Other10Total float64 `xml:"Other10Total,omitempty"`

	ShippingTax float64 `xml:"ShippingTax,omitempty"`

	OrderTax float64 `xml:"OrderTax,omitempty"`

	Details *ArrayOfOrderDetailResponse `xml:"Details,omitempty"`

	Warnings *ArrayOfString `xml:"Warnings,omitempty"`
}

type CreateCustomerResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateCustomerResponse"`

	*ApiResponse

	CustomerID int32 `xml:"CustomerID,omitempty"`
}

type SetAccountResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetAccountResponse"`

	*ApiResponse
}

type CreateOrderImportResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateOrderImportResponse"`

	*ApiResponse

	OrderID int32 `xml:"OrderID,omitempty"`
}

type TransactionalResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ TransactionalResponse"`

	*ApiResponse

	TransactionResponses *ArrayOfApiResponse `xml:"TransactionResponses,omitempty"`
}

type ArrayOfApiResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfApiResponse"`

	ApiResponse []*ApiResponse `xml:"ApiResponse,omitempty"`
}

type BaseCreateExpectedPaymentResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ BaseCreateExpectedPaymentResponse"`

	*ApiResponse
}

type CreateExpectedBankWireResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateExpectedBankWireResponse"`

	*BaseCreateExpectedPaymentResponse

	ExpectedPaymentID int32 `xml:"ExpectedPaymentID,omitempty"`
}

type CreateExpectedCODResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateExpectedCODResponse"`

	*BaseCreateExpectedPaymentResponse

	ExpectedPaymentID int32 `xml:"ExpectedPaymentID,omitempty"`
}

type CreateExpectedPaymentResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateExpectedPaymentResponse"`

	*BaseCreateExpectedPaymentResponse

	ExpectedPaymentID int32 `xml:"ExpectedPaymentID,omitempty"`
}

type AuthorizeOnlyCreditCardResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ AuthorizeOnlyCreditCardResponse"`

	*ApiResponse

	AuthorizationCode string `xml:"AuthorizationCode,omitempty"`

	MerchantTransactionKey string `xml:"MerchantTransactionKey,omitempty"`
}

type BaseCreatePaymentResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ BaseCreatePaymentResponse"`

	*ApiResponse
}

type CreatePaymentCreditCardResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreatePaymentCreditCardResponse"`

	*BaseCreatePaymentResponse

	PaymentID int32 `xml:"PaymentID,omitempty"`
}

type ChargeCreditCardResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ChargeCreditCardResponse"`

	*CreatePaymentCreditCardResponse

	Amount float64 `xml:"Amount,omitempty"`

	AuthorizationCode string `xml:"AuthorizationCode,omitempty"`
}

type RefundPriorCreditCardChargeResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ RefundPriorCreditCardChargeResponse"`

	*CreatePaymentCreditCardResponse

	Amount float64 `xml:"Amount,omitempty"`

	AuthorizationCode string `xml:"AuthorizationCode,omitempty"`
}

type CreatePaymentResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreatePaymentResponse"`

	*BaseCreatePaymentResponse

	PaymentID int32 `xml:"PaymentID,omitempty"`
}

type DebitBankAccountResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DebitBankAccountResponse"`

	*CreatePaymentResponse

	Amount float64 `xml:"Amount,omitempty"`

	AuthorizationCode string `xml:"AuthorizationCode,omitempty"`
}

type RefundPriorWalletChargeResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ RefundPriorWalletChargeResponse"`

	*CreatePaymentResponse

	Amount float64 `xml:"Amount,omitempty"`

	AuthorizationCode string `xml:"AuthorizationCode,omitempty"`
}

type UpdateCustomerLeadResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateCustomerLeadResponse"`

	*ApiResponse
}

type UpdateItemResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateItemResponse"`

	*ApiResponse
}

type SetCustomerLeadSocialNetworksResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetCustomerLeadSocialNetworksResponse"`

	*ApiResponse
}

type ApiAuthentication struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ApiAuthentication"`

	LoginName string `xml:"LoginName,omitempty"`

	Password string `xml:"Password,omitempty"`

	Company string `xml:"Company,omitempty"`

	Identity string `xml:"Identity,omitempty"`

	RequestTimeUtc time.Time `xml:"RequestTimeUtc,omitempty"`

	Signature string `xml:"Signature,omitempty"`
}

type MoveEmailRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ MoveEmailRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	MailID int32 `xml:"MailID,omitempty"`

	ToMailFolderID int32 `xml:"ToMailFolderID,omitempty"`
}

type MoveEmailResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ MoveEmailResponse"`

	*ApiResponse
}

type UpdateEmailStatusRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateEmailStatusRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	MailID int32 `xml:"MailID,omitempty"`

	MailStatusType *MailStatusType `xml:"MailStatusType,omitempty"`
}

type UpdateEmailStatusResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateEmailStatusResponse"`

	*ApiResponse
}

type GetEmailAttachmentRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetEmailAttachmentRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	MailID int32 `xml:"MailID,omitempty"`

	AttachmentID int32 `xml:"AttachmentID,omitempty"`
}

type GetEmailAttachmentResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetEmailAttachmentResponse"`

	*ApiResponse

	CustomerID int32 `xml:"CustomerID,omitempty"`

	MailID int32 `xml:"MailID,omitempty"`

	Attachment *EmailAttachment `xml:"Attachment,omitempty"`
}

type DeleteEmailRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteEmailRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	MailID int32 `xml:"MailID,omitempty"`
}

type DeleteEmailResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteEmailResponse"`

	*ApiResponse
}

type CreateEmailTemplateRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateEmailTemplateRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	Description string `xml:"Description,omitempty"`

	Content string `xml:"Content,omitempty"`
}

type CreateEmailTemplateResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateEmailTemplateResponse"`

	*ApiResponse

	TemplateID int32 `xml:"TemplateID,omitempty"`
}

type UpdateEmailTemplateRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateEmailTemplateRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	TemplateID int32 `xml:"TemplateID,omitempty"`

	Description string `xml:"Description,omitempty"`

	Content string `xml:"Content,omitempty"`
}

type UpdateEmailTemplateResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateEmailTemplateResponse"`

	*ApiResponse
}

type DeleteEmailTemplateRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteEmailTemplateRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	TemplateID int32 `xml:"TemplateID,omitempty"`
}

type DeleteEmailTemplateResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteEmailTemplateResponse"`

	*ApiResponse
}

type EnsureMailFoldersRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ EnsureMailFoldersRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`
}

type EnsureMailFoldersResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ EnsureMailFoldersResponse"`

	*ApiResponse
}

type CreateMailFolderRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateMailFolderRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	MailFolderName string `xml:"MailFolderName,omitempty"`
}

type CreateMailFolderResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateMailFolderResponse"`

	*ApiResponse
}

type UpdateMailFolderRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateMailFolderRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	MailFolderID int32 `xml:"MailFolderID,omitempty"`

	MailFolderName string `xml:"MailFolderName,omitempty"`
}

type UpdateMailFolderResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateMailFolderResponse"`

	*ApiResponse
}

type DeleteMailFolderRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteMailFolderRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	MailFolderID int32 `xml:"MailFolderID,omitempty"`
}

type DeleteMailFolderResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteMailFolderResponse"`

	*ApiResponse
}

type EmptyMailFolderRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ EmptyMailFolderRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	MailFolderID int32 `xml:"MailFolderID,omitempty"`
}

type EmptyMailFolderResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ EmptyMailFolderResponse"`

	*ApiResponse
}

type UpdateItemRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateItemRequest"`

	*ApiRequest

	ItemCode string `xml:"ItemCode,omitempty"`

	Description string `xml:"Description,omitempty"`

	Weight float64 `xml:"Weight,omitempty"`

	Notes string `xml:"Notes,omitempty"`

	AvailableInAllCountryRegions bool `xml:"AvailableInAllCountryRegions,omitempty"`

	TaxedInAllCountryRegions bool `xml:"TaxedInAllCountryRegions,omitempty"`

	AvailableInAllWarehouses bool `xml:"AvailableInAllWarehouses,omitempty"`

	IsVirtual bool `xml:"IsVirtual,omitempty"`

	ItemTypeID int32 `xml:"ItemTypeID,omitempty"`

	ShortDetail string `xml:"ShortDetail,omitempty"`

	ShortDetail2 string `xml:"ShortDetail2,omitempty"`

	ShortDetail3 string `xml:"ShortDetail3,omitempty"`

	ShortDetail4 string `xml:"ShortDetail4,omitempty"`

	LongDetail string `xml:"LongDetail,omitempty"`

	LongDetail2 string `xml:"LongDetail2,omitempty"`

	LongDetail3 string `xml:"LongDetail3,omitempty"`

	LongDetail4 string `xml:"LongDetail4,omitempty"`

	OtherCheck1 bool `xml:"OtherCheck1,omitempty"`

	OtherCheck2 bool `xml:"OtherCheck2,omitempty"`

	OtherCheck3 bool `xml:"OtherCheck3,omitempty"`

	OtherCheck4 bool `xml:"OtherCheck4,omitempty"`

	OtherCheck5 bool `xml:"OtherCheck5,omitempty"`

	Field1 string `xml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty"`

	HideFromSearch bool `xml:"HideFromSearch,omitempty"`
}

type SetItemPriceResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetItemPriceResponse"`

	*ApiResponse
}

type SetItemWarehouseResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetItemWarehouseResponse"`

	*ApiResponse
}

type GetItemCountryRegionRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetItemCountryRegionRequest"`

	*ApiRequest

	ItemCode string `xml:"ItemCode,omitempty"`
}

type SetItemImageRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetItemImageRequest"`

	*ApiRequest

	ItemCode string `xml:"ItemCode,omitempty"`

	TinyImageName string `xml:"TinyImageName,omitempty"`

	TinyImageData []byte `xml:"TinyImageData,omitempty"`

	SmallImageName string `xml:"SmallImageName,omitempty"`

	SmallImageData []byte `xml:"SmallImageData,omitempty"`

	LargeImageName string `xml:"LargeImageName,omitempty"`

	LargeImageData []byte `xml:"LargeImageData,omitempty"`
}

type SetItemImageResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetItemImageResponse"`

	*ApiResponse
}

type SetImageFileRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetImageFileRequest"`

	*ApiRequest

	Path string `xml:"Path,omitempty"`

	Name string `xml:"Name,omitempty"`

	ImageData []byte `xml:"ImageData,omitempty"`
}

type SetImageFileResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetImageFileResponse"`

	*ApiResponse
}

type GetCustomerBalancesRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomerBalancesRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`
}

type GetCustomerBalancesResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomerBalancesResponse"`

	*ApiResponse

	CustomerBalances *ArrayOfCustomerBalanceResponse `xml:"CustomerBalances,omitempty"`
}

type ArrayOfCustomerBalanceResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfCustomerBalanceResponse"`

	CustomerBalanceResponse []*CustomerBalanceResponse `xml:"CustomerBalanceResponse,omitempty"`
}

type CustomerBalanceResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CustomerBalanceResponse"`

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	CurrencyDescription string `xml:"CurrencyDescription,omitempty"`

	Balance float64 `xml:"Balance,omitempty"`
}

type CreateCustomerBalanceAdjustmentRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateCustomerBalanceAdjustmentRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	CustomerTransactionTypeID int32 `xml:"CustomerTransactionTypeID,omitempty"`

	TransactionDate time.Time `xml:"TransactionDate,omitempty"`

	Amount float64 `xml:"Amount,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	Notes string `xml:"Notes,omitempty"`
}

type CreateCustomerBalanceAdjustmentResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateCustomerBalanceAdjustmentResponse"`

	*ApiResponse

	TransactionID int32 `xml:"TransactionID,omitempty"`
}

type GetPartiesRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetPartiesRequest"`

	*ApiRequest

	PartyID int32 `xml:"PartyID,omitempty"`

	HostID int32 `xml:"HostID,omitempty"`

	DistributorID int32 `xml:"DistributorID,omitempty"`

	PartyStatusType int32 `xml:"PartyStatusType,omitempty"`

	BookingPartyID int32 `xml:"BookingPartyID,omitempty"`

	Field1 string `xml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty"`
}

type GetPartiesResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetPartiesResponse"`

	*ApiResponse

	Parties *ArrayOfPartyResponse `xml:"Parties,omitempty"`
}

type ArrayOfPartyResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfPartyResponse"`

	PartyResponse []*PartyResponse `xml:"PartyResponse,omitempty"`
}

type PartyResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ PartyResponse"`

	PartyID int32 `xml:"PartyID,omitempty"`

	PartyType int32 `xml:"PartyType,omitempty"`

	PartyStatusType int32 `xml:"PartyStatusType,omitempty"`

	HostID int32 `xml:"HostID,omitempty"`

	DistributorID int32 `xml:"DistributorID,omitempty"`

	StartDate time.Time `xml:"StartDate,omitempty"`

	CloseDate time.Time `xml:"CloseDate,omitempty"`

	Description string `xml:"Description,omitempty"`

	EventStart time.Time `xml:"EventStart,omitempty"`

	EventEnd time.Time `xml:"EventEnd,omitempty"`

	LanguageID int32 `xml:"LanguageID,omitempty"`

	Information string `xml:"Information,omitempty"`

	Address *PartyAddress `xml:"Address,omitempty"`

	BookingPartyID int32 `xml:"BookingPartyID,omitempty"`

	Field1 string `xml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty"`
}

type GetGuestsRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetGuestsRequest"`

	*ApiRequest

	GuestID int32 `xml:"GuestID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	HostID int32 `xml:"HostID,omitempty"`

	GuestStatuses *ArrayOfInt `xml:"GuestStatuses,omitempty"`

	LanguageID int32 `xml:"LanguageID,omitempty"`

	FirstName string `xml:"FirstName,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty"`

	LastName string `xml:"LastName,omitempty"`

	NameSuffix string `xml:"NameSuffix,omitempty"`

	Company string `xml:"Company,omitempty"`

	Email string `xml:"Email,omitempty"`

	Phone string `xml:"Phone,omitempty"`

	Phone2 string `xml:"Phone2,omitempty"`

	MobilePhone string `xml:"MobilePhone,omitempty"`

	Address1 string `xml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty"`

	Address3 string `xml:"Address3,omitempty"`

	City string `xml:"City,omitempty"`

	County string `xml:"County,omitempty"`

	State string `xml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty"`

	Field1 string `xml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty"`

	Field11 string `xml:"Field11,omitempty"`

	Field12 string `xml:"Field12,omitempty"`

	Field13 string `xml:"Field13,omitempty"`

	Field14 string `xml:"Field14,omitempty"`

	Field15 string `xml:"Field15,omitempty"`

	Date1 time.Time `xml:"Date1,omitempty"`

	Date2 time.Time `xml:"Date2,omitempty"`

	Date3 time.Time `xml:"Date3,omitempty"`

	Date4 time.Time `xml:"Date4,omitempty"`

	Date5 time.Time `xml:"Date5,omitempty"`

	CreatedDateStart time.Time `xml:"CreatedDateStart,omitempty"`

	CreatedDateEnd time.Time `xml:"CreatedDateEnd,omitempty"`
}

type GetGuestsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetGuestsResponse"`

	*ApiResponse

	Guests *ArrayOfGuestResponse `xml:"Guests,omitempty"`
}

type ArrayOfGuestResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfGuestResponse"`

	GuestResponse []*GuestResponse `xml:"GuestResponse,omitempty"`
}

type GuestResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GuestResponse"`

	GuestID int32 `xml:"GuestID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	HostID int32 `xml:"HostID,omitempty"`

	FirstName string `xml:"FirstName,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty"`

	LastName string `xml:"LastName,omitempty"`

	NameSuffix string `xml:"NameSuffix,omitempty"`

	Company string `xml:"Company,omitempty"`

	Gender *Gender `xml:"Gender,omitempty"`

	GuestStatus int32 `xml:"GuestStatus,omitempty"`

	LanguageID int32 `xml:"LanguageID,omitempty"`

	Address1 string `xml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty"`

	Address3 string `xml:"Address3,omitempty"`

	City string `xml:"City,omitempty"`

	County string `xml:"County,omitempty"`

	State string `xml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty"`

	Phone string `xml:"Phone,omitempty"`

	Phone2 string `xml:"Phone2,omitempty"`

	MobilePhone string `xml:"MobilePhone,omitempty"`

	Email string `xml:"Email,omitempty"`

	Field1 string `xml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty"`

	Field11 string `xml:"Field11,omitempty"`

	Field12 string `xml:"Field12,omitempty"`

	Field13 string `xml:"Field13,omitempty"`

	Field14 string `xml:"Field14,omitempty"`

	Field15 string `xml:"Field15,omitempty"`

	Date1 time.Time `xml:"Date1,omitempty"`

	Date2 time.Time `xml:"Date2,omitempty"`

	Date3 time.Time `xml:"Date3,omitempty"`

	Date4 time.Time `xml:"Date4,omitempty"`

	Date5 time.Time `xml:"Date5,omitempty"`

	CreatedDate time.Time `xml:"CreatedDate,omitempty"`

	ModifiedDate time.Time `xml:"ModifiedDate,omitempty"`
}

type GetPartyGuestsRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetPartyGuestsRequest"`

	*ApiRequest

	PartyID int32 `xml:"PartyID,omitempty"`
}

type GetPartyGuestsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetPartyGuestsResponse"`

	*ApiResponse

	Guests *ArrayOfGuestResponse `xml:"Guests,omitempty"`
}

type GetGuestSocialNetworksRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetGuestSocialNetworksRequest"`

	*ApiRequest

	GuestID int32 `xml:"GuestID,omitempty"`
}

type GetGuestSocialNetworksResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetGuestSocialNetworksResponse"`

	*ApiResponse

	GuestSocialNetworks *ArrayOfGuestSocialNetworksResponse `xml:"GuestSocialNetworks,omitempty"`
}

type ArrayOfGuestSocialNetworksResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfGuestSocialNetworksResponse"`

	GuestSocialNetworksResponse []*GuestSocialNetworksResponse `xml:"GuestSocialNetworksResponse,omitempty"`
}

type GuestSocialNetworksResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GuestSocialNetworksResponse"`

	SocialNetworkID int32 `xml:"SocialNetworkID,omitempty"`

	SocialNetworkDescription string `xml:"SocialNetworkDescription,omitempty"`

	Url string `xml:"Url,omitempty"`
}

type SetGuestSocialNetworksRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetGuestSocialNetworksRequest"`

	*ApiRequest

	GuestID int32 `xml:"GuestID,omitempty"`

	GuestSocialNetworks *ArrayOfGuestSocialNetworkRequest `xml:"GuestSocialNetworks,omitempty"`
}

type ArrayOfGuestSocialNetworkRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfGuestSocialNetworkRequest"`

	GuestSocialNetworkRequest []*GuestSocialNetworkRequest `xml:"GuestSocialNetworkRequest,omitempty"`
}

type GuestSocialNetworkRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GuestSocialNetworkRequest"`

	SocialNetworkID int32 `xml:"SocialNetworkID,omitempty"`

	Url string `xml:"Url,omitempty"`
}

type SetGuestSocialNetworksResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetGuestSocialNetworksResponse"`

	*ApiResponse
}

type CreateGuestRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateGuestRequest"`

	*ApiRequest

	Date1 time.Time `xml:"Date1,omitempty"`

	Date2 time.Time `xml:"Date2,omitempty"`

	Date3 time.Time `xml:"Date3,omitempty"`

	Date4 time.Time `xml:"Date4,omitempty"`

	Date5 time.Time `xml:"Date5,omitempty"`

	HostID int32 `xml:"HostID,omitempty"`

	PartyID int32 `xml:"PartyID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	FirstName string `xml:"FirstName,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty"`

	LastName string `xml:"LastName,omitempty"`

	NameSuffix string `xml:"NameSuffix,omitempty"`

	Company string `xml:"Company,omitempty"`

	Gender *Gender `xml:"Gender,omitempty"`

	LanguageID int32 `xml:"LanguageID,omitempty"`

	GuestStatus int32 `xml:"GuestStatus,omitempty"`

	Address1 string `xml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty"`

	Address3 string `xml:"Address3,omitempty"`

	City string `xml:"City,omitempty"`

	County string `xml:"County,omitempty"`

	State string `xml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty"`

	Phone string `xml:"Phone,omitempty"`

	Phone2 string `xml:"Phone2,omitempty"`

	MobilePhone string `xml:"MobilePhone,omitempty"`

	Email string `xml:"Email,omitempty"`

	Field1 string `xml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty"`

	Field11 string `xml:"Field11,omitempty"`

	Field12 string `xml:"Field12,omitempty"`

	Field13 string `xml:"Field13,omitempty"`

	Field14 string `xml:"Field14,omitempty"`

	Field15 string `xml:"Field15,omitempty"`

	Notes string `xml:"Notes,omitempty"`

	EntryDate time.Time `xml:"EntryDate,omitempty"`
}

type CreateGuestResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateGuestResponse"`

	*ApiResponse

	GuestID int32 `xml:"GuestID,omitempty"`
}

type UpdateGuestRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateGuestRequest"`

	*ApiRequest

	GuestID int32 `xml:"GuestID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	FirstName string `xml:"FirstName,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty"`

	LastName string `xml:"LastName,omitempty"`

	NameSuffix string `xml:"NameSuffix,omitempty"`

	Company string `xml:"Company,omitempty"`

	Gender *Gender `xml:"Gender,omitempty"`

	GuestStatus int32 `xml:"GuestStatus,omitempty"`

	LanguageID int32 `xml:"LanguageID,omitempty"`

	Address1 string `xml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty"`

	Address3 string `xml:"Address3,omitempty"`

	City string `xml:"City,omitempty"`

	County string `xml:"County,omitempty"`

	State string `xml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty"`

	Phone string `xml:"Phone,omitempty"`

	Phone2 string `xml:"Phone2,omitempty"`

	MobilePhone string `xml:"MobilePhone,omitempty"`

	Email string `xml:"Email,omitempty"`

	Field1 string `xml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty"`

	Field11 string `xml:"Field11,omitempty"`

	Field12 string `xml:"Field12,omitempty"`

	Field13 string `xml:"Field13,omitempty"`

	Field14 string `xml:"Field14,omitempty"`

	Field15 string `xml:"Field15,omitempty"`

	Date1 time.Time `xml:"Date1,omitempty"`

	Date2 time.Time `xml:"Date2,omitempty"`

	Date3 time.Time `xml:"Date3,omitempty"`

	Date4 time.Time `xml:"Date4,omitempty"`

	Date5 time.Time `xml:"Date5,omitempty"`

	Notes string `xml:"Notes,omitempty"`
}

type UpdateGuestResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateGuestResponse"`

	*ApiResponse
}

type AddPartyGuestsRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ AddPartyGuestsRequest"`

	*ApiRequest

	PartyID int32 `xml:"PartyID,omitempty"`

	GuestIDs *ArrayOfInt `xml:"GuestIDs,omitempty"`
}

type AddPartyGuestsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ AddPartyGuestsResponse"`

	*ApiResponse
}

type RemovePartyGuestsRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ RemovePartyGuestsRequest"`

	*ApiRequest

	PartyID int32 `xml:"PartyID,omitempty"`

	GuestIDs *ArrayOfInt `xml:"GuestIDs,omitempty"`
}

type RemovePartyGuestsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ RemovePartyGuestsResponse"`

	*ApiResponse
}

type CreateExtendedDbSchemeRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateExtendedDbSchemeRequest"`

	*ApiRequest

	Schema *Schema `xml:"Schema,omitempty"`
}

type Schema struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ Schema"`

	Name string `xml:"Name,omitempty"`

	Entities *ArrayOfEntity `xml:"Entities,omitempty"`
}

type ArrayOfEntity struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfEntity"`

	Entity []*Entity `xml:"Entity,omitempty"`
}

type Entity struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ Entity"`

	SchemaName string `xml:"SchemaName,omitempty"`

	DbSchema string `xml:"DbSchema,omitempty"`

	EntityName string `xml:"EntityName,omitempty"`

	EntitySetName string `xml:"EntitySetName,omitempty"`

	Properties *ArrayOfProperty `xml:"Properties,omitempty"`

	Navigations *ArrayOfNavigation `xml:"Navigations,omitempty"`

	SyncTypeID int32 `xml:"SyncTypeID,omitempty"`
}

type ArrayOfProperty struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfProperty"`

	Property []*Property `xml:"Property,omitempty"`
}

type Property struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ Property"`

	Name string `xml:"Name,omitempty"`

	IsKey bool `xml:"IsKey,omitempty"`

	IsNew bool `xml:"IsNew,omitempty"`

	IsAutoNumber bool `xml:"IsAutoNumber,omitempty"`

	AllowDbNull bool `xml:"AllowDbNull,omitempty"`

	Type *PropertyType `xml:"Type,omitempty"`

	DefaultName string `xml:"DefaultName,omitempty"`

	DefaultValue string `xml:"DefaultValue,omitempty"`

	Size int32 `xml:"Size,omitempty"`
}

type ArrayOfNavigation struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfNavigation"`

	Navigation []*Navigation `xml:"Navigation,omitempty"`
}

type Navigation struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ Navigation"`

	Name string `xml:"Name,omitempty"`
}

type CreateExtendedDbSchemaResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateExtendedDbSchemaResponse"`

	*ApiResponse

	Schema *Schema `xml:"Schema,omitempty"`
}

type GetSchemaRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetSchemaRequest"`

	*ApiRequest

	SchemaName string `xml:"SchemaName,omitempty"`
}

type GetSchemaResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetSchemaResponse"`

	*ApiResponse

	Schema *Schema `xml:"Schema,omitempty"`
}

type DeleteSchemaRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteSchemaRequest"`

	*ApiRequest

	SchemaName string `xml:"SchemaName,omitempty"`
}

type DeleteSchemaResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteSchemaResponse"`

	*ApiResponse

	SchemaName string `xml:"SchemaName,omitempty"`
}

type CreateEntityRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateEntityRequest"`

	*ApiRequest

	Entity *Entity `xml:"Entity,omitempty"`
}

type CreateEntityResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateEntityResponse"`

	*ApiResponse

	Entity *Entity `xml:"Entity,omitempty"`
}

type GetEntityRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetEntityRequest"`

	*ApiRequest

	SchemaName string `xml:"SchemaName,omitempty"`

	EntityName string `xml:"EntityName,omitempty"`
}

type GetEntityResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetEntityResponse"`

	*ApiResponse

	Entity *Entity `xml:"Entity,omitempty"`
}

type UpdateEntityRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateEntityRequest"`

	*ApiRequest

	Entity *Entity `xml:"Entity,omitempty"`

	EntityName string `xml:"EntityName,omitempty"`
}

type UpdateEntityResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateEntityResponse"`

	*ApiResponse

	Entity *Entity `xml:"Entity,omitempty"`
}

type DeleteEntityRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteEntityRequest"`

	*ApiRequest

	SchemaName string `xml:"SchemaName,omitempty"`

	EntityName string `xml:"EntityName,omitempty"`
}

type DeleteEntityResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteEntityResponse"`

	*ApiResponse

	SchemaName string `xml:"schemaName,omitempty"`

	EntityName string `xml:"EntityName,omitempty"`
}

type StartSandboxRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ StartSandboxRequest"`

	*ApiRequest

	SandboxID int32 `xml:"SandboxID,omitempty"`

	EnableRevolvingCommissionRun bool `xml:"EnableRevolvingCommissionRun,omitempty"`

	EnableBiSync bool `xml:"EnableBiSync,omitempty"`

	UseRealTimeBackup bool `xml:"UseRealTimeBackup,omitempty"`

	SyncFilterDays int32 `xml:"SyncFilterDays,omitempty"`

	SyncSettingsEnable string `xml:"SyncSettingsEnable,omitempty"`

	PremiumSandbox bool `xml:"PremiumSandbox,omitempty"`
}

type StartSandboxResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ StartSandboxResponse"`

	*ApiResponse

	Sandbox *Sandbox `xml:"Sandbox,omitempty"`
}

type Sandbox struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ Sandbox"`

	CompanyID int32 `xml:"CompanyID,omitempty"`

	SandboxID int32 `xml:"SandboxID,omitempty"`

	Description string `xml:"Description,omitempty"`

	Type *SandboxType `xml:"Type,omitempty"`

	Status string `xml:"Status,omitempty"`

	StartDate time.Time `xml:"StartDate,omitempty"`

	PercentComplete float64 `xml:"PercentComplete,omitempty"`

	Hours float64 `xml:"Hours,omitempty"`

	AllowVolumePush bool `xml:"AllowVolumePush,omitempty"`

	AllowBiSync bool `xml:"AllowBiSync,omitempty"`

	SyncFilterDays int32 `xml:"SyncFilterDays,omitempty"`

	UseRealTimeBackup bool `xml:"UseRealTimeBackup,omitempty"`
}

type StopSandboxRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ StopSandboxRequest"`

	*ApiRequest

	SandboxID int32 `xml:"SandboxID,omitempty"`
}

type StopSandboxResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ StopSandboxResponse"`

	*ApiResponse

	SandboxID int32 `xml:"SandboxID,omitempty"`
}

type RefreshSandboxRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ RefreshSandboxRequest"`

	*ApiRequest

	SandboxID int32 `xml:"SandboxID,omitempty"`
}

type RefreshSandboxResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ RefreshSandboxResponse"`

	*ApiResponse

	Sandbox *Sandbox `xml:"Sandbox,omitempty"`
}

type GetSandboxRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetSandboxRequest"`

	*ApiRequest

	SandboxID int32 `xml:"SandboxID,omitempty"`
}

type GetSandboxResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetSandboxResponse"`

	*ApiResponse

	Sandbox *Sandbox `xml:"Sandbox,omitempty"`

	Sandboxes *ArrayOfSandbox `xml:"Sandboxes,omitempty"`
}

type ArrayOfSandbox struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfSandbox"`

	Sandbox []*Sandbox `xml:"Sandbox,omitempty"`
}

type GetCustomersRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomersRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	LoginName string `xml:"LoginName,omitempty"`

	FirstName string `xml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty"`

	Company string `xml:"Company,omitempty"`

	Email string `xml:"Email,omitempty"`

	Phone string `xml:"Phone,omitempty"`

	Phone2 string `xml:"Phone2,omitempty"`

	MobilePhone string `xml:"MobilePhone,omitempty"`

	Fax string `xml:"Fax,omitempty"`

	MainAddress1 string `xml:"MainAddress1,omitempty"`

	MainAddress2 string `xml:"MainAddress2,omitempty"`

	MainAddress3 string `xml:"MainAddress3,omitempty"`

	MainCity string `xml:"MainCity,omitempty"`

	MainState string `xml:"MainState,omitempty"`

	MainZip string `xml:"MainZip,omitempty"`

	MainCountry string `xml:"MainCountry,omitempty"`

	TaxID string `xml:"TaxID,omitempty"`

	//CustomerTypes *ArrayOfInt `xml:"CustomerTypes,omitempty"`

	//CustomerStatuses *ArrayOfInt `xml:"CustomerStatuses,omitempty"`

	EnrollerID int32 `xml:"EnrollerID,omitempty"`

	SponsorID int32 `xml:"SponsorID,omitempty"`

	Field1 string `xml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty"`

	Field11 string `xml:"Field11,omitempty"`

	Field12 string `xml:"Field12,omitempty"`

	Field13 string `xml:"Field13,omitempty"`

	Field14 string `xml:"Field14,omitempty"`

	Field15 string `xml:"Field15,omitempty"`

	//CreatedDateStart time.Time `xml:"CreatedDateStart,omitempty"`

	//CreatedDateEnd time.Time `xml:"CreatedDateEnd,omitempty"`

	GreaterThanCustomerID int32 `xml:"GreaterThanCustomerID,omitempty"`

	//GreaterThanModifiedDate time.Time `xml:"GreaterThanModifiedDate,omitempty"`

	BatchSize int32 `xml:"BatchSize,omitempty"`
}

type GetCustomersResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomersResult"`

	*ApiResponse

	Customers *ArrayOfCustomerResponse `xml:"Customers,omitempty"`
	//Customers []CustomerResponse `xml:"CustomerResponse,omitempty"`

	RecordCount int32 `xml:"RecordCount,omitempty"`
}

type ArrayOfCustomerResponse struct {
	//XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfCustomerResponse"`

	CustomerResponse []*CustomerResponse `xml:"CustomerResponse,omitempty"`
}

type CustomerResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CustomerResponse"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	FirstName string `xml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty"`

	Company string `xml:"Company,omitempty"`

	CustomerType int32 `xml:"CustomerType,omitempty"`

	CustomerStatus int32 `xml:"CustomerStatus,omitempty"`

	Email string `xml:"Email,omitempty"`

	Phone string `xml:"Phone,omitempty"`

	Phone2 string `xml:"Phone2,omitempty"`

	MobilePhone string `xml:"MobilePhone,omitempty"`

	Fax string `xml:"Fax,omitempty"`

	MainAddress1 string `xml:"MainAddress1,omitempty"`

	MainAddress2 string `xml:"MainAddress2,omitempty"`

	MainCity string `xml:"MainCity,omitempty"`

	MainState string `xml:"MainState,omitempty"`

	MainZip string `xml:"MainZip,omitempty"`

	MainCountry string `xml:"MainCountry,omitempty"`

	MainCounty string `xml:"MainCounty,omitempty"`

	MailAddress1 string `xml:"MailAddress1,omitempty"`

	MailAddress2 string `xml:"MailAddress2,omitempty"`

	MailCity string `xml:"MailCity,omitempty"`

	MailState string `xml:"MailState,omitempty"`

	MailZip string `xml:"MailZip,omitempty"`

	MailCountry string `xml:"MailCountry,omitempty"`

	MailCounty string `xml:"MailCounty,omitempty"`

	OtherAddress1 string `xml:"OtherAddress1,omitempty"`

	OtherAddress2 string `xml:"OtherAddress2,omitempty"`

	OtherCity string `xml:"OtherCity,omitempty"`

	OtherState string `xml:"OtherState,omitempty"`

	OtherZip string `xml:"OtherZip,omitempty"`

	OtherCountry string `xml:"OtherCountry,omitempty"`

	OtherCounty string `xml:"OtherCounty,omitempty"`

	LoginName string `xml:"LoginName,omitempty"`

	EnrollerID int32 `xml:"EnrollerID,omitempty"`

	SponsorID int32 `xml:"SponsorID,omitempty"`

	RankID int32 `xml:"RankID,omitempty"`

	//BirthDate time.Time `xml:"BirthDate,omitempty"`

	Field1 string `xml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty"`

	Field11 string `xml:"Field11,omitempty"`

	Field12 string `xml:"Field12,omitempty"`

	Field13 string `xml:"Field13,omitempty"`

	Field14 string `xml:"Field14,omitempty"`

	Field15 string `xml:"Field15,omitempty"`

	//Date1 time.Time `xml:"Date1,omitempty"`

	//Date2 time.Time `xml:"Date2,omitempty"`

	//Date3 time.Time `xml:"Date3,omitempty"`

	//Date4 time.Time `xml:"Date4,omitempty"`

	//Date5 time.Time `xml:"Date5,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	PayableToName string `xml:"PayableToName,omitempty"`

	DefaultWarehouseID int32 `xml:"DefaultWarehouseID,omitempty"`

	PayableType *PayableType `xml:"PayableType,omitempty"`

	CheckThreshold float64 `xml:"CheckThreshold,omitempty"`

	PriceType int32 `xml:"PriceType,omitempty"`

	LanguageID int32 `xml:"LanguageID,omitempty"`

	Gender *Gender `xml:"Gender,omitempty"`

	SalesTaxID string `xml:"SalesTaxID,omitempty"`

	VatRegistration string `xml:"VatRegistration,omitempty"`

	IsSalesTaxExempt bool `xml:"IsSalesTaxExempt,omitempty"`

	IsSubscribedToBroadcasts bool `xml:"IsSubscribedToBroadcasts,omitempty"`

	//CreatedDate time.Time `xml:"CreatedDate,omitempty"`

	//ModifiedDate time.Time `xml:"ModifiedDate,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty"`

	NameSuffix string `xml:"NameSuffix,omitempty"`

	MainAddress3 string `xml:"MainAddress3,omitempty"`

	MailAddress3 string `xml:"MailAddress3,omitempty"`

	OtherAddress3 string `xml:"OtherAddress3,omitempty"`

	BinaryPlacementPreference int32 `xml:"BinaryPlacementPreference,omitempty"`

	UseBinaryHoldingTank bool `xml:"UseBinaryHoldingTank,omitempty"`

	MainAddressVerified bool `xml:"MainAddressVerified,omitempty"`

	MailAddressVerified bool `xml:"MailAddressVerified,omitempty"`

	OtherAddressVerified bool `xml:"OtherAddressVerified,omitempty"`
}

type CreateWarehouseRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateWarehouseRequest"`

	*ApiRequest

	Description string `xml:"Description,omitempty"`

	Address1 string `xml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty"`

	City string `xml:"City,omitempty"`

	State string `xml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty"`

	Currencies *ArrayOfString `xml:"Currencies,omitempty"`
}

type CreateWarehouseResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateWarehouseResponse"`

	*ApiResponse

	Warehouse *WarehouseResponse `xml:"Warehouse,omitempty"`
}

type WarehouseResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ WarehouseResponse"`

	WarehouseID int32 `xml:"WarehouseID,omitempty"`

	Description string `xml:"Description,omitempty"`

	Address1 string `xml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty"`

	City string `xml:"City,omitempty"`

	State string `xml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty"`
}

type GetCustomerNotesRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomerNotesRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`
}

type GetCustomerNotesResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomerNotesResponse"`

	*ApiResponse

	CustomerNotes *ArrayOfCustomerNotesResponse `xml:"CustomerNotes,omitempty"`
}

type ArrayOfCustomerNotesResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfCustomerNotesResponse"`

	CustomerNotesResponse []*CustomerNotesResponse `xml:"CustomerNotesResponse,omitempty"`
}

type CustomerNotesResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CustomerNotesResponse"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	Notes string `xml:"Notes,omitempty"`
}

type AppendCustomerNotesRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ AppendCustomerNotesRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	Notes string `xml:"Notes,omitempty"`
}

type AppendCustomerNotesResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ AppendCustomerNotesResponse"`

	*ApiResponse

	CustomerID int32 `xml:"CustomerID,omitempty"`
}

type GetVolumesRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetVolumesRequest"`

	*ApiRequest

	PeriodType int32 `xml:"PeriodType,omitempty"`

	PeriodID int32 `xml:"PeriodID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty"`
}

type GetVolumesResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetVolumesResponse"`

	*ApiResponse

	Volumes *ArrayOfVolumeResponse `xml:"Volumes,omitempty"`
}

type ArrayOfVolumeResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfVolumeResponse"`

	VolumeResponse []*VolumeResponse `xml:"VolumeResponse,omitempty"`
}

type VolumeResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ VolumeResponse"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	PeriodType int32 `xml:"PeriodType,omitempty"`

	PeriodID int32 `xml:"PeriodID,omitempty"`

	PeriodDescription string `xml:"PeriodDescription,omitempty"`

	Volume1 float64 `xml:"Volume1,omitempty"`

	Volume2 float64 `xml:"Volume2,omitempty"`

	Volume3 float64 `xml:"Volume3,omitempty"`

	Volume4 float64 `xml:"Volume4,omitempty"`

	Volume5 float64 `xml:"Volume5,omitempty"`

	Volume6 float64 `xml:"Volume6,omitempty"`

	Volume7 float64 `xml:"Volume7,omitempty"`

	Volume8 float64 `xml:"Volume8,omitempty"`

	Volume9 float64 `xml:"Volume9,omitempty"`

	Volume10 float64 `xml:"Volume10,omitempty"`

	Volume11 float64 `xml:"Volume11,omitempty"`

	Volume12 float64 `xml:"Volume12,omitempty"`

	Volume13 float64 `xml:"Volume13,omitempty"`

	Volume14 float64 `xml:"Volume14,omitempty"`

	Volume15 float64 `xml:"Volume15,omitempty"`

	Volume16 float64 `xml:"Volume16,omitempty"`

	Volume17 float64 `xml:"Volume17,omitempty"`

	Volume18 float64 `xml:"Volume18,omitempty"`

	Volume19 float64 `xml:"Volume19,omitempty"`

	Volume20 float64 `xml:"Volume20,omitempty"`

	Volume21 float64 `xml:"Volume21,omitempty"`

	Volume22 float64 `xml:"Volume22,omitempty"`

	Volume23 float64 `xml:"Volume23,omitempty"`

	Volume24 float64 `xml:"Volume24,omitempty"`

	Volume25 float64 `xml:"Volume25,omitempty"`

	Volume26 float64 `xml:"Volume26,omitempty"`

	Volume27 float64 `xml:"Volume27,omitempty"`

	Volume28 float64 `xml:"Volume28,omitempty"`

	Volume29 float64 `xml:"Volume29,omitempty"`

	Volume30 float64 `xml:"Volume30,omitempty"`

	Volume31 float64 `xml:"Volume31,omitempty"`

	Volume32 float64 `xml:"Volume32,omitempty"`

	Volume33 float64 `xml:"Volume33,omitempty"`

	Volume34 float64 `xml:"Volume34,omitempty"`

	Volume35 float64 `xml:"Volume35,omitempty"`

	Volume36 float64 `xml:"Volume36,omitempty"`

	Volume37 float64 `xml:"Volume37,omitempty"`

	Volume38 float64 `xml:"Volume38,omitempty"`

	Volume39 float64 `xml:"Volume39,omitempty"`

	Volume40 float64 `xml:"Volume40,omitempty"`

	Volume41 float64 `xml:"Volume41,omitempty"`

	Volume42 float64 `xml:"Volume42,omitempty"`

	Volume43 float64 `xml:"Volume43,omitempty"`

	Volume44 float64 `xml:"Volume44,omitempty"`

	Volume45 float64 `xml:"Volume45,omitempty"`

	Volume46 float64 `xml:"Volume46,omitempty"`

	Volume47 float64 `xml:"Volume47,omitempty"`

	Volume48 float64 `xml:"Volume48,omitempty"`

	Volume49 float64 `xml:"Volume49,omitempty"`

	Volume50 float64 `xml:"Volume50,omitempty"`

	Volume51 float64 `xml:"Volume51,omitempty"`

	Volume52 float64 `xml:"Volume52,omitempty"`

	Volume53 float64 `xml:"Volume53,omitempty"`

	Volume54 float64 `xml:"Volume54,omitempty"`

	Volume55 float64 `xml:"Volume55,omitempty"`

	Volume56 float64 `xml:"Volume56,omitempty"`

	Volume57 float64 `xml:"Volume57,omitempty"`

	Volume58 float64 `xml:"Volume58,omitempty"`

	Volume59 float64 `xml:"Volume59,omitempty"`

	Volume60 float64 `xml:"Volume60,omitempty"`

	Volume61 float64 `xml:"Volume61,omitempty"`

	Volume62 float64 `xml:"Volume62,omitempty"`

	Volume63 float64 `xml:"Volume63,omitempty"`

	Volume64 float64 `xml:"Volume64,omitempty"`

	Volume65 float64 `xml:"Volume65,omitempty"`

	Volume66 float64 `xml:"Volume66,omitempty"`

	Volume67 float64 `xml:"Volume67,omitempty"`

	Volume68 float64 `xml:"Volume68,omitempty"`

	Volume69 float64 `xml:"Volume69,omitempty"`

	Volume70 float64 `xml:"Volume70,omitempty"`

	Volume71 float64 `xml:"Volume71,omitempty"`

	Volume72 float64 `xml:"Volume72,omitempty"`

	Volume73 float64 `xml:"Volume73,omitempty"`

	Volume74 float64 `xml:"Volume74,omitempty"`

	Volume75 float64 `xml:"Volume75,omitempty"`

	Volume76 float64 `xml:"Volume76,omitempty"`

	Volume77 float64 `xml:"Volume77,omitempty"`

	Volume78 float64 `xml:"Volume78,omitempty"`

	Volume79 float64 `xml:"Volume79,omitempty"`

	Volume80 float64 `xml:"Volume80,omitempty"`

	Volume81 float64 `xml:"Volume81,omitempty"`

	Volume82 float64 `xml:"Volume82,omitempty"`

	Volume83 float64 `xml:"Volume83,omitempty"`

	Volume84 float64 `xml:"Volume84,omitempty"`

	Volume85 float64 `xml:"Volume85,omitempty"`

	Volume86 float64 `xml:"Volume86,omitempty"`

	Volume87 float64 `xml:"Volume87,omitempty"`

	Volume88 float64 `xml:"Volume88,omitempty"`

	Volume89 float64 `xml:"Volume89,omitempty"`

	Volume90 float64 `xml:"Volume90,omitempty"`

	Volume91 float64 `xml:"Volume91,omitempty"`

	Volume92 float64 `xml:"Volume92,omitempty"`

	Volume93 float64 `xml:"Volume93,omitempty"`

	Volume94 float64 `xml:"Volume94,omitempty"`

	Volume95 float64 `xml:"Volume95,omitempty"`

	Volume96 float64 `xml:"Volume96,omitempty"`

	Volume97 float64 `xml:"Volume97,omitempty"`

	Volume98 float64 `xml:"Volume98,omitempty"`

	Volume99 float64 `xml:"Volume99,omitempty"`

	Volume100 float64 `xml:"Volume100,omitempty"`

	RankID int32 `xml:"RankID,omitempty"`

	PaidRankID int32 `xml:"PaidRankID,omitempty"`

	Volume101 float64 `xml:"Volume101,omitempty"`

	Volume102 float64 `xml:"Volume102,omitempty"`

	Volume103 float64 `xml:"Volume103,omitempty"`

	Volume104 float64 `xml:"Volume104,omitempty"`

	Volume105 float64 `xml:"Volume105,omitempty"`

	Volume106 float64 `xml:"Volume106,omitempty"`

	Volume107 float64 `xml:"Volume107,omitempty"`

	Volume108 float64 `xml:"Volume108,omitempty"`

	Volume109 float64 `xml:"Volume109,omitempty"`

	Volume110 float64 `xml:"Volume110,omitempty"`

	Volume111 float64 `xml:"Volume111,omitempty"`

	Volume112 float64 `xml:"Volume112,omitempty"`

	Volume113 float64 `xml:"Volume113,omitempty"`

	Volume114 float64 `xml:"Volume114,omitempty"`

	Volume115 float64 `xml:"Volume115,omitempty"`

	Volume116 float64 `xml:"Volume116,omitempty"`

	Volume117 float64 `xml:"Volume117,omitempty"`

	Volume118 float64 `xml:"Volume118,omitempty"`

	Volume119 float64 `xml:"Volume119,omitempty"`

	Volume120 float64 `xml:"Volume120,omitempty"`

	Volume121 float64 `xml:"Volume121,omitempty"`

	Volume122 float64 `xml:"Volume122,omitempty"`

	Volume123 float64 `xml:"Volume123,omitempty"`

	Volume124 float64 `xml:"Volume124,omitempty"`

	Volume125 float64 `xml:"Volume125,omitempty"`

	Volume126 float64 `xml:"Volume126,omitempty"`

	Volume127 float64 `xml:"Volume127,omitempty"`

	Volume128 float64 `xml:"Volume128,omitempty"`

	Volume129 float64 `xml:"Volume129,omitempty"`

	Volume130 float64 `xml:"Volume130,omitempty"`

	Volume131 float64 `xml:"Volume131,omitempty"`

	Volume132 float64 `xml:"Volume132,omitempty"`

	Volume133 float64 `xml:"Volume133,omitempty"`

	Volume134 float64 `xml:"Volume134,omitempty"`

	Volume135 float64 `xml:"Volume135,omitempty"`

	Volume136 float64 `xml:"Volume136,omitempty"`

	Volume137 float64 `xml:"Volume137,omitempty"`

	Volume138 float64 `xml:"Volume138,omitempty"`

	Volume139 float64 `xml:"Volume139,omitempty"`

	Volume140 float64 `xml:"Volume140,omitempty"`

	Volume141 float64 `xml:"Volume141,omitempty"`

	Volume142 float64 `xml:"Volume142,omitempty"`

	Volume143 float64 `xml:"Volume143,omitempty"`

	Volume144 float64 `xml:"Volume144,omitempty"`

	Volume145 float64 `xml:"Volume145,omitempty"`

	Volume146 float64 `xml:"Volume146,omitempty"`

	Volume147 float64 `xml:"Volume147,omitempty"`

	Volume148 float64 `xml:"Volume148,omitempty"`

	Volume149 float64 `xml:"Volume149,omitempty"`

	Volume150 float64 `xml:"Volume150,omitempty"`

	Volume151 float64 `xml:"Volume151,omitempty"`

	Volume152 float64 `xml:"Volume152,omitempty"`

	Volume153 float64 `xml:"Volume153,omitempty"`

	Volume154 float64 `xml:"Volume154,omitempty"`

	Volume155 float64 `xml:"Volume155,omitempty"`

	Volume156 float64 `xml:"Volume156,omitempty"`

	Volume157 float64 `xml:"Volume157,omitempty"`

	Volume158 float64 `xml:"Volume158,omitempty"`

	Volume159 float64 `xml:"Volume159,omitempty"`

	Volume160 float64 `xml:"Volume160,omitempty"`

	Volume161 float64 `xml:"Volume161,omitempty"`

	Volume162 float64 `xml:"Volume162,omitempty"`

	Volume163 float64 `xml:"Volume163,omitempty"`

	Volume164 float64 `xml:"Volume164,omitempty"`

	Volume165 float64 `xml:"Volume165,omitempty"`

	Volume166 float64 `xml:"Volume166,omitempty"`

	Volume167 float64 `xml:"Volume167,omitempty"`

	Volume168 float64 `xml:"Volume168,omitempty"`

	Volume169 float64 `xml:"Volume169,omitempty"`

	Volume170 float64 `xml:"Volume170,omitempty"`

	Volume171 float64 `xml:"Volume171,omitempty"`

	Volume172 float64 `xml:"Volume172,omitempty"`

	Volume173 float64 `xml:"Volume173,omitempty"`

	Volume174 float64 `xml:"Volume174,omitempty"`

	Volume175 float64 `xml:"Volume175,omitempty"`

	Volume176 float64 `xml:"Volume176,omitempty"`

	Volume177 float64 `xml:"Volume177,omitempty"`

	Volume178 float64 `xml:"Volume178,omitempty"`

	Volume179 float64 `xml:"Volume179,omitempty"`

	Volume180 float64 `xml:"Volume180,omitempty"`

	Volume181 float64 `xml:"Volume181,omitempty"`

	Volume182 float64 `xml:"Volume182,omitempty"`

	Volume183 float64 `xml:"Volume183,omitempty"`

	Volume184 float64 `xml:"Volume184,omitempty"`

	Volume185 float64 `xml:"Volume185,omitempty"`

	Volume186 float64 `xml:"Volume186,omitempty"`

	Volume187 float64 `xml:"Volume187,omitempty"`

	Volume188 float64 `xml:"Volume188,omitempty"`

	Volume189 float64 `xml:"Volume189,omitempty"`

	Volume190 float64 `xml:"Volume190,omitempty"`

	Volume191 float64 `xml:"Volume191,omitempty"`

	Volume192 float64 `xml:"Volume192,omitempty"`

	Volume193 float64 `xml:"Volume193,omitempty"`

	Volume194 float64 `xml:"Volume194,omitempty"`

	Volume195 float64 `xml:"Volume195,omitempty"`

	Volume196 float64 `xml:"Volume196,omitempty"`

	Volume197 float64 `xml:"Volume197,omitempty"`

	Volume198 float64 `xml:"Volume198,omitempty"`

	Volume199 float64 `xml:"Volume199,omitempty"`

	Volume200 float64 `xml:"Volume200,omitempty"`
}

type GetRealTimeCommissionsRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetRealTimeCommissionsRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`
}

type GetRealTimeCommissionsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetRealTimeCommissionsResponse"`

	*ApiResponse

	Commissions *ArrayOfCommissionResponse `xml:"Commissions,omitempty"`
}

type ArrayOfCommissionResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfCommissionResponse"`

	CommissionResponse []*CommissionResponse `xml:"CommissionResponse,omitempty"`
}

type CommissionResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CommissionResponse"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	PeriodType int32 `xml:"PeriodType,omitempty"`

	PeriodID int32 `xml:"PeriodID,omitempty"`

	PeriodDescription string `xml:"PeriodDescription,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	CommissionTotal float64 `xml:"CommissionTotal,omitempty"`

	Bonuses *ArrayOfCommissionBonusResponse `xml:"Bonuses,omitempty"`
}

type ArrayOfCommissionBonusResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfCommissionBonusResponse"`

	CommissionBonusResponse []*CommissionBonusResponse `xml:"CommissionBonusResponse,omitempty"`
}

type CommissionBonusResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CommissionBonusResponse"`

	Description string `xml:"Description,omitempty"`

	Amount float64 `xml:"Amount,omitempty"`

	BonusID int32 `xml:"BonusID,omitempty"`
}

type GetRealTimeCommissionDetailRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetRealTimeCommissionDetailRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	PeriodType int32 `xml:"PeriodType,omitempty"`

	PeriodID int32 `xml:"PeriodID,omitempty"`

	BonusID int32 `xml:"BonusID,omitempty"`
}

type GetRealTimeCommissionDetailResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetRealTimeCommissionDetailResponse"`

	*ApiResponse

	CommissionDetails *ArrayOfCommissionDetailResponse `xml:"CommissionDetails,omitempty"`
}

type ArrayOfCommissionDetailResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfCommissionDetailResponse"`

	CommissionDetailResponse []*CommissionDetailResponse `xml:"CommissionDetailResponse,omitempty"`
}

type CommissionDetailResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CommissionDetailResponse"`

	FromCustomerID int32 `xml:"FromCustomerID,omitempty"`

	FromCustomerName string `xml:"FromCustomerName,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty"`

	Level int32 `xml:"Level,omitempty"`

	PaidLevel int32 `xml:"PaidLevel,omitempty"`

	SourceAmount float64 `xml:"SourceAmount,omitempty"`

	Percentage float64 `xml:"Percentage,omitempty"`

	CommissionAmount float64 `xml:"CommissionAmount,omitempty"`
}

type GetRankQualificationsRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetRankQualificationsRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	RankID int32 `xml:"RankID,omitempty"`

	PeriodType int32 `xml:"PeriodType,omitempty"`

	PeriodID int32 `xml:"PeriodID,omitempty"`

	CultureCode string `xml:"CultureCode,omitempty"`

	RankGroupID int32 `xml:"RankGroupID,omitempty"`
}

type GetRankQualificationsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetRankQualificationsResponse"`

	*ApiResponse

	CustomerID int32 `xml:"CustomerID,omitempty"`

	RankID int32 `xml:"RankID,omitempty"`

	RankDescription string `xml:"RankDescription,omitempty"`

	Qualifies bool `xml:"Qualifies,omitempty"`

	QualifiesOverride bool `xml:"QualifiesOverride,omitempty"`

	PayeeQualificationLegs *ArrayOfArrayOfQualificationResponse `xml:"PayeeQualificationLegs,omitempty"`

	BackRankID int32 `xml:"BackRankID,omitempty"`

	BackRankDescription string `xml:"BackRankDescription,omitempty"`

	NextRankID int32 `xml:"NextRankID,omitempty"`

	NextRankDescription string `xml:"NextRankDescription,omitempty"`

	Score float64 `xml:"Score,omitempty"`
}

type ArrayOfArrayOfQualificationResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfArrayOfQualificationResponse"`

	ArrayOfQualificationResponse []*ArrayOfQualificationResponse `xml:"ArrayOfQualificationResponse,omitempty"`
}

type ArrayOfQualificationResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfQualificationResponse"`

	QualificationResponse []*QualificationResponse `xml:"QualificationResponse,omitempty"`
}

type QualificationResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ QualificationResponse"`

	QualificationDescription string `xml:"QualificationDescription,omitempty"`

	Required string `xml:"Required,omitempty"`

	Actual string `xml:"Actual,omitempty"`

	Qualifies bool `xml:"Qualifies,omitempty"`

	QualifiesOverride bool `xml:"QualifiesOverride,omitempty"`

	SupportingTable struct {
	} `xml:"SupportingTable,omitempty"`

	Completed float64 `xml:"Completed,omitempty"`

	Weight float64 `xml:"Weight,omitempty"`

	Score float64 `xml:"Score,omitempty"`
}

type GetQualificationOverridesRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetQualificationOverridesRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	OverrideID int32 `xml:"OverrideID,omitempty"`

	BatchSize int32 `xml:"BatchSize,omitempty"`
}

type GetQualitificationOverridesResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetQualitificationOverridesResponse"`

	*ApiResponse

	QualificationOverrides *ArrayOfGetQualificationOverrideResponse `xml:"QualificationOverrides,omitempty"`
}

type ArrayOfGetQualificationOverrideResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfGetQualificationOverrideResponse"`

	GetQualificationOverrideResponse []*GetQualificationOverrideResponse `xml:"GetQualificationOverrideResponse,omitempty"`
}

type GetQualificationOverrideResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetQualificationOverrideResponse"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	OverrideID int32 `xml:"OverrideID,omitempty"`

	Qualifies bool `xml:"Qualifies,omitempty"`

	PeriodType int32 `xml:"PeriodType,omitempty"`

	StartPeriodID int32 `xml:"StartPeriodID,omitempty"`

	EndPeriodID int32 `xml:"EndPeriodID,omitempty"`

	ModifiedDate time.Time `xml:"ModifiedDate,omitempty"`

	ModifiedBy string `xml:"ModifiedBy,omitempty"`

	Amount float64 `xml:"Amount,omitempty"`
}

type SetQualificationOverrideRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetQualificationOverrideRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	OverrideID int32 `xml:"OverrideID,omitempty"`

	Qualifies bool `xml:"Qualifies,omitempty"`

	PeriodType int32 `xml:"PeriodType,omitempty"`

	StartPeriodID int32 `xml:"StartPeriodID,omitempty"`

	EndPeriodID int32 `xml:"EndPeriodID,omitempty"`

	Amount float64 `xml:"Amount,omitempty"`
}

type SetQualificationOverrideResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetQualificationOverrideResponse"`

	*ApiResponse
}

type DeleteQualificationOverrideRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteQualificationOverrideRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	OverrideID int32 `xml:"OverrideID,omitempty"`
}

type DeleteQualificationOverrideResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteQualificationOverrideResponse"`

	*ApiResponse
}

type AdjustInventoryRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ AdjustInventoryRequest"`

	*ApiRequest

	WarehouseID int32 `xml:"WarehouseID,omitempty"`

	ItemCode string `xml:"ItemCode,omitempty"`

	Quantity int32 `xml:"Quantity,omitempty"`

	Notes string `xml:"Notes,omitempty"`
}

type AdjustInventoryResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ AdjustInventoryResponse"`

	*ApiResponse
}

type GetCustomerSiteRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomerSiteRequest"`

	*ApiRequest

	WebAlias string `xml:"WebAlias,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty"`
}

type GetCustomerSiteResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomerSiteResponse"`

	*ApiResponse

	CustomerID int32 `xml:"CustomerID,omitempty"`

	WebAlias string `xml:"WebAlias,omitempty"`

	FirstName string `xml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty"`

	Company string `xml:"Company,omitempty"`

	Address1 string `xml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty"`

	City string `xml:"City,omitempty"`

	State string `xml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty"`

	Email string `xml:"Email,omitempty"`

	Phone string `xml:"Phone,omitempty"`

	Phone2 string `xml:"Phone2,omitempty"`

	Fax string `xml:"Fax,omitempty"`

	Notes1 string `xml:"Notes1,omitempty"`

	Notes2 string `xml:"Notes2,omitempty"`

	Notes3 string `xml:"Notes3,omitempty"`

	Notes4 string `xml:"Notes4,omitempty"`

	Url1 string `xml:"Url1,omitempty"`

	Url2 string `xml:"Url2,omitempty"`

	Url3 string `xml:"Url3,omitempty"`

	Url4 string `xml:"Url4,omitempty"`

	Url5 string `xml:"Url5,omitempty"`

	Url6 string `xml:"Url6,omitempty"`

	Url7 string `xml:"Url7,omitempty"`

	Url8 string `xml:"Url8,omitempty"`

	Url9 string `xml:"Url9,omitempty"`

	Url10 string `xml:"Url10,omitempty"`

	Url1Description string `xml:"Url1Description,omitempty"`

	Url2Description string `xml:"Url2Description,omitempty"`

	Url3Description string `xml:"Url3Description,omitempty"`

	Url4Description string `xml:"Url4Description,omitempty"`

	Url5Description string `xml:"Url5Description,omitempty"`

	Url6Description string `xml:"Url6Description,omitempty"`

	Url7Description string `xml:"Url7Description,omitempty"`

	Url8Description string `xml:"Url8Description,omitempty"`

	Url9Description string `xml:"Url9Description,omitempty"`

	Url10Description string `xml:"Url10Description,omitempty"`

	Image1 string `xml:"Image1,omitempty"`

	Image2 string `xml:"Image2,omitempty"`

	ImageUrl1 string `xml:"ImageUrl1,omitempty"`

	ImageUrl2 string `xml:"ImageUrl2,omitempty"`
}

type GetCustomerExtendedRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomerExtendedRequest"`

	*ApiRequest

	ExtendedGroupID int32 `xml:"ExtendedGroupID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	CustomerExtendedID int32 `xml:"CustomerExtendedID,omitempty"`

	Field1 string `xml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty"`

	Field11 string `xml:"Field11,omitempty"`

	Field12 string `xml:"Field12,omitempty"`

	Field13 string `xml:"Field13,omitempty"`

	Field14 string `xml:"Field14,omitempty"`

	Field15 string `xml:"Field15,omitempty"`

	Field16 string `xml:"Field16,omitempty"`

	Field17 string `xml:"Field17,omitempty"`

	Field18 string `xml:"Field18,omitempty"`

	Field19 string `xml:"Field19,omitempty"`

	Field20 string `xml:"Field20,omitempty"`

	Field21 string `xml:"Field21,omitempty"`

	Field22 string `xml:"Field22,omitempty"`

	Field23 string `xml:"Field23,omitempty"`

	Field24 string `xml:"Field24,omitempty"`

	Field25 string `xml:"Field25,omitempty"`

	Field26 string `xml:"Field26,omitempty"`

	Field27 string `xml:"Field27,omitempty"`

	Field28 string `xml:"Field28,omitempty"`

	Field29 string `xml:"Field29,omitempty"`

	Field30 string `xml:"Field30,omitempty"`

	GreaterThanCustomerExtendedID int32 `xml:"GreaterThanCustomerExtendedID,omitempty"`

	GreaterThanModifiedDate time.Time `xml:"GreaterThanModifiedDate,omitempty"`

	BatchSize int32 `xml:"BatchSize,omitempty"`
}

type GetCustomerExtendedResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomerExtendedResponse"`

	*ApiResponse

	Items *ArrayOfCustomerExtendedResponse `xml:"Items,omitempty"`
}

type ArrayOfCustomerExtendedResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfCustomerExtendedResponse"`

	CustomerExtendedResponse []*CustomerExtendedResponse `xml:"CustomerExtendedResponse,omitempty"`
}

type CustomerExtendedResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CustomerExtendedResponse"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	ExtendedGroupID int32 `xml:"ExtendedGroupID,omitempty"`

	CustomerExtendedID int32 `xml:"CustomerExtendedID,omitempty"`

	Field1 string `xml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty"`

	Field11 string `xml:"Field11,omitempty"`

	Field12 string `xml:"Field12,omitempty"`

	Field13 string `xml:"Field13,omitempty"`

	Field14 string `xml:"Field14,omitempty"`

	Field15 string `xml:"Field15,omitempty"`

	Field16 string `xml:"Field16,omitempty"`

	Field17 string `xml:"Field17,omitempty"`

	Field18 string `xml:"Field18,omitempty"`

	Field19 string `xml:"Field19,omitempty"`

	Field20 string `xml:"Field20,omitempty"`

	Field21 string `xml:"Field21,omitempty"`

	Field22 string `xml:"Field22,omitempty"`

	Field23 string `xml:"Field23,omitempty"`

	Field24 string `xml:"Field24,omitempty"`

	Field25 string `xml:"Field25,omitempty"`

	Field26 string `xml:"Field26,omitempty"`

	Field27 string `xml:"Field27,omitempty"`

	Field28 string `xml:"Field28,omitempty"`

	Field29 string `xml:"Field29,omitempty"`

	Field30 string `xml:"Field30,omitempty"`

	ModifiedDate time.Time `xml:"ModifiedDate,omitempty"`
}

type GetCustomerBillingRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomerBillingRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`
}

type GetCustomerBillingResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomerBillingResponse"`

	*ApiResponse

	PrimaryCreditCard *CreditCardAccountResponse `xml:"PrimaryCreditCard,omitempty"`

	SecondaryCreditCard *CreditCardAccountResponse `xml:"SecondaryCreditCard,omitempty"`

	BankAccount *BankAccountResponse `xml:"BankAccount,omitempty"`

	PrimaryWalletAccount *WalletAccountResponse `xml:"PrimaryWalletAccount,omitempty"`

	SecondaryWallletAccount *WalletAccountResponse `xml:"SecondaryWallletAccount,omitempty"`
}

type CreditCardAccountResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreditCardAccountResponse"`

	CreditCardNumberDisplay string `xml:"CreditCardNumberDisplay,omitempty"`

	ExpirationMonth int32 `xml:"ExpirationMonth,omitempty"`

	ExpirationYear int32 `xml:"ExpirationYear,omitempty"`

	CreditCardType int32 `xml:"CreditCardType,omitempty"`

	CreditCardTypeDescription string `xml:"CreditCardTypeDescription,omitempty"`

	BillingName string `xml:"BillingName,omitempty"`

	BillingAddress string `xml:"BillingAddress,omitempty"`

	BillingCity string `xml:"BillingCity,omitempty"`

	BillingState string `xml:"BillingState,omitempty"`

	BillingZip string `xml:"BillingZip,omitempty"`

	BillingCountry string `xml:"BillingCountry,omitempty"`

	BillingAddress2 string `xml:"BillingAddress2,omitempty"`
}

type BankAccountResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ BankAccountResponse"`

	BankAccountNumberDisplay string `xml:"BankAccountNumberDisplay,omitempty"`

	BankRoutingNumber string `xml:"BankRoutingNumber,omitempty"`

	BankName string `xml:"BankName,omitempty"`

	BankAccountType *BankAccountType `xml:"BankAccountType,omitempty"`

	NameOnAccount string `xml:"NameOnAccount,omitempty"`

	BillingAddress string `xml:"BillingAddress,omitempty"`

	BillingCity string `xml:"BillingCity,omitempty"`

	BillingState string `xml:"BillingState,omitempty"`

	BillingZip string `xml:"BillingZip,omitempty"`

	BillingCountry string `xml:"BillingCountry,omitempty"`

	BillingAddress2 string `xml:"BillingAddress2,omitempty"`
}

type WalletAccountResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ WalletAccountResponse"`

	WalletType int32 `xml:"WalletType,omitempty"`

	WalletAccountDisplay string `xml:"WalletAccountDisplay,omitempty"`
}

type GetPaymentsRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetPaymentsRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty"`
}

type GetPaymentsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetPaymentsResponse"`

	*ApiResponse

	Payments *ArrayOfPaymentResponse `xml:"Payments,omitempty"`
}

type ArrayOfPaymentResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfPaymentResponse"`

	PaymentResponse []*PaymentResponse `xml:"PaymentResponse,omitempty"`
}

type PaymentResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ PaymentResponse"`

	PaymentID int32 `xml:"PaymentID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	PaymentType *PaymentType `xml:"PaymentType,omitempty"`

	PaymentDate time.Time `xml:"PaymentDate,omitempty"`

	Amount float64 `xml:"Amount,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	BillingName string `xml:"BillingName,omitempty"`

	BillingAddress1 string `xml:"BillingAddress1,omitempty"`

	BillingAddress2 string `xml:"BillingAddress2,omitempty"`

	BillingCity string `xml:"BillingCity,omitempty"`

	BillingState string `xml:"BillingState,omitempty"`

	BillingZip string `xml:"BillingZip,omitempty"`

	BillingCountry string `xml:"BillingCountry,omitempty"`

	BankName string `xml:"BankName,omitempty"`

	Memo string `xml:"Memo,omitempty"`

	CreditCardNumberDisplay string `xml:"CreditCardNumberDisplay,omitempty"`

	AuthorizationCode string `xml:"AuthorizationCode,omitempty"`

	CreditCardType int32 `xml:"CreditCardType,omitempty"`

	CreditCardTypeDescription string `xml:"CreditCardTypeDescription,omitempty"`
}

type FundPaymentCardRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ FundPaymentCardRequest"`

	*BaseCreatePayoutRequest

	PaymentCardTypeID int32 `xml:"PaymentCardTypeID,omitempty"`

	BillIDList *ArrayOfInt `xml:"BillIDList,omitempty"`
}

type FundPaymentCardResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ FundPaymentCardResponse"`

	*CreatePayoutResponse
}

type CreatePaymentPointAccountRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreatePaymentPointAccountRequest"`

	*BaseCreatePaymentRequest

	OrderID int32 `xml:"OrderID,omitempty"`

	PointAccountID int32 `xml:"PointAccountID,omitempty"`

	PaymentDate time.Time `xml:"PaymentDate,omitempty"`

	Amount float64 `xml:"Amount,omitempty"`

	Memo string `xml:"Memo,omitempty"`
}

type CreatePaymentPointAccountResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreatePaymentPointAccountResponse"`

	*BaseCreatePaymentResponse

	PaymentID int32 `xml:"PaymentID,omitempty"`
}

type CreatePaymentCheckRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreatePaymentCheckRequest"`

	*BaseCreatePaymentRequest

	OrderID int32 `xml:"OrderID,omitempty"`

	PaymentDate time.Time `xml:"PaymentDate,omitempty"`

	Amount float64 `xml:"Amount,omitempty"`

	Memo string `xml:"Memo,omitempty"`

	CheckNumber string `xml:"CheckNumber,omitempty"`

	CheckAccountNumber string `xml:"CheckAccountNumber,omitempty"`

	CheckRoutingNumber string `xml:"CheckRoutingNumber,omitempty"`

	CheckDate time.Time `xml:"CheckDate,omitempty"`
}

type CreatePaymentCheckResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreatePaymentCheckResponse"`

	*BaseCreatePaymentResponse

	PaymentID int32 `xml:"PaymentID,omitempty"`
}

type GetCustomReportRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomReportRequest"`

	*ApiRequest

	ReportID int32 `xml:"ReportID,omitempty"`

	Parameters *ArrayOfParameterRequest `xml:"Parameters,omitempty"`
}

type ArrayOfParameterRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfParameterRequest"`

	ParameterRequest []*ParameterRequest `xml:"ParameterRequest,omitempty"`
}

type ParameterRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ParameterRequest"`

	ParameterName string `xml:"ParameterName,omitempty"`

	Value struct {
	} `xml:"Value,omitempty"`
}

type GetCustomReportResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCustomReportResponse"`

	*ApiResponse

	ReportData struct {
		Schema *Schema `xml:"schema,omitempty"`
	} `xml:"ReportData,omitempty"`
}

type GetReportRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetReportRequest"`

	*ApiRequest

	ReportID int32 `xml:"ReportID,omitempty"`

	Parameters *ArrayOfReportParameterRequest `xml:"Parameters,omitempty"`
}

type ArrayOfReportParameterRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfReportParameterRequest"`

	ReportParameterRequest []*ReportParameterRequest `xml:"ReportParameterRequest,omitempty"`
}

type ReportParameterRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ReportParameterRequest"`

	ParameterName string `xml:"ParameterName,omitempty"`

	Value string `xml:"Value,omitempty"`
}

type GetReportResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetReportResponse"`

	*ApiResponse

	ReportData struct {
		Schema *Schema `xml:"schema,omitempty"`
	} `xml:"ReportData,omitempty"`
}

type ChargeWalletAccountResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ChargeWalletAccountResponse"`

	*CreatePaymentResponse

	Amount float64 `xml:"Amount,omitempty"`

	AuthorizationCode string `xml:"AuthorizationCode,omitempty"`
}

type GetAccountDirectDepositRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetAccountDirectDepositRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`
}

type GetAccountDirectDepositResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetAccountDirectDepositResponse"`

	*ApiResponse

	NameOnAccount string `xml:"NameOnAccount,omitempty"`

	BankAccountNumberDisplay string `xml:"BankAccountNumberDisplay,omitempty"`

	BankRoutingNumber string `xml:"BankRoutingNumber,omitempty"`

	DepositAccountType *DepositAccountType `xml:"DepositAccountType,omitempty"`

	BankName string `xml:"BankName,omitempty"`

	BankAddress string `xml:"BankAddress,omitempty"`

	BankCity string `xml:"BankCity,omitempty"`

	BankState string `xml:"BankState,omitempty"`

	BankZip string `xml:"BankZip,omitempty"`

	BankCountry string `xml:"BankCountry,omitempty"`

	BankAccountType *BankAccountType `xml:"BankAccountType,omitempty"`

	Iban string `xml:"Iban,omitempty"`

	SwiftCode string `xml:"SwiftCode,omitempty"`
}

type OptInEmailResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ OptInEmailResponse"`

	*ApiResponse

	RecordsAffected int32 `xml:"RecordsAffected,omitempty"`
}

type OptInSmsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ OptInSmsResponse"`

	*ApiResponse

	RecordsAffected int32 `xml:"RecordsAffected,omitempty"`
}

type SetCustomerSiteImageRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetCustomerSiteImageRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	ImageName string `xml:"ImageName,omitempty"`

	ImageData []byte `xml:"ImageData,omitempty"`

	CustomerSiteImageType *CustomerSiteImageType `xml:"CustomerSiteImageType,omitempty"`
}

type SetCustomerSiteImageResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetCustomerSiteImageResponse"`

	*ApiResponse
}

type LoginCustomerRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ LoginCustomerRequest"`

	*ApiRequest

	LoginName string `xml:"LoginName,omitempty"`

	Password string `xml:"Password,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty"`
}

type LoginCustomerResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ LoginCustomerResponse"`

	*ApiResponse

	SessionID string `xml:"SessionID,omitempty"`
}

type GetLoginSessionRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetLoginSessionRequest"`

	*ApiRequest

	SessionID string `xml:"SessionID,omitempty"`
}

type GetLoginSessionResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetLoginSessionResponse"`

	*ApiResponse

	CustomerID int32 `xml:"CustomerID,omitempty"`

	FirstName string `xml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty"`
}

type AuthenticateCustomerRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ AuthenticateCustomerRequest"`

	*ApiRequest

	LoginName string `xml:"LoginName,omitempty"`

	Password string `xml:"Password,omitempty"`
}

type AuthenticateCustomerResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ AuthenticateCustomerResponse"`

	*ApiResponse

	CustomerID int32 `xml:"CustomerID,omitempty"`

	FirstName string `xml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty"`
}

type AuthenticateUserRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ AuthenticateUserRequest"`

	*ApiRequest

	LoginName string `xml:"LoginName,omitempty"`

	Password string `xml:"Password,omitempty"`
}

type AuthenticateUserResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ AuthenticateUserResponse"`

	*ApiResponse

	FirstName string `xml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty"`
}

type GetUserPermissionsRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetUserPermissionsRequest"`

	*ApiRequest

	LoginName string `xml:"LoginName,omitempty"`
}

type GetUserPermissionsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetUserPermissionsResponse"`

	*ApiResponse

	RestrictToCustomerTypes *ArrayOfInt `xml:"RestrictToCustomerTypes,omitempty"`

	RestrictToCustomerStatuses *ArrayOfInt `xml:"RestrictToCustomerStatuses,omitempty"`

	RestrictToWarehouses *ArrayOfInt `xml:"RestrictToWarehouses,omitempty"`

	RestrictToCountries *ArrayOfString `xml:"RestrictToCountries,omitempty"`

	RestrictToCurrencies *ArrayOfString `xml:"RestrictToCurrencies,omitempty"`

	ViewDeletedCustomers bool `xml:"ViewDeletedCustomers,omitempty"`

	AllowRemoteCheckPrint bool `xml:"AllowRemoteCheckPrint,omitempty"`

	AllowOverrideItemPrice bool `xml:"AllowOverrideItemPrice,omitempty"`

	AllowStatementPrint bool `xml:"AllowStatementPrint,omitempty"`

	DefaultWarehouseID int32 `xml:"DefaultWarehouseID,omitempty"`

	LanguageID int32 `xml:"LanguageID,omitempty"`

	CultureCode string `xml:"CultureCode,omitempty"`
}

type ChangeOrderStatusRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ChangeOrderStatusRequest"`

	*ApiRequest

	OrderID int32 `xml:"OrderID,omitempty"`

	OrderStatus *OrderStatusType `xml:"OrderStatus,omitempty"`
}

type ChangeOrderStatusResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ChangeOrderStatusResponse"`

	*ApiResponse
}

type ChangeAutoOrderStatusRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ChangeAutoOrderStatusRequest"`

	*ApiRequest

	AutoOrderID int32 `xml:"AutoOrderID,omitempty"`

	AutoOrderStatus *AutoOrderStatusType `xml:"AutoOrderStatus,omitempty"`
}

type ChangeAutoOrderStatusResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ChangeAutoOrderStatusResponse"`

	*ApiResponse
}

type GetShipMethodsRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetShipMethodsRequest"`

	*ApiRequest

	WarehouseID int32 `xml:"WarehouseID,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	OrderSubTotal float64 `xml:"OrderSubTotal,omitempty"`

	OrderWieght float64 `xml:"OrderWieght,omitempty"`
}

type GetShipMethodsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetShipMethodsResponse"`

	*ApiResponse

	ShipMethods *ArrayOfShipMethodResponse `xml:"ShipMethods,omitempty"`
}

type GetOrdersRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetOrdersRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty"`

	OrderIDs *ArrayOfInt `xml:"OrderIDs,omitempty"`

	OrderDateStart time.Time `xml:"OrderDateStart,omitempty"`

	OrderDateEnd time.Time `xml:"OrderDateEnd,omitempty"`

	OrderStatus *OrderStatusType `xml:"OrderStatus,omitempty"`

	WarehouseID int32 `xml:"WarehouseID,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	ReturnCustomer bool `xml:"ReturnCustomer,omitempty"`

	ReturnKitDetails bool `xml:"ReturnKitDetails,omitempty"`

	GreaterThanOrderID int32 `xml:"GreaterThanOrderID,omitempty"`

	GreaterThanModifiedDate time.Time `xml:"GreaterThanModifiedDate,omitempty"`

	BatchSize int32 `xml:"BatchSize,omitempty"`

	Other11 string `xml:"Other11,omitempty"`

	Other12 string `xml:"Other12,omitempty"`

	Other13 string `xml:"Other13,omitempty"`

	Other14 string `xml:"Other14,omitempty"`

	Other15 string `xml:"Other15,omitempty"`

	Other16 string `xml:"Other16,omitempty"`

	Other17 string `xml:"Other17,omitempty"`

	Other18 string `xml:"Other18,omitempty"`

	Other19 string `xml:"Other19,omitempty"`

	Other20 string `xml:"Other20,omitempty"`

	PartyID int32 `xml:"PartyID,omitempty"`
}

type OrderResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ OrderResponse"`

	OrderID int32 `xml:"OrderID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	OrderStatus *OrderStatusType `xml:"OrderStatus,omitempty"`

	OrderDate time.Time `xml:"OrderDate,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	WarehouseID int32 `xml:"WarehouseID,omitempty"`

	ShipMethodID int32 `xml:"ShipMethodID,omitempty"`

	PriceType int32 `xml:"PriceType,omitempty"`

	FirstName string `xml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty"`

	Company string `xml:"Company,omitempty"`

	Address1 string `xml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty"`

	Address3 string `xml:"Address3,omitempty"`

	City string `xml:"City,omitempty"`

	State string `xml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty"`

	County string `xml:"County,omitempty"`

	Email string `xml:"Email,omitempty"`

	Phone string `xml:"Phone,omitempty"`

	Notes string `xml:"Notes,omitempty"`

	Total float64 `xml:"Total,omitempty"`

	SubTotal float64 `xml:"SubTotal,omitempty"`

	TaxTotal float64 `xml:"TaxTotal,omitempty"`

	ShippingTotal float64 `xml:"ShippingTotal,omitempty"`

	DiscountTotal float64 `xml:"DiscountTotal,omitempty"`

	DiscountPercent float64 `xml:"DiscountPercent,omitempty"`

	WeightTotal float64 `xml:"WeightTotal,omitempty"`

	BusinessVolumeTotal float64 `xml:"BusinessVolumeTotal,omitempty"`

	CommissionableVolumeTotal float64 `xml:"CommissionableVolumeTotal,omitempty"`

	TrackingNumber1 string `xml:"TrackingNumber1,omitempty"`

	TrackingNumber2 string `xml:"TrackingNumber2,omitempty"`

	TrackingNumber3 string `xml:"TrackingNumber3,omitempty"`

	TrackingNumber4 string `xml:"TrackingNumber4,omitempty"`

	TrackingNumber5 string `xml:"TrackingNumber5,omitempty"`

	Other1Total float64 `xml:"Other1Total,omitempty"`

	Other2Total float64 `xml:"Other2Total,omitempty"`

	Other3Total float64 `xml:"Other3Total,omitempty"`

	Other4Total float64 `xml:"Other4Total,omitempty"`

	Other5Total float64 `xml:"Other5Total,omitempty"`

	Other6Total float64 `xml:"Other6Total,omitempty"`

	Other7Total float64 `xml:"Other7Total,omitempty"`

	Other8Total float64 `xml:"Other8Total,omitempty"`

	Other9Total float64 `xml:"Other9Total,omitempty"`

	Other10Total float64 `xml:"Other10Total,omitempty"`

	ShippingTax float64 `xml:"ShippingTax,omitempty"`

	OrderTax float64 `xml:"OrderTax,omitempty"`

	FedTaxTotal float64 `xml:"FedTaxTotal,omitempty"`

	StateTaxTotal float64 `xml:"StateTaxTotal,omitempty"`

	FedShippingTax float64 `xml:"FedShippingTax,omitempty"`

	StateShippingTax float64 `xml:"StateShippingTax,omitempty"`

	CityShippingTax float64 `xml:"CityShippingTax,omitempty"`

	CityLocalShippingTax float64 `xml:"CityLocalShippingTax,omitempty"`

	CountyShippingTax float64 `xml:"CountyShippingTax,omitempty"`

	CountyLocalShippingTax float64 `xml:"CountyLocalShippingTax,omitempty"`

	Other11 string `xml:"Other11,omitempty"`

	Other12 string `xml:"Other12,omitempty"`

	Other13 string `xml:"Other13,omitempty"`

	Other14 string `xml:"Other14,omitempty"`

	Other15 string `xml:"Other15,omitempty"`

	Other16 string `xml:"Other16,omitempty"`

	Other17 string `xml:"Other17,omitempty"`

	Other18 string `xml:"Other18,omitempty"`

	Other19 string `xml:"Other19,omitempty"`

	Other20 string `xml:"Other20,omitempty"`

	ModifiedDate time.Time `xml:"ModifiedDate,omitempty"`

	OrderType *OrderType `xml:"OrderType,omitempty"`

	ShippedDate time.Time `xml:"ShippedDate,omitempty"`

	CreatedDate time.Time `xml:"CreatedDate,omitempty"`

	CreatedBy string `xml:"CreatedBy,omitempty"`

	ModifiedBy string `xml:"ModifiedBy,omitempty"`

	TaxFedRate float64 `xml:"TaxFedRate,omitempty"`

	TaxStateRate float64 `xml:"TaxStateRate,omitempty"`

	TaxCityRate float64 `xml:"TaxCityRate,omitempty"`

	TaxCityLocalRate float64 `xml:"TaxCityLocalRate,omitempty"`

	TaxCountyRate float64 `xml:"TaxCountyRate,omitempty"`

	TaxCountyLocalRate float64 `xml:"TaxCountyLocalRate,omitempty"`

	TaxManualRate float64 `xml:"TaxManualRate,omitempty"`

	TaxCity string `xml:"TaxCity,omitempty"`

	TaxCounty string `xml:"TaxCounty,omitempty"`

	TaxState string `xml:"TaxState,omitempty"`

	TaxZip string `xml:"TaxZip,omitempty"`

	TaxCountry string `xml:"TaxCountry,omitempty"`

	TaxIsExempt bool `xml:"TaxIsExempt,omitempty"`

	TaxIsOverRide bool `xml:"TaxIsOverRide,omitempty"`

	Details *ArrayOfOrderDetailResponse `xml:"Details,omitempty"`

	Payments *ArrayOfPaymentResponse `xml:"Payments,omitempty"`

	ExpectedPayments *ArrayOfExpectedPaymentResponse `xml:"ExpectedPayments,omitempty"`

	Customer *CustomerResponse `xml:"Customer,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty"`

	NameSuffix string `xml:"NameSuffix,omitempty"`

	AutoOrderID int32 `xml:"AutoOrderID,omitempty"`

	PartyID int32 `xml:"PartyID,omitempty"`

	Reference1 string `xml:"Reference1,omitempty"`

	IsRMA bool `xml:"IsRMA,omitempty"`
}

type ArrayOfExpectedPaymentResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfExpectedPaymentResponse"`

	ExpectedPaymentResponse []*ExpectedPaymentResponse `xml:"ExpectedPaymentResponse,omitempty"`
}

type ExpectedPaymentResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ExpectedPaymentResponse"`

	ExpectedPaymentID int32 `xml:"ExpectedPaymentID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	PaymentType *PaymentType `xml:"PaymentType,omitempty"`

	PaymentDate time.Time `xml:"PaymentDate,omitempty"`

	Amount float64 `xml:"Amount,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	BillingName string `xml:"BillingName,omitempty"`

	AuthorizationCode string `xml:"AuthorizationCode,omitempty"`
}

type ArrayOfOrderResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfOrderResponse"`

	OrderResponse []*OrderResponse `xml:"OrderResponse,omitempty"`
}

type GetOrdersResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetOrdersResponse"`

	*ApiResponse

	Orders *ArrayOfOrderResponse `xml:"Orders,omitempty"`

	RecordCount int32 `xml:"RecordCount,omitempty"`
}

type GetOrderTotalsRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetOrderTotalsRequest"`

	*ApiRequest

	StartDate time.Time `xml:"StartDate,omitempty"`

	EndDate time.Time `xml:"EndDate,omitempty"`
}

type OrderTotalByCurrency struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ OrderTotalByCurrency"`

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	Amount float64 `xml:"Amount,omitempty"`

	Count int32 `xml:"Count,omitempty"`
}

type ArrayOfOrderTotalByCurrency struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfOrderTotalByCurrency"`

	OrderTotalByCurrency []*OrderTotalByCurrency `xml:"OrderTotalByCurrency,omitempty"`
}

type GetOrderTotalsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetOrderTotalsResponse"`

	*ApiResponse

	StartDate time.Time `xml:"StartDate,omitempty"`

	EndDate time.Time `xml:"EndDate,omitempty"`

	AcceptedByCurrency *ArrayOfOrderTotalByCurrency `xml:"AcceptedByCurrency,omitempty"`
}

type GetAutoOrdersRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetAutoOrdersRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	AutoOrderID int32 `xml:"AutoOrderID,omitempty"`

	AutoOrderStatus *AutoOrderStatusType `xml:"AutoOrderStatus,omitempty"`
}

type AutoOrderResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ AutoOrderResponse"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	AutoOrderID int32 `xml:"AutoOrderID,omitempty"`

	AutoOrderStatus *AutoOrderStatusType `xml:"AutoOrderStatus,omitempty"`

	Frequency *FrequencyType `xml:"Frequency,omitempty"`

	StartDate time.Time `xml:"StartDate,omitempty"`

	StopDate time.Time `xml:"StopDate,omitempty"`

	LastRunDate time.Time `xml:"LastRunDate,omitempty"`

	NextRunDate time.Time `xml:"NextRunDate,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	WarehouseID int32 `xml:"WarehouseID,omitempty"`

	ShipMethodID int32 `xml:"ShipMethodID,omitempty"`

	PaymentType *AutoOrderPaymentType `xml:"PaymentType,omitempty"`

	ProcessType *AutoOrderProcessType `xml:"ProcessType,omitempty"`

	FirstName string `xml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty"`

	Company string `xml:"Company,omitempty"`

	Address1 string `xml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty"`

	Address3 string `xml:"Address3,omitempty"`

	City string `xml:"City,omitempty"`

	State string `xml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty"`

	County string `xml:"County,omitempty"`

	Email string `xml:"Email,omitempty"`

	Phone string `xml:"Phone,omitempty"`

	Notes string `xml:"Notes,omitempty"`

	Total float64 `xml:"Total,omitempty"`

	SubTotal float64 `xml:"SubTotal,omitempty"`

	TaxTotal float64 `xml:"TaxTotal,omitempty"`

	ShippingTotal float64 `xml:"ShippingTotal,omitempty"`

	DiscountTotal float64 `xml:"DiscountTotal,omitempty"`

	BusinessVolumeTotal float64 `xml:"BusinessVolumeTotal,omitempty"`

	CommissionableVolumeTotal float64 `xml:"CommissionableVolumeTotal,omitempty"`

	Description string `xml:"Description,omitempty"`

	Other11 string `xml:"Other11,omitempty"`

	Other12 string `xml:"Other12,omitempty"`

	Other13 string `xml:"Other13,omitempty"`

	Other14 string `xml:"Other14,omitempty"`

	Other15 string `xml:"Other15,omitempty"`

	Other16 string `xml:"Other16,omitempty"`

	Other17 string `xml:"Other17,omitempty"`

	Other18 string `xml:"Other18,omitempty"`

	Other19 string `xml:"Other19,omitempty"`

	Other20 string `xml:"Other20,omitempty"`

	Details *ArrayOfAutoOrderDetailResponse `xml:"Details,omitempty"`

	ModifiedDate time.Time `xml:"ModifiedDate,omitempty"`

	ModifiedBy string `xml:"ModifiedBy,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty"`

	NameSuffix string `xml:"NameSuffix,omitempty"`

	SpecificDayInterval int32 `xml:"SpecificDayInterval,omitempty"`
}

type ArrayOfAutoOrderDetailResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfAutoOrderDetailResponse"`

	AutoOrderDetailResponse []*AutoOrderDetailResponse `xml:"AutoOrderDetailResponse,omitempty"`
}

type AutoOrderDetailResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ AutoOrderDetailResponse"`

	ItemCode string `xml:"ItemCode,omitempty"`

	Description string `xml:"Description,omitempty"`

	Quantity float64 `xml:"Quantity,omitempty"`

	PriceEach float64 `xml:"PriceEach,omitempty"`

	PriceTotal float64 `xml:"PriceTotal,omitempty"`

	BusinessVolumeEach float64 `xml:"BusinessVolumeEach,omitempty"`

	BusinesVolume float64 `xml:"BusinesVolume,omitempty"`

	CommissionableVolumeEach float64 `xml:"CommissionableVolumeEach,omitempty"`

	CommissionableVolume float64 `xml:"CommissionableVolume,omitempty"`

	ParentItemCode string `xml:"ParentItemCode,omitempty"`

	PriceEachOverride float64 `xml:"PriceEachOverride,omitempty"`

	TaxableEachOverride float64 `xml:"TaxableEachOverride,omitempty"`

	ShippingPriceEachOverride float64 `xml:"ShippingPriceEachOverride,omitempty"`

	BusinessVolumeEachOverride float64 `xml:"BusinessVolumeEachOverride,omitempty"`

	CommissionableVolumeEachOverride float64 `xml:"CommissionableVolumeEachOverride,omitempty"`

	Reference1 string `xml:"Reference1,omitempty"`

	ProcessWhileDate time.Time `xml:"ProcessWhileDate,omitempty"`

	SkipUntilDate time.Time `xml:"SkipUntilDate,omitempty"`
}

type ArrayOfAutoOrderResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfAutoOrderResponse"`

	AutoOrderResponse []*AutoOrderResponse `xml:"AutoOrderResponse,omitempty"`
}

type GetAutoOrdersResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetAutoOrdersResponse"`

	*ApiResponse

	AutoOrders *ArrayOfAutoOrderResponse `xml:"AutoOrders,omitempty"`
}

type ChangeOrderStatusBatchRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ChangeOrderStatusBatchRequest"`

	*ApiRequest

	OrderStatus *OrderStatusType `xml:"OrderStatus,omitempty"`

	Details *ArrayOfOrderBatchDetailRequest `xml:"Details,omitempty"`
}

type ArrayOfOrderBatchDetailRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfOrderBatchDetailRequest"`

	OrderBatchDetailRequest []*OrderBatchDetailRequest `xml:"OrderBatchDetailRequest,omitempty"`
}

type OrderBatchDetailRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ OrderBatchDetailRequest"`

	OrderID int32 `xml:"OrderID,omitempty"`

	TrackingNumber1 string `xml:"TrackingNumber1,omitempty"`

	TrackingNumber2 string `xml:"TrackingNumber2,omitempty"`

	TrackingNumber3 string `xml:"TrackingNumber3,omitempty"`

	TrackingNumber4 string `xml:"TrackingNumber4,omitempty"`

	TrackingNumber5 string `xml:"TrackingNumber5,omitempty"`

	ShippedDate time.Time `xml:"ShippedDate,omitempty"`
}

type ChangeOrderStatusBatchResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ChangeOrderStatusBatchResponse"`

	*ApiResponse
}

type MergeCustomerRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ MergeCustomerRequest"`

	*ApiRequest

	ToCustomerID int32 `xml:"ToCustomerID,omitempty"`

	FromCustomerID int32 `xml:"FromCustomerID,omitempty"`
}

type MergeCustomerResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ MergeCustomerResponse"`

	*ApiResponse
}

type PlaceEnrollerNodeRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ PlaceEnrollerNodeRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	ToEnrollerID int32 `xml:"ToEnrollerID,omitempty"`

	Reason string `xml:"Reason,omitempty"`
}

type PlaceEnrollerNodeResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ PlaceEnrollerNodeResponse"`

	*ApiResponse
}

type PlaceStackNodeRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ PlaceStackNodeRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	ToParentID int32 `xml:"ToParentID,omitempty"`

	Reason string `xml:"Reason,omitempty"`
}

type PlaceStackNodeResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ PlaceStackNodeResponse"`

	*ApiResponse
}

type PlaceUniLevelNodeRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ PlaceUniLevelNodeRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	ToSponsorID int32 `xml:"ToSponsorID,omitempty"`

	Reason string `xml:"Reason,omitempty"`

	OptionalPlacement int32 `xml:"OptionalPlacement,omitempty"`

	OptionalFindAvailable bool `xml:"OptionalFindAvailable,omitempty"`

	OptionalUnilevelBuildTypeID int32 `xml:"OptionalUnilevelBuildTypeID,omitempty"`
}

type PlaceUniLevelNodeResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ PlaceUniLevelNodeResponse"`

	*ApiResponse
}

type PlaceBinaryNodeRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ PlaceBinaryNodeRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	ToParentID int32 `xml:"ToParentID,omitempty"`

	PlacementType *BinaryPlacementType `xml:"PlacementType,omitempty"`

	Reason string `xml:"Reason,omitempty"`
}

type PlaceBinaryNodeResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ PlaceBinaryNodeResponse"`

	*ApiResponse
}

type GetBinaryPreferenceRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetBinaryPreferenceRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`
}

type GetBinaryPreferenceResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetBinaryPreferenceResponse"`

	*ApiResponse

	CustomerID int32 `xml:"CustomerID,omitempty"`

	PlacementType *BinaryPlacementType `xml:"PlacementType,omitempty"`
}

type SetBinaryPreferenceRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetBinaryPreferenceRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	PlacementType *BinaryPlacementType `xml:"PlacementType,omitempty"`
}

type SetBinaryPreferenceResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetBinaryPreferenceResponse"`

	*ApiResponse
}

type PlaceMatrixNodeRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ PlaceMatrixNodeRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	ToParentCustomerID int32 `xml:"ToParentCustomerID,omitempty"`

	ToParentMatrixID int32 `xml:"ToParentMatrixID,omitempty"`

	Reason string `xml:"Reason,omitempty"`

	Placement int32 `xml:"Placement,omitempty"`
}

type PlaceMatrixNodeResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ PlaceMatrixNodeResponse"`

	*ApiResponse
}

type GetCountryRegionsRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCountryRegionsRequest"`

	*ApiRequest

	CountryCode string `xml:"CountryCode,omitempty"`
}

type RegionResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ RegionResponse"`

	RegionCode string `xml:"RegionCode,omitempty"`

	RegionName string `xml:"RegionName,omitempty"`
}

type CountryResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CountryResponse"`

	CountryCode string `xml:"CountryCode,omitempty"`

	CountryName string `xml:"CountryName,omitempty"`
}

type ArrayOfCountryResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfCountryResponse"`

	CountryResponse []*CountryResponse `xml:"CountryResponse,omitempty"`
}

type ArrayOfRegionResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfRegionResponse"`

	RegionResponse []*RegionResponse `xml:"RegionResponse,omitempty"`
}

type GetCountryRegionsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCountryRegionsResponse"`

	*ApiResponse

	Countries *ArrayOfCountryResponse `xml:"Countries,omitempty"`

	SelectedCountry string `xml:"SelectedCountry,omitempty"`

	Regions *ArrayOfRegionResponse `xml:"Regions,omitempty"`
}

type GetDownlineRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetDownlineRequest"`

	*ApiRequest

	TreeType *TreeType `xml:"TreeType,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	PeriodType int32 `xml:"PeriodType,omitempty"`

	PeriodID int32 `xml:"PeriodID,omitempty"`

	MaxLevelDepth int32 `xml:"MaxLevelDepth,omitempty"`

	CustomerTypes *ArrayOfInt `xml:"CustomerTypes,omitempty"`

	Ranks *ArrayOfInt `xml:"Ranks,omitempty"`

	PayRanks *ArrayOfInt `xml:"PayRanks,omitempty"`

	VolumeFilters *ArrayOfVolumeFilter `xml:"VolumeFilters,omitempty"`

	CustomerStatusTypes *ArrayOfInt `xml:"CustomerStatusTypes,omitempty"`

	BatchSize int32 `xml:"BatchSize,omitempty"`

	SortByLevel bool `xml:"SortByLevel,omitempty"`

	BatchOffset int32 `xml:"BatchOffset,omitempty"`
}

type ArrayOfVolumeFilter struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfVolumeFilter"`

	VolumeFilter []*VolumeFilter `xml:"VolumeFilter,omitempty"`
}

type VolumeFilter struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ VolumeFilter"`

	VolumeID int32 `xml:"VolumeID,omitempty"`

	Compare *NumericCompareType `xml:"Compare,omitempty"`

	Value float64 `xml:"Value,omitempty"`
}

type NodeResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ NodeResponse"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	NodeID int32 `xml:"NodeID,omitempty"`

	ParentID int32 `xml:"ParentID,omitempty"`

	Level int32 `xml:"Level,omitempty"`

	Position int32 `xml:"Position,omitempty"`

	CustomerType int32 `xml:"CustomerType,omitempty"`

	CustomerStatus int32 `xml:"CustomerStatus,omitempty"`

	RankID int32 `xml:"RankID,omitempty"`

	PayRankID int32 `xml:"PayRankID,omitempty"`

	FirstName string `xml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty"`

	Company string `xml:"Company,omitempty"`

	Email string `xml:"Email,omitempty"`

	Phone string `xml:"Phone,omitempty"`

	CreatedDate time.Time `xml:"CreatedDate,omitempty"`

	Volume1 float64 `xml:"Volume1,omitempty"`

	Volume2 float64 `xml:"Volume2,omitempty"`

	Volume3 float64 `xml:"Volume3,omitempty"`

	Volume4 float64 `xml:"Volume4,omitempty"`

	Volume5 float64 `xml:"Volume5,omitempty"`

	Volume6 float64 `xml:"Volume6,omitempty"`

	Volume7 float64 `xml:"Volume7,omitempty"`

	Volume8 float64 `xml:"Volume8,omitempty"`

	Volume9 float64 `xml:"Volume9,omitempty"`

	Volume10 float64 `xml:"Volume10,omitempty"`

	Volume11 float64 `xml:"Volume11,omitempty"`

	Volume12 float64 `xml:"Volume12,omitempty"`

	Volume13 float64 `xml:"Volume13,omitempty"`

	Volume14 float64 `xml:"Volume14,omitempty"`

	Volume15 float64 `xml:"Volume15,omitempty"`

	Volume16 float64 `xml:"Volume16,omitempty"`

	Volume17 float64 `xml:"Volume17,omitempty"`

	Volume18 float64 `xml:"Volume18,omitempty"`

	Volume19 float64 `xml:"Volume19,omitempty"`

	Volume20 float64 `xml:"Volume20,omitempty"`

	Volume21 float64 `xml:"Volume21,omitempty"`

	Volume22 float64 `xml:"Volume22,omitempty"`

	Volume23 float64 `xml:"Volume23,omitempty"`

	Volume24 float64 `xml:"Volume24,omitempty"`

	Volume25 float64 `xml:"Volume25,omitempty"`

	Volume26 float64 `xml:"Volume26,omitempty"`

	Volume27 float64 `xml:"Volume27,omitempty"`

	Volume28 float64 `xml:"Volume28,omitempty"`

	Volume29 float64 `xml:"Volume29,omitempty"`

	Volume30 float64 `xml:"Volume30,omitempty"`

	Volume31 float64 `xml:"Volume31,omitempty"`

	Volume32 float64 `xml:"Volume32,omitempty"`

	Volume33 float64 `xml:"Volume33,omitempty"`

	Volume34 float64 `xml:"Volume34,omitempty"`

	Volume35 float64 `xml:"Volume35,omitempty"`

	Volume36 float64 `xml:"Volume36,omitempty"`

	Volume37 float64 `xml:"Volume37,omitempty"`

	Volume38 float64 `xml:"Volume38,omitempty"`

	Volume39 float64 `xml:"Volume39,omitempty"`

	Volume40 float64 `xml:"Volume40,omitempty"`

	Volume41 float64 `xml:"Volume41,omitempty"`

	Volume42 float64 `xml:"Volume42,omitempty"`

	Volume43 float64 `xml:"Volume43,omitempty"`

	Volume44 float64 `xml:"Volume44,omitempty"`

	Volume45 float64 `xml:"Volume45,omitempty"`

	Volume46 float64 `xml:"Volume46,omitempty"`

	Volume47 float64 `xml:"Volume47,omitempty"`

	Volume48 float64 `xml:"Volume48,omitempty"`

	Volume49 float64 `xml:"Volume49,omitempty"`

	Volume50 float64 `xml:"Volume50,omitempty"`

	Volume51 float64 `xml:"Volume51,omitempty"`

	Volume52 float64 `xml:"Volume52,omitempty"`

	Volume53 float64 `xml:"Volume53,omitempty"`

	Volume54 float64 `xml:"Volume54,omitempty"`

	Volume55 float64 `xml:"Volume55,omitempty"`

	Volume56 float64 `xml:"Volume56,omitempty"`

	Volume57 float64 `xml:"Volume57,omitempty"`

	Volume58 float64 `xml:"Volume58,omitempty"`

	Volume59 float64 `xml:"Volume59,omitempty"`

	Volume60 float64 `xml:"Volume60,omitempty"`

	Volume61 float64 `xml:"Volume61,omitempty"`

	Volume62 float64 `xml:"Volume62,omitempty"`

	Volume63 float64 `xml:"Volume63,omitempty"`

	Volume64 float64 `xml:"Volume64,omitempty"`

	Volume65 float64 `xml:"Volume65,omitempty"`

	Volume66 float64 `xml:"Volume66,omitempty"`

	Volume67 float64 `xml:"Volume67,omitempty"`

	Volume68 float64 `xml:"Volume68,omitempty"`

	Volume69 float64 `xml:"Volume69,omitempty"`

	Volume70 float64 `xml:"Volume70,omitempty"`

	Volume71 float64 `xml:"Volume71,omitempty"`

	Volume72 float64 `xml:"Volume72,omitempty"`

	Volume73 float64 `xml:"Volume73,omitempty"`

	Volume74 float64 `xml:"Volume74,omitempty"`

	Volume75 float64 `xml:"Volume75,omitempty"`

	Volume76 float64 `xml:"Volume76,omitempty"`

	Volume77 float64 `xml:"Volume77,omitempty"`

	Volume78 float64 `xml:"Volume78,omitempty"`

	Volume79 float64 `xml:"Volume79,omitempty"`

	Volume80 float64 `xml:"Volume80,omitempty"`

	Volume81 float64 `xml:"Volume81,omitempty"`

	Volume82 float64 `xml:"Volume82,omitempty"`

	Volume83 float64 `xml:"Volume83,omitempty"`

	Volume84 float64 `xml:"Volume84,omitempty"`

	Volume85 float64 `xml:"Volume85,omitempty"`

	Volume86 float64 `xml:"Volume86,omitempty"`

	Volume87 float64 `xml:"Volume87,omitempty"`

	Volume88 float64 `xml:"Volume88,omitempty"`

	Volume89 float64 `xml:"Volume89,omitempty"`

	Volume90 float64 `xml:"Volume90,omitempty"`

	Volume91 float64 `xml:"Volume91,omitempty"`

	Volume92 float64 `xml:"Volume92,omitempty"`

	Volume93 float64 `xml:"Volume93,omitempty"`

	Volume94 float64 `xml:"Volume94,omitempty"`

	Volume95 float64 `xml:"Volume95,omitempty"`

	Volume96 float64 `xml:"Volume96,omitempty"`

	Volume97 float64 `xml:"Volume97,omitempty"`

	Volume98 float64 `xml:"Volume98,omitempty"`

	Volume99 float64 `xml:"Volume99,omitempty"`

	Volume100 float64 `xml:"Volume100,omitempty"`

	Volume101 float64 `xml:"Volume101,omitempty"`

	Volume102 float64 `xml:"Volume102,omitempty"`

	Volume103 float64 `xml:"Volume103,omitempty"`

	Volume104 float64 `xml:"Volume104,omitempty"`

	Volume105 float64 `xml:"Volume105,omitempty"`

	Volume106 float64 `xml:"Volume106,omitempty"`

	Volume107 float64 `xml:"Volume107,omitempty"`

	Volume108 float64 `xml:"Volume108,omitempty"`

	Volume109 float64 `xml:"Volume109,omitempty"`

	Volume110 float64 `xml:"Volume110,omitempty"`

	Volume111 float64 `xml:"Volume111,omitempty"`

	Volume112 float64 `xml:"Volume112,omitempty"`

	Volume113 float64 `xml:"Volume113,omitempty"`

	Volume114 float64 `xml:"Volume114,omitempty"`

	Volume115 float64 `xml:"Volume115,omitempty"`

	Volume116 float64 `xml:"Volume116,omitempty"`

	Volume117 float64 `xml:"Volume117,omitempty"`

	Volume118 float64 `xml:"Volume118,omitempty"`

	Volume119 float64 `xml:"Volume119,omitempty"`

	Volume120 float64 `xml:"Volume120,omitempty"`

	Volume121 float64 `xml:"Volume121,omitempty"`

	Volume122 float64 `xml:"Volume122,omitempty"`

	Volume123 float64 `xml:"Volume123,omitempty"`

	Volume124 float64 `xml:"Volume124,omitempty"`

	Volume125 float64 `xml:"Volume125,omitempty"`

	Volume126 float64 `xml:"Volume126,omitempty"`

	Volume127 float64 `xml:"Volume127,omitempty"`

	Volume128 float64 `xml:"Volume128,omitempty"`

	Volume129 float64 `xml:"Volume129,omitempty"`

	Volume130 float64 `xml:"Volume130,omitempty"`

	Volume131 float64 `xml:"Volume131,omitempty"`

	Volume132 float64 `xml:"Volume132,omitempty"`

	Volume133 float64 `xml:"Volume133,omitempty"`

	Volume134 float64 `xml:"Volume134,omitempty"`

	Volume135 float64 `xml:"Volume135,omitempty"`

	Volume136 float64 `xml:"Volume136,omitempty"`

	Volume137 float64 `xml:"Volume137,omitempty"`

	Volume138 float64 `xml:"Volume138,omitempty"`

	Volume139 float64 `xml:"Volume139,omitempty"`

	Volume140 float64 `xml:"Volume140,omitempty"`

	Volume141 float64 `xml:"Volume141,omitempty"`

	Volume142 float64 `xml:"Volume142,omitempty"`

	Volume143 float64 `xml:"Volume143,omitempty"`

	Volume144 float64 `xml:"Volume144,omitempty"`

	Volume145 float64 `xml:"Volume145,omitempty"`

	Volume146 float64 `xml:"Volume146,omitempty"`

	Volume147 float64 `xml:"Volume147,omitempty"`

	Volume148 float64 `xml:"Volume148,omitempty"`

	Volume149 float64 `xml:"Volume149,omitempty"`

	Volume150 float64 `xml:"Volume150,omitempty"`

	Volume151 float64 `xml:"Volume151,omitempty"`

	Volume152 float64 `xml:"Volume152,omitempty"`

	Volume153 float64 `xml:"Volume153,omitempty"`

	Volume154 float64 `xml:"Volume154,omitempty"`

	Volume155 float64 `xml:"Volume155,omitempty"`

	Volume156 float64 `xml:"Volume156,omitempty"`

	Volume157 float64 `xml:"Volume157,omitempty"`

	Volume158 float64 `xml:"Volume158,omitempty"`

	Volume159 float64 `xml:"Volume159,omitempty"`

	Volume160 float64 `xml:"Volume160,omitempty"`

	Volume161 float64 `xml:"Volume161,omitempty"`

	Volume162 float64 `xml:"Volume162,omitempty"`

	Volume163 float64 `xml:"Volume163,omitempty"`

	Volume164 float64 `xml:"Volume164,omitempty"`

	Volume165 float64 `xml:"Volume165,omitempty"`

	Volume166 float64 `xml:"Volume166,omitempty"`

	Volume167 float64 `xml:"Volume167,omitempty"`

	Volume168 float64 `xml:"Volume168,omitempty"`

	Volume169 float64 `xml:"Volume169,omitempty"`

	Volume170 float64 `xml:"Volume170,omitempty"`

	Volume171 float64 `xml:"Volume171,omitempty"`

	Volume172 float64 `xml:"Volume172,omitempty"`

	Volume173 float64 `xml:"Volume173,omitempty"`

	Volume174 float64 `xml:"Volume174,omitempty"`

	Volume175 float64 `xml:"Volume175,omitempty"`

	Volume176 float64 `xml:"Volume176,omitempty"`

	Volume177 float64 `xml:"Volume177,omitempty"`

	Volume178 float64 `xml:"Volume178,omitempty"`

	Volume179 float64 `xml:"Volume179,omitempty"`

	Volume180 float64 `xml:"Volume180,omitempty"`

	Volume181 float64 `xml:"Volume181,omitempty"`

	Volume182 float64 `xml:"Volume182,omitempty"`

	Volume183 float64 `xml:"Volume183,omitempty"`

	Volume184 float64 `xml:"Volume184,omitempty"`

	Volume185 float64 `xml:"Volume185,omitempty"`

	Volume186 float64 `xml:"Volume186,omitempty"`

	Volume187 float64 `xml:"Volume187,omitempty"`

	Volume188 float64 `xml:"Volume188,omitempty"`

	Volume189 float64 `xml:"Volume189,omitempty"`

	Volume190 float64 `xml:"Volume190,omitempty"`

	Volume191 float64 `xml:"Volume191,omitempty"`

	Volume192 float64 `xml:"Volume192,omitempty"`

	Volume193 float64 `xml:"Volume193,omitempty"`

	Volume194 float64 `xml:"Volume194,omitempty"`

	Volume195 float64 `xml:"Volume195,omitempty"`

	Volume196 float64 `xml:"Volume196,omitempty"`

	Volume197 float64 `xml:"Volume197,omitempty"`

	Volume198 float64 `xml:"Volume198,omitempty"`

	Volume199 float64 `xml:"Volume199,omitempty"`

	Volume200 float64 `xml:"Volume200,omitempty"`
}

type ArrayOfNodeResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfNodeResponse"`

	NodeResponse []*NodeResponse `xml:"NodeResponse,omitempty"`
}

type GetDownlineResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetDownlineResponse"`

	*ApiResponse

	Nodes *ArrayOfNodeResponse `xml:"Nodes,omitempty"`

	RecordCount int32 `xml:"RecordCount,omitempty"`
}

type GetUplineRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetUplineRequest"`

	*ApiRequest

	TreeType *TreeType `xml:"TreeType,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	PeriodType int32 `xml:"PeriodType,omitempty"`

	PeriodID int32 `xml:"PeriodID,omitempty"`
}

type GetUplineResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetUplineResponse"`

	*ApiResponse

	Nodes *ArrayOfNodeResponse `xml:"Nodes,omitempty"`
}

type DequeueCustomerEventsRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DequeueCustomerEventsRequest"`

	*ApiRequest
}

type CustomerEventResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CustomerEventResponse"`

	EventID int32 `xml:"EventID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	EventDescription string `xml:"EventDescription,omitempty"`

	Fields *ArrayOfCustomerEventField `xml:"Fields,omitempty"`

	EventDate time.Time `xml:"EventDate,omitempty"`
}

type ArrayOfCustomerEventField struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfCustomerEventField"`

	CustomerEventField []*CustomerEventField `xml:"CustomerEventField,omitempty"`
}

type CustomerEventField struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CustomerEventField"`

	Name string `xml:"Name,omitempty"`

	Value int32 `xml:"Value,omitempty"`
}

type ArrayOfCustomerEventResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfCustomerEventResponse"`

	CustomerEventResponse []*CustomerEventResponse `xml:"CustomerEventResponse,omitempty"`
}

type DequeueCustomerEventsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DequeueCustomerEventsResponse"`

	*ApiResponse

	CustomerEvents *ArrayOfCustomerEventResponse `xml:"CustomerEvents,omitempty"`
}

type CreatePointTransactionRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreatePointTransactionRequest"`

	*ApiRequest

	PointAccountID int32 `xml:"PointAccountID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	Amount float64 `xml:"Amount,omitempty"`

	Reference string `xml:"Reference,omitempty"`

	TransactionType *PointTransactionType `xml:"TransactionType,omitempty"`
}

type CreatePointTransactionResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreatePointTransactionResponse"`

	*ApiResponse
}

type GetPointAccountRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetPointAccountRequest"`

	*ApiRequest

	PointAccountID int32 `xml:"PointAccountID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty"`
}

type GetPointAccountResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetPointAccountResponse"`

	*ApiResponse

	Balance float64 `xml:"Balance,omitempty"`
}

type GetSubscriptionRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetSubscriptionRequest"`

	*ApiRequest

	SubscriptionID int32 `xml:"SubscriptionID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty"`
}

type GetSubscriptionResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetSubscriptionResponse"`

	*ApiResponse

	Status *SubscriptionStatus `xml:"Status,omitempty"`

	StartDate time.Time `xml:"StartDate,omitempty"`

	ExpireDate time.Time `xml:"ExpireDate,omitempty"`
}

type ValidateRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ValidateRequest"`

	*ApiRequest
}

type IsEnrollerChildValidateRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ IsEnrollerChildValidateRequest"`

	*ValidateRequest

	ParentID int32 `xml:"ParentID,omitempty"`

	ChildID int32 `xml:"ChildID,omitempty"`
}

type IsTaxIDAvailableValidateRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ IsTaxIDAvailableValidateRequest"`

	*ValidateRequest

	TaxID string `xml:"TaxID,omitempty"`

	TaxTypeID int32 `xml:"TaxTypeID,omitempty"`

	ExcludeCustomerID int32 `xml:"ExcludeCustomerID,omitempty"`
}

type IsMatrixChildValidateRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ IsMatrixChildValidateRequest"`

	*ValidateRequest

	ParentID int32 `xml:"ParentID,omitempty"`

	ChildID int32 `xml:"ChildID,omitempty"`
}

type IsUniLevelChildValidateRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ IsUniLevelChildValidateRequest"`

	*ValidateRequest

	ParentID int32 `xml:"ParentID,omitempty"`

	ChildID int32 `xml:"ChildID,omitempty"`
}

type IsLoginNameAvailableValidateRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ IsLoginNameAvailableValidateRequest"`

	*ValidateRequest

	LoginName string `xml:"LoginName,omitempty"`
}

type ValidateResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ValidateResponse"`

	*ApiResponse

	IsValid bool `xml:"IsValid,omitempty"`
}

type VerifyAddressRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ VerifyAddressRequest"`

	*ApiRequest

	Address string `xml:"Address,omitempty"`

	City string `xml:"City,omitempty"`

	State string `xml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty"`
}

type VerifyAddressResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ VerifyAddressResponse"`

	*ApiResponse

	Address string `xml:"Address,omitempty"`

	City string `xml:"City,omitempty"`

	County string `xml:"County,omitempty"`

	State string `xml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty"`
}

type OptOutEmailRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ OptOutEmailRequest"`

	*ApiRequest

	Email string `xml:"Email,omitempty"`
}

type OptOutEmailResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ OptOutEmailResponse"`

	*ApiResponse

	RecordsAffected int32 `xml:"RecordsAffected,omitempty"`
}

type OptOutSmsRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ OptOutSmsRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	PhoneNumber string `xml:"PhoneNumber,omitempty"`
}

type OptOutSmsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ OptOutSmsResponse"`

	*ApiResponse

	RecordsAffected int32 `xml:"RecordsAffected,omitempty"`
}

type GetShoppingCartRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetShoppingCartRequest"`

	*ApiRequest

	ShoppingID string `xml:"ShoppingID,omitempty"`
}

type GetShoppingCartResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetShoppingCartResponse"`

	*ApiResponse

	ExistingOrderID int32 `xml:"ExistingOrderID,omitempty"`

	ExistingAutoOrderID int32 `xml:"ExistingAutoOrderID,omitempty"`

	Details *ArrayOfOrderDetailResponse `xml:"Details,omitempty"`
}

type GetWarehousesRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetWarehousesRequest"`

	*ApiRequest
}

type ArrayOfWarehouseResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfWarehouseResponse"`

	WarehouseResponse []*WarehouseResponse `xml:"WarehouseResponse,omitempty"`
}

type GetWarehousesResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetWarehousesResponse"`

	*ApiResponse

	Warehouses *ArrayOfWarehouseResponse `xml:"Warehouses,omitempty"`
}

type GetSessionRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetSessionRequest"`

	*ApiRequest

	SessionID string `xml:"SessionID,omitempty"`
}

type GetSessionResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetSessionResponse"`

	*ApiResponse

	SessionData string `xml:"SessionData,omitempty"`
}

type SetSessionRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetSessionRequest"`

	*ApiRequest

	SessionID string `xml:"SessionID,omitempty"`

	SessionData string `xml:"SessionData,omitempty"`
}

type SetSessionResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SetSessionResponse"`

	*ApiResponse
}

type GetItemsRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetItemsRequest"`

	*ApiRequest

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	PriceType int32 `xml:"PriceType,omitempty"`

	WarehouseID int32 `xml:"WarehouseID,omitempty"`

	ItemCodes *ArrayOfString `xml:"ItemCodes,omitempty"`

	WebID int32 `xml:"WebID,omitempty"`

	WebCategoryID int32 `xml:"WebCategoryID,omitempty"`

	ReturnLongDetail bool `xml:"ReturnLongDetail,omitempty"`

	RestrictToWarehouse bool `xml:"RestrictToWarehouse,omitempty"`

	LanguageID int32 `xml:"LanguageID,omitempty"`

	ExcludeHideFromSearch bool `xml:"ExcludeHideFromSearch,omitempty"`
}

type ItemResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ItemResponse"`

	ItemCode string `xml:"ItemCode,omitempty"`

	Description string `xml:"Description,omitempty"`

	Price float64 `xml:"Price,omitempty"`

	CommissionableVolume float64 `xml:"CommissionableVolume,omitempty"`

	BusinessVolume float64 `xml:"BusinessVolume,omitempty"`

	Other1Price float64 `xml:"Other1Price,omitempty"`

	Other2Price float64 `xml:"Other2Price,omitempty"`

	Other3Price float64 `xml:"Other3Price,omitempty"`

	Other4Price float64 `xml:"Other4Price,omitempty"`

	Other5Price float64 `xml:"Other5Price,omitempty"`

	Other6Price float64 `xml:"Other6Price,omitempty"`

	Other7Price float64 `xml:"Other7Price,omitempty"`

	Other8Price float64 `xml:"Other8Price,omitempty"`

	Other9Price float64 `xml:"Other9Price,omitempty"`

	Other10Price float64 `xml:"Other10Price,omitempty"`

	Category string `xml:"Category,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty"`

	TinyPicture string `xml:"TinyPicture,omitempty"`

	SmallPicture string `xml:"SmallPicture,omitempty"`

	LargePicture string `xml:"LargePicture,omitempty"`

	ShortDetail string `xml:"ShortDetail,omitempty"`

	ShortDetail2 string `xml:"ShortDetail2,omitempty"`

	ShortDetail3 string `xml:"ShortDetail3,omitempty"`

	ShortDetail4 string `xml:"ShortDetail4,omitempty"`

	LongDetail string `xml:"LongDetail,omitempty"`

	LongDetail2 string `xml:"LongDetail2,omitempty"`

	LongDetail3 string `xml:"LongDetail3,omitempty"`

	LongDetail4 string `xml:"LongDetail4,omitempty"`

	InventoryStatus *InventoryStatusType `xml:"InventoryStatus,omitempty"`

	StockLevel int32 `xml:"StockLevel,omitempty"`

	AvailableStockLevel int32 `xml:"AvailableStockLevel,omitempty"`

	MaxAllowedOnOrder int32 `xml:"MaxAllowedOnOrder,omitempty"`

	Field1 string `xml:"Field1,omitempty"`

	Field2 string `xml:"Field2,omitempty"`

	Field3 string `xml:"Field3,omitempty"`

	Field4 string `xml:"Field4,omitempty"`

	Field5 string `xml:"Field5,omitempty"`

	Field6 string `xml:"Field6,omitempty"`

	Field7 string `xml:"Field7,omitempty"`

	Field8 string `xml:"Field8,omitempty"`

	Field9 string `xml:"Field9,omitempty"`

	Field10 string `xml:"Field10,omitempty"`

	OtherCheck1 bool `xml:"OtherCheck1,omitempty"`

	OtherCheck2 bool `xml:"OtherCheck2,omitempty"`

	OtherCheck3 bool `xml:"OtherCheck3,omitempty"`

	OtherCheck4 bool `xml:"OtherCheck4,omitempty"`

	OtherCheck5 bool `xml:"OtherCheck5,omitempty"`

	IsVirtual bool `xml:"IsVirtual,omitempty"`

	AllowOnAutoOrder bool `xml:"AllowOnAutoOrder,omitempty"`

	IsGroupMaster bool `xml:"IsGroupMaster,omitempty"`

	GroupDescription string `xml:"GroupDescription,omitempty"`

	GroupMembersDescription string `xml:"GroupMembersDescription,omitempty"`

	GroupMembers *ArrayOfItemMemberResponse `xml:"GroupMembers,omitempty"`

	IsDynamicKitMaster bool `xml:"IsDynamicKitMaster,omitempty"`

	HideFromSearch bool `xml:"HideFromSearch,omitempty"`

	KitMembers *ArrayOfKitMemberResponse `xml:"KitMembers,omitempty"`

	TaxablePrice float64 `xml:"TaxablePrice,omitempty"`

	ShippingPrice float64 `xml:"ShippingPrice,omitempty"`
}

type ArrayOfItemMemberResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfItemMemberResponse"`

	ItemMemberResponse []*ItemMemberResponse `xml:"ItemMemberResponse,omitempty"`
}

type ItemMemberResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ItemMemberResponse"`

	ItemCode string `xml:"ItemCode,omitempty"`

	MemberDescription string `xml:"MemberDescription,omitempty"`

	ItemDescription string `xml:"ItemDescription,omitempty"`

	InventoryStatus *InventoryStatusType `xml:"InventoryStatus,omitempty"`

	StockLevel int32 `xml:"StockLevel,omitempty"`

	AvailableStockLevel int32 `xml:"AvailableStockLevel,omitempty"`
}

type ArrayOfKitMemberResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfKitMemberResponse"`

	KitMemberResponse []*KitMemberResponse `xml:"KitMemberResponse,omitempty"`
}

type KitMemberResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ KitMemberResponse"`

	Description string `xml:"Description,omitempty"`

	KitMemberItems *ArrayOfKitMemberItemResponse `xml:"KitMemberItems,omitempty"`
}

type ArrayOfKitMemberItemResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfKitMemberItemResponse"`

	KitMemberItemResponse []*KitMemberItemResponse `xml:"KitMemberItemResponse,omitempty"`
}

type KitMemberItemResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ KitMemberItemResponse"`

	ItemCode string `xml:"ItemCode,omitempty"`

	Description string `xml:"Description,omitempty"`

	InventoryStatus *InventoryStatusType `xml:"InventoryStatus,omitempty"`
}

type ArrayOfItemResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfItemResponse"`

	ItemResponse []*ItemResponse `xml:"ItemResponse,omitempty"`
}

type GetItemsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetItemsResponse"`

	*ApiResponse

	Items *ArrayOfItemResponse `xml:"Items,omitempty"`
}

type GetLanguagesRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetLanguagesRequest"`

	*ApiRequest
}

type LanguageResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ LanguageResponse"`

	LanguageID int32 `xml:"LanguageID,omitempty"`

	Description string `xml:"Description,omitempty"`

	CultureCode string `xml:"CultureCode,omitempty"`
}

type ArrayOfLanguageResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfLanguageResponse"`

	LanguageResponse []*LanguageResponse `xml:"LanguageResponse,omitempty"`
}

type GetLanguagesResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetLanguagesResponse"`

	*ApiResponse

	CompanyLanguages *ArrayOfLanguageResponse `xml:"CompanyLanguages,omitempty"`
}

type CreateWebCategoryRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateWebCategoryRequest"`

	*ApiRequest

	WebID int32 `xml:"WebID,omitempty"`

	ParentID int32 `xml:"ParentID,omitempty"`

	Description string `xml:"Description,omitempty"`
}

type CreateWebCategoryResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateWebCategoryResponse"`

	*ApiResponse

	CategoryID int32 `xml:"CategoryID,omitempty"`

	WebID int32 `xml:"WebID,omitempty"`
}

type UpdateWebCategoryRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateWebCategoryRequest"`

	*ApiRequest

	WebID int32 `xml:"WebID,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty"`

	Description string `xml:"Description,omitempty"`
}

type UpdateWebCategoryResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateWebCategoryResponse"`

	*ApiResponse

	CategoryID int32 `xml:"CategoryID,omitempty"`

	WebID int32 `xml:"WebID,omitempty"`

	Description string `xml:"Description,omitempty"`
}

type DeleteWebCategoryRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteWebCategoryRequest"`

	*ApiRequest

	WebID int32 `xml:"WebID,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty"`
}

type DeleteWebCategoryResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteWebCategoryResponse"`

	*ApiResponse

	CategoryID int32 `xml:"CategoryID,omitempty"`

	WebID int32 `xml:"WebID,omitempty"`
}

type AddProductsToCategoryRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ AddProductsToCategoryRequest"`

	*ApiRequest

	WebID int32 `xml:"WebID,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty"`

	ItemCodes *ArrayOfString `xml:"ItemCodes,omitempty"`
}

type AddProductsToCategoryResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ AddProductsToCategoryResponse"`

	*ApiResponse

	CategoryID int32 `xml:"CategoryID,omitempty"`

	WebID int32 `xml:"WebID,omitempty"`
}

type DeleteProductFromCategoryRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteProductFromCategoryRequest"`

	*ApiRequest

	WebID int32 `xml:"WebID,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty"`

	ItemCode string `xml:"ItemCode,omitempty"`
}

type DeleteProductFromCategoryResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteProductFromCategoryResponse"`

	*ApiResponse

	CategoryID int32 `xml:"CategoryID,omitempty"`

	WebID int32 `xml:"WebID,omitempty"`

	ItemCode string `xml:"ItemCode,omitempty"`
}

type GetCompanyNewsRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCompanyNewsRequest"`

	*ApiRequest

	StartDate time.Time `xml:"StartDate,omitempty"`

	EndDate time.Time `xml:"EndDate,omitempty"`

	DepartmentType int32 `xml:"DepartmentType,omitempty"`
}

type CompanyNewsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CompanyNewsResponse"`

	Description string `xml:"Description,omitempty"`

	NewsID int32 `xml:"NewsID,omitempty"`

	CreatedDate time.Time `xml:"CreatedDate,omitempty"`

	WebSettings *NewsWebSettings `xml:"WebSettings,omitempty"`

	CompanySettings *NewsCompanySettings `xml:"CompanySettings,omitempty"`
}

type ArrayOfCompanyNewsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfCompanyNewsResponse"`

	CompanyNewsResponse []*CompanyNewsResponse `xml:"CompanyNewsResponse,omitempty"`
}

type GetCompanyNewsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCompanyNewsResponse"`

	*ApiResponse

	CompanyNews *ArrayOfCompanyNewsResponse `xml:"CompanyNews,omitempty"`
}

type GetCompanyNewsItemRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCompanyNewsItemRequest"`

	*ApiRequest

	NewsID int32 `xml:"NewsID,omitempty"`
}

type DepartmentInfo struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DepartmentInfo"`

	Description string `xml:"Description,omitempty"`

	DepartmentType int32 `xml:"DepartmentType,omitempty"`
}

type ArrayOfDepartmentInfo struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ ArrayOfDepartmentInfo"`

	DepartmentInfo []*DepartmentInfo `xml:"DepartmentInfo,omitempty"`
}

type GetCompanyNewsItemResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetCompanyNewsItemResponse"`

	*ApiResponse

	Description string `xml:"Description,omitempty"`

	NewsID int32 `xml:"NewsID,omitempty"`

	CreatedDate time.Time `xml:"CreatedDate,omitempty"`

	WebSettings *NewsWebSettings `xml:"WebSettings,omitempty"`

	CompanySettings *NewsCompanySettings `xml:"CompanySettings,omitempty"`

	Content string `xml:"Content,omitempty"`

	Departments *ArrayOfDepartmentInfo `xml:"Departments,omitempty"`
}

type GetRandomMessageRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetRandomMessageRequest"`

	*ApiRequest

	LanguageID int32 `xml:"LanguageID,omitempty"`
}

type GetRandomMessageResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ GetRandomMessageResponse"`

	*ApiResponse

	RandomMessageID int32 `xml:"RandomMessageID,omitempty"`

	Content string `xml:"Content,omitempty"`
}

type FireResponderRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ FireResponderRequest"`

	*ApiRequest

	ResponderID int32 `xml:"ResponderID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	OrderID int32 `xml:"OrderID,omitempty"`
}

type FireResponderResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ FireResponderResponse"`

	*ApiResponse
}

type SendSmsRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SendSmsRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	Message string `xml:"Message,omitempty"`

	Phone string `xml:"Phone,omitempty"`
}

type SendSmsResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ SendSmsResponse"`

	*ApiResponse
}

type CreateVendorBillRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateVendorBillRequest"`

	*ApiRequest

	VendorBillStatusTypeID int32 `xml:"VendorBillStatusTypeID,omitempty"`

	CustomerID int32 `xml:"CustomerID,omitempty"`

	CurrencyCode string `xml:"CurrencyCode,omitempty"`

	DueDate time.Time `xml:"DueDate,omitempty"`

	Amount float64 `xml:"Amount,omitempty"`

	Reference string `xml:"Reference,omitempty"`

	Notes string `xml:"Notes,omitempty"`
}

type CreateVendorBillResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateVendorBillResponse"`

	*ApiResponse

	VendorBillID int32 `xml:"VendorBillID,omitempty"`
}

type CreateCustomerContactRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateCustomerContactRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	FirstName string `xml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty"`

	Company string `xml:"Company,omitempty"`

	BusinessPhone string `xml:"BusinessPhone,omitempty"`

	HomePhone string `xml:"HomePhone,omitempty"`

	Mobile string `xml:"Mobile,omitempty"`

	Fax string `xml:"Fax,omitempty"`

	Email string `xml:"Email,omitempty"`

	Address1 string `xml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty"`

	City string `xml:"City,omitempty"`

	State string `xml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty"`

	BirthDate time.Time `xml:"BirthDate,omitempty"`

	Notes string `xml:"Notes,omitempty"`

	LinkedIn string `xml:"LinkedIn,omitempty"`

	Facebook string `xml:"Facebook,omitempty"`

	Blog string `xml:"Blog,omitempty"`

	MySpace string `xml:"MySpace,omitempty"`

	GooglePlus string `xml:"GooglePlus,omitempty"`

	Twitter string `xml:"Twitter,omitempty"`
}

type CreateCustomerContactResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateCustomerContactResponse"`

	*ApiResponse

	CustomerID int32 `xml:"CustomerID,omitempty"`

	CustomerContactID int32 `xml:"CustomerContactID,omitempty"`
}

type UpdateCustomerContactRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateCustomerContactRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	CustomerContactID int32 `xml:"CustomerContactID,omitempty"`

	FirstName string `xml:"FirstName,omitempty"`

	LastName string `xml:"LastName,omitempty"`

	MiddleName string `xml:"MiddleName,omitempty"`

	Company string `xml:"Company,omitempty"`

	BusinessPhone string `xml:"BusinessPhone,omitempty"`

	HomePhone string `xml:"HomePhone,omitempty"`

	Mobile string `xml:"Mobile,omitempty"`

	Fax string `xml:"Fax,omitempty"`

	Email string `xml:"Email,omitempty"`

	Address1 string `xml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty"`

	City string `xml:"City,omitempty"`

	State string `xml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty"`

	BirthDate time.Time `xml:"BirthDate,omitempty"`

	Notes string `xml:"Notes,omitempty"`

	LinkedIn string `xml:"LinkedIn,omitempty"`

	Facebook string `xml:"Facebook,omitempty"`

	Blog string `xml:"Blog,omitempty"`

	MySpace string `xml:"MySpace,omitempty"`

	GooglePlus string `xml:"GooglePlus,omitempty"`

	Twitter string `xml:"Twitter,omitempty"`
}

type UpdateCustomerContactResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateCustomerContactResponse"`

	*ApiResponse

	CustomerID int32 `xml:"CustomerID,omitempty"`

	CustomerContactID int32 `xml:"CustomerContactID,omitempty"`
}

type DeleteCustomerContactRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteCustomerContactRequest"`

	*ApiRequest

	CustomerID int32 `xml:"CustomerID,omitempty"`

	CustomerContactID int32 `xml:"CustomerContactID,omitempty"`
}

type DeleteCustomerContactResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteCustomerContactResponse"`

	*ApiResponse

	CustomerID int32 `xml:"CustomerID,omitempty"`

	CustomerContactID int32 `xml:"CustomerContactID,omitempty"`
}

type CreateCalendarItemRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateCalendarItemRequest"`

	*ApiRequest

	UserID int32 `xml:"UserID,omitempty"`

	CalendarID int32 `xml:"CalendarID,omitempty"`

	CalendarItemType *CalendarItemType `xml:"CalendarItemType,omitempty"`

	CalendarItemStatusType *CalendarItemStatusType `xml:"CalendarItemStatusType,omitempty"`

	CalendarItemPriorityType *CalendarItemPriorityType `xml:"CalendarItemPriorityType,omitempty"`

	Subject string `xml:"Subject,omitempty"`

	Location string `xml:"Location,omitempty"`

	Notes string `xml:"Notes,omitempty"`

	StartDate time.Time `xml:"StartDate,omitempty"`

	EndDate time.Time `xml:"EndDate,omitempty"`

	TimeZone int32 `xml:"TimeZone,omitempty"`

	Address1 string `xml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty"`

	City string `xml:"City,omitempty"`

	State string `xml:"State,omitempty"`

	Country string `xml:"Country,omitempty"`

	PostalCode string `xml:"PostalCode,omitempty"`

	ContactInfo string `xml:"ContactInfo,omitempty"`

	ContactPhone string `xml:"ContactPhone,omitempty"`

	ContactPhoneType *ContactPhoneType `xml:"ContactPhoneType,omitempty"`

	ContactEmail string `xml:"ContactEmail,omitempty"`

	EventHost string `xml:"EventHost,omitempty"`

	SpecialGuests string `xml:"SpecialGuests,omitempty"`

	EventFlyer string `xml:"EventFlyer,omitempty"`

	EventCostInfo string `xml:"EventCostInfo,omitempty"`

	EventConferenceCallOrWebinar string `xml:"EventConferenceCallOrWebinar,omitempty"`

	EventRegistrationInfo string `xml:"EventRegistrationInfo,omitempty"`

	EventTags string `xml:"EventTags,omitempty"`

	IsShared bool `xml:"IsShared,omitempty"`
}

type CreateCalendarItemResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ CreateCalendarItemResponse"`

	*ApiResponse

	CalendarID int32 `xml:"CalendarID,omitempty"`

	CalendarItemID int32 `xml:"CalendarItemID,omitempty"`
}

type UpdateCalendarItemRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateCalendarItemRequest"`

	*ApiRequest

	UserID int32 `xml:"UserID,omitempty"`

	CalendarID int32 `xml:"CalendarID,omitempty"`

	CalendarItemID int32 `xml:"CalendarItemID,omitempty"`

	CalendarItemType *CalendarItemType `xml:"CalendarItemType,omitempty"`

	CalendarItemStatusType *CalendarItemStatusType `xml:"CalendarItemStatusType,omitempty"`

	CalendarItemPriorityType *CalendarItemPriorityType `xml:"CalendarItemPriorityType,omitempty"`

	Subject string `xml:"Subject,omitempty"`

	Location string `xml:"Location,omitempty"`

	Notes string `xml:"Notes,omitempty"`

	StartDate time.Time `xml:"StartDate,omitempty"`

	EndDate time.Time `xml:"EndDate,omitempty"`

	TimeZone int32 `xml:"TimeZone,omitempty"`

	Address1 string `xml:"Address1,omitempty"`

	Address2 string `xml:"Address2,omitempty"`

	City string `xml:"City,omitempty"`

	State string `xml:"State,omitempty"`

	Country string `xml:"Country,omitempty"`

	PostalCode string `xml:"PostalCode,omitempty"`

	ContactInfo string `xml:"ContactInfo,omitempty"`

	ContactPhone string `xml:"ContactPhone,omitempty"`

	ContactPhoneType *ContactPhoneType `xml:"ContactPhoneType,omitempty"`

	ContactEmail string `xml:"ContactEmail,omitempty"`

	EventHost string `xml:"EventHost,omitempty"`

	SpecialGuests string `xml:"SpecialGuests,omitempty"`

	EventFlyer string `xml:"EventFlyer,omitempty"`

	EventCostInfo string `xml:"EventCostInfo,omitempty"`

	EventConferenceCallOrWebinar string `xml:"EventConferenceCallOrWebinar,omitempty"`

	EventRegistrationInfo string `xml:"EventRegistrationInfo,omitempty"`

	EventTags string `xml:"EventTags,omitempty"`

	IsShared bool `xml:"IsShared,omitempty"`
}

type UpdateCalendarItemResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ UpdateCalendarItemResponse"`

	*ApiResponse

	CalendarID int32 `xml:"CalendarID,omitempty"`

	CalendarItemID int32 `xml:"CalendarItemID,omitempty"`
}

type DeleteCalendarItemRequest struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteCalendarItemRequest"`

	*ApiRequest

	UserID int32 `xml:"UserID,omitempty"`

	CalendarID int32 `xml:"CalendarID,omitempty"`

	CalendarItemID int32 `xml:"CalendarItemID,omitempty"`
}

type DeleteCalendarItemResponse struct {
	XMLName xml.Name `xml:"http://api.exigo.com/ DeleteCalendarItemResponse"`

	*ApiResponse

	CalendarID int32 `xml:"CalendarID,omitempty"`

	CalendarItemID int32 `xml:"CalendarItemID,omitempty"`
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

func (service *ExigoApiSoap) AddHeader(header interface{}) {
	service.client.AddHeader(header)
}

// Backwards-compatible function: use AddHeader instead
func (service *ExigoApiSoap) SetHeader(header interface{}) {
	service.client.AddHeader(header)
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

	Items []interface{} `xml:",omitempty"`
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

const (
	// Predefined WSS namespaces to be used in
	WssNsWSSE string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	WssNsWSU  string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
	WssNsType string = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText"
)

type WSSSecurityHeader struct {
	XMLName   xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ wsse:Security"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	MustUnderstand string `xml:"mustUnderstand,attr,omitempty"`

	Token *WSSUsernameToken `xml:",omitempty"`
}

type WSSUsernameToken struct {
	XMLName   xml.Name `xml:"wsse:UsernameToken"`
	XmlNSWsu  string   `xml:"xmlns:wsu,attr"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	Id string `xml:"wsu:Id,attr,omitempty"`

	Username *WSSUsername `xml:",omitempty"`
	Password *WSSPassword `xml:",omitempty"`
}

type WSSUsername struct {
	XMLName   xml.Name `xml:"wsse:Username"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`

	Data string `xml:",chardata"`
}

type WSSPassword struct {
	XMLName   xml.Name `xml:"wsse:Password"`
	XmlNSWsse string   `xml:"xmlns:wsse,attr"`
	XmlNSType string   `xml:"Type,attr"`

	Data string `xml:",chardata"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url     string
	tls     bool
	auth    *BasicAuth
	headers []interface{}
}

// **********
// Accepted solution from http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
// Author: Icza - http://stackoverflow.com/users/1705598/icza

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randStringBytesMaskImprSrc(n int) string {
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// **********

func NewWSSSecurityHeader(user, pass, mustUnderstand string) *WSSSecurityHeader {
	hdr := &WSSSecurityHeader{XmlNSWsse: WssNsWSSE, MustUnderstand: mustUnderstand}
	hdr.Token = &WSSUsernameToken{XmlNSWsu: WssNsWSU, XmlNSWsse: WssNsWSSE, Id: "UsernameToken-" + randStringBytesMaskImprSrc(9)}
	hdr.Token.Username = &WSSUsername{XmlNSWsse: WssNsWSSE, Data: user}
	hdr.Token.Password = &WSSPassword{XmlNSWsse: WssNsWSSE, XmlNSType: WssNsType, Data: pass}
	return hdr
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

func (s *SOAPClient) AddHeader(header interface{}) {
	s.headers = append(s.headers, header)
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{}

	if s.headers != nil && len(s.headers) > 0 {
		soapHeader := &SOAPHeader{Items: make([]interface{}, len(s.headers))}
		copy(soapHeader.Items, s.headers)
		envelope.Header = soapHeader
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

	log.Println(buffer.String())

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	req.Header.Add("SOAPAction", soapAction)

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

	log.Println(string(rawbody))
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}

	return nil
}
