package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jokertennis/Payroll-Software/server/gen/models"

	"github.com/jokertennis/Payroll-Software/src/main/domain-model/deduction"
	"github.com/jokertennis/Payroll-Software/src/main/domain-model/earning"
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
		{ID: 1, CompanyID: 1, Name: "YumiAdministrator", MailAddress: "yumi@example.com", Password: "testpass"},
		{ID: 2, CompanyID: 1, Name: "TestForAdministrator", MailAddress: "test.administrator@example.com", Password: "testpass"}}

	for _, administrator := range administrators {
		if err := administrator.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create datas of administrator. err:%s", err)
		}
	}
	
	earnings := []models.Earning{
		{ID: 1, Nominal: "支給", Amount: 300000, EarningType: string(earning_domain_model.IndividualEarning)},
		{ID: 2, Nominal: "支給", Amount: 350000, EarningType: string(earning_domain_model.IndividualEarning)},
		{ID: 3, Nominal: "支給", Amount: 250000, EarningType: string(earning_domain_model.IndividualEarning)},
		{ID: 4, Nominal: "支給", Amount: 200000, EarningType: string(earning_domain_model.IndividualEarning)},
		{ID: 5, Nominal: "支給", Amount: 400000, EarningType: string(earning_domain_model.IndividualEarning)},
		{ID: 6, Nominal: "支給", Amount: 450000, EarningType: string(earning_domain_model.IndividualEarning)},
		{ID: 7, Nominal: "支給", Amount: 500000, EarningType: string(earning_domain_model.IndividualEarning)},
		{ID: 8, Nominal: "支給", Amount: 350000, EarningType: string(earning_domain_model.IndividualEarning)},
		{ID: 9, Nominal: "支給", Amount: 300000, EarningType: string(earning_domain_model.IndividualEarning)},
		{ID: 51, Nominal: "エンジニアLv1家族手当0人", Amount: 300000, EarningType: string(earning_domain_model.FixedEarning)},
		{ID: 52, Nominal: "エンジニアLv1家族手当1人", Amount: 350000, EarningType: string(earning_domain_model.FixedEarning)},
		{ID: 53, Nominal: "エンジニアLv1家族手当2人", Amount: 400000, EarningType: string(earning_domain_model.FixedEarning)},
		{ID: 54, Nominal: "エンジニアLv1家族手当3人", Amount: 450000, EarningType: string(earning_domain_model.FixedEarning)},
		{ID: 55, Nominal: "年俸360万１ヶ月分", Amount: 300000, EarningType: string(earning_domain_model.FixedEarning)},
		{ID: 56, Nominal: "年俸420万１ヶ月分", Amount: 350000, EarningType: string(earning_domain_model.FixedEarning)},
		{ID: 57, Nominal: "年俸600万１ヶ月分", Amount: 500000, EarningType: string(earning_domain_model.FixedEarning)},
		{ID: 58, Nominal: "Jimさんの給料１ヶ月分", Amount: 300000, EarningType: string(earning_domain_model.FixedEarning)},
		{ID: 59, Nominal: "Steveさんの給料１ヶ月分", Amount: 350000, EarningType: string(earning_domain_model.FixedEarning)},
	}
	
	for _, earning := range earnings {
		if err := earning.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create datas of earning. err:%s", err)
		}
	}

	earning_details := []models.EarningDetail{
		{ID: 1, EarningID: 1, Nominal: "基本給", Amount: 250000},
		{ID: 2, EarningID: 1, Nominal: "住宅手当", Amount: 20000},
		{ID: 3, EarningID: 1, Nominal: "通勤手当/非課税", Amount: 30000},
		{ID: 4, EarningID: 2, Nominal: "基本給", Amount: 300000},
		{ID: 5, EarningID: 2, Nominal: "残業手当", Amount: 50000},
		{ID: 6, EarningID: 3, Nominal: "基本給", Amount: 250000},
		{ID: 7, EarningID: 4, Nominal: "基本給", Amount: 200000},
		{ID: 8, EarningID: 5, Nominal: "基本給", Amount: 250000},
		{ID: 9, EarningID: 5, Nominal: "残業手当", Amount: 50000},
		{ID: 10, EarningID: 5, Nominal: "エンジニア手当", Amount: 100000},
		{ID: 11, EarningID: 6, Nominal: "基本給", Amount: 250000},
		{ID: 12, EarningID: 6, Nominal: "残業手当", Amount: 150000},
		{ID: 13, EarningID: 7, Nominal: "10月賞与", Amount: 500000},
		{ID: 14, EarningID: 8, Nominal: "10月賞与", Amount: 350000},
		{ID: 15, EarningID: 9, Nominal: "基本給", Amount: 200000},
		{ID: 16, EarningID: 9, Nominal: "残業手当", Amount: 50000},
		{ID: 17, EarningID: 9, Nominal: "通勤手当/非課税", Amount: 50000},
		{ID: 101, EarningID: 51, Nominal: "基本給", Amount: 250000},
		{ID: 102, EarningID: 51, Nominal: "固定残業代", Amount: 50000},
		{ID: 103, EarningID: 52, Nominal: "基本給", Amount: 250000},
		{ID: 104, EarningID: 52, Nominal: "固定残業代", Amount: 50000},
		{ID: 105, EarningID: 52, Nominal: "家族手当(1人分)", Amount: 50000},
		{ID: 106, EarningID: 53, Nominal: "基本給", Amount: 250000},
		{ID: 107, EarningID: 53, Nominal: "固定残業代", Amount: 50000},
		{ID: 108, EarningID: 53, Nominal: "家族手当(2人分)", Amount: 100000},
		{ID: 109, EarningID: 54, Nominal: "基本給", Amount: 250000},
		{ID: 110, EarningID: 54, Nominal: "固定残業代", Amount: 50000},
		{ID: 111, EarningID: 54, Nominal: "家族手当(3人分)", Amount: 150000},
		{ID: 112, EarningID: 55, Nominal: "年俸360万１ヶ月分", Amount: 300000},
		{ID: 113, EarningID: 56, Nominal: "年俸420万１ヶ月分", Amount: 350000},
		{ID: 114, EarningID: 57, Nominal: "年俸600万１ヶ月分", Amount: 500000},
		{ID: 115, EarningID: 58, Nominal: "基本給", Amount: 250000},
		{ID: 116, EarningID: 58, Nominal: "固定残業代", Amount: 50000},
		{ID: 117, EarningID: 59, Nominal: "基本給", Amount: 250000},
		{ID: 118, EarningID: 59, Nominal: "固定残業代", Amount: 50000},
		{ID: 119, EarningID: 59, Nominal: "家族手当(1人分)", Amount: 50000},
	}

	for _, earning_detail := range earning_details {
		if err := earning_detail.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create datas of earning_detail. err:%s", err)
		}
	}

	deductions := []models.Deduction{
		{ID: 1, Amount: 60000, Nominal: "控除", DeductionType: string(deduction_domain_model.IndividualDeduction)},
		{ID: 2, Amount: 70000, Nominal: "控除", DeductionType: string(deduction_domain_model.IndividualDeduction)},
		{ID: 3, Amount: 50000, Nominal: "控除", DeductionType: string(deduction_domain_model.IndividualDeduction)},
		{ID: 4, Amount: 40000, Nominal: "控除", DeductionType: string(deduction_domain_model.IndividualDeduction)},
		{ID: 5, Amount: 80000, Nominal: "控除", DeductionType: string(deduction_domain_model.IndividualDeduction)},
		{ID: 6, Amount: 90000, Nominal: "控除", DeductionType: string(deduction_domain_model.IndividualDeduction)},
		{ID: 7, Amount: 100000, Nominal: "控除", DeductionType: string(deduction_domain_model.IndividualDeduction)},
		{ID: 8, Amount: 70000, Nominal: "控除", DeductionType: string(deduction_domain_model.IndividualDeduction)},
		{ID: 9, Amount: 60000, Nominal: "控除", DeductionType: string(deduction_domain_model.IndividualDeduction)},
		{ID: 51, Amount: 60000, Nominal: "エンジニアLv1家族手当0人控除", DeductionType: string(deduction_domain_model.FixedDeduction)},
		{ID: 52, Amount: 70000, Nominal: "エンジニアLv1家族手当1人控除", DeductionType: string(deduction_domain_model.FixedDeduction)},
		{ID: 53, Amount: 80000, Nominal: "エンジニアLv1家族手当2人控除", DeductionType: string(deduction_domain_model.FixedDeduction)},
		{ID: 54, Amount: 90000, Nominal: "エンジニアLv1家族手当3人控除", DeductionType: string(deduction_domain_model.FixedDeduction)},
		{ID: 55, Amount: 60000, Nominal: "年俸360万１ヶ月分控除", DeductionType: string(deduction_domain_model.FixedDeduction)},
		{ID: 56, Amount: 70000, Nominal: "年俸420万１ヶ月分控除", DeductionType: string(deduction_domain_model.FixedDeduction)},
		{ID: 57, Amount: 100000, Nominal: "年俸600万１ヶ月分控除", DeductionType: string(deduction_domain_model.FixedDeduction)},
		{ID: 58, Amount: 60000, Nominal: "Jimさんの給料1ヶ月分控除", DeductionType: string(deduction_domain_model.FixedDeduction)},
		{ID: 59, Amount: 70000, Nominal: "Steveさんの給料1ヶ月分控除", DeductionType: string(deduction_domain_model.FixedDeduction)},
	}
	
	for _, deduction := range deductions {
		if err := deduction.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create datas of deduction. err:%s", err)
		}
	}

	deduction_details := []models.DeductionDetail{
		{ID: 1, DeductionID: 1, Nominal: "健康保険料", Amount: 12000},
		{ID: 2, DeductionID: 1, Nominal: "厚生年金保険料", Amount: 20000},
		{ID: 3, DeductionID: 1, Nominal: "雇用保険料", Amount: 300},
		{ID: 4, DeductionID: 1, Nominal: "所得税", Amount: 5000},
		{ID: 5, DeductionID: 1, Nominal: "住民税", Amount: 22700},
		{ID: 6, DeductionID: 2, Nominal: "健康保険料", Amount: 13000},
		{ID: 7, DeductionID: 2, Nominal: "厚生年金保険料", Amount: 21000},
		{ID: 8, DeductionID: 2, Nominal: "雇用保険料", Amount: 500},
		{ID: 9, DeductionID: 2, Nominal: "所得税", Amount: 6000},
		{ID: 10, DeductionID: 2, Nominal: "住民税", Amount: 29500},
		{ID: 11, DeductionID: 3, Nominal: "健康保険料", Amount: 10000},
		{ID: 12, DeductionID: 3, Nominal: "厚生年金保険料", Amount: 18000},
		{ID: 13, DeductionID: 3, Nominal: "雇用保険料", Amount: 300},
		{ID: 14, DeductionID: 3, Nominal: "所得税", Amount: 4500},
		{ID: 15, DeductionID: 3, Nominal: "住民税", Amount: 17200},
		{ID: 16, DeductionID: 4, Nominal: "健康保険料", Amount: 9000},
		{ID: 17, DeductionID: 4, Nominal: "厚生年金保険料", Amount: 15000},
		{ID: 18, DeductionID: 4, Nominal: "雇用保険料", Amount: 300},
		{ID: 19, DeductionID: 4, Nominal: "所得税", Amount: 4000},
		{ID: 20, DeductionID: 4, Nominal: "住民税", Amount: 11700},
		{ID: 21, DeductionID: 5, Nominal: "健康保険料", Amount: 18000},
		{ID: 22, DeductionID: 5, Nominal: "厚生年金保険料", Amount: 30000},
		{ID: 23, DeductionID: 5, Nominal: "雇用保険料", Amount: 600},
		{ID: 24, DeductionID: 5, Nominal: "所得税", Amount: 8000},
		{ID: 25, DeductionID: 5, Nominal: "住民税", Amount: 23400},
		{ID: 26, DeductionID: 6, Nominal: "健康保険料", Amount: 18000},
		{ID: 27, DeductionID: 6, Nominal: "厚生年金保険料", Amount: 25000},
		{ID: 28, DeductionID: 6, Nominal: "雇用保険料", Amount: 600},
		{ID: 29, DeductionID: 6, Nominal: "所得税", Amount: 1000},
		{ID: 30, DeductionID: 6, Nominal: "住民税", Amount: 45400},
		{ID: 31, DeductionID: 7, Nominal: "健康保険料", Amount: 20000},
		{ID: 32, DeductionID: 7, Nominal: "厚生年金保険料", Amount: 20000},
		{ID: 33, DeductionID: 7, Nominal: "雇用保険料", Amount: 30000},
		{ID: 34, DeductionID: 7, Nominal: "所得税", Amount: 20000},
		{ID: 35, DeductionID: 7, Nominal: "住民税", Amount: 30000},
		{ID: 36, DeductionID: 8, Nominal: "健康保険料", Amount: 13000},
		{ID: 37, DeductionID: 8, Nominal: "厚生年金保険料", Amount: 21000},
		{ID: 38, DeductionID: 8, Nominal: "雇用保険料", Amount: 500},
		{ID: 39, DeductionID: 8, Nominal: "所得税", Amount: 6000},
		{ID: 40, DeductionID: 8, Nominal: "住民税", Amount: 29500},
		{ID: 41, DeductionID: 9, Nominal: "健康保険料", Amount: 12000},
		{ID: 42, DeductionID: 9, Nominal: "厚生年金保険料", Amount: 20000},
		{ID: 43, DeductionID: 9, Nominal: "雇用保険料", Amount: 300},
		{ID: 44, DeductionID: 9, Nominal: "所得税", Amount: 5000},
		{ID: 45, DeductionID: 9, Nominal: "住民税", Amount: 22700},
		{ID: 101, DeductionID: 51, Nominal: "健康保険料", Amount: 12000},
		{ID: 102, DeductionID: 51, Nominal: "厚生年金保険料", Amount: 20000},
		{ID: 103, DeductionID: 51, Nominal: "雇用保険料", Amount: 300},
		{ID: 104, DeductionID: 51, Nominal: "所得税", Amount: 5000},
		{ID: 105, DeductionID: 51, Nominal: "住民税", Amount: 22700},
		{ID: 106, DeductionID: 52, Nominal: "健康保険料", Amount: 13000},
		{ID: 107, DeductionID: 52, Nominal: "厚生年金保険料", Amount: 21000},
		{ID: 108, DeductionID: 52, Nominal: "雇用保険料", Amount: 500},
		{ID: 109, DeductionID: 52, Nominal: "所得税", Amount: 6000},
		{ID: 110, DeductionID: 52, Nominal: "住民税", Amount: 29500},
		{ID: 111, DeductionID: 53, Nominal: "健康保険料", Amount: 18000},
		{ID: 112, DeductionID: 53, Nominal: "厚生年金保険料", Amount: 30000},
		{ID: 113, DeductionID: 53, Nominal: "雇用保険料", Amount: 600},
		{ID: 114, DeductionID: 53, Nominal: "所得税", Amount: 8000},
		{ID: 115, DeductionID: 53, Nominal: "住民税", Amount: 23400},
		{ID: 116, DeductionID: 54, Nominal: "健康保険料", Amount: 18000},
		{ID: 117, DeductionID: 54, Nominal: "厚生年金保険料", Amount: 25000},
		{ID: 118, DeductionID: 54, Nominal: "雇用保険料", Amount: 600},
		{ID: 119, DeductionID: 54, Nominal: "所得税", Amount: 1000},
		{ID: 120, DeductionID: 54, Nominal: "住民税", Amount: 45400},
		{ID: 121, DeductionID: 55, Nominal: "健康保険料", Amount: 12000},
		{ID: 122, DeductionID: 55, Nominal: "厚生年金保険料", Amount: 20000},
		{ID: 123, DeductionID: 55, Nominal: "雇用保険料", Amount: 300},
		{ID: 124, DeductionID: 55, Nominal: "所得税", Amount: 5000},
		{ID: 125, DeductionID: 55, Nominal: "住民税", Amount: 22700},
		{ID: 126, DeductionID: 56, Nominal: "健康保険料", Amount: 13000},
		{ID: 127, DeductionID: 56, Nominal: "厚生年金保険料", Amount: 21000},
		{ID: 128, DeductionID: 56, Nominal: "雇用保険料", Amount: 500},
		{ID: 129, DeductionID: 56, Nominal: "所得税", Amount: 6000},
		{ID: 130, DeductionID: 56, Nominal: "住民税", Amount: 29500},
		{ID: 131, DeductionID: 57, Nominal: "健康保険料", Amount: 20000},
		{ID: 132, DeductionID: 57, Nominal: "厚生年金保険料", Amount: 20000},
		{ID: 133, DeductionID: 57, Nominal: "雇用保険料", Amount: 30000},
		{ID: 134, DeductionID: 57, Nominal: "所得税", Amount: 20000},
		{ID: 135, DeductionID: 57, Nominal: "住民税", Amount: 30000},
		{ID: 136, DeductionID: 58, Nominal: "健康保険料", Amount: 12000},
		{ID: 137, DeductionID: 58, Nominal: "厚生年金保険料", Amount: 20000},
		{ID: 138, DeductionID: 58, Nominal: "雇用保険料", Amount: 300},
		{ID: 139, DeductionID: 58, Nominal: "所得税", Amount: 5000},
		{ID: 140, DeductionID: 58, Nominal: "住民税", Amount: 22700},
		{ID: 141, DeductionID: 59, Nominal: "健康保険料", Amount: 13000},
		{ID: 142, DeductionID: 59, Nominal: "厚生年金保険料", Amount: 21000},
		{ID: 143, DeductionID: 59, Nominal: "雇用保険料", Amount: 500},
		{ID: 144, DeductionID: 59, Nominal: "所得税", Amount: 6000},
		{ID: 145, DeductionID: 59, Nominal: "住民税", Amount: 29500},
	}

	for _, deduction_detail := range deduction_details {
		if err := deduction_detail.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create datas of deduction_detail. err:%s", err)
		}
	}

	salary_statements := []models.SalaryStatement{
		{ID: 1, EarningID: 1, DeductionID: 1, EmployeeID: 1, Nominal: "2021年10月 給与明細", Payday: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年10月1日~2022年10月31日"},
		{ID: 2, EarningID: 2, DeductionID: 2, EmployeeID: 2, Nominal: "2021年10月 給与明細", Payday: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年10月1日~2022年10月31日"},
		{ID: 3, EarningID: 3, DeductionID: 3, EmployeeID: 3, Nominal: "2021年10月 給与明細", Payday: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年10月1日~2022年10月31日"},
		{ID: 4, EarningID: 4, DeductionID: 4, EmployeeID: 4, Nominal: "2021年10月 給与明細", Payday: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年10月1日~2022年10月31日"},
		{ID: 5, EarningID: 5, DeductionID: 5, EmployeeID: 5, Nominal: "2021年10月 給与明細", Payday: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年10月1日~2022年10月31日"},
		{ID: 6, EarningID: 6, DeductionID: 6, EmployeeID: 6, Nominal: "2021年10月 給与明細", Payday: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年10月1日~2022年10月31日"},
		{ID: 7, EarningID: 7, DeductionID: 7, EmployeeID: 7, Nominal: "2021年10月 給与明細", Payday: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年10月1日~2022年10月31日"},
		{ID: 8, EarningID: 8, DeductionID: 8, EmployeeID: 8, Nominal: "2021年10月 給与明細", Payday: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年10月1日~2022年10月31日"},
		{ID: 9, EarningID: 9, DeductionID: 9, EmployeeID: 9, Nominal: "2021年10月 給与明細", Payday: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年10月1日~2022年10月31日"},
		{ID: 10, EarningID: 51, DeductionID: 51, EmployeeID: 9, Nominal: "2021年11月 給与明細", Payday: time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年11月1日~2022年11月30日"},
		{ID: 11, EarningID: 52, DeductionID: 52, EmployeeID: 9, Nominal: "2021年11月 給与明細", Payday: time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年11月1日~2022年11月30日"},
		{ID: 12, EarningID: 53, DeductionID: 53, EmployeeID: 9, Nominal: "2021年11月 給与明細", Payday: time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年11月1日~2022年11月30日"},
		{ID: 13, EarningID: 54, DeductionID: 54, EmployeeID: 9, Nominal: "2021年11月 給与明細", Payday: time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年11月1日~2022年11月30日"},
		{ID: 14, EarningID: 55, DeductionID: 55, EmployeeID: 9, Nominal: "2021年11月 給与明細", Payday: time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年11月1日~2022年11月30日"},
		{ID: 15, EarningID: 56, DeductionID: 56, EmployeeID: 9, Nominal: "2021年11月 給与明細", Payday: time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年11月1日~2022年11月30日"},
		{ID: 16, EarningID: 57, DeductionID: 57, EmployeeID: 9, Nominal: "2021年11月 給与明細", Payday: time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年11月1日~2022年11月30日"},
		{ID: 17, EarningID: 58, DeductionID: 58, EmployeeID: 9, Nominal: "2021年11月 給与明細", Payday: time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年11月1日~2022年11月30日"},
		{ID: 18, EarningID: 59, DeductionID: 59, EmployeeID: 9, Nominal: "2021年11月 給与明細", Payday: time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年11月1日~2022年11月30日"},
		{ID: 19, EarningID: 51, DeductionID: 51, EmployeeID: 9, Nominal: "2021年12月 給与明細", Payday: time.Date(2021, 12, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"},
		{ID: 20, EarningID: 52, DeductionID: 52, EmployeeID: 9, Nominal: "2021年12月 給与明細", Payday: time.Date(2021, 12, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"},
		{ID: 21, EarningID: 53, DeductionID: 53, EmployeeID: 9, Nominal: "2021年12月 給与明細", Payday: time.Date(2021, 12, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"},
		{ID: 22, EarningID: 54, DeductionID: 54, EmployeeID: 9, Nominal: "2021年12月 給与明細", Payday: time.Date(2021, 12, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"},
		{ID: 23, EarningID: 55, DeductionID: 55, EmployeeID: 9, Nominal: "2021年12月 給与明細", Payday: time.Date(2021, 12, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"},
		{ID: 24, EarningID: 56, DeductionID: 56, EmployeeID: 9, Nominal: "2021年12月 給与明細", Payday: time.Date(2021, 12, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"},
		{ID: 25, EarningID: 57, DeductionID: 57, EmployeeID: 9, Nominal: "2021年12月 給与明細", Payday: time.Date(2021, 12, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"},
		{ID: 26, EarningID: 58, DeductionID: 58, EmployeeID: 9, Nominal: "2021年12月 給与明細", Payday: time.Date(2021, 12, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"},
		{ID: 27, EarningID: 59, DeductionID: 59, EmployeeID: 9, Nominal: "2021年12月 給与明細", Payday: time.Date(2021, 12, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"}}

	for _, salary_statement := range salary_statements {
		if err := salary_statement.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create datas of salary_statement. err:%s", err)
		}
	}

	return nil
}
