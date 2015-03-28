//  Scanner  -  A parser to parse a configuration file.
//
//  E.B.Smith  -  November, 2014


package Scanner


import (
    "os"
    "fmt"
    "path"
    "bytes"
    "bufio"
    "unicode"
    "errors"
    "strings"
    "strconv"
)


type Scanner struct {
    file        *os.File
    reader      *bufio.Reader
    lineNumber  int
    error       error
    token       string
}


func NewScanner(file *os.File) *Scanner {
    if file == nil { return nil }
    scanner := new(Scanner)
    scanner.file = file
    scanner.reader = bufio.NewReader(file)
    scanner.lineNumber = 1
    scanner.token = ""
    return scanner
}


func (scanner *Scanner) FileName() string {
    return scanner.file.Name()
}


func (scanner *Scanner) LineNumber() int {
    return scanner.lineNumber
}


func (scanner *Scanner) IsAtEnd() bool {
    return scanner.error != nil
}


func (scanner *Scanner) Token() string {
    return scanner.token
}


func (scanner *Scanner) LastError() error {
    return scanner.error
}


func (scanner *Scanner) SetErrorMessage(message string) error {
    basename := path.Base(scanner.FileName())
    message = fmt.Sprintf("%s:%d Scanned '%s'. %s",
            basename, scanner.LineNumber(), scanner.Token(), message)
    scanner.error = errors.New(message)
    return scanner.error
}


func (scanner *Scanner) SetError(error error) error {
    basename := path.Base(scanner.FileName())
    message := fmt.Sprintf("%s:%d Scanned '%s'. %v",
            basename, scanner.LineNumber(), scanner.Token(), error)
    scanner.error = errors.New(message)
    return scanner.error
}


//  Scan Routines --


func IsValidIdentifierStartRune(r rune) bool {
    return unicode.IsLetter(r) || r == '_'
}


func IsValidIdentifierRune(r rune) bool {
    return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || r == '_'
}

func IsOctalDigit(r rune) bool {
    return unicode.IsDigit(r) && r != '8' && r != '9'
}

func ZIsSpace(r rune) bool {
    return unicode.IsSpace(r) || r == '#'
}


func ZIsLineFeed(r rune) bool {
    return r == '\n' || r == '\u0085'
}


func (scanner *Scanner) ScanSpaces() error {
    scanner.token = ""
    for ! scanner.IsAtEnd() {
        var r rune
        r, _, scanner.error = scanner.reader.ReadRune()

        if r == '#' {
            for !scanner.IsAtEnd() && !ZIsLineFeed(r) {
                r, _, scanner.error = scanner.reader.ReadRune()
            }
        }
        if ZIsLineFeed(r) {
            scanner.lineNumber++
            continue
        }
        if ZIsSpace(r) {
            continue
        }

        scanner.reader.UnreadRune()
        return nil
    }

    return scanner.error;
}


func IsValidStringRune(r rune) bool {
    if r == ';' || r == ',' || ZIsSpace(r) { return false }
    return unicode.IsGraphic(r)
}


func (scanner *Scanner) ScanString() (next string, error error) {
    error = scanner.ScanSpaces()

    var (r rune; buffer bytes.Buffer)
    r, _, scanner.error = scanner.reader.ReadRune()

    for IsValidStringRune(r) {
        buffer.WriteRune(r)
        r, _, scanner.error = scanner.reader.ReadRune()
        }
    scanner.reader.UnreadRune()

    scanner.token = buffer.String()
    return scanner.token, nil
}


func (scanner *Scanner) ScanInt() (int int, error error) {
    scanner.ScanSpaces()
    var r rune
    r, _, scanner.error = scanner.reader.ReadRune()

    if ! unicode.IsDigit(r) {
        scanner.reader.UnreadRune()
        scanner.token, _ = scanner.ScanNext()
        return 0, scanner.SetErrorMessage("Integer expected")
        }

    var buffer bytes.Buffer
    for unicode.IsDigit(r) {
        buffer.WriteRune(r)
        r, _, scanner.error = scanner.reader.ReadRune()
        }
    scanner.reader.UnreadRune()

    scanner.token = buffer.String()
    return  strconv.Atoi(scanner.token)
}


func (scanner *Scanner) ScanBool() (value bool, error error) {
    var s string
    s, scanner.error = scanner.ScanNext()
    if scanner.error != nil { return false, scanner.error }

    s = strings.ToLower(s)
    if s == "true" || s == "yes" || s == "t" || s == "y" || s == "1" {
        return true, nil
    }
    if s == "false" || s == "no" || s == "f" || s == "n" || s == "0" {
        return false, nil
    }
    scanner.SetErrorMessage("Expected a boolean value")
    return false, scanner.error
}


func (scanner *Scanner) ScanIdentifier() (identifier string, error error) {
    scanner.ScanSpaces()
    var r rune
    r, _, scanner.error = scanner.reader.ReadRune()
    if scanner.error != nil {
        return "", scanner.error
        }

    if ! IsValidIdentifierStartRune(r) {
        scanner.reader.UnreadRune()
        scanner.ScanNext()
        return "", scanner.SetErrorMessage("Expected an identifier")
        }

    var buffer bytes.Buffer
    for IsValidIdentifierRune(r) {
        buffer.WriteRune(r)
        r, _, scanner.error = scanner.reader.ReadRune()
        }
    scanner.reader.UnreadRune()

    scanner.token = buffer.String()
    return scanner.token, nil
}


func (scanner *Scanner) ScanOctal() (Integer int, error error) {
    scanner.ScanSpaces()
    var r rune
    r, _, scanner.error = scanner.reader.ReadRune()

    if ! IsOctalDigit(r) {
        scanner.reader.UnreadRune()
        scanner.token, _ = scanner.ScanNext()
        return 0, scanner.SetErrorMessage("Octal number expected")
    }

    var buffer bytes.Buffer
    for IsOctalDigit(r) {
        buffer.WriteRune(r)
        r, _, scanner.error = scanner.reader.ReadRune()
    }
    scanner.reader.UnreadRune()

    scanner.token = buffer.String()
    val, error := strconv.ParseInt(scanner.token, 8, 0)
    return int(val), error
}


func (scanner *Scanner) ScanQuotedString() (string, error) {
    scanner.ScanSpaces()
    var r rune
    r, _, scanner.error = scanner.reader.ReadRune()

    if r != '"' {
        scanner.token = string(r)
        return "", scanner.SetErrorMessage("Quoted string expected")
    }
    scanner.reader.UnreadRune()

    parseCount, error := fmt.Fscanf(scanner.reader, "%q", &scanner.token)
    if error != nil { return "", scanner.SetError(error) }
    if parseCount != 1 { return "", scanner.SetErrorMessage("Quoted string expected") }

    return scanner.token, nil
}


func (scanner *Scanner) ScanNext() (next string, error error) {
    scanner.ScanSpaces()
    var r rune
    r, _, scanner.error = scanner.reader.ReadRune()

    if r == '"' {
        scanner.reader.UnreadRune()
        return scanner.ScanQuotedString()
    }
    if unicode.IsPunct(r) {
        var buffer bytes.Buffer
        buffer.WriteRune(r)
        scanner.token = buffer.String()
        return scanner.token, nil
    }
    if unicode.IsDigit(r) {
        scanner.reader.UnreadRune()
        scanner.ScanInt()
        return scanner.token, scanner.error
    }

    scanner.reader.UnreadRune()
    return scanner.ScanString()
}


func (scanner *Scanner) ScanToEOL() (string, error) {
    var (r rune; buffer bytes.Buffer)
    r, _, scanner.error = scanner.reader.ReadRune()

    for !scanner.IsAtEnd()  &&  !ZIsLineFeed(r) {
        buffer.WriteRune(r)
        r, _, scanner.error = scanner.reader.ReadRune()
    }
    scanner.reader.UnreadRune()

    scanner.token = buffer.String()
    scanner.token = strings.TrimSpace(scanner.token)
    return scanner.token, nil
}

