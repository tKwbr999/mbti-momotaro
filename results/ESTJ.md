# ESTJ

## Type Information
| Category | Value |
|----------|-------|
| MBTI Type | ESTJ |
| Type Name | 幹部 |
| Description | 実務的で事実重視の管理者 |

## Parameters
| Parameter | Value |
|-----------|-------|
| Temperature | 1.10 |
| Frequency Penalty | 0.80 |
| Presence Penalty | 1.20 |
| Max Tokens | 200 |

## Standard Prompt Response
### Prompt

- 童話「ももたろう」がどんな話か教えてください。
- 文字数:50文字
- 改行:句点ごとに改行


### Response
「ももたろう」は、日本の童話です。  
ある日、桃から生まれた男の子がいました。  
彼は鬼を退治するため、犬、猿、キジを仲間にします。  
共に力を合わせ、鬼の島へ向かい、宝を持ち帰ります。  

### Token Usage
| Category | Count |
|----------|-------|
| Prompt Tokens | 48 |
| Completion Tokens | 76 |
| Total Tokens | 124 |

## MBTI Prompt with Parameters
### Prompt

- ペルソナに従い、童話「ももたろう」がどんな話か教えてください。
- ペルソナ:MBTIタイプ>ESTJ 幹部, 性格>実務的で事実重視の管理者
- 文字数:50文字
- 改行:句点ごとに改行


### Response
「ももたろう」は、桃から生まれた男の子が鬼退治をする物語です。  
犬、猿、キジを仲間にし、団結して戦い勝利します。  
勇気と友愛の大切さを教えてくれます。

### Token Usage
| Category | Count |
|----------|-------|
| Prompt Tokens | 87 |
| Completion Tokens | 68 |
| Total Tokens | 155 |

## MBTI Prompt without Parameters
### Prompt

- ペルソナに従い、童話「ももたろう」がどんな話か教えてください。
- ペルソナ:MBTIタイプ>ESTJ 幹部, 性格>実務的で事実重視の管理者
- 文字数:50文字
- 改行:句点ごとに改行


### Response
ももたろうは、日本の伝説的な童話です。  
桃から生まれた男の子が鬼を退治し、  
村を救う物語です。  
実行力があり、仲間と協力して  
目的を達成する姿が描かれています。  

### Token Usage
| Category | Count |
|----------|-------|
| Prompt Tokens | 87 |
| Completion Tokens | 70 |
| Total Tokens | 157 |
