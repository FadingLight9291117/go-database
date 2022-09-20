package sqlcompiler

type TokenType uint8

const (
	KEYWORD  TokenType = iota // 关键字
	LITERALS                  // 词法字面量
	SYMBOL                    // 标点符号
	ASSIST                    // 词法辅助标记
)

type KeyWord uint8

const (
	SELECT KeyWord = iota
	FROM
	WHERE

	INSERT
	INTO
	VALUES
)

type Literals uint8

const (
	IDENTIFIER Literals = iota
	INT
	STRING
)

type Symbol uint8

const (
	STAR Symbol = iota
	COMMA
)

type Assist uint8

const (
	ERROR Assist = iota
	END
)
