package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/sashabaranov/go-openai"
)

type MBTIConfig struct {
	MBTIType         string
	Temperature      float32
	FrequencyPenalty float32
	PresencePenalty  float32
	MaxTokens        int
	Character        string
	Description      string
}

// var mbtiConfigs = []MBTIConfig{
// 	{"ENFJ", 1.2, 0.6, 0.8, 250, "主人公", "カリスマ的で思いやりのある指導者型"},
// 	{"ENFP", 1.3, 0.5, 0.6, 300, "運動家", "情熱的で創造的な自由人"},
// 	{"ENTJ", 1.2, 0.6, 0.8, 250, "指揮官", "大胆不敵な想像力豊かな指導者型"},
// 	{"ENTP", 1.3, 0.5, 0.6, 300, "討論者", "知的で好奇心旺盛な思想家"},
// 	{"ESFJ", 1.2, 0.7, 1.1, 220, "領事", "気配り上手な思いやりあるサポーター"},
// 	{"ESFP", 1.3, 0.5, 0.5, 300, "エンターテイナー", "自由奔放で人生を楽しむ芸能人気質"},
// 	{"ESTJ", 1.1, 0.8, 1.2, 200, "幹部", "実務的で事実重視の管理者"},
// 	{"ESTP", 1.3, 0.6, 0.5, 250, "起業家", "機転の利く危険を恐れない実践者"},
// 	{"INFJ", 0.6, 0.8, 0.7, 150, "提唱者", "静かな理想主義者で思想家"},
// 	{"INFP", 0.7, 0.6, 0.5, 200, "仲介者", "理想主義的で誠実な調停者"},
// 	{"INTJ", 0.7, 0.8, 0.7, 150, "建築家", "想像力豊かな戦略的思考の持ち主"},
// 	{"INTP", 0.8, 0.7, 0.6, 200, "論理学者", "革新的な発明家肌の思想家"},
// 	{"ISFJ", 0.6, 0.9, 1.0, 120, "擁護者", "献身的で心優しい守護者"},
// 	{"ISFP", 0.7, 0.8, 0.7, 180, "冒険家", "柔軟で魅力的な芸術家"},
// 	{"ISTJ", 0.5, 1.0, 1.0, 100, "管理者", "実践的で事実に基づく思考の持ち主"},
// 	{"ISTP", 0.6, 0.9, 0.8, 150, "巨匠", "大胆で実践的な実験者"},
// }

var mbtiConfigs_refactor = []MBTIConfig{
	{"ENFJ", 1.1, 0.7, 0.8, 200, "主人公", "カリスマ的で思いやりのある指導者型"},
	{"ENFP", 1.2, 0.6, 0.7, 250, "運動家", "情熱的で創造的な自由人"},
	{"ENTJ", 1.1, 0.7, 0.8, 200, "指揮官", "大胆不敵な想像力豊かな指導者型"},
	{"ENTP", 1.2, 0.6, 0.7, 250, "討論者", "知的で好奇心旺盛な思想家"},
	{"ESFJ", 1.0, 0.8, 0.9, 180, "領事", "気配り上手な思いやりあるサポーター"},
	{"ESFP", 1.2, 0.6, 0.6, 250, "エンターテイナー", "自由奔放で人生を楽しむ芸能人気質"},
	{"ESTJ", 0.9, 0.9, 1.0, 180, "幹部", "実務的で事実重視の管理者"},
	{"ESTP", 1.2, 0.7, 0.6, 220, "起業家", "機転の利く危険を恐れない実践者"},
	{"INFJ", 0.7, 0.8, 0.8, 180, "提唱者", "静かな理想主義者で思想家"},
	{"INFP", 0.8, 0.7, 0.6, 200, "仲介者", "理想主義的で誠実な調停者"},
	{"INTJ", 0.8, 0.8, 0.8, 180, "建築家", "想像力豊かな戦略的思考の持ち主"},
	{"INTP", 0.9, 0.7, 0.7, 200, "論理学者", "革新的な発明家肌の思想家"},
	{"ISFJ", 0.7, 0.9, 0.9, 150, "擁護者", "献身的で心優しい守護者"},
	{"ISFP", 0.8, 0.8, 0.7, 180, "冒険家", "柔軟で魅力的な芸術家"},
	{"ISTJ", 0.6, 1.0, 0.9, 150, "管理者", "実践的で事実に基づく思考の持ち主"},
	{"ISTP", 0.7, 0.9, 0.8, 180, "巨匠", "大胆で実践的な実験者"},
}

const prompt = `
- 童話「ももたろう」がどんな話か教えてください。
- 文字数:50文字
- 改行:句点ごとに改行
`

const promptMBTI = `
- ペルソナに従い、童話「ももたろう」がどんな話か教えてください。
- ペルソナ:MBTIタイプ>%s %s, 性格>%s
- 文字数:50文字
- 改行:句点ごとに改行
`
const template = `## %s

| 項目 | 詳細 |
|------|------|
| MBTI タイプ | %s |
| タイプ名 | %s |
| 基本性格 | %s |

### AIパラメーター

| パラメーター | 値 |
|-------------|-----|
| Temperature | %.2f |
| Frequency Penalty | %.2f |
| Presence Penalty | %.2f |
| Max Tokens | %d |

### 生成結果

| プロンプト | 生成内容 |
|------------|----------|
| Standard | %s |
| MBTI with Parameters | %s |
| MBTI without Parameters | %s |

### トークン使用量

| メトリック | Standard | MBTI with Params | MBTI without Params |
|------------|----------|-------------------|---------------------|
| Prompt Tokens | %d | %d | %d |
| Completion Tokens | %d | %d | %d |
| Total Tokens | %d | %d | %d |
`

func main() {
	// 現在の日付を取得する
	// OpenAI APIを使用し商品レビューを生成
	key := os.Getenv("OPENAI_API_KEY")
	if key == "" {
		fmt.Println("OPENAI_API_KEY is not set")
		return
	}
	client := openai.NewClient(key)

	// 結果を保存するための文字列ビルダー
	var results strings.Builder

	intro := fmt.Sprintf(`# ももたろう感想生成比較
## 利用プロンプト
| Type           | Content |
|----------------|---------|
| Standard       | %s      |
| MBTI           | %s      |

## 利用AIパラメーター
MBTIタイプごとに異なるAIパラメーターを使用してるため各セクションにて記載

`,
		formatMarkdown(prompt),
		formatMarkdown(fmt.Sprintf(promptMBTI, "[MBTI]", "[名称]", "[基本性格]")))

	results.WriteString(intro)

	//

	for _, mbti := range mbtiConfigs_refactor {
		fmt.Printf("Running for %s\n", mbti.MBTIType)
		// 1. 通常のプロンプト実行
		normalResp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT4oMini,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    "user",
						Content: prompt,
					},
				},
			},
		)
		if err != nil {
			log.Printf("Error with normal prompt for %s: %v", mbti.MBTIType, err)
			continue
		}

		// 2. MBTIパラメータありのプロンプト実行
		mbtiParamResp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT4oMini,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    "user",
						Content: fmt.Sprintf(promptMBTI, mbti.MBTIType, mbti.Character, mbti.Description),
					},
				},
				Temperature:      mbti.Temperature,
				FrequencyPenalty: mbti.FrequencyPenalty,
				PresencePenalty:  mbti.PresencePenalty,
				MaxTokens:        mbti.MaxTokens,
			},
		)
		if err != nil {
			log.Printf("Error with MBTI param prompt for %s: %v", mbti.MBTIType, err)
			continue
		}

		// 3. MBTIパラメータなしのプロンプト実行
		mbtiNoParamResp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT4oMini,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    "user",
						Content: fmt.Sprintf(promptMBTI, mbti.MBTIType, mbti.Character, mbti.Description),
					},
				},
			},
		)
		if err != nil {
			log.Printf("Error with MBTI no-param prompt for %s: %v", mbti.MBTIType, err)
			continue
		}

		// マークダウン形式で結果を生成
		result := fmt.Sprintf(template,
			mbti.MBTIType,
			mbti.MBTIType,
			mbti.Character,
			mbti.Description,
			mbti.Temperature,
			mbti.FrequencyPenalty,
			mbti.PresencePenalty,
			mbti.MaxTokens,
			// 応答内容
			formatMarkdown(normalResp.Choices[0].Message.Content),
			formatMarkdown(mbtiParamResp.Choices[0].Message.Content),
			formatMarkdown(mbtiNoParamResp.Choices[0].Message.Content),
			// トークン使用量
			normalResp.Usage.PromptTokens,
			mbtiParamResp.Usage.PromptTokens,
			mbtiNoParamResp.Usage.PromptTokens,
			normalResp.Usage.CompletionTokens,
			mbtiParamResp.Usage.CompletionTokens,
			mbtiNoParamResp.Usage.CompletionTokens,
			normalResp.Usage.TotalTokens,
			mbtiParamResp.Usage.TotalTokens,
			mbtiNoParamResp.Usage.TotalTokens,
		)

		results.WriteString(result)
	}

	// 結果をファイルに保存
	if err := saveResults(results.String()); err != nil {
		log.Fatalf("Failed to save results: %v", err)
	}
}

// Save results to a single Markdown file
func saveResults(content string) error {
	rootDir, err := findProjectRoot()
	if err != nil {
		return fmt.Errorf("failed to find project root: %v", err)
	}

	// result.mdのパスを生成
	resultPath := filepath.Join(rootDir, "RESULT.md")

	// ファイルに書き込み
	if err := os.WriteFile(resultPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write results to file: %v", err)
	}

	// result.md に目次を最上部につける

	log.Printf("Results saved to: %s", resultPath)
	return nil
}

// プロジェクトルートディレクトリを見つける関数
func findProjectRoot() (string, error) {
	// カレントディレクトリから開始
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// go.modファイルを探してプロジェクトルートを特定
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("could not find project root (no go.mod file found)")
		}
		dir = parent
	}
}

// 改行を適切に処理する関数
func formatMarkdown(content string) string {
	// 改行を<br>に変換
	return strings.ReplaceAll(content, "\n", "<br>")
}
