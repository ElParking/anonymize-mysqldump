package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/xwb1989/sqlparser"
	"syreclabs.com/go/faker"
)

var (
	usedEmails map[string]bool
	mu         sync.Mutex
)

func init() {
	usedEmails = make(map[string]bool)
}

// GetFakerFuncs creates a map of faker helper functions.
func GetFakerFuncs() map[string]func(*sqlparser.SQLVal) *sqlparser.SQLVal {
	fakerHelpers := map[string]func(*sqlparser.SQLVal) *sqlparser.SQLVal{
		"username":             generateUsername,
		"password":             generatePassword,
		"email":                generateEmail,
		"url":                  generateURL,
		"name":                 generateName,
		"firstName":            generateFirstName,
		"lastName":             generateLastName,
		"personPrefix":         generatePersonPrefix,
		"personTitle":          generatePersonTitle,
		"phoneNumber":          generatePhoneNumber,
		"billingAddressFull":   generateBillingAddress,
		"addressFull":          generateAddress,
		"addressStreet":        generateStreetAddress,
		"addressSecondary":     generateSecondaryStreetAddress,
		"addressCity":          generateCity,
		"addressState":         generateAddressState,
		"addressPostCode":      generatePostcode,
		"addressCountry":       generateCountry,
		"addressCountryCode":   generateCountryCode,
		"paragraph":            generateParagraph,
		"shortString":          generateShortString,
		"ipv4":                 generateIPv4,
		"companyName":          generateCompanyName,
		"companySuffix":        generateCompanySuffix,
		"companyNumber":        generateCompanyNumber,
		"creditCardNumber":     generateCreditCardNumber,
		"creditCardExpiryDate": generateCreditCardExpiryDate,
		"creditCardType":       generateCreditCardType,
		"norwegianSSN":         generateNorwegianSSN,
		"WPDateTime":           generateWPDateTime,
		"WPFutureDateTime":     generateWPFutureDateTime,
		"purge":                generateEmptyString,
		"spanishDNI":           generateSpanishDniString,
	}

	return fakerHelpers
}

func generateUsername(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Internet().UserName()))
}

func generatePassword(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	password := md5.Sum([]byte(faker.Internet().Password(8, 14)))
	return sqlparser.NewStrVal([]byte(hex.EncodeToString(password[:])))
}

func generateEmail(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	mu.Lock()
	defer mu.Unlock()

	var email string
	var second_email_option string
	var third_email_option string
	var used_email string
	for {
		email = faker.Internet().SafeEmail()
		second_email_option = faker.Internet().FreeEmail()
		third_email_option = faker.Internet().Email()
		if !usedEmails[email] {
			usedEmails[email] = true
			used_email = email
			break
		} else if !usedEmails[second_email_option] {
			usedEmails[second_email_option] = true
			used_email = second_email_option
			break
		} else if !usedEmails[third_email_option] {
			usedEmails[third_email_option] = true
			used_email = third_email_option
			break
		}
	}

	return sqlparser.NewStrVal([]byte(used_email))
}

func generatePhoneNumber(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.PhoneNumber().CellPhone()))
}

func generateURL(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Internet().Url()))
}

func generateName(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Name().Name()))
}

func generateFirstName(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Name().FirstName()))
}

func generateLastName(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Name().LastName()))
}

func generatePersonPrefix(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Name().Prefix()))
}

func generatePersonTitle(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Name().Title()))
}

func generateParagraph(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Lorem().Sentence(3)))
}

func generateIPv4(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Internet().IpV4Address()))
}

func generateBillingAddress(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	address := ""
	address += " " + faker.Name().FirstName()
	address += " " + faker.Name().LastName()
	address += " " + faker.Address().String()
	address += " " + faker.Address().CountryCode()
	address += " " + faker.Internet().SafeEmail()
	address += " " + faker.PhoneNumber().CellPhone()

	return sqlparser.NewStrVal([]byte(address))
}

func generateAddress(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Address().String()))
}

func generateStreetAddress(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Address().StreetAddress()))
}

func generateSecondaryStreetAddress(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Address().SecondaryAddress()))
}

func generateAddressState(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Address().State()))
}

func generateCity(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Address().City()))
}

func generatePostcode(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Address().Postcode()))
}

func generateCountry(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Address().Country()))
}

func generateCountryCode(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Address().CountryCode()))
}

func generateCreditCardNumber(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Business().CreditCardNumber()))
}

func generateCreditCardExpiryDate(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Business().CreditCardExpiryDate()))
}

func generateCreditCardType(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Business().CreditCardType()))
}

func generateCompanyName(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Company().Name()))
}

func generateCompanySuffix(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Company().Suffix()))
}

func generateCompanyNumber(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Number().Number(9)))
}

func generateShortString(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Lorem().Characters(30)))
}

func generateEmptyString(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(""))
}

func generateNorwegianSSN(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(generateFakeNorwegianSSN(faker.Date().Birthday(18, 90))))
}

func generateWPDateTime(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	// Define the maximum duration in the future/past.
	// 11 years, 365 days, 24 hours, etc.
	maxSeconds := int64(11 * 365 * 24 * 60 * 60)

	// Generate a random number of seconds to add
	randomSeconds := rand.Int63n(maxSeconds)

	var futureTime time.Time

	// Randomize if we add or remove time.
	if rand.Intn(2) == 0 {
		futureTime = time.Now().Add(time.Duration(randomSeconds) * time.Second)
	} else {
		futureTime = time.Now().Add(-time.Duration(randomSeconds) * time.Second)
	}

	return sqlparser.NewStrVal([]byte(futureTime.Format("2006-01-02 15:04:05")))
}
func generateWPFutureDateTime(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	// Define the maximum duration in the future.
	// 11 years, 365 days, 24 hours, etc.
	maxSeconds := int64(11 * 365 * 24 * 60 * 60)

	// Generate a random number of seconds to add
	randomSeconds := rand.Int63n(maxSeconds)

	// Add the random duration to the current time
	futureTime := time.Now().Add(time.Duration(randomSeconds) * time.Second)

	return sqlparser.NewStrVal([]byte(futureTime.Format("2006-01-02 15:04:05")))
}

func generateSpanishDniString(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	number := rand.Intn(99999999)
	nif_seq := "TRWAGMYFPDXBNJZSQVHLCKE"
	letter := string(nif_seq[(number % len(nif_seq))])
	dni := fmt.Sprintf("%08d", number) + letter
	return sqlparser.NewStrVal([]byte(dni))
}
