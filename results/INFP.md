# INFP

## Type Information
| Category | Value |
|----------|-------|
| MBTI Type | INFP |
| Type Name | 仲介者 |
| Description | 理想主義的で誠実な調停者 |

## Parameters
| Parameter | Value |
|-----------|-------|
| Temperature | 0.70 |
| Frequency Penalty | 0.60 |
| Presence Penalty | 0.50 |
| Max Tokens | 200 |

## Standard Prompt Response
### Prompt

- 童話「ももたろう」がどんな話か教えてください。
- 文字数:50文字
- 改行:句点ごとに改行


### Response
ももたろうは、日本の妖怪退治の物語です。  
桃から生まれた男の子が、犬、猿、キジと仲間になり、  
鬼ヶ島へ向かって鬼を倒し、村人を救います。  

### Token Usage
| Category | Count |
|----------|-------|
| Prompt Tokens | 48 |
| Completion Tokens | 62 |
| Total Tokens | 110 |

## MBTI Prompt with Parameters
### Prompt

- ペルソナに従い、童話「ももたろう」がどんな話か教えてください。
- ペルソナ:MBTIタイプ>INFP 仲介者, 性格>理想主義的で誠実な調停者
- 文字数:50文字
- 改行:句点ごとに改行


### Response
桃から生まれた桃太郎が、  
犬、猿、キジと共に鬼ヶ島へ。  
悪い鬼を退治し、宝物を取り戻す。  
友情と勇気の物語です。

### Token Usage
| Category | Count |
|----------|-------|
| Prompt Tokens | 89 |
| Completion Tokens | 54 |
| Total Tokens | 143 |

## MBTI Prompt without Parameters
### Prompt

- ペルソナに従い、童話「ももたろう」がどんな話か教えてください。
- ペルソナ:MBTIタイプ>INFP 仲介者, 性格>理想主義的で誠実な調停者
- 文字数:50文字
- 改行:句点ごとに改行


### Response
ももたろうは、桃から生まれた男の子のお話です。  
彼は鬼退治のため、犬、猿、鳥と共に旅をします。  
友情や勇気を通じて、悪を討つ理想的な物語です。

### Token Usage
| Category | Count |
|----------|-------|
| Prompt Tokens | 89 |
| Completion Tokens | 65 |
| Total Tokens | 154 |
