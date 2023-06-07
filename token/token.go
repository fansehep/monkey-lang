package token


const (
    // 未知的词法单元或者字符
    LLLEGAL = "LLLEGAL"
    EOF = "EOF"

    // 标识符
    IDENT = "IDENT"
    // 字面量
    INT = "INT"

    // 运算符
    ASSIGN = "="
    PLUS = "+"
    MINUS = "-"
    BANG = "!"
    ASTERISK = "*"
    SLASH = "/"

    LESS_THAN = "<"
    GREATER_THAN = ">"

    // 分隔符
    COMMA = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    FUNCTION = "FUNCTION"
    LET = "LET"
    TRUE = "TRUE"
    FALSE = "FALSE"
    IF = "IF"
    ELSE = "ELSE"
    RETURN = "RETURN"

    EQ = "=="
    NOT_EQ = "!="


)

var keywords = map[string]TokenType {
    "fn": FUNCTION,
    "let": LET,
    "true": TRUE,
    "false": FALSE,
    "if": IF,
    "else": ELSE,
    "return": RETURN,
}

func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    // LookupIdent 来检查是否时语言的关键字,
    // 如果是。则返回 TokenType
    // 否则则认为是用户定义的标识符
    return IDENT
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}


func New(ty TokenType, li string) Token {
    return Token {
        Type: ty,
        Literal: li,
    }
}
