bybit.go.api 1.1.0

Rest
Account
  Set Spot Hedge Mode (fixed: GET → POST)
  Demo Apply Money (new; replaces deprecated RequestTestFund)
  Query DCP info (consolidated from ins_account.go)
  SMP group ID query by UID (consolidated from ins_account_smp_group.go)

Affiliate
  Get Affiliate User List (new path: /v5/affiliate/aff-user-list)
  Get Affiliate Customer List (backward-compat alias for legacy path)

Spot Margin
  Query Borrow Liability (fixed path casing: Liability → liability)

Market
  Get RPI Order Book (fixed path: rpi-orderbook → rpi_orderbook)
  Get Spread Trade Instruments Info (set to public, no auth required)
  Get Spread Trade Order Book (set to public, no auth required)
  Get Spread Trade Tickers (set to public, no auth required)
  Get Spread Recent Trade (set to public, no auth required)

API Limit
  Set API Rate Limit (consolidated from apilimit.go)
  Get API Rate Limit (consolidated from apilimit.go)
  Query Cap
  Query All
  Query Broker All UID Details
  Query Broker Cap
  Set Broker API Limit

Bot (Grid/DCA/Combo/FMart)
  Create Combo Bot
  Close Combo Bot
  Create FMart Bot
  Close FMart Bot
  Close FGrid Bot
  Validate FGrid Input
  Spot DCA, Spot Grid endpoints

Ins Loan / Crypto Loan
  Ins Loan Coin Delta Amount
  Ins Loan Ensure Tokens
  (Deprecated crypto loan and C2C lending methods retained for backward compatibility)

Bug fixes
  Fixed InterestBearingBorrowSize JSON tag casing (PascalCase → camelCase)
  Fixed CoinApy.Coin type (int → string) and ApyE8 tag (apy_e8 → apyE8)
  Fixed minOrderAmt JSON tag typo (jsoN → json)
  Removed duplicate functions: DcpSetTimewindow, CancelSpreadOrder, and voucher endpoints

Code quality
  Renamed 16 model files from snake_case to camelCase
  Applied gofmt to all source files
  Removed empty placeholder files (bot.go, models/botResponse.go, models/crypto_loan_fixedResponse.go)

---

bybit.go.api 1.0.7

Rest
Get closed option positions
Get price limit
Pre check order

Websocket
Order price limit
Insurance pool
Order price limit
Rpi order book

New crypto loan
Get borrowable coins
Get collateral coins
Get max allowed collateral reduction amount
Adjust collateral amount
Get collateral adjustment history
Get crypto loan position

Flexible loan
Borrow
Repay
Get flexible loan
Get borrow history
Get repay history

Fixed loan
Get supplying market
Get borrowing market
Create borrow order
Create supply order
Cancel borrow order
Cancel supply order
Get borrow contact info
Get supply contact info
Get borrow order info
Get supply order info
Repay
Get repayment history

Deprecated leverage token endpoints and legacy crypto loan endpoints 