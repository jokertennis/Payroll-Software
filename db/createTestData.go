package db

import (
	"context"
	"fmt"
	"time"
	"usr/local/go/server/gen/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func CreateTestData(ctx context.Context, db boil.ContextExecutor) error {
	companies := []models.Company{
		{Name: "株式会社川田"},
		{Name: "株式会社kawata"}}

	for _, company := range companies {
		if err := company.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create demo datas of company. err:%s", err)
		}
	}

	employees := []models.Employee{
		{CompanyID: 1, Name: "Yumi"},
		{CompanyID: 1, Name: "James"},
		{CompanyID: 1, Name: "Potter"},
		{CompanyID: 1, Name: "Kevin"},
		{CompanyID: 1, Name: "Tom"},
		{CompanyID: 1, Name: "Ayumi"},
		{CompanyID: 1, Name: "Bob"},
		{CompanyID: 1, Name: "Yuji"},
		{CompanyID: 1, Name: "Charo"},
		{CompanyID: 1, Name: "Eron"},
		{CompanyID: 1, Name: "Jim"},
		{CompanyID: 1, Name: "Steve"}}

	for _, employee := range employees {
		if err := employee.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create demo datas of employee. err:%s", err)
		}
	}

	earnings := []models.Earning{
		{Nominal: "支給", Amount: 300000},
		{Nominal: "支給", Amount: 350000},
		{Nominal: "支給", Amount: 250000},
		{Nominal: "支給", Amount: 200000},
		{Nominal: "支給", Amount: 400000},
		{Nominal: "支給", Amount: 450000},
		{Nominal: "支給", Amount: 500000},
		{Nominal: "支給", Amount: 350000},
		{Nominal: "支給", Amount: 300000}}
	
	for _, earning := range earnings {
		if err := earning.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create demo datas of earning. err:%s", err)
		}
	}

	earning_details := []models.EarningDetail{
		{EarningID: 1, Nominal: "基本給", Amount: 250000},
		{EarningID: 1, Nominal: "住宅手当", Amount: 20000},
		{EarningID: 1, Nominal: "通勤手当/非課税", Amount: 30000},
		{EarningID: 2, Nominal: "基本給", Amount: 300000},
		{EarningID: 2, Nominal: "残業手当", Amount: 50000},
		{EarningID: 3, Nominal: "基本給", Amount: 250000},
		{EarningID: 4, Nominal: "基本給", Amount: 200000},
		{EarningID: 5, Nominal: "基本給", Amount: 250000},
		{EarningID: 5, Nominal: "残業手当", Amount: 50000},
		{EarningID: 5, Nominal: "エンジニア手当", Amount: 100000},
		{EarningID: 6, Nominal: "基本給", Amount: 250000},
		{EarningID: 6, Nominal: "残業手当", Amount: 150000},
		{EarningID: 7, Nominal: "10月賞与", Amount: 500000},
		{EarningID: 8, Nominal: "10月賞与", Amount: 350000},
		{EarningID: 9, Nominal: "基本給", Amount: 200000},
		{EarningID: 9, Nominal: "残業手当", Amount: 50000},
		{EarningID: 9, Nominal: "通勤手当/非課税", Amount: 50000}}

	for _, earning_detail := range earning_details {
		if err := earning_detail.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create demo datas of earning_detail. err:%s", err)
		}
	}

	fixed_earnings := []models.FixedEarning{
		{Nominal: "エンジニアLv1家族手当0人", Amount: 300000},
		{Nominal: "エンジニアLv1家族手当1人", Amount: 350000},
		{Nominal: "エンジニアLv1家族手当2人", Amount: 400000},
		{Nominal: "エンジニアLv1家族手当3人", Amount: 450000},
		{Nominal: "年俸360万１ヶ月分", Amount: 300000},
		{Nominal: "年俸420万１ヶ月分", Amount: 350000},
		{Nominal: "年俸600万１ヶ月分", Amount: 500000},
		{Nominal: "Jimさんの給料１ヶ月分", Amount: 300000},
		{Nominal: "Steveさんの給料１ヶ月分", Amount: 350000}}
	
	for _, fixed_earning := range fixed_earnings {
		if err := fixed_earning.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create demo datas of fixed_earning. err:%s", err)
		}
	}

	fixed_earning_details := []models.FixedEarningDetail{
		{FixedEarningID: 1, Nominal: "基本給", Amount: 250000},
		{FixedEarningID: 1, Nominal: "固定残業代", Amount: 50000},
		{FixedEarningID: 2, Nominal: "基本給", Amount: 250000},
		{FixedEarningID: 2, Nominal: "固定残業代", Amount: 50000},
		{FixedEarningID: 2, Nominal: "家族手当(1人分)", Amount: 50000},
		{FixedEarningID: 3, Nominal: "基本給", Amount: 250000},
		{FixedEarningID: 3, Nominal: "固定残業代", Amount: 50000},
		{FixedEarningID: 3, Nominal: "家族手当(2人分)", Amount: 100000},
		{FixedEarningID: 4, Nominal: "基本給", Amount: 250000},
		{FixedEarningID: 4, Nominal: "固定残業代", Amount: 50000},
		{FixedEarningID: 4, Nominal: "家族手当(3人分)", Amount: 150000},
		{FixedEarningID: 5, Nominal: "年俸360万１ヶ月分", Amount: 300000},
		{FixedEarningID: 6, Nominal: "年俸420万１ヶ月分", Amount: 350000},
		{FixedEarningID: 7, Nominal: "年俸600万１ヶ月分", Amount: 500000},
		{FixedEarningID: 8, Nominal: "基本給", Amount: 250000},
		{FixedEarningID: 8, Nominal: "固定残業代", Amount: 50000},
		{FixedEarningID: 9, Nominal: "基本給", Amount: 250000},
		{FixedEarningID: 9, Nominal: "固定残業代", Amount: 50000},
		{FixedEarningID: 9, Nominal: "家族手当(1人分)", Amount: 50000}}
	
	for _, fixed_earning_detail := range fixed_earning_details {
		if err := fixed_earning_detail.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create demo datas of fixed_earning_detail. err:%s", err)
		}
	}

	deductions := []models.Deduction{
		{Amount: 60000, Nominal: "控除"},
		{Amount: 70000, Nominal: "控除"},
		{Amount: 50000, Nominal: "控除"},
		{Amount: 40000, Nominal: "控除"},
		{Amount: 80000, Nominal: "控除"},
		{Amount: 90000, Nominal: "控除"},
		{Amount: 100000, Nominal: "控除"},
		{Amount: 70000, Nominal: "控除"},
		{Amount: 60000, Nominal: "控除"}}
	
	for _, deduction := range deductions {
		if err := deduction.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create demo datas of deduction. err:%s", err)
		}
	}

	deduction_details := []models.DeductionDetail{
		{DeductionID: 1, Nominal: "健康保険料", Amount: 12000},
		{DeductionID: 1, Nominal: "厚生年金保険料", Amount: 20000},
		{DeductionID: 1, Nominal: "雇用保険料", Amount: 300},
		{DeductionID: 1, Nominal: "所得税", Amount: 5000},
		{DeductionID: 1, Nominal: "住民税", Amount: 22700},
		{DeductionID: 2, Nominal: "健康保険料", Amount: 13000},
		{DeductionID: 2, Nominal: "厚生年金保険料", Amount: 21000},
		{DeductionID: 2, Nominal: "雇用保険料", Amount: 500},
		{DeductionID: 2, Nominal: "所得税", Amount: 6000},
		{DeductionID: 2, Nominal: "住民税", Amount: 29500},
		{DeductionID: 3, Nominal: "健康保険料", Amount: 10000},
		{DeductionID: 3, Nominal: "厚生年金保険料", Amount: 18000},
		{DeductionID: 3, Nominal: "雇用保険料", Amount: 300},
		{DeductionID: 3, Nominal: "所得税", Amount: 4500},
		{DeductionID: 3, Nominal: "住民税", Amount: 17200},
		{DeductionID: 4, Nominal: "健康保険料", Amount: 9000},
		{DeductionID: 4, Nominal: "厚生年金保険料", Amount: 15000},
		{DeductionID: 4, Nominal: "雇用保険料", Amount: 300},
		{DeductionID: 4, Nominal: "所得税", Amount: 4000},
		{DeductionID: 4, Nominal: "住民税", Amount: 11700},
		{DeductionID: 5, Nominal: "健康保険料", Amount: 18000},
		{DeductionID: 5, Nominal: "厚生年金保険料", Amount: 30000},
		{DeductionID: 5, Nominal: "雇用保険料", Amount: 600},
		{DeductionID: 5, Nominal: "所得税", Amount: 8000},
		{DeductionID: 5, Nominal: "住民税", Amount: 23400},
		{DeductionID: 6, Nominal: "健康保険料", Amount: 18000},
		{DeductionID: 6, Nominal: "厚生年金保険料", Amount: 25000},
		{DeductionID: 6, Nominal: "雇用保険料", Amount: 600},
		{DeductionID: 6, Nominal: "所得税", Amount: 1000},
		{DeductionID: 6, Nominal: "住民税", Amount: 45400},
		{DeductionID: 7, Nominal: "健康保険料", Amount: 20000},
		{DeductionID: 7, Nominal: "厚生年金保険料", Amount: 20000},
		{DeductionID: 7, Nominal: "雇用保険料", Amount: 30000},
		{DeductionID: 7, Nominal: "所得税", Amount: 20000},
		{DeductionID: 7, Nominal: "住民税", Amount: 30000},
		{DeductionID: 8, Nominal: "健康保険料", Amount: 13000},
		{DeductionID: 8, Nominal: "厚生年金保険料", Amount: 21000},
		{DeductionID: 8, Nominal: "雇用保険料", Amount: 500},
		{DeductionID: 8, Nominal: "所得税", Amount: 6000},
		{DeductionID: 8, Nominal: "住民税", Amount: 29500},
		{DeductionID: 9, Nominal: "健康保険料", Amount: 12000},
		{DeductionID: 9, Nominal: "厚生年金保険料", Amount: 20000},
		{DeductionID: 9, Nominal: "雇用保険料", Amount: 300},
		{DeductionID: 9, Nominal: "所得税", Amount: 5000},
		{DeductionID: 9, Nominal: "住民税", Amount: 22700}}

	for _, deduction_detail := range deduction_details {
		if err := deduction_detail.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create demo datas of deduction_detail. err:%s", err)
		}
	}

	fixed_deductions := []models.FixedDeduction{
		{Nominal: "エンジニアLv1家族手当0人控除", Amount: 60000},
		{Nominal: "エンジニアLv1家族手当1人控除", Amount: 70000},
		{Nominal: "エンジニアLv1家族手当2人控除", Amount: 80000},
		{Nominal: "エンジニアLv1家族手当3人控除", Amount: 90000},
		{Nominal: "年俸360万１ヶ月分控除", Amount: 60000},
		{Nominal: "年俸420万１ヶ月分控除", Amount: 70000},
		{Nominal: "年俸600万１ヶ月分控除", Amount: 100000},
		{Nominal: "Jimさんの給料1ヶ月分控除", Amount: 60000},
		{Nominal: "Steveさんの給料1ヶ月分控除", Amount: 70000}}
		
	for _, fixed_deduction := range fixed_deductions {
		if err := fixed_deduction.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create demo datas of fixed_deduction. err:%s", err)
		}
	}

	fixed_deduction_details := []models.FixedDeductionDetail{
		{FixedDeductionID: 1, Nominal: "健康保険料", Amount: 12000},
		{FixedDeductionID: 1, Nominal: "厚生年金保険料", Amount: 20000},
		{FixedDeductionID: 1, Nominal: "雇用保険料", Amount: 300},
		{FixedDeductionID: 1, Nominal: "所得税", Amount: 5000},
		{FixedDeductionID: 1, Nominal: "住民税", Amount: 22700},
		{FixedDeductionID: 2, Nominal: "健康保険料", Amount: 13000},
		{FixedDeductionID: 2, Nominal: "厚生年金保険料", Amount: 21000},
		{FixedDeductionID: 2, Nominal: "雇用保険料", Amount: 500},
		{FixedDeductionID: 2, Nominal: "所得税", Amount: 6000},
		{FixedDeductionID: 2, Nominal: "住民税", Amount: 29500},
		{FixedDeductionID: 3, Nominal: "健康保険料", Amount: 18000},
		{FixedDeductionID: 3, Nominal: "厚生年金保険料", Amount: 30000},
		{FixedDeductionID: 3, Nominal: "雇用保険料", Amount: 600},
		{FixedDeductionID: 3, Nominal: "所得税", Amount: 8000},
		{FixedDeductionID: 3, Nominal: "住民税", Amount: 23400},
		{FixedDeductionID: 4, Nominal: "健康保険料", Amount: 18000},
		{FixedDeductionID: 4, Nominal: "厚生年金保険料", Amount: 25000},
		{FixedDeductionID: 4, Nominal: "雇用保険料", Amount: 600},
		{FixedDeductionID: 4, Nominal: "所得税", Amount: 1000},
		{FixedDeductionID: 4, Nominal: "住民税", Amount: 45400},
		{FixedDeductionID: 5, Nominal: "健康保険料", Amount: 12000},
		{FixedDeductionID: 5, Nominal: "厚生年金保険料", Amount: 20000},
		{FixedDeductionID: 5, Nominal: "雇用保険料", Amount: 300},
		{FixedDeductionID: 5, Nominal: "所得税", Amount: 5000},
		{FixedDeductionID: 5, Nominal: "住民税", Amount: 22700},
		{FixedDeductionID: 6, Nominal: "健康保険料", Amount: 13000},
		{FixedDeductionID: 6, Nominal: "厚生年金保険料", Amount: 21000},
		{FixedDeductionID: 6, Nominal: "雇用保険料", Amount: 500},
		{FixedDeductionID: 6, Nominal: "所得税", Amount: 6000},
		{FixedDeductionID: 6, Nominal: "住民税", Amount: 29500},
		{FixedDeductionID: 7, Nominal: "健康保険料", Amount: 20000},
		{FixedDeductionID: 7, Nominal: "厚生年金保険料", Amount: 20000},
		{FixedDeductionID: 7, Nominal: "雇用保険料", Amount: 30000},
		{FixedDeductionID: 7, Nominal: "所得税", Amount: 20000},
		{FixedDeductionID: 7, Nominal: "住民税", Amount: 30000},
		{FixedDeductionID: 8, Nominal: "健康保険料", Amount: 12000},
		{FixedDeductionID: 8, Nominal: "厚生年金保険料", Amount: 20000},
		{FixedDeductionID: 8, Nominal: "雇用保険料", Amount: 300},
		{FixedDeductionID: 8, Nominal: "所得税", Amount: 5000},
		{FixedDeductionID: 8, Nominal: "住民税", Amount: 22700},
		{FixedDeductionID: 9, Nominal: "健康保険料", Amount: 13000},
		{FixedDeductionID: 9, Nominal: "厚生年金保険料", Amount: 21000},
		{FixedDeductionID: 9, Nominal: "雇用保険料", Amount: 500},
		{FixedDeductionID: 9, Nominal: "所得税", Amount: 6000},
		{FixedDeductionID: 9, Nominal: "住民税", Amount: 29500}}

	for _, fixed_deduction_detail := range fixed_deduction_details {
		if err := fixed_deduction_detail.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create demo datas of fixed_deduction_detail. err:%s", err)
		}
	}

	salary_statements := []models.SalaryStatement{
		{EarningID: null.Uint32{Uint32: 1, Valid: true}, DeductionID: null.Uint32{Uint32: 1, Valid: true}, EmployeeID: 1, Nominal: "2021年10月 給与明細", Payday: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年10月1日~2022年10月31日"},
		{EarningID: null.Uint32{Uint32: 2, Valid: true}, DeductionID: null.Uint32{Uint32: 2, Valid: true}, EmployeeID: 2, Nominal: "2021年10月 給与明細", Payday: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年10月1日~2022年10月31日"},
		{EarningID: null.Uint32{Uint32: 3, Valid: true}, DeductionID: null.Uint32{Uint32: 3, Valid: true}, EmployeeID: 3, Nominal: "2021年10月 給与明細", Payday: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年10月1日~2022年10月31日"},
		{EarningID: null.Uint32{Uint32: 4, Valid: true}, DeductionID: null.Uint32{Uint32: 4, Valid: true}, EmployeeID: 4, Nominal: "2021年10月 給与明細", Payday: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年10月1日~2022年10月31日"},
		{EarningID: null.Uint32{Uint32: 5, Valid: true}, DeductionID: null.Uint32{Uint32: 5, Valid: true}, EmployeeID: 5, Nominal: "2021年10月 給与明細", Payday: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年10月1日~2022年10月31日"},
		{EarningID: null.Uint32{Uint32: 6, Valid: true}, DeductionID: null.Uint32{Uint32: 6, Valid: true}, EmployeeID: 6, Nominal: "2021年10月 給与明細", Payday: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年10月1日~2022年10月31日"},
		{EarningID: null.Uint32{Uint32: 7, Valid: true}, DeductionID: null.Uint32{Uint32: 7, Valid: true}, EmployeeID: 7, Nominal: "2021年10月 給与明細", Payday: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年10月1日~2022年10月31日"},
		{EarningID: null.Uint32{Uint32: 8, Valid: true}, DeductionID: null.Uint32{Uint32: 8, Valid: true}, EmployeeID: 8, Nominal: "2021年10月 給与明細", Payday: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年10月1日~2022年10月31日"},
		{EarningID: null.Uint32{Uint32: 9, Valid: true}, DeductionID: null.Uint32{Uint32: 9, Valid: true}, EmployeeID: 9, Nominal: "2021年10月 給与明細", Payday: time.Date(2021, 10, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年10月1日~2022年10月31日"},
		{FixedEarningID: null.Uint32{Uint32: 1, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 1, Valid: true}, EmployeeID: 9, Nominal: "2021年11月 給与明細", Payday: time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年11月1日~2022年11月30日"},
		{FixedEarningID: null.Uint32{Uint32: 2, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 2, Valid: true}, EmployeeID: 9, Nominal: "2021年11月 給与明細", Payday: time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年11月1日~2022年11月30日"},
		{FixedEarningID: null.Uint32{Uint32: 3, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 3, Valid: true}, EmployeeID: 9, Nominal: "2021年11月 給与明細", Payday: time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年11月1日~2022年11月30日"},
		{FixedEarningID: null.Uint32{Uint32: 4, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 4, Valid: true}, EmployeeID: 9, Nominal: "2021年11月 給与明細", Payday: time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年11月1日~2022年11月30日"},
		{FixedEarningID: null.Uint32{Uint32: 5, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 5, Valid: true}, EmployeeID: 9, Nominal: "2021年11月 給与明細", Payday: time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年11月1日~2022年11月30日"},
		{FixedEarningID: null.Uint32{Uint32: 6, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 6, Valid: true}, EmployeeID: 9, Nominal: "2021年11月 給与明細", Payday: time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年11月1日~2022年11月30日"},
		{FixedEarningID: null.Uint32{Uint32: 7, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 7, Valid: true}, EmployeeID: 9, Nominal: "2021年11月 給与明細", Payday: time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年11月1日~2022年11月30日"},
		{FixedEarningID: null.Uint32{Uint32: 8, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 8, Valid: true}, EmployeeID: 9, Nominal: "2021年11月 給与明細", Payday: time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年11月1日~2022年11月30日"},
		{FixedEarningID: null.Uint32{Uint32: 9, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 9, Valid: true}, EmployeeID: 9, Nominal: "2021年11月 給与明細", Payday: time.Date(2021, 11, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年11月1日~2022年11月30日"},
		{FixedEarningID: null.Uint32{Uint32: 1, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 1, Valid: true}, EmployeeID: 9, Nominal: "2021年12月 給与明細", Payday: time.Date(2021, 12, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"},
		{FixedEarningID: null.Uint32{Uint32: 2, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 2, Valid: true}, EmployeeID: 9, Nominal: "2021年12月 給与明細", Payday: time.Date(2021, 12, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"},
		{FixedEarningID: null.Uint32{Uint32: 3, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 3, Valid: true}, EmployeeID: 9, Nominal: "2021年12月 給与明細", Payday: time.Date(2021, 12, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"},
		{FixedEarningID: null.Uint32{Uint32: 4, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 4, Valid: true}, EmployeeID: 9, Nominal: "2021年12月 給与明細", Payday: time.Date(2021, 12, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"},
		{FixedEarningID: null.Uint32{Uint32: 5, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 5, Valid: true}, EmployeeID: 9, Nominal: "2021年12月 給与明細", Payday: time.Date(2021, 12, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"},
		{FixedEarningID: null.Uint32{Uint32: 6, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 6, Valid: true}, EmployeeID: 9, Nominal: "2021年12月 給与明細", Payday: time.Date(2021, 12, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"},
		{FixedEarningID: null.Uint32{Uint32: 7, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 7, Valid: true}, EmployeeID: 9, Nominal: "2021年12月 給与明細", Payday: time.Date(2021, 12, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"},
		{FixedEarningID: null.Uint32{Uint32: 8, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 8, Valid: true}, EmployeeID: 9, Nominal: "2021年12月 給与明細", Payday: time.Date(2021, 12, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"},
		{FixedEarningID: null.Uint32{Uint32: 9, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 9, Valid: true}, EmployeeID: 9, Nominal: "2021年12月 給与明細", Payday: time.Date(2021, 12, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"}}
		
	for _, salary_statement := range salary_statements {
		if err := salary_statement.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create demo datas of salary_statement. err:%s", err)
		}
	}

	return nil
}