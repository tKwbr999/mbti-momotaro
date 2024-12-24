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
ももたろうは、桃から生まれた男の子です。  
成長した彼は鬼を退治するために旅に出ます。  
犬、猿、キジを仲間にし、鬼ヶ島で鬼を倒し、  
村に平和をもたらす物語です。  

### Token Usage
| Category | Count |
|----------|-------|
| Prompt Tokens | 48 |
| Completion Tokens | 75 |
| Total Tokens | 123 |

## MBTI Prompt with Parameters
### Prompt

- MBTIタイプの性格をもったあなたが童話「ももたろう」がどんな話か教えてください。
- MBTIタイプ:ISTJ
- 文字数:50文字
- 改行:句点ごとに改行


### Response
ももたろうは、桃から生まれた男の子です。  
彼は鬼退治に出かけます。  
犬、猿、キジを仲間にし、  
協力して鬼を倒します。  
村人を救い、平和が訪れます。  

### Token Usage
| Category | Count |
|----------|-------|
| Prompt Tokens | 66 |
| Completion Tokens | 69 |
| Total Tokens | 135 |

## MBTI Prompt without Parameters
### Prompt

- MBTIタイプの性格をもったあなたが童話「ももたろう」がどんな話か教えてください。
- MBTIタイプ:ISTJ
- 文字数:50文字
- 改行:句点ごとに改行


### Response
ももたろうは、桃から生まれた男の子です。  
彼は鬼を退治するため仲間を集めます。  
犬、猿、キジと共に、鬼ヶ島へ向かいます。  
鬼を倒して、村に平和をもたらします。

### Token Usage
| Category | Count |
|----------|-------|
| Prompt Tokens | 66 |
| Completion Tokens | 70 |
| Total Tokens | 136 |
