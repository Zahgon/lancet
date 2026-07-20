package validator

import (
	"regexp"
)

var (
	alphaMatcher        *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z]+$`)
	letterRegexMatcher  *regexp.Regexp = regexp.MustCompile(`[a-zA-Z]`)
	alphaNumericMatcher *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z0-9-]+$`)
	numberRegexMatcher  *regexp.Regexp = regexp.MustCompile(`\d`)
	intStrMatcher       *regexp.Regexp = regexp.MustCompile(`^[\+-]?\d+$`)

	dnsMatcher               *regexp.Regexp = regexp.MustCompile(`^(?:[a-zA-Z0-9](?:[a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)*(?:xn--[a-zA-Z0-9\-]{1,59}|[a-zA-Z0-9](?:[a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?)$`)
	emailMatcher             *regexp.Regexp = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	chineseMobileMatcher     *regexp.Regexp = regexp.MustCompile(`^1(?:3\d|4[4-9]|5[0-35-9]|6[67]|7[013-8]|8\d|9\d)\d{8}$`)
	chineseIdMatcher         *regexp.Regexp = regexp.MustCompile(`([1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx])|([1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{2}[0-9Xx])`)
	chineseMatcher           *regexp.Regexp = regexp.MustCompile("[\u4e00-\u9fa5]")
	chinesePhoneMatcher      *regexp.Regexp = regexp.MustCompile(`\d{3}-\d{8}|\d{4}-\d{7}|\d{4}-\d{8}`)
	creditCardMatcher        *regexp.Regexp = regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|(222[1-9]|22[3-9][0-9]|2[3-6][0-9]{2}|27[01][0-9]|2720)[0-9]{12}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11}|6[27][0-9]{14})$`)
	base64Matcher            *regexp.Regexp = regexp.MustCompile(`^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$`)
	base64URLMatcher         *regexp.Regexp = regexp.MustCompile(`^([A-Za-z0-9_-]{4})*([A-Za-z0-9_-]{2}(==)?|[A-Za-z0-9_-]{3}=?)?$`)
	binMatcher               *regexp.Regexp = regexp.MustCompile(`^(0b)?[01]+$`)
	hexMatcher               *regexp.Regexp = regexp.MustCompile(`^(#|0x|0X)?[0-9a-fA-F]+$`)
	visaMatcher              *regexp.Regexp = regexp.MustCompile(`^4[0-9]{12}(?:[0-9]{3})?$`)
	masterCardMatcher        *regexp.Regexp = regexp.MustCompile(`^5[1-5][0-9]{14}$`)
	americanExpressMatcher   *regexp.Regexp = regexp.MustCompile(`^3[47][0-9]{13}$`)
	chinaUnionPayMatcher     *regexp.Regexp = regexp.MustCompile(`^62[0-9]{14,17}$`)
	chineseHMPassportMatcher *regexp.Regexp = regexp.MustCompile(`^[CM]\d{8}$`)
)

var passportMatcher = map[string]*regexp.Regexp{
	"CN": regexp.MustCompile(`^P\d{9}$`),
	"US": regexp.MustCompile(`^\d{9}$`),
	"GB": regexp.MustCompile(`^[A-Z0-9]{9}$`),
	"RU": regexp.MustCompile(`^[A-Z]{2}\d{7}$`),
	"DE": regexp.MustCompile(`^\d{9}$`),
	"FR": regexp.MustCompile(`^[A-Z]{2}\d{7}$`),
	"JP": regexp.MustCompile(`^\d{8}$`),
	"IT": regexp.MustCompile(`^\d{8}$`),
	"AU": regexp.MustCompile(`^[A-Z]{1}\d{8}$`),
	"BR": regexp.MustCompile(`^\d{9}$`),
	"IN": regexp.MustCompile(`^[A-Z]{1,2}\d{7}$`),
	"HK": regexp.MustCompile(`^M\d{8}$`),
	"MO": regexp.MustCompile(`^[A-Z]\d{8}$`),
}

var (
	factor = [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}

	verifyStr = [11]string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}

	birthStartYear = 1900

	provinceKv = map[string]struct{}{
		"11": {},
		"12": {},
		"13": {},
		"14": {},
		"15": {},
		"21": {},
		"22": {},
		"23": {},
		"31": {},
		"32": {},
		"33": {},
		"34": {},
		"35": {},
		"36": {},
		"37": {},
		"41": {},
		"42": {},
		"43": {},
		"44": {},
		"45": {},
		"46": {},
		"50": {},
		"51": {},
		"52": {},
		"53": {},
		"54": {},
		"61": {},
		"62": {},
		"63": {},
		"64": {},
		"65": {},
	}
)

