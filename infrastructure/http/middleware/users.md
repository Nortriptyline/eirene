users
-----
- id (PK)
- username
- email
- password_hash
- created_at
- updated_at

financial_institutions
----------------------
- id (PK)
- name
- website
- created_at
- updated_at

bank_accounts
-------------
- id (PK)
- user_id (FK)
- institution_id (FK)
- account_number
- account_type
- balance
- currency
- created_at
- updated_at

transactions
------------
- id (PK)
- account_id (FK)
- amount
- transaction_type
- description
- transaction_date
- created_at
- updated_at


Account types
--------------
- checking
- savings
- credit_card
- loan
- investment
- mortgage
- business
- joint
- fixed_deposit
- recurring_deposit


payment_type
--------------
- deposit
- withdrawal
- transfer
- payment
- fee
- interest
- refund
- purchase
- loan_repayment
- dividend
- deposit
- withdrawal
- transfer
- payment
- fee
- interest
- refund
- purchase
- loan_repayment
- dividend