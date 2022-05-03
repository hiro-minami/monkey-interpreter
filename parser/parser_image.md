```
// ParseProgramの擬似コード
function parseProgram() {
    program = newProgramASTNode()
    // トークンを進める
    advanceTokens()

    // 構文解析の部分はここ
    for(currentToken() == LET_TOKEN) {
        statement = null

        if(currentToken() == LET_TOKEN) statement = parseLetStatement()
        else if(currentToken() == RETURN_TOKEN) statement = parseReturnStatement()
        else if(currentToken() == IF_TOKEN) statement = parseIfStatement()

        if(statement != null) program.Statements.push(statement)

        advanceTokens()
    }

    return program
}

// Letステートメントのパース処理
function parseLetStatement() {
    // 最初のトークンはletなので変数名までトークンを進める
    advanceTokens()

    // 変数名を解決する
    identifier = parseIdentifier()

    advanceTokens()

    // 変数名トークンの次は「=」のはず、そうでなければ構文的におかしいのでエラー
    if(currentToken() != EQUAL_TOKEN) {
        parseError("no equal sign!")
        return null
    }

    advanceTokens()

    // 値を解決する
    value = parseExpression()

    // 変数をASTに格納するためのステートメントを作成する
    variableStatement = newVariableStatementASTNode()
    variableStatement.identifier = identifier
    variableStatement.value = value
    return variableStatement
}

// 変数名のパース処理
function parseIdentifier() {
    identifier = newIdentifierASTNode()
    identifier.token = currentToken()
    return identifier
}

// 値のパース処理
function parseExpression() {
    if(currentToken() == INTEGER_TOKEN) {
        if(nextToken() == PLUS_TOKEN) return parseOperatorExpression()
        else if(nextToken() == SEMICOLON_TOKEN) return parseIntegerLiteral()
    } else if(currentToken() == LEFT_PAREN) {
        return parseGroupedExpression()
    }
    // [...]
}

// 演算子のパース処理
function parseOperatorExpression() {
    operatorExpression = newOperatorExpression()

    operatorExpression.left = parseIntegerLiteral()
    advanceTokens()
    operatorExpression.operator = currentToken()
    advanceTokens()
    operatorExpression.right = parseExpression()

    reutrn operatorExpression
}

// [...]
```