func IsAlpha(str string) bool { _ = "STUB: not implemented"; return false }

func IsAllUpper(str string) bool { _ = "STUB: not implemented"; return false }

func IsAllLower(str string) bool { _ = "STUB: not implemented"; return false }

func IsASCII(str string) bool { _ = "STUB: not implemented"; return false }

func IsPrintable(str string) bool { _ = "STUB: not implemented"; return false }

func ContainUpper(str string) bool { _ = "STUB: not implemented"; return false }

func ContainLower(str string) bool { _ = "STUB: not implemented"; return false }

func ContainLetter(str string) bool { _ = "STUB: not implemented"; return false }

func ContainNumber(input string) bool { _ = "STUB: not implemented"; return false }

func IsJSON(str string) bool { _ = "STUB: not implemented"; return false }

func IsAlphaNumeric(s string) bool { _ = "STUB: not implemented"; return false }

func IsNumberStr(s string) bool { _ = "STUB: not implemented"; return false }

func IsFloatStr(str string) bool { _ = "STUB: not implemented"; return false }

func IsIntStr(str string) bool { _ = "STUB: not implemented"; return false }

func IsIp(ipstr string) bool { _ = "STUB: not implemented"; return false }

func IsIpPort(str string) bool { _ = "STUB: not implemented"; return false }

func IsIpV4(ipstr string) bool { _ = "STUB: not implemented"; return false }

func IsIpV6(ipstr string) bool { _ = "STUB: not implemented"; return false }

func IsPort(str string) bool { _ = "STUB: not implemented"; return false }

func IsUrl(str string) bool { _ = "STUB: not implemented"; return false }

func IsDns(dns string) bool { _ = "STUB: not implemented"; return false }

func IsEmail(email string) bool { _ = "STUB: not implemented"; return false }

func IsChineseMobile(mobileNum string) bool { _ = "STUB: not implemented"; return false }

func IsChineseIdNum(id string) bool { _ = "STUB: not implemented"; return false }

func ContainChinese(s string) bool { _ = "STUB: not implemented"; return false }

func IsChinesePhone(phone string) bool { _ = "STUB: not implemented"; return false }

func IsCreditCard(creditCart string) bool { _ = "STUB: not implemented"; return false }

func IsBase64(base64 string) bool { _ = "STUB: not implemented"; return false }

func IsEmptyString(str string) bool { _ = "STUB: not implemented"; return false }

func IsRegexMatch(str, regex string) bool { _ = "STUB: not implemented"; return false }

func IsStrongPassword(password string, length int) bool { _ = "STUB: not implemented"; return false }

func IsWeakPassword(password string) bool { _ = "STUB: not implemented"; return false }

func IsZeroValue(value any) bool { _ = "STUB: not implemented"; return false }

func IsGBK(data []byte) bool { _ = "STUB: not implemented"; return false }

func IsNumber(v any) bool { _ = "STUB: not implemented"; return false }

func IsFloat(v any) bool { _ = "STUB: not implemented"; return false }

func IsInt(v any) bool { _ = "STUB: not implemented"; return false }

func IsBin(v string) bool { _ = "STUB: not implemented"; return false }

func IsHex(v string) bool { _ = "STUB: not implemented"; return false }

func IsBase64URL(v string) bool { _ = "STUB: not implemented"; return false }

func IsJWT(v string) bool { _ = "STUB: not implemented"; return false }

func IsVisa(v string) bool { _ = "STUB: not implemented"; return false }

func IsMasterCard(v string) bool { _ = "STUB: not implemented"; return false }

func IsAmericanExpress(v string) bool { _ = "STUB: not implemented"; return false }

func IsUnionPay(cardNo string) bool { _ = "STUB: not implemented"; return false }

func IsChinaUnionPay(cardNo string) bool { _ = "STUB: not implemented"; return false }

func IsPassport(passport, country string) bool { _ = "STUB: not implemented"; return false }

func IsChineseHMPassport(hmPassport string) bool { _ = "STUB: not implemented"; return false }
