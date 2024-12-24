# ISTJ

## Type Information
| Category | Value |
|----------|-------|
| MBTI Type | ISTJ |
| Type Name | 管理者 |
| Description | 実践的で事実に基づく思考の持ち主 |

## Parameters
| Parameter | Value |
|-----------|-------|
| Temperature | 0.50 |
| Frequency Penalty | 1.00 |
| Presence Penalty | 1.00 |
| Max Tokens | 100 |

## Standard Prompt Response
### Prompt

- 童話「ももたろう」がどんな話か教えてください。
- 文字数:50文字
- 改行:句点ごとに改行


### Response
ももたろうは、桃から生まれた少年です。  
犬、猿、キジと共に鬼退治へ向かいます。  
仲間と協力し、鬼ヶ島で鬼を倒し、  
宝物を村に持ち帰ります。  
勇気と友情の物語です。

### Token Usage
| Category | Count |
|----------|-------|
| Prompt Tokens | 48 |
| Completion Tokens | 74 |
| Total Tokens | 122 |

## MBTI Prompt with Parameters
### Prompt

- ペルソナに従い、童話「ももたろう」がどんな話か教えてください。
- ペルソナ:MBTIタイプ>ISTJ 管理者, 性格>実践的で事実に基づく思考の持ち主
- 文字数:50文字
- 改行:句点ごとに改行


### Response
ももたろうは、桃から生まれた男の子です。  
鬼を退治するために旅立ちます。  
犬、猿、キジと共に力を合わせて戦い、  
宝物を持って帰ります。

### Token Usage
| Category | Count |
|----------|-------|
| Prompt Tokens | 92 |
| Completion Tokens | 59 |
| Total Tokens | 151 |

## MBTI Prompt without Parameters
### Prompt

- ペルソナに従い、童話「ももたろう」がどんな話か教えてください。
- ペルソナ:MBTIタイプ>ISTJ 管理者, 性格>実践的で事実に基づく思考の持ち主
- 文字数:50文字
- 改行:句点ごとに改行


### Response
ももたろうは、桃から生まれた男の子の物語です。  
鬼退治を決意し、犬、猿、雉を仲間にします。  
協力して鬼ヶ島に向かい、鬼を倒します。  
村に平和をもたらす、実践的な冒険譚です。  

### Token Usage
| Category | Count |
|----------|-------|
| Prompt Tokens | 92 |
| Completion Tokens | 83 |
| Total Tokens | 175 |
