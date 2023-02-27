package db

import (
	"context"
	"fmt"
	"time"
	"usr/local/go/server/gen/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// This function is used when we want to create test/demo data/
func CreateData(ctx context.Context, db boil.ContextExecutor) error {
	companies := []models.Company{
		{ID: 1, Name: "株式会社川田"},
		{ID: 2, Name: "株式会社kawata"}}

	for _, company := range companies {
		if err := company.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create demo datas of company. err:%s", err)
		}
	}

	employees := []models.Employee{
		{ID: 1, CompanyID: 1, Name: "Yumi", MailAddress: "yumi@example.com", Password: "testpass"},
		{ID: 2, CompanyID: 1, Name: "James", MailAddress: "james@example.com", Password: "testpass"},
		{ID: 3, CompanyID: 1, Name: "Potter", MailAddress: "potter@example.com", Password: "testpass"},
		{ID: 4, CompanyID: 1, Name: "Kevin", MailAddress: "kevin@example.com", Password: "testpass"},
		{ID: 5, CompanyID: 1, Name: "Tom", MailAddress: "tom@example.com", Password: "testpass"},
		{ID: 6, CompanyID: 1, Name: "Ayumi", MailAddress: "ayumi@example.com", Password: "testpass"},
		{ID: 7, CompanyID: 1, Name: "Bob", MailAddress: "bob@example.com", Password: "testpass"},
		{ID: 8, CompanyID: 1, Name: "Yuji", MailAddress: "yuji@example.com", Password: "testpass"},
		{ID: 9, CompanyID: 1, Name: "Charo", MailAddress: "charo@example.com", Password: "testpass"},
		{ID: 10, CompanyID: 1, Name: "Eron", MailAddress: "eron@example.com", Password: "testpass"},
		{ID: 11, CompanyID: 1, Name: "Jim", MailAddress: "jim@example.com", Password: "testpass"},
		{ID: 12, CompanyID: 1, Name: "Steve", MailAddress: "steve@example.com", Password: "testpass"}}

	for _, employee := range employees {
		if err := employee.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create datas of employee. err:%s", err)
		}
	}

	administrators := []models.Administrator{
		{ID: 1, CompanyID: 1, Name: "administrator", Password: "testpass"}}

	for _, administrator := range administrators {
		if err := administrator.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create datas of administrator. err:%s", err)
		}
	}

	individual_earnings := []models.IndividualEarning{
		{ID: 1, Nominal: "支給", Amount: 300000},
		{ID: 2, Nominal: "支給", Amount: 350000},
		{ID: 3, Nominal: "支給", Amount: 250000},
		{ID: 4, Nominal: "支給", Amount: 200000},
		{ID: 5, Nominal: "支給", Amount: 400000},
		{ID: 6, Nominal: "支給", Amount: 450000},
		{ID: 7, Nominal: "支給", Amount: 500000},
		{ID: 8, Nominal: "支給", Amount: 350000},
		{ID: 9, Nominal: "支給", Amount: 300000}}
	
	for _, individual_earning := range individual_earnings {
		if err := individual_earning.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create datas of individual_earning. err:%s", err)
		}
	}

	individual_earning_details := []models.IndividualEarningDetail{
		{ID: 1, IndividualEarningID: 1, Nominal: "基本給", Amount: 250000},
		{ID: 2, IndividualEarningID: 1, Nominal: "住宅手当", Amount: 20000},
		{ID: 3, IndividualEarningID: 1, Nominal: "通勤手当/非課税", Amount: 30000},
		{ID: 4, IndividualEarningID: 2, Nominal: "基本給", Amount: 300000},
		{ID: 5, IndividualEarningID: 2, Nominal: "残業手当", Amount: 50000},
		{ID: 6, IndividualEarningID: 3, Nominal: "基本給", Amount: 250000},
		{ID: 7, IndividualEarningID: 4, Nominal: "基本給", Amount: 200000},
		{ID: 8, IndividualEarningID: 5, Nominal: "基本給", Amount: 250000},
		{ID: 9, IndividualEarningID: 5, Nominal: "残業手当", Amount: 50000},
		{ID: 10, IndividualEarningID: 5, Nominal: "エンジニア手当", Amount: 100000},
		{ID: 11, IndividualEarningID: 6, Nominal: "基本給", Amount: 250000},
		{ID: 12, IndividualEarningID: 6, Nominal: "残業手当", Amount: 150000},
		{ID: 13, IndividualEarningID: 7, Nominal: "10月賞与", Amount: 500000},
		{ID: 14, IndividualEarningID: 8, Nominal: "10月賞与", Amount: 350000},
		{ID: 15, IndividualEarningID: 9, Nominal: "基本給", Amount: 200000},
		{ID: 16, IndividualEarningID: 9, Nominal: "残業手当", Amount: 50000},
		{ID: 17, IndividualEarningID: 9, Nominal: "通勤手当/非課税", Amount: 50000}}

	for _, individual_earning_detail := range individual_earning_details {
		if err := individual_earning_detail.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create datas of individual_earning_detail. err:%s", err)
		}
	}

	fixed_earnings := []models.FixedEarning{
		{ID: 1, Nominal: "エンジニアLv1家族手当0人", Amount: 300000},
		{ID: 2, Nominal: "エンジニアLv1家族手当1人", Amount: 350000},
		{ID: 3, Nominal: "エンジニアLv1家族手当2人", Amount: 400000},
		{ID: 4, Nominal: "エンジニアLv1家族手当3人", Amount: 450000},
		{ID: 5, Nominal: "年俸360万１ヶ月分", Amount: 300000},
		{ID: 6, Nominal: "年俸420万１ヶ月分", Amount: 350000},
		{ID: 7, Nominal: "年俸600万１ヶ月分", Amount: 500000},
		{ID: 8, Nominal: "Jimさんの給料１ヶ月分", Amount: 300000},
		{ID: 9, Nominal: "Steveさんの給料１ヶ月分", Amount: 350000}}
	
	for _, fixed_earning := range fixed_earnings {
		if err := fixed_earning.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create datas of fixed_earning. err:%s", err)
		}
	}

	fixed_earning_details := []models.FixedEarningDetail{
		{ID: 1, FixedEarningID: 1, Nominal: "基本給", Amount: 250000},
		{ID: 2, FixedEarningID: 1, Nominal: "固定残業代", Amount: 50000},
		{ID: 3, FixedEarningID: 2, Nominal: "基本給", Amount: 250000},
		{ID: 4, FixedEarningID: 2, Nominal: "固定残業代", Amount: 50000},
		{ID: 5, FixedEarningID: 2, Nominal: "家族手当(1人分)", Amount: 50000},
		{ID: 6, FixedEarningID: 3, Nominal: "基本給", Amount: 250000},
		{ID: 7, FixedEarningID: 3, Nominal: "固定残業代", Amount: 50000},
		{ID: 8, FixedEarningID: 3, Nominal: "家族手当(2人分)", Amount: 100000},
		{ID: 9, FixedEarningID: 4, Nominal: "基本給", Amount: 250000},
		{ID: 10, FixedEarningID: 4, Nominal: "固定残業代", Amount: 50000},
		{ID: 11, FixedEarningID: 4, Nominal: "家族手当(3人分)", Amount: 150000},
		{ID: 12, FixedEarningID: 5, Nominal: "年俸360万１ヶ月分", Amount: 300000},
		{ID: 13, FixedEarningID: 6, Nominal: "年俸420万１ヶ月分", Amount: 350000},
		{ID: 14, FixedEarningID: 7, Nominal: "年俸600万１ヶ月分", Amount: 500000},
		{ID: 15, FixedEarningID: 8, Nominal: "基本給", Amount: 250000},
		{ID: 16, FixedEarningID: 8, Nominal: "固定残業代", Amount: 50000},
		{ID: 17, FixedEarningID: 9, Nominal: "基本給", Amount: 250000},
		{ID: 18, FixedEarningID: 9, Nominal: "固定残業代", Amount: 50000},
		{ID: 19, FixedEarningID: 9, Nominal: "家族手当(1人分)", Amount: 50000}}
	
	for _, fixed_earning_detail := range fixed_earning_details {
		if err := fixed_earning_detail.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create datas of fixed_earning_detail. err:%s", err)
		}
	}

	individual_deductions := []models.IndividualDeduction{
		{ID: 1, Amount: 60000, Nominal: "控除"},
		{ID: 2, Amount: 70000, Nominal: "控除"},
		{ID: 3, Amount: 50000, Nominal: "控除"},
		{ID: 4, Amount: 40000, Nominal: "控除"},
		{ID: 5, Amount: 80000, Nominal: "控除"},
		{ID: 6, Amount: 90000, Nominal: "控除"},
		{ID: 7, Amount: 100000, Nominal: "控除"},
		{ID: 8, Amount: 70000, Nominal: "控除"},
		{ID: 9, Amount: 60000, Nominal: "控除"}}
	
	for _, individual_deduction := range individual_deductions {
		if err := individual_deduction.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create datas of individual_deduction. err:%s", err)
		}
	}

	individual_deduction_details := []models.IndividualDeductionDetail{
		{ID: 1, IndividualDeductionID: 1, Nominal: "健康保険料", Amount: 12000},
		{ID: 2, IndividualDeductionID: 1, Nominal: "厚生年金保険料", Amount: 20000},
		{ID: 3, IndividualDeductionID: 1, Nominal: "雇用保険料", Amount: 300},
		{ID: 4, IndividualDeductionID: 1, Nominal: "所得税", Amount: 5000},
		{ID: 5, IndividualDeductionID: 1, Nominal: "住民税", Amount: 22700},
		{ID: 6, IndividualDeductionID: 2, Nominal: "健康保険料", Amount: 13000},
		{ID: 7, IndividualDeductionID: 2, Nominal: "厚生年金保険料", Amount: 21000},
		{ID: 8, IndividualDeductionID: 2, Nominal: "雇用保険料", Amount: 500},
		{ID: 9, IndividualDeductionID: 2, Nominal: "所得税", Amount: 6000},
		{ID: 10, IndividualDeductionID: 2, Nominal: "住民税", Amount: 29500},
		{ID: 11, IndividualDeductionID: 3, Nominal: "健康保険料", Amount: 10000},
		{ID: 12, IndividualDeductionID: 3, Nominal: "厚生年金保険料", Amount: 18000},
		{ID: 13, IndividualDeductionID: 3, Nominal: "雇用保険料", Amount: 300},
		{ID: 14, IndividualDeductionID: 3, Nominal: "所得税", Amount: 4500},
		{ID: 15, IndividualDeductionID: 3, Nominal: "住民税", Amount: 17200},
		{ID: 16, IndividualDeductionID: 4, Nominal: "健康保険料", Amount: 9000},
		{ID: 17, IndividualDeductionID: 4, Nominal: "厚生年金保険料", Amount: 15000},
		{ID: 18, IndividualDeductionID: 4, Nominal: "雇用保険料", Amount: 300},
		{ID: 19, IndividualDeductionID: 4, Nominal: "所得税", Amount: 4000},
		{ID: 20, IndividualDeductionID: 4, Nominal: "住民税", Amount: 11700},
		{ID: 21, IndividualDeductionID: 5, Nominal: "健康保険料", Amount: 18000},
		{ID: 22, IndividualDeductionID: 5, Nominal: "厚生年金保険料", Amount: 30000},
		{ID: 23, IndividualDeductionID: 5, Nominal: "雇用保険料", Amount: 600},
		{ID: 24, IndividualDeductionID: 5, Nominal: "所得税", Amount: 8000},
		{ID: 25, IndividualDeductionID: 5, Nominal: "住民税", Amount: 23400},
		{ID: 26, IndividualDeductionID: 6, Nominal: "健康保険料", Amount: 18000},
		{ID: 27, IndividualDeductionID: 6, Nominal: "厚生年金保険料", Amount: 25000},
		{ID: 28, IndividualDeductionID: 6, Nominal: "雇用保険料", Amount: 600},
		{ID: 29, IndividualDeductionID: 6, Nominal: "所得税", Amount: 1000},
		{ID: 30, IndividualDeductionID: 6, Nominal: "住民税", Amount: 45400},
		{ID: 31, IndividualDeductionID: 7, Nominal: "健康保険料", Amount: 20000},
		{ID: 32, IndividualDeductionID: 7, Nominal: "厚生年金保険料", Amount: 20000},
		{ID: 33, IndividualDeductionID: 7, Nominal: "雇用保険料", Amount: 30000},
		{ID: 34, IndividualDeductionID: 7, Nominal: "所得税", Amount: 20000},
		{ID: 35, IndividualDeductionID: 7, Nominal: "住民税", Amount: 30000},
		{ID: 36, IndividualDeductionID: 8, Nominal: "健康保険料", Amount: 13000},
		{ID: 37, IndividualDeductionID: 8, Nominal: "厚生年金保険料", Amount: 21000},
		{ID: 38, IndividualDeductionID: 8, Nominal: "雇用保険料", Amount: 500},
		{ID: 39, IndividualDeductionID: 8, Nominal: "所得税", Amount: 6000},
		{ID: 40, IndividualDeductionID: 8, Nominal: "住民税", Amount: 29500},
		{ID: 41, IndividualDeductionID: 9, Nominal: "健康保険料", Amount: 12000},
		{ID: 42, IndividualDeductionID: 9, Nominal: "厚生年金保険料", Amount: 20000},
		{ID: 43, IndividualDeductionID: 9, Nominal: "雇用保険料", Amount: 300},
		{ID: 44, IndividualDeductionID: 9, Nominal: "所得税", Amount: 5000},
		{ID: 45, IndividualDeductionID: 9, Nominal: "住民税", Amount: 22700}}

	for _, individual_deduction_detail := range individual_deduction_details {
		if err := individual_deduction_detail.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create datas of individual_deduction_detail. err:%s", err)
		}
	}

	fixed_deductions := []models.FixedDeduction{
		{ID: 1, Nominal: "エンジニアLv1家族手当0人控除", Amount: 60000},
		{ID: 2, Nominal: "エンジニアLv1家族手当1人控除", Amount: 70000},
		{ID: 3, Nominal: "エンジニアLv1家族手当2人控除", Amount: 80000},
		{ID: 4, Nominal: "エンジニアLv1家族手当3人控除", Amount: 90000},
		{ID: 5, Nominal: "年俸360万１ヶ月分控除", Amount: 60000},
		{ID: 6, Nominal: "年俸420万１ヶ月分控除", Amount: 70000},
		{ID: 7, Nominal: "年俸600万１ヶ月分控除", Amount: 100000},
		{ID: 8, Nominal: "Jimさんの給料1ヶ月分控除", Amount: 60000},
		{ID: 9, Nominal: "Steveさんの給料1ヶ月分控除", Amount: 70000}}
		
	for _, fixed_deduction := range fixed_deductions {
		if err := fixed_deduction.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create datas of fixed_deduction. err:%s", err)
		}
	}

	fixed_deduction_details := []models.FixedDeductionDetail{
		{ID: 1, FixedDeductionID: 1, Nominal: "健康保険料", Amount: 12000},
		{ID: 2, FixedDeductionID: 1, Nominal: "厚生年金保険料", Amount: 20000},
		{ID: 3, FixedDeductionID: 1, Nominal: "雇用保険料", Amount: 300},
		{ID: 4, FixedDeductionID: 1, Nominal: "所得税", Amount: 5000},
		{ID: 5, FixedDeductionID: 1, Nominal: "住民税", Amount: 22700},
		{ID: 6, FixedDeductionID: 2, Nominal: "健康保険料", Amount: 13000},
		{ID: 7, FixedDeductionID: 2, Nominal: "厚生年金保険料", Amount: 21000},
		{ID: 8, FixedDeductionID: 2, Nominal: "雇用保険料", Amount: 500},
		{ID: 9, FixedDeductionID: 2, Nominal: "所得税", Amount: 6000},
		{ID: 10, FixedDeductionID: 2, Nominal: "住民税", Amount: 29500},
		{ID: 11, FixedDeductionID: 3, Nominal: "健康保険料", Amount: 18000},
		{ID: 12, FixedDeductionID: 3, Nominal: "厚生年金保険料", Amount: 30000},
		{ID: 13, FixedDeductionID: 3, Nominal: "雇用保険料", Amount: 600},
		{ID: 14, FixedDeductionID: 3, Nominal: "所得税", Amount: 8000},
		{ID: 15, FixedDeductionID: 3, Nominal: "住民税", Amount: 23400},
		{ID: 16, FixedDeductionID: 4, Nominal: "健康保険料", Amount: 18000},
		{ID: 17, FixedDeductionID: 4, Nominal: "厚生年金保険料", Amount: 25000},
		{ID: 18, FixedDeductionID: 4, Nominal: "雇用保険料", Amount: 600},
		{ID: 19, FixedDeductionID: 4, Nominal: "所得税", Amount: 1000},
		{ID: 20, FixedDeductionID: 4, Nominal: "住民税", Amount: 45400},
		{ID: 21, FixedDeductionID: 5, Nominal: "健康保険料", Amount: 12000},
		{ID: 22, FixedDeductionID: 5, Nominal: "厚生年金保険料", Amount: 20000},
		{ID: 23, FixedDeductionID: 5, Nominal: "雇用保険料", Amount: 300},
		{ID: 24, FixedDeductionID: 5, Nominal: "所得税", Amount: 5000},
		{ID: 25, FixedDeductionID: 5, Nominal: "住民税", Amount: 22700},
		{ID: 26, FixedDeductionID: 6, Nominal: "健康保険料", Amount: 13000},
		{ID: 27, FixedDeductionID: 6, Nominal: "厚生年金保険料", Amount: 21000},
		{ID: 28, FixedDeductionID: 6, Nominal: "雇用保険料", Amount: 500},
		{ID: 29, FixedDeductionID: 6, Nominal: "所得税", Amount: 6000},
		{ID: 30, FixedDeductionID: 6, Nominal: "住民税", Amount: 29500},
		{ID: 31, FixedDeductionID: 7, Nominal: "健康保険料", Amount: 20000},
		{ID: 32, FixedDeductionID: 7, Nominal: "厚生年金保険料", Amount: 20000},
		{ID: 33, FixedDeductionID: 7, Nominal: "雇用保険料", Amount: 30000},
		{ID: 34, FixedDeductionID: 7, Nominal: "所得税", Amount: 20000},
		{ID: 35, FixedDeductionID: 7, Nominal: "住民税", Amount: 30000},
		{ID: 36, FixedDeductionID: 8, Nominal: "健康保険料", Amount: 12000},
		{ID: 37, FixedDeductionID: 8, Nominal: "厚生年金保険料", Amount: 20000},
		{ID: 38, FixedDeductionID: 8, Nominal: "雇用保険料", Amount: 300},
		{ID: 39, FixedDeductionID: 8, Nominal: "所得税", Amount: 5000},
		{ID: 40, FixedDeductionID: 8, Nominal: "住民税", Amount: 22700},
		{ID: 41, FixedDeductionID: 9, Nominal: "健康保険料", Amount: 13000},
		{ID: 42, FixedDeductionID: 9, Nominal: "厚生年金保険料", Amount: 21000},
		{ID: 43, FixedDeductionID: 9, Nominal: "雇用保険料", Amount: 500},
		{ID: 44, FixedDeductionID: 9, Nominal: "所得税", Amount: 6000},
		{ID: 45, FixedDeductionID: 9, Nominal: "住民税", Amount: 29500}}

	for _, fixed_deduction_detail := range fixed_deduction_details {
		if err := fixed_deduction_detail.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create datas of fixed_deduction_detail. err:%s", err)
		}
	}

	salary_statements := []models.SalaryStatement{
		{ID: 1, IndividualEarningID: null.Uint32{Uint32: 1, Valid: true}, IndividualDeductionID: null.Uint32{Uint32: 1, Valid: true}, EmployeeID: 1, Nominal: "2021年10月 給与明細", Payday: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年10月1日~2022年10月31日"},
		{ID: 2, IndividualEarningID: null.Uint32{Uint32: 2, Valid: true}, IndividualDeductionID: null.Uint32{Uint32: 2, Valid: true}, EmployeeID: 2, Nominal: "2021年10月 給与明細", Payday: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年10月1日~2022年10月31日"},
		{ID: 3, IndividualEarningID: null.Uint32{Uint32: 3, Valid: true}, IndividualDeductionID: null.Uint32{Uint32: 3, Valid: true}, EmployeeID: 3, Nominal: "2021年10月 給与明細", Payday: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年10月1日~2022年10月31日"},
		{ID: 4, IndividualEarningID: null.Uint32{Uint32: 4, Valid: true}, IndividualDeductionID: null.Uint32{Uint32: 4, Valid: true}, EmployeeID: 4, Nominal: "2021年10月 給与明細", Payday: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年10月1日~2022年10月31日"},
		{ID: 5, IndividualEarningID: null.Uint32{Uint32: 5, Valid: true}, IndividualDeductionID: null.Uint32{Uint32: 5, Valid: true}, EmployeeID: 5, Nominal: "2021年10月 給与明細", Payday: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年10月1日~2022年10月31日"},
		{ID: 6, IndividualEarningID: null.Uint32{Uint32: 6, Valid: true}, IndividualDeductionID: null.Uint32{Uint32: 6, Valid: true}, EmployeeID: 6, Nominal: "2021年10月 給与明細", Payday: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年10月1日~2022年10月31日"},
		{ID: 7, IndividualEarningID: null.Uint32{Uint32: 7, Valid: true}, IndividualDeductionID: null.Uint32{Uint32: 7, Valid: true}, EmployeeID: 7, Nominal: "2021年10月 給与明細", Payday: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年10月1日~2022年10月31日"},
		{ID: 8, IndividualEarningID: null.Uint32{Uint32: 8, Valid: true}, IndividualDeductionID: null.Uint32{Uint32: 8, Valid: true}, EmployeeID: 8, Nominal: "2021年10月 給与明細", Payday: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年10月1日~2022年10月31日"},
		{ID: 9, IndividualEarningID: null.Uint32{Uint32: 9, Valid: true}, IndividualDeductionID: null.Uint32{Uint32: 9, Valid: true}, EmployeeID: 9, Nominal: "2021年10月 給与明細", Payday: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年10月1日~2022年10月31日"},
		{ID: 10, FixedEarningID: null.Uint32{Uint32: 1, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 1, Valid: true}, EmployeeID: 9, Nominal: "2021年11月 給与明細", Payday: time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年11月1日~2022年11月30日"},
		{ID: 11, FixedEarningID: null.Uint32{Uint32: 2, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 2, Valid: true}, EmployeeID: 9, Nominal: "2021年11月 給与明細", Payday: time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年11月1日~2022年11月30日"},
		{ID: 12, FixedEarningID: null.Uint32{Uint32: 3, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 3, Valid: true}, EmployeeID: 9, Nominal: "2021年11月 給与明細", Payday: time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年11月1日~2022年11月30日"},
		{ID: 13, FixedEarningID: null.Uint32{Uint32: 4, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 4, Valid: true}, EmployeeID: 9, Nominal: "2021年11月 給与明細", Payday: time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年11月1日~2022年11月30日"},
		{ID: 14, FixedEarningID: null.Uint32{Uint32: 5, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 5, Valid: true}, EmployeeID: 9, Nominal: "2021年11月 給与明細", Payday: time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年11月1日~2022年11月30日"},
		{ID: 15, FixedEarningID: null.Uint32{Uint32: 6, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 6, Valid: true}, EmployeeID: 9, Nominal: "2021年11月 給与明細", Payday: time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年11月1日~2022年11月30日"},
		{ID: 16, FixedEarningID: null.Uint32{Uint32: 7, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 7, Valid: true}, EmployeeID: 9, Nominal: "2021年11月 給与明細", Payday: time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年11月1日~2022年11月30日"},
		{ID: 17, FixedEarningID: null.Uint32{Uint32: 8, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 8, Valid: true}, EmployeeID: 9, Nominal: "2021年11月 給与明細", Payday: time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年11月1日~2022年11月30日"},
		{ID: 18, FixedEarningID: null.Uint32{Uint32: 9, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 9, Valid: true}, EmployeeID: 9, Nominal: "2021年11月 給与明細", Payday: time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年11月1日~2022年11月30日"},
		{ID: 19, FixedEarningID: null.Uint32{Uint32: 1, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 1, Valid: true}, EmployeeID: 9, Nominal: "2021年12月 給与明細", Payday: time.Date(2021, 12, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"},
		{ID: 20, FixedEarningID: null.Uint32{Uint32: 2, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 2, Valid: true}, EmployeeID: 9, Nominal: "2021年12月 給与明細", Payday: time.Date(2021, 12, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"},
		{ID: 21, FixedEarningID: null.Uint32{Uint32: 3, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 3, Valid: true}, EmployeeID: 9, Nominal: "2021年12月 給与明細", Payday: time.Date(2021, 12, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"},
		{ID: 22, FixedEarningID: null.Uint32{Uint32: 4, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 4, Valid: true}, EmployeeID: 9, Nominal: "2021年12月 給与明細", Payday: time.Date(2021, 12, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"},
		{ID: 23, FixedEarningID: null.Uint32{Uint32: 5, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 5, Valid: true}, EmployeeID: 9, Nominal: "2021年12月 給与明細", Payday: time.Date(2021, 12, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"},
		{ID: 24, FixedEarningID: null.Uint32{Uint32: 6, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 6, Valid: true}, EmployeeID: 9, Nominal: "2021年12月 給与明細", Payday: time.Date(2021, 12, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"},
		{ID: 25, FixedEarningID: null.Uint32{Uint32: 7, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 7, Valid: true}, EmployeeID: 9, Nominal: "2021年12月 給与明細", Payday: time.Date(2021, 12, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"},
		{ID: 26, FixedEarningID: null.Uint32{Uint32: 8, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 8, Valid: true}, EmployeeID: 9, Nominal: "2021年12月 給与明細", Payday: time.Date(2021, 12, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"},
		{ID: 27, FixedEarningID: null.Uint32{Uint32: 9, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 9, Valid: true}, EmployeeID: 9, Nominal: "2021年12月 給与明細", Payday: time.Date(2021, 12, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"}}
		
	for _, salary_statement := range salary_statements {
		if err := salary_statement.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create datas of salary_statement. err:%s", err)
		}
	}

	return nil
}