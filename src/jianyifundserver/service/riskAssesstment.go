package service

import (
	simplejson "github.com/bitly/go-simplejson"
	"github.com/astaxie/beego"
)

var InstitutionScoreMatrix = map[string]map[string]int{"question1": map[string]int{"A": 2, "B": 1, "C": 3, "D": 4},
"question2": map[string]int{"A": 0, "B": 1, "C": 2, "D": 3, "E": 4},
"question3": map[string]int{"A": 4, "B": 3, "C": 2, "D": 1, "E": 0},
"question4": map[string]int{"A": 0, "B": 2},
"question5": map[string]int{"A": 3, "B": 2, "C": 1, "D": 0},
"question6": map[string]int{"A": 1, "B": 2, "C": 3, "D": 4},
"question7": map[string]int{"A": 1, "B": 2, "C": 3, "D": 4},
"question8": map[string]int{"A": 0, "B": 1, "C": 2, "D": 3, "E": 4},
"question9": map[string]int{"A": 1, "B": 3, "C": 4, "D": 5},
"question10": map[string]int{"A": 1, "B": 2, "C": 4},
"question11": map[string]int{"A": 1, "B": 2, "C": 4, "D": 8},
"question12": map[string]int{"A": 1, "B": 3, "C": 7},
"question13": map[string]int{"A": 1, "B": 2, "C": 4, "D": 6, "E": 8},
"question14": map[string]int{"A": 1, "B": 2, "C": 3, "D": 4},
"question15": map[string]int{"A": 1, "B": 3, "C": 4, "D": 6},
"question16": map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 0},
"question17": map[string]int{"A": 1, "B": 2, "C": 8, "D": 10},
"question18": map[string]int{"A": 1, "B": 2, "C": 8, "D": 10},
"question19": map[string]int{"A": 1, "B": 4, "C": 8, "D": 10, "E": 10},
"question20": map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5},
"question21": map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5},
"question22": map[string]int{"A": 3, "B": 5, "C": 4, "D": 1},
}

func GetInstitutionRiskGrade (score int) (grade string) {
	switch {
	case score >= 16 && score <= 42:
		grade = "C1 保守型(适合低风险产品或服务)"
	case score >= 43 && score <= 66:
		grade = "C2 谨慎型(适合低风险及中低风险产品或服务) "
	case score >= 67 && score <= 88:
		grade = "C3 稳健型(适合低风险、中低风险及中风险产品或服务) "
	case score >= 89 && score <= 100:
		grade = "C4 积极型(适合低风险、中低风险、中风险及中高风险产品或服务) "
	case score >= 101 && score <= 120:
		grade = "C5 进取型(适合低风险、中低风险、中风险、中高风险及高风险产品或服务) "
	default:
		grade = "特别保护型(适合低风险产品或服务且不得购买或接受其他风险等级产品和服务)" 
	}
	return
}

var PersonalScoreMatrix = map[string]map[string]int{"question1": map[string]int{"A": 0, "B": 2},
"question2": map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5},
"question3": map[string]int{"A": 3, "B": 2, "C": 1, "D": 1, "E": 0},
"question4": map[string]int{"A": 1, "B": 2, "C": 3, "D": 4},
"question5": map[string]int{"A": 5, "B": 4, "C": 3, "D": 2, "E": 1},
"question6": map[string]int{"A": 1, "B": 2, "C": 3, "D": 4},
"question7": map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5},
"question8": map[string]int{"A": 1, "B": 2, "C": 4},
"question9": map[string]int{"A": 1, "B": 2, "C": 4, "D": 8},
"question10": map[string]int{"A": 1, "B": 2, "C": 3, "D": 4},
"question11": map[string]int{"A": 1, "B": 2, "C": 4, "D": 6},
"question12": map[string]int{"A": 0, "B": 1, "C": 2, "D": 3, "E": 4},
"question13": map[string]int{"A": 1, "B": 2, "C": 3, "D": 4},
"question14": map[string]int{"A": 1, "B": 2, "C": 4, "D": 8},
"question15": map[string]int{"A": 1, "B": 2, "C": 4, "D": 8, "E": 8},
"question16": map[string]int{"A": 0, "B": 1, "C": 3, "D": 5, "E": 6},
"question17": map[string]int{"A": 1, "B": 2, "C": 4, "D": 8, "E": 10},
"question18": map[string]int{"A": 1, "B": 2, "C": 4, "D": 8},
"question19": map[string]int{"A": 0, "B": 2, "C": 1},
"question20": map[string]int{"A": 6, "B": 6, "C": 4, "D": 4, "E": 2, "F": 2, "G": 1, "H": 1},
"question21": map[string]int{"A": 8, "B": 6, "C": 4, "D": 2, "E": 1},
"question22": map[string]int{"A": 6, "B": 4, "C": 2, "D": 1},
}

func GetPersonalRiskGrade (score int) (grade string) {
	switch {
	case score >= 17 && score <= 34:
		grade = "C1 保守型(适合低风险产品或服务)"
	case score >= 35 && score <= 52:
		grade = "C2 谨慎型(适合低风险及中低风险产品或服务) "
	case score >= 53 && score <= 82:
		grade = "C3 稳健型(适合低风险、中低风险及中风险产品或服务) "
	case score >= 83 && score <= 98:
		grade = "C4 积极型(适合低风险、中低风险、中风险及中高风险产品或服务) "
	case score >= 99 && score <= 120:
		grade = "C5 进取型(适合低风险、中低风险、中风险、中高风险及高风险产品或服务) "
	default:
		grade = "特别保护型(适合低风险产品或服务且不得购买或接受其他风险等级产品和服务)" 
	}
	return
}



func CalRiskAssessment (AnswerDetail string, InvestorType int) (Score int, ScoreGrade string) {
	res, err := simplejson.NewJson([]byte(AnswerDetail))
		if err != nil {
			beego.Error(err.Error())
			return 
		}
	beego.Debug(res)
	switch InvestorType {
	case 1:
		beego.Debug("个人风险测评")
		
		for k, v := range PersonalScoreMatrix {
			if data, ok := res.CheckGet(k); ok {
				Score = Score + v[data.MustString()]
			}
		}
		ScoreGrade = GetPersonalRiskGrade(Score)
		beego.Debug("score: %d", Score)
		beego.Debug("scoreGrade: %s", ScoreGrade)

	case 2:
		beego.Debug("机构风险测评")
		for k, v := range PersonalScoreMatrix {
			if data, ok := res.CheckGet(k); ok {
				Score = Score + v[data.MustString()]
			}
		}
		ScoreGrade = GetPersonalRiskGrade(Score)
		beego.Debug("score: %d", Score)
		beego.Debug("scoreGrade: %s", ScoreGrade)
	}
	return 
}