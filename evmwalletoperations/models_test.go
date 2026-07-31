package evmwalletoperations

import "testing"

func TestWalletSessionReservationConsumptionIsInternal(t *testing.T) {
	t.Parallel()

	operation := WalletOperation{
		WalletSessionReservationConsumed: true,
	}
	if got := operation.GetValueByTag(
		"wallet_session_reservation_consumed",
	); got != true {
		t.Fatalf(
			"GetValueByTag(wallet_session_reservation_consumed) = %#v, want true",
			got,
		)
	}
	if _, exposed := operation.ToJSON()["wallet_session_reservation_consumed"]; exposed {
		t.Fatal(
			"WalletOperation.ToJSON() exposed internal reservation consumption",
		)
	}
}

func TestTransactionFromAddressProjection(t *testing.T) {
	t.Parallel()

	fromAddress := "0x1111111111111111111111111111111111111111"
	operation := WalletOperation{TransactionFromAddress: &fromAddress}

	if got := operation.GetValueByTag(
		"transaction_from_address",
	); got != operation.TransactionFromAddress {
		t.Fatalf(
			"GetValueByTag(transaction_from_address) = %#v, want %#v",
			got,
			operation.TransactionFromAddress,
		)
	}
	if got := operation.ToJSON()["transaction_from_address"]; got != operation.TransactionFromAddress {
		t.Fatalf(
			"ToJSON transaction_from_address = %#v, want %#v",
			got,
			operation.TransactionFromAddress,
		)
	}
}
