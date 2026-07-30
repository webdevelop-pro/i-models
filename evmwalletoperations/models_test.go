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
