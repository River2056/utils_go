package misc

import "testing"

func TestSetMethodCaps(t *testing.T) {
	line := `amendData.setcrRootId();
		amendData.setmainlineId();
		amendData.setfacilityId();
		amendData.setportfolioCode();
		amendData.setbusinessEntity();
		amendData.setcrName();
		amendData.setlineName();
		amendData.setlimitCcyCode();
		amendData.setlimitAmt();
		amendData.setavailableAmt();
		amendData.setoutstandingAmt();
		amendData.setstartDate();
		amendData.setexpiryDate();
		amendData.setavailFlag();
		amendData.setannualReviewDate();
		amendData.setprocessDate();
		amendData.setfacilityLimitAmt();
		amendData.setfacilityCollateralAmt();
		amendData.setfacilityAvailableAmt();
		amendData.setfreezeDate();
		amendData.setfreezeReasonCode();
		amendData.setinterestMargin();
		amendData.setstdInterestMargin();
		amendData.setinternalRemark();
		amendData.setnewLoanIndustrialCode();
		amendData.setloanPurpose();
		amendData.setactivationDate();
		amendData.setorrAdjust();
		amendData.setstatus();
		amendData.setactiveFlow();
		amendData.setflowId();`
	SetMethodCaps(line)
}
