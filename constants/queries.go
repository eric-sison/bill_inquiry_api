package constants

const GetCustomerInfoQuery = `SELECT account_no AS account, 
			consumer_name As consumer,
			ISNULL(address, '') AS address,
			ISNULL(meter_no, '') AS meterNo,
			CAST(check_digit AS INT) AS checkDigit,
			ISNULL(description, '') AS description,
			status,
			ISNULL(CAST(AverageUsage AS FLOAT), 0.0) AS averageUsage
		FROM ViewConsumer_Info WHERE account_no = @account AND check_digit = @digit`

const GetBalanceQuery = `SELECT balance, 
			due_date AS dueDate, 
			sc AS seniorDiscount, 
			advance_payment AS advancePayment, 
			penalty,
			total,
			water_used
		FROM ViewConsumer_Balance
		INNER JOIN ViewConsumer_Info ON ViewConsumer_Info.account_no = ViewConsumer_Balance.account_no
		WHERE ViewConsumer_Balance.account_no = @account AND ViewConsumer_Info.check_digit = @digit`

const GetOtherChargesQuery = `SELECT ISNULL(charges_description, '') AS remarks, Monthly_due AS due, balance
		FROM ViewConsumer_other_charges
		INNER JOIN ViewConsumer_Info ON ViewConsumer_Info.account_no = ViewConsumer_other_charges.account_no
		WHERE ViewConsumer_other_charges.account_no = @account AND ViewConsumer_Info.check_digit = @digit`

const GetConsumerArrears = `SELECT ViewConsumer_withArrears.account_no, arrears
		FROM ViewConsumer_withArrears
		INNER JOIN ViewConsumer_Info ON ViewConsumer_Info.account_no = ViewConsumer_withArrears.account_no
		WHERE ViewConsumer_withArrears.account_no = @account AND ViewConsumer_Info.check_digit = @digit`
