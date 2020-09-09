package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	//unicommon "github.com/unidoc/unidoc/common"

	"github.com/unidoc/unidoc/pdf/creator"
	pdf "github.com/unidoc/unidoc/pdf/model"
	"github.com/unidoc/unidoc/pdf/model/fonts"
)

var (
	// fpre ....
	// FpreCode ...
	FpreCode = "50-26-0"
	// FpreName ...
	FpreName = "235-317-0"
	// FpreService ...
	FpreService = "235-345-0"
	// FpreBranch ...
	FpreBranch = "235-373-0"
	// FpreSales ...
	FpreSales = "235-403-0"
	// FpreEventCode ...
	FpreEventCode = "235-432-0"
	// FpreClientID ...
	FpreClientID = "235-461-0"
	// FpreCheckBoxDocument01 ...
	FpreCheckBoxDocument01 = "125-526-0"
	// FpreCheckBoxDocument02 ...
	FpreCheckBoxDocument02 = "125-551-0"
	// FpreCheckBoxDocument03 ...
	FpreCheckBoxDocument03 = "125-570-0"
	// FpreCheckBoxDocument04 ...
	FpreCheckBoxDocument04 = "125-570-0"
	// FpreCheckBoxDocument05 ...
	FpreCheckBoxDocument05 = "125-570-0"
	// FpreCheckBoxDocument06 ...
	FpreCheckBoxDocument06 = "125-570-0"
	// FpreCheckBoxDocument07 ...
	FpreCheckBoxDocument07 = "125-570-0"
	// FpreCheckBoxDocument08 ...
	FpreCheckBoxDocument08 = "125-570-0"
	// FpreCheckBoxDocument09 ...
	FpreCheckBoxDocument09 = "125-570-0"
	// FpreDate01 ...
	FpreDate01 = "65-65-1"
	// FpreDate02 ...
	FpreDate02 = "77-65-1"
	// FpreDate03 ...
	FpreDate03 = "89-65-1"
	// FpreCheckBoxRRB01 ...
	FpreCheckBoxRRB01 = "c002"
	// FpreCheckBoxRRB02 ...
	FpreCheckBoxRRB02 = "c003"
	// FpreCheckBoxAccountType01 ...
	FpreCheckBoxAccountType01 = "c004"
	// FpreCheckBoxAccountType02 ...
	FpreCheckBoxAccountType02 = "160-81-1"
	// FpreCheckBoxAccountType03 ...
	FpreCheckBoxAccountType03 = "196-81-1"
	// FpreCheckBoxBehalfOf01 ...
	FpreCheckBoxBehalfOf01 = "78-99-1"
	// FpreCheckBoxBehalfOf02 ...
	FpreCheckBoxBehalfOf02 = "c008"
	// FpreBehalfOfOther ...
	FpreBehalfOfOther = "c009"
	// FpreFullName ...
	FpreFullName = "13-158-1"
	// FpreKtpName ...
	FpreKtpName = "13-178-1"
	// FpreCheckBoxGender01 ...
	FpreCheckBoxGender01 = "144-190-1"
	// FpreCheckBoxGender02 ...
	FpreCheckBoxGender02 = "201-190-1"
	// FpreCheckBoxIdentity01 ...
	FpreCheckBoxIdentity01 = "144-205-1"
	// FpreCheckBoxIdentity02 ...
	FpreCheckBoxIdentity02 = "d006"
	// FpreIdentityNumber ...
	FpreIdentityNumber = "60-220-1"
	// FpreCheckBoxValidity ...
	FpreCheckBoxValidity = "14-248-1"
	// FprePlaceOfBirth ...
	FprePlaceOfBirth = "105-268-1"
	// FpreDateOfBirth01 ...
	FpreDateOfBirth01 = "108-285-1"
	// FpreDateOfBirth02 ...
	FpreDateOfBirth02 = "120-285-1"
	// FpreDateOfBirth03 ...
	FpreDateOfBirth03 = "132-285-1"
	// FpreCheckBoxMaritalStatusSingel ...
	FpreCheckBoxMaritalStatusSingel = "14-311-1"
	// FpreCheckBoxMaritalStatusWindow ...
	FpreCheckBoxMaritalStatusWindow = "e005"
	// FpreCheckBoxMaritalStatusMarried ...
	FpreCheckBoxMaritalStatusMarried = "14-327-1"
	// FpreCheckBoxMaritalStatusWindower ...
	FpreCheckBoxMaritalStatusWindower = "e007"
	// FpreCheckBoxCitizenShip01 ...
	FpreCheckBoxCitizenShip01 = "140-311-1"
	// FpreCheckBoxCitizenShip02 ...
	FpreCheckBoxCitizenShip02 = "e009"
	// FpreCheckBoxCitizenShipOther ...
	// FpreCheckBoxCitizenShipOther = "f001"
	// FpreMotherName ...
	FpreMotherName = "14-355-1"
	// FpreAddressID ...
	FpreAddressID = "14-378-1"
	// FpreRT ...
	FpreRT = "40-394-1"
	// FpreRW ...
	FpreRW = "85-394-1"
	// FprePostalCode ...
	FprePostalCode = "218-394-1"
	// FpreVillage ...
	FpreVillage = "99-408-1"
	// FpreSubDistrict ...
	FpreSubDistrict = "94-422-1"
	// FpreCity ...
	FpreCity = "120-438-1"
	// FpreProvince ...
	FpreProvince = "70-452-1"
	// FpreCheckBoxID ...
	FpreCheckBoxID = "137-470-1"
	// FpreAddress ...
	FpreAddress = "14-480-1"
	// FpreCurrentRt ..
	FpreCurrentRt = "40-497-1"
	// FpreCurrentRW ...
	FpreCurrentRW = "85-497-1"
	// FpreCurrentPostalCode ...
	FpreCurrentPostalCode = "218-497-1"
	// FpreCurrentVillage ...
	FpreCurrentVillage = "99-512-1"
	// FpreCurrentSubDistrict ...
	FpreCurrentSubDistrict = "94-526-1"
	// FpreCurrentCity ...
	FpreCurrentCity = "120-541-1"
	// FpreCurrentProvince ...
	FpreCurrentProvince = "70-556-1"
	// FpreResidentialStatus01 ...
	FpreResidentialStatus01 = "14-582-1"
	// FpreResidentialStatus02 ...
	FpreResidentialStatus02 = "106-582-1"
	// FpreResidentialStatus03 ...
	FpreResidentialStatus03 = "201-582-1"
	// FpreResidentialStatus04 ...
	FpreResidentialStatus04 = "14-598-1"
	// FpreResidentialStatus05 ...
	FpreResidentialStatus05 = "106-598-1"
	// FpreResidentialStatus06 ...
	FpreResidentialStatus06 = "201-598-1"
	// FpreHomePhone ...
	FpreHomePhone = "70-627-1"
	// FprePhoneCellular ...
	FprePhoneCellular = "205-627-1"
	// FpreEmail ...
	FpreEmail = "40-655-1"
	// FpreNpwpCountryID ...
	FpreNpwpCountryID = "14-682-1"
	// FpreNpwpCountryOther ...
	FpreNpwpCountryOther = "i001"
	// FpreNpwpCountryOtherID ...
	FpreNpwpCountryOtherID = "i002"
	// FpreNpwpNumber ...
	FpreNpwpNumber = "106-707-1"
	// FpreNpwpName ...
	FpreNpwpName = "214-707-1"
	// FpreEdu01 ...
	FpreEdu01 = "14-756-1"
	// FpreEdu02 ...
	FpreEdu02 = "80-756-1"
	// FpreEdu03 ...
	FpreEdu03 = "144-756-1"
	// FpreEdu04 ...
	FpreEdu04 = "208-756-1"
	// FpreEdu05 ...
	FpreEdu05 = "14-773-1"
	// FpreEdu06 ...
	FpreEdu06 = "80-773-1"
	// FpreEdu07 ...
	FpreEdu07 = "144-773-1"
	// FpreEdu08 ...
	FpreEdu08 = "208-773-1"
	// FpreEdu09 ...
	FpreEdu09 = "j004"
	// FpreReligion01 ...
	FpreReligion01 = "14-803-1"
	// FpreReligion02 ...
	FpreReligion02 = "47-803-1"
	// FpreReligion03 ...
	FpreReligion03 = "89-803-1"
	// FpreReligion04 ...
	FpreReligion04 = "127-803-1"
	// FpreReligion05 ...
	FpreReligion05 = "164-803-1"
	// FpreReligion06 ...
	FpreReligion06 = "196-803-1"
	// FpreReligion07 ...
	FpreReligion07 = "245-803-1"
	// FpreOce01 ...
	FpreOce01 = "311-73-1"
	// FpreOce02 ...
	FpreOce02 = "382-73-1"
	// FpreOce03 ...
	FpreOce03 = "467-73-1"
	// FpreOce04 ...
	FpreOce04 = "311-89-1"
	// FpreOce05 ...
	FpreOce05 = "382-89-1"
	// FpreOce06 ...
	FpreOce06 = "467-89-1"
	// FpreOce07 ...
	FpreOce07 = "311-106-1"
	// FpreOce08 ...
	FpreOce08 = "382-106-1"
	// FpreOce09 ...
	FpreOce09 = "467-106-1"
	// FpreOce10 ...
	FpreOce10 = "311-122-1"
	// FpreOce11 ...
	FpreOce11 = "382-122-1"
	// FpreOce12 ...
	FpreOce12 = "467-124-1"
	// FpreOce13 ...
	// FpreOce13 = "l007"
	// FpreOce14 ...
	FpreOce14 = "533-111-1"
	// FpreBusinessName ...
	FpreBusinessName = "412-146-1"
	// FprePeriodeService01 ...
	FprePeriodeService01 = "443-158-1"
	// FprePeriodeService02 ...
	FprePeriodeService02 = "530-158-1"
	// FpreLineBusiness01 ...
	FpreLineBusiness01 = "311-184-1"
	// FpreLineBusiness02 ...
	FpreLineBusiness02 = "368-184-1"
	// FpreLineBusiness03 ...
	FpreLineBusiness03 = "421-184-1"
	// FpreLineBusiness04 ...
	FpreLineBusiness04 = "499-184-1"
	// FpreLineBusiness05 ...
	FpreLineBusiness05 = "311-200-1"
	// FpreLineBusiness06 ...
	FpreLineBusiness06 = "368-200-1"
	// FpreLineBusiness07 ...
	FpreLineBusiness07 = "421-200-1"
	// FpreLineBusiness08 ...
	FpreLineBusiness08 = "499-200-1"
	// FpreLineBusiness09 ...
	FpreLineBusiness09 = "311-216-1"
	// FpreLineBusiness10 ...
	FpreLineBusiness10 = "368-216-1"
	// FpreLineBusiness11 ...
	FpreLineBusiness11 = "421-216-1"
	// FpreLineBusiness12 ...
	FpreLineBusiness12 = "499-216-1"
	// FpreLineBusiness13 ...
	FpreLineBusiness13 = "311-232-1"
	// FpreLineBusiness14 ...
	FpreLineBusiness14 = "368-232-1"
	// FpreLineBusiness15 ...
	FpreLineBusiness15 = "421-232-1"
	// FpreLineBusiness16 ...
	FpreLineBusiness16 = "499-232-1"
	// FpreLineBusiness17 ...
	FpreLineBusiness17 = "311-246-1"
	// FpreLineBusiness18 ...
	FpreLineBusiness18 = "368-248-1"
	// FpreLineBusiness19 ...
	FpreLineBusiness19 = "421-248-1"
	// FpreLineBusiness20 ...
	FpreLineBusiness20 = "499-248-1"
	// FpreDesignation01 ...
	FpreDesignation01 = "311-276-1"
	// FpreDesignation02 ...
	FpreDesignation02 = "379-276-1"
	// FpreDesignation03 ...
	FpreDesignation03 = "440-276-1"
	// FpreDesignation04 ...
	FpreDesignation04 = "499-276-1"
	// FpreDesignation05 ...
	FpreDesignation05 = "311-292-1"
	// FpreDesignation06 ...
	FpreDesignation06 = "379-292-1"
	// FpreDesignation07 ...
	FpreDesignation07 = "440-292-1"
	// FpreDesignation08 ...
	FpreDesignation08 = "499-292-1"
	// FpreBusinessAddress ...
	FpreBusinessAddress = "311-320-1"
	// FpreBusinessCity ...
	FpreBusinessCity = "379-340-1"
	// FpreBusinessPostalCode ...
	FpreBusinessPostalCode = "530-340-1"
	// FpreBusinessProvice ...
	FpreBusinessProvice = "379-353-1"
	// FpreBusinessPhone ...
	FpreBusinessPhone = "388-374-1"
	// FpreSourceOfFound01 ...
	FpreSourceOfFound01 = "311-411-1"
	// FpreSourceOfFound02 ...
	FpreSourceOfFound02 = "379-411-1"
	// FpreSourceOfFound03 ...
	FpreSourceOfFound03 = "440-411-1"
	// FpreSourceOfFound04 ...
	FpreSourceOfFound04 = "499-411-1"
	// FpreSourceOfFound05 ...
	FpreSourceOfFound05 = "311-426-1" //start input coordinat shob
	// FpreSourceOfFound06 ...
	FpreSourceOfFound06 = "379-426-1"
	// FpreSourceOfFound07 ...
	FpreSourceOfFound07 = "440-426-1"
	// FpreSourceOfFound08 ...
	FpreSourceOfFound08 = "499-426-1"
	// FpreSourceOfFound09 ...
	FpreSourceOfFound09 = "311-443-1"
	// FpreSourceOfFound10 ...
	FpreSourceOfFound10 = "379-443-1"
	// FpreMainSourceIncome01 ...
	FpreMainSourceIncome01 = "312-472-1"
	// FpreMainSourceIncome02 ...
	FpreMainSourceIncome02 = "379-472-1"
	// FpreMainSourceIncome03 ...
	FpreMainSourceIncome03 = "441-472-1"
	// FpreMainSourceIncome04 ...
	FpreMainSourceIncome04 = "499-472-1"
	// FpreMainSourceIncome05 ...
	FpreMainSourceIncome05 = "312-488-1"
	// FpreMainSourceIncome06 ...
	FpreMainSourceIncome06 = "379-488-1"
	// FpreMainSourceIncome07 ...
	FpreMainSourceIncome07 = "441-488-1"
	// FpreMainSourceIncome08 ...
	FpreMainSourceIncome08 = "499-488-1"
	// FpreMainSourceIncome09 ...
	FpreMainSourceIncome09 = "312-504-1"
	// FpreSalary ...
	FpreSalary = "379-504-1"
	// FprAditionalSourceIncome01 ...
	FprAditionalSourceIncome01 = "312-558-1"
	// FprAditionalSourceIncome02 ...
	FprAditionalSourceIncome02 = "379-558-1"
	// FprAditionalSourceIncome03 ...
	FprAditionalSourceIncome03 = "441-558-1"
	// FprAditionalSourceIncome04 ...
	FprAditionalSourceIncome04 = "499-558-1"
	// FprAditionalSourceIncome05 ...
	FprAditionalSourceIncome05 = "312-573-1"
	// FprAditionalSourceIncome06 ...
	FprAditionalSourceIncome06 = "379-573-1"
	// FprAditionalSourceIncome07 ...
	FprAditionalSourceIncome07 = "441-573-1"
	// FprAditionalSourceIncome08 ...
	FprAditionalSourceIncome08 = "499-573-1"
	// FprAditionalSourceIncome09 ...
	FprAditionalSourceIncome09 = "312-590-1"
	// FpreAditionalIncome ...
	FpreAditionalIncome = "322-620-1"
	// FpreMallingAddress01 ...
	FpreMallingAddress01 = "407-643-1"
	// FpreMallingAddress02 ...
	FpreMallingAddress02 = "408-660-1"
	// FpreMallingAddress03 ...
	FpreMallingAddress03 = "516-660-1"
	// FpreRelationAccountName ...
	FpreRelationAccountName = "13-70-2"
	// FpreRelationBankName ...
	FpreRelationBankName = "96-82-2"
	// FpreRelationBranch ...
	FpreRelationBranch = "66-96-2"
	// FpreRelationAccountNumber ...
	FpreRelationAccountNumber = "116-110-2"
	// FpreRelationName ...
	FpreRelationName = "13-165-2"
	// FpreRelationOccupation01 ...
	FpreRelationOccupation01 = "14-190-2"
	// FpreRelationOccupation02 ...
	FpreRelationOccupation02 = "84-190-2"
	// FpreRelationOccupation03 ...
	FpreRelationOccupation03 = "169-190-2"
	// FpreRelationOccupation04 ...
	FpreRelationOccupation04 = "14-205-2"
	// FpreRelationOccupation05 ...
	FpreRelationOccupation05 = "84-205-2"
	// FpreRelationOccupation06 ...
	FpreRelationOccupation06 = "169-205-2"
	// FpreRelationOccupation07 ...
	FpreRelationOccupation07 = "14-223-2"
	// FpreRelationOccupation08 ...
	FpreRelationOccupation08 = "84-223-2"
	// FpreRelationOccupation09 ...
	FpreRelationOccupation09 = "169-224-2"
	// FpreRelationOccupation10 ...
	FpreRelationOccupation10 = "14-239-2"
	// FpreRelationOccupation11 ...
	FpreRelationOccupation11 = "84-239-2"
	// FpreRelationOccupation12 ...
	FpreRelationOccupation12 = "169-241-2"
	// FpreRelationOccupation13 ...
	FpreRelationOccupation13 = "205-241-2"
	// FpreRelationOccupation14 ...
	FpreRelationOccupation14 = "235-228-2"
	// FpreRelationBusinessName ...
	FpreRelationBusinessName = "114-261-2"
	// FpreRelationPeriodeServiceYear ...
	FpreRelationPeriodeServiceYear = "v003"
	// FpreRelationPeriodeServiceMonth ...
	FpreRelationPeriodeServiceMonth = "v004"
	// FpreRelationLineOfBusiness01 ...
	FpreRelationLineOfBusiness01 = "v008"
	// FpreRelationLineOfBusiness02 ...
	FpreRelationLineOfBusiness02 = "v009"
	// FpreRelationLineOfBusiness03 ...
	FpreRelationLineOfBusiness03 = "w001"
	// FpreRelationLineOfBusiness04 ...
	FpreRelationLineOfBusiness04 = "w002"
	// FpreRelationLineOfBusiness05 ...
	FpreRelationLineOfBusiness05 = "w003"
	// FpreRelationLineOfBusiness06 ...
	FpreRelationLineOfBusiness06 = "w004"
	// FpreRelationLineOfBusiness07 ...
	FpreRelationLineOfBusiness07 = "w005"
	// FpreRelationLineOfBusiness08 ...
	FpreRelationLineOfBusiness08 = "w006"
	// FpreRelationLineOfBusiness09 ...
	FpreRelationLineOfBusiness09 = "w007"
	// FpreRelationLineOfBusiness10 ...
	FpreRelationLineOfBusiness10 = "w008"
	// FpreRelationLineOfBusiness11 ...
	FpreRelationLineOfBusiness11 = "w009"
	// FpreRelationLineOfBusiness12 ...
	FpreRelationLineOfBusiness12 = "x001"
	// FpreRelationLineOfBusiness13 ...
	FpreRelationLineOfBusiness13 = "x002"
	// FpreRelationLineOfBusiness14 ...
	FpreRelationLineOfBusiness14 = "x003"
	// FpreRelationLineOfBusiness15 ...
	FpreRelationLineOfBusiness15 = "x004"
	// FpreRelationLineOfBusiness16 ...
	FpreRelationLineOfBusiness16 = "x005"
	// FpreRelationLineOfBusiness17 ...
	FpreRelationLineOfBusiness17 = "x006"
	// FpreRelationLineOfBusiness18 ...
	FpreRelationLineOfBusiness18 = "x007"
	// FpreRelationLineOfBusiness19 ...
	FpreRelationLineOfBusiness19 = "x008"
	// FpreRelationLineOfBusiness20 ...
	FpreRelationLineOfBusiness20 = "x009"
	// FpreRelationJobs01 ...
	FpreRelationJobs01 = "y001"
	// FpreRelationJobs02 ...
	FpreRelationJobs02 = "y002"
	// FpreRelationJobs03 ...
	FpreRelationJobs03 = "y003"
	// FpreRelationJobs04 ...
	FpreRelationJobs04 = "y004"
	// FpreRelationJobs05 ...
	FpreRelationJobs05 = "y005"
	// FpreRelationJobs06 ...
	FpreRelationJobs06 = "y006"
	// FpreRelationJobs07 ...
	FpreRelationJobs07 = "y007"
	// FpreRelationBusinessAdderess ...
	FpreRelationBusinessAdderess = "y008"
	// FpreRelationBusinessCity ...
	FpreRelationBusinessCity = "y009"
	// FpreRelationBusinessPostalCode ...
	FpreRelationBusinessPostalCode = "z001"
	// FpreRelationBusinessProvince ...
	FpreRelationBusinessProvince = "z002"
	// FpreRelationBusinessPhone ...
	FpreRelationBusinessPhone = "z003"
	// FpreRelationAmanualGross ...
	FpreRelationAmanualGross = "z004"
	// FpreRelationAditionalIncome ...
	FpreRelationAditionalIncome = "z005"
	//FpreAditionalInformasi01Yes ...
	FpreAditionalInformasi01Yes = "a111"
	//FpreAditionalInformasi01No ...
	FpreAditionalInformasi01No = "a112"
	//FpreAditionalInformasi01FreeText ...
	FpreAditionalInformasi01FreeText = "a113"
	//FpreAditionalInformasi02Yes ...
	FpreAditionalInformasi02Yes = "a114"
	//FpreAditionalInformasi02No ...
	FpreAditionalInformasi02No = "a115"
	// FpreAditionalInformasiName ...
	FpreAditionalInformasiName = "a116"
	// FpreAditionalInformasiHubungan ...
	FpreAditionalInformasiHubungan = "a117"
	//FpreAditionalInformasi03Yes ...
	FpreAditionalInformasi03Yes = "a118"
	//FpreAditionalInformasi03No ...
	FpreAditionalInformasi03No = "a119"
	// FpreAditionalInformasi03Name ...
	FpreAditionalInformasi03Name = "b111"
	// FpreAditionalInformasi03accountNumber ...
	FpreAditionalInformasi03accountNumber = "b112"
	//FpreAditionalInformasi04Yes ...
	FpreAditionalInformasi04Yes = "b113"
	//FpreAditionalInformasi04No ...
	FpreAditionalInformasi04No = "b114"
	// FpreAditionalInformasi04Name ...
	FpreAditionalInformasi04Name = "b115"
	// FpreAditionalInformasi04Hubungan ...
	FpreAditionalInformasi04Hubungan = "b116"
	// FpreAditionalInformasi04accountNumber ...
	FpreAditionalInformasi04accountNumber = "b117"
	// FpreAditionalInformasi05Yes ...
	FpreAditionalInformasi05Yes = "b118"
	//FpreAditionalInformasi05No ...
	FpreAditionalInformasi05No = "b119"
	// FpreAditionalInformasi05CompanyName ...
	FpreAditionalInformasi05CompanyName = "c111"
	// FpreAditionalInformasi06Yes ...
	FpreAditionalInformasi06Yes = "c112"
	//FpreAditionalInformasi06No ...
	FpreAditionalInformasi06No = "c113"
	// FpreAditionalInformasi06CompanyName ...
	FpreAditionalInformasi06CompanyName = "c114"
	// FpreAditionalInformasi06Position ...
	FpreAditionalInformasi06Position = "c115"
	// FpreAditionalInformasi07Yes ...
	FpreAditionalInformasi07Yes = "c116"
	//FpreAditionalInformasi07No ...
	FpreAditionalInformasi07No = "c117"
	// FpreAditionalInformasi07CompanyName ...
	FpreAditionalInformasi07CompanyName = "c118"
	// FpreAditionalInformasi07Position ...
	FpreAditionalInformasi07Position = "c119"
	// FpreAditionalInformasi07Saham ...
	FpreAditionalInformasi07Saham = "d111"
	// FpreAditionalInformasi08Yes ...
	FpreAditionalInformasi08Yes = "d112"
	//FpreAditionalInformasi08No ...
	FpreAditionalInformasi08No = "d113"
	// FpreAditionalInformasi08Position ...
	FpreAditionalInformasi08Position = "d114"
	// FpreAditionalInformasi09Yes ...
	FpreAditionalInformasi09Yes = "d115"
	//FpreAditionalInformasi09No ...
	FpreAditionalInformasi09No = "d116"
	// FpreAditionalInformasi09Position ...
	FpreAditionalInformasi09Position = "d117"
	// FpreAditionalInformasi10Yes ...
	FpreAditionalInformasi10Yes = "d118"
	//FpreAditionalInformasi10No ...
	FpreAditionalInformasi10No = "d119"
	// FpreAditionalInformasi10CompanyName ...
	FpreAditionalInformasi10CompanyName = "e111"
	// FpreAditionalInformasi11Yes ...
	FpreAditionalInformasi11Yes = "e112"
	//FpreAditionalInformasi11No ...
	FpreAditionalInformasi11No = "e113"
	//FpreAditionalInformasi11Other ...
	FpreAditionalInformasi11Other = "e114"
	// FpreAditionalInformasi12Yes ...
	FpreAditionalInformasi12Yes = "e115"
	//FpreAditionalInformasi12No ...
	FpreAditionalInformasi12No = "e116"
	// FpreAditionalInformasi13Yes ...
	FpreAditionalInformasi13Yes = "e117"
	//FpreAditionalInformasi13No ...
	FpreAditionalInformasi13No = "e118"
	//FpreAditionalInformasi14Yes ...
	FpreAditionalInformasi14Yes = "e119"
	//FpreAditionalInformasi15Yes ...
	FpreAditionalInformasi15Yes = "f111"
	// FpreAditionalInformasi16Yes ...
	FpreAditionalInformasi16Yes = "f112"
	//FpreAditionalInformasi16No ...
	FpreAditionalInformasi16No = "f113"
	// FpreAditionalInformasi17Yes ...
	FpreAditionalInformasi17Yes = "f114"
	//FpreAditionalInformasi17No ...
	FpreAditionalInformasi17No = "f115"
	// FpreAditionalInformasi18Yes ...
	FpreAditionalInformasi18Yes = "f116"
	// FpreTujuanIvestasi01 ...
	FpreTujuanIvestasi01 = "f117"
	// FpreTujuanIvestasi02 ...
	FpreTujuanIvestasi02 = "f118"
	// FpreTujuanIvestasi03 ...
	FpreTujuanIvestasi03 = "f119"
	// FpreTujuanIvestasi04 ...
	FpreTujuanIvestasi04 = "g111"
	// FpreTujuanIvestasi05 ...
	FpreTujuanIvestasi05 = "g112"
	// FpreTujuanIvestasi06 ...
	FpreTujuanIvestasi06 = "g113"
)

func main() {
	inputPath := "fpre.pdf"
	outputPath := "output.pdf"

	userData := map[string]string{
		FpreCode:                          "L8MSSNU2",
		FpreName:                          "Thoriq Aziz Asuro",
		FpreService:                       "Online",
		FpreBranch:                        "Mandiri",
		FpreSales:                         "",
		FpreEventCode:                     "",
		FpreClientID:                      "",
		FpreCheckBoxDocument01:            "V",
		FpreCheckBoxDocument02:            "V",
		FpreCheckBoxDocument03:            "V",
		FpreCheckBoxDocument04:            "",
		FpreCheckBoxDocument05:            "",
		FpreCheckBoxDocument06:            "",
		FpreCheckBoxDocument07:            "",
		FpreCheckBoxDocument08:            "",
		FpreCheckBoxDocument09:            "",
		FpreDate01:                        "01",
		FpreDate02:                        "08",
		FpreDate03:                        "2020",
		FpreCheckBoxRRB01:                 "V",
		FpreCheckBoxRRB02:                 "",
		FpreCheckBoxAccountType01:         "",
		FpreCheckBoxAccountType02:         "V",
		FpreCheckBoxAccountType03:         "V",
		FpreCheckBoxBehalfOf01:            "V",
		FpreCheckBoxBehalfOf02:            "",
		FpreBehalfOfOther:                 "",
		FpreFullName:                      "Thoriq Aziz Asuro",
		FpreKtpName:                       "Thoriq Aziz Asuro",
		FpreCheckBoxGender01:              "V",
		FpreCheckBoxGender02:              "V",
		FpreCheckBoxIdentity01:            "V",
		FpreCheckBoxIdentity02:            "",
		FpreIdentityNumber:                "7382738273",
		FpreCheckBoxValidity:              "V",
		FprePlaceOfBirth:                  "Surabaya",
		FpreDateOfBirth01:                 "05",
		FpreDateOfBirth02:                 "08",
		FpreDateOfBirth03:                 "1998",
		FpreCheckBoxMaritalStatusSingel:   "V",
		FpreCheckBoxMaritalStatusWindow:   "",
		FpreCheckBoxMaritalStatusMarried:  "V",
		FpreCheckBoxMaritalStatusWindower: "",
		FpreCheckBoxCitizenShip01:         "V",
		FpreCheckBoxCitizenShip02:         "",
		FpreMotherName:                    "Siti",
		FpreAddressID:                     "Jl Kayon",
		FpreRT:                            "08",
		FpreRW:                            "06",
		FprePostalCode:                    "0452",
		FpreVillage:                       "Ampel",
		FpreSubDistrict:                   "semampir",
		FpreCity:                          "Surabaya",
		FpreProvince:                      "Jawa Timur",
		FpreCheckBoxID:                    "V",
		FpreAddress:                       "Jl Kayon 2",
		FpreCurrentRt:                     "08",
		FpreCurrentRW:                     "05",
		FpreCurrentPostalCode:             "0349",
		FpreCurrentVillage:                "surabaya",
		FpreCurrentSubDistrict:            "Semampir",
		FpreCurrentCity:                   "Suarabaya",
		FpreCurrentProvince:               "Jawa Timur",
		FpreResidentialStatus01:           "V",
		FpreResidentialStatus02:           "V",
		FpreResidentialStatus03:           "V",
		FpreResidentialStatus04:           "V",
		FpreResidentialStatus05:           "V",
		FpreResidentialStatus06:           "V",
		FpreHomePhone:                     "9239283",
		FprePhoneCellular:                 "938298382",
		FpreEmail:                         "Thoriq.azizasuro@gmail.com",
		FpreNpwpCountryID:                 "V",
		FpreNpwpCountryOther:              "",
		FpreNpwpCountryOtherID:            "",
		FpreNpwpNumber:                    "02932093",
		FpreNpwpName:                      "thoriq",
		FpreEdu01:                         "V",
		FpreEdu02:                         "V",
		FpreEdu03:                         "V",
		FpreEdu04:                         "V",
		FpreEdu05:                         "V",
		FpreEdu06:                         "V",
		FpreEdu07:                         "V",
		FpreEdu08:                         "V",
		FpreEdu09:                         "",
		FpreReligion01:                    "V",
		FpreReligion02:                    "V",
		FpreReligion03:                    "V",
		FpreReligion04:                    "V",
		FpreReligion05:                    "V",
		FpreReligion06:                    "V",
		FpreReligion07:                    "V",
		FpreOce01:                         "V",
		FpreOce02:                         "V",
		FpreOce03:                         "V",
		FpreOce04:                         "V",
		FpreOce05:                         "V",
		FpreOce06:                         "V",
		FpreOce07:                         "V",
		FpreOce08:                         "V",
		FpreOce09:                         "V",
		FpreOce10:                         "V",
		FpreOce11:                         "V",
		FpreOce12:                         "V",
		FpreOce14:                         "27372",
		FpreBusinessName:                  "Twiscode",
		FprePeriodeService01:              "2020",
		FprePeriodeService02:              "01",
		FpreLineBusiness01:                "V",
		FpreLineBusiness02:                "V",
		FpreLineBusiness03:                "V",
		FpreLineBusiness04:                "V",
		FpreLineBusiness05:                "V",
		FpreLineBusiness06:                "V",
		FpreLineBusiness07:                "V",
		FpreLineBusiness08:                "V",
		FpreLineBusiness09:                "V",
		FpreLineBusiness10:                "V",
		FpreLineBusiness11:                "V",
		FpreLineBusiness12:                "V",
		FpreLineBusiness13:                "V",
		FpreLineBusiness14:                "V",
		FpreLineBusiness15:                "V",
		FpreLineBusiness16:                "V",
		FpreLineBusiness17:                "V",
		FpreLineBusiness18:                "V",
		FpreLineBusiness19:                "V",
		FpreLineBusiness20:                "V",
		FpreDesignation01:                 "V",
		FpreDesignation02:                 "V",
		FpreDesignation03:                 "V",
		FpreDesignation04:                 "V",
		FpreDesignation05:                 "V",
		FpreDesignation06:                 "V",
		FpreDesignation07:                 "V",
		FpreDesignation08:                 "",
		FpreBusinessAddress:               "Jl Kayon 3",
		FpreBusinessCity:                  "surabaya",
		FpreBusinessPostalCode:            "0328",
		FpreBusinessProvice:               "Jawa Timur",
		FpreBusinessPhone:                 "09323",
		FpreSourceOfFound01:               "V",
		FpreSourceOfFound02:               "V",
		FpreSourceOfFound03:               "V",
		FpreSourceOfFound04:               "V",
		FpreSourceOfFound05:               "V",
		FpreSourceOfFound06:               "V",
		FpreSourceOfFound07:               "V",
		FpreSourceOfFound08:               "V",
		FpreSourceOfFound09:               "V",
		FpreSourceOfFound10:               "Lainnyaa",
		FpreMainSourceIncome01:            "V",
		FpreMainSourceIncome02:            "V",
		FpreMainSourceIncome03:            "V",
		FpreMainSourceIncome04:            "V",
		FpreMainSourceIncome05:            "V",
		FpreMainSourceIncome06:            "V",
		FpreMainSourceIncome07:            "V",
		FpreMainSourceIncome08:            "V",
		FpreMainSourceIncome09:            "V",
		FpreSalary:                        "10000",
		FprAditionalSourceIncome01:        "V",
		FprAditionalSourceIncome02:        "V",
		FprAditionalSourceIncome03:        "V",
		FprAditionalSourceIncome04:        "V",
		FprAditionalSourceIncome05:        "V",
		FprAditionalSourceIncome06:        "V",
		FprAditionalSourceIncome07:        "V",
		FprAditionalSourceIncome08:        "V",
		FprAditionalSourceIncome09:        "V",
		FpreAditionalIncome:               "100.000,00",
		FpreMallingAddress01:              "V",
		FpreMallingAddress02:              "V",
		FpreMallingAddress03:              "V",
		FpreRelationAccountName:           "SROBUWU",
		FpreRelationBankName:              "XYZ",
		FpreRelationBranch:                "AWEF",
		FpreRelationAccountNumber:         "08080811",
		FpreRelationName:                  "MASON GREENWOOD",
		FpreRelationOccupation01:          "V",
		FpreRelationOccupation02:          "V",
		FpreRelationOccupation03:          "V",
		FpreRelationOccupation04:          "V",
		FpreRelationOccupation05:          "V",
		FpreRelationOccupation06:          "V",
		FpreRelationOccupation07:          "V",
		FpreRelationOccupation08:          "V",
		FpreRelationOccupation09:          "V",
		FpreRelationOccupation10:          "V",
		FpreRelationOccupation11:          "V",
		FpreRelationOccupation12:          "V",
		FpreRelationOccupation13:          "FREE TEXT",
		FpreRelationOccupation14:          "Employee",
		FpreRelationBusinessName:          "TEST",
		// FpreRelationPeriodeServiceYear:        strconv.Itoa(rpJobTitleStartAtYear),
		// FpreRelationPeriodeServiceMonth:       strconv.Itoa(rpJobTitleStartAtMonth),
		// FpreRelationLineOfBusiness01:          CheckBoxChecking(user.RpLineOfBusinessMappingName, "05"),
		// FpreRelationLineOfBusiness02:          CheckBoxChecking(user.RpLineOfBusinessMappingName, "37"),
		// FpreRelationLineOfBusiness03:          CheckBoxChecking(user.RpLineOfBusinessMappingName, "16"),
		// FpreRelationLineOfBusiness04:          CheckBoxChecking(user.RpLineOfBusinessMappingName, "10"),
		// FpreRelationLineOfBusiness05:          CheckBoxChecking(user.RpLineOfBusinessMappingName, "24"),
		// FpreRelationLineOfBusiness06:          CheckBoxChecking(user.RpLineOfBusinessMappingName, "15"),
		// FpreRelationLineOfBusiness07:          CheckBoxChecking(user.RpLineOfBusinessMappingName, "08"),
		// FpreRelationLineOfBusiness08:          CheckBoxChecking(user.RpLineOfBusinessMappingName, "09"),
		// FpreRelationLineOfBusiness09:          CheckBoxChecking(user.RpLineOfBusinessMappingName, "07"),
		// FpreRelationLineOfBusiness10:          CheckBoxChecking(user.RpLineOfBusinessMappingName, "00"),
		// FpreRelationLineOfBusiness11:          CheckBoxChecking(user.RpLineOfBusinessMappingName, "06"),
		// FpreRelationLineOfBusiness12:          CheckBoxChecking(user.RpLineOfBusinessMappingName, "34"),
		// FpreRelationLineOfBusiness13:          CheckBoxChecking(user.RpLineOfBusinessMappingName, "26"),
		// FpreRelationLineOfBusiness14:          CheckBoxChecking(user.RpLineOfBusinessMappingName, "14"),
		// FpreRelationLineOfBusiness15:          CheckBoxChecking(user.RpLineOfBusinessMappingName, "08"),
		// FpreRelationLineOfBusiness16:          CheckBoxChecking(user.RpLineOfBusinessMappingName, "04"),
		// FpreRelationLineOfBusiness17:          CheckBoxChecking(user.RpLineOfBusinessMappingName, "40"),
		// FpreRelationLineOfBusiness18:          CheckBoxChecking(user.RpLineOfBusinessMappingName, "28"),
		// FpreRelationLineOfBusiness19:          CheckBoxChecking(user.RpLineOfBusinessMappingName, "03"),
		// FpreRelationLineOfBusiness20:          CheckBoxChecking(user.RpLineOfBusinessMappingName, "--lainnya--"),
		// FpreRelationJobs01:                    CheckBoxChecking(user.RpJobTitleMappingName, "06"),
		// FpreRelationJobs02:                    CheckBoxChecking(user.RpJobTitleMappingName, "01"),
		// FpreRelationJobs03:                    CheckBoxChecking(user.RpJobTitleMappingName, "07"),
		// FpreRelationJobs04:                    CheckBoxChecking(user.RpJobTitleMappingName, "05"),
		// FpreRelationJobs05:                    CheckBoxChecking(user.RpJobTitleMappingName, "03"),
		// FpreRelationJobs06:                    CheckBoxChecking(user.RpJobTitleMappingName, "04"),
		// FpreRelationJobs07:                    CheckBoxChecking(user.RpJobTitleMappingName, "--lainnya--"),
		// FpreRelationBusinessAdderess:          user.RpCompanyAddress,
		// FpreRelationBusinessCity:              user.RpCompanyCityMappingName,
		// FpreRelationBusinessPostalCode:        user.RpCompanySubDistrictPostalCode,
		// FpreRelationBusinessProvince:          user.RpCompanySubDistrictName,
		// FpreRelationBusinessPhone:             user.RpCompanyPhone,
		// FpreRelationAmanualGross:              number.FormatCurrency(user.RpPrimaryIncomeMaxValue, "IDR", ".", ",", 2),
		// FpreRelationAditionalIncome:           "",
		// FpreAditionalInformasi01Yes:           CheckBoxChekingBool(true, user.HasRelationToFinancialInstitution),
		// FpreAditionalInformasi01No:            CheckBoxChekingBool(false, user.HasRelationToFinancialInstitution),
		// FpreAditionalInformasi01FreeText:      TmClientAddInfoRelationToFinancialInstitutionDetail(user.HasRelationToFinancialInstitution, user.RelationToFinancialInstitutionDetails, user.LineOfBusinessMappingName, user.RpLineOfBusinessMappingName, user.CompanyName, user.RpCompanyName),
		// FpreAditionalInformasi02Yes:           CheckBoxChekingBool(true, user.HasRelationToMansek),
		// FpreAditionalInformasi02No:            CheckBoxChekingBool(false, user.HasRelationToMansek),
		// FpreAditionalInformasiName:            TmClientAddInfoRelationToMansekDetail(user.HasRelationToMansek, user.RelationToMansekDetails, 0),
		// FpreAditionalInformasiHubungan:        TmClientAddInfoRelationToMansekDetail(user.HasRelationToMansek, user.RelationToMansekDetails, 1),
		// FpreAditionalInformasi03Yes:           CheckBoxChekingBool(true, user.RelationHasMansekAccount),
		// FpreAditionalInformasi03No:            CheckBoxChekingBool(false, user.RelationHasMansekAccount),
		// FpreAditionalInformasi03Name:          TmClientAddInfoRelationHasMansekAccountDetail(user.RelationHasMansekAccount, user.RelationHasMansekAccountDetails, 0),
		// FpreAditionalInformasi03accountNumber: TmClientAddInfoRelationHasMansekAccountDetail(user.RelationHasMansekAccount, user.RelationHasMansekAccountDetails, 1),
		// FpreAditionalInformasi04Yes:           CheckBoxChekingBool(true, user.HasMansekAccount),
		// FpreAditionalInformasi04No:            CheckBoxChekingBool(false, user.HasMansekAccount),
		// FpreAditionalInformasi04Name:          TmClientAddInfoMansekAccountDetail(user.HasMansekAccount, user.MansekAccountDetails, 0),
		// FpreAditionalInformasi04Hubungan:      TmClientAddInfoMansekAccountDetail(user.HasMansekAccount, user.MansekAccountDetails, 1),
		// FpreAditionalInformasi04accountNumber: TmClientAddInfoMansekAccountDetail(user.HasMansekAccount, user.MansekAccountDetails, 2),
		// FpreAditionalInformasi05Yes:           CheckBoxChekingBool(true, user.HasControllingStakeAccount),
		// FpreAditionalInformasi05No:            CheckBoxChekingBool(false, user.HasControllingStakeAccount),
		// FpreAditionalInformasi05CompanyName:   TmClientAddInfoControllingStakeAccountDetail(user.HasControllingStakeAccount, user.ControllingStakeAccountDetails, 0),
		// FpreAditionalInformasi06Yes:           CheckBoxChekingBool(true, user.HasControllingStakeAccount),
		// FpreAditionalInformasi06No:            CheckBoxChekingBool(false, user.HasControllingStakeAccount),
		// FpreAditionalInformasi06CompanyName:   TmClientAddInfoControllingStakeAccountDetail(user.HasControllingStakeAccount, user.ControllingStakeAccountDetails, 0),
		// FpreAditionalInformasi06Position:      "",
		// FpreAditionalInformasi07Yes:           CheckBoxChekingBool(true, user.IsPeStakeholder),
		// FpreAditionalInformasi07No:            CheckBoxChekingBool(false, user.IsPeStakeholder),
		// FpreAditionalInformasi07CompanyName:   TmClientAddInfoPeStakeholderDetail(user.IsPeStakeholder, user.PeStakeholderDetails, 0),
		// FpreAditionalInformasi07Position:      TmClientAddInfoPeStakeholderDetail(user.IsPeStakeholder, user.PeStakeholderDetails, 1),
		// FpreAditionalInformasi07Saham:         TmClientAddInfoPeStakeholderDetail(user.IsPeStakeholder, user.PeStakeholderDetails, 2),
		// FpreAditionalInformasi08Yes:           CheckBoxChekingBool(true, user.HasPoliticallyExposedRelation),
		// FpreAditionalInformasi08No:            CheckBoxChekingBool(false, user.HasPoliticallyExposedRelation),
		// FpreAditionalInformasi08Position:      TmClientAddInfoPoliticallyExposedRelationDetail(user.HasPoliticallyExposedRelation, user.PoliticallyExposedRelationDetails, 0),
		// FpreAditionalInformasi09Yes:           CheckBoxChekingBool(true, user.HasPoliticallyExposedRelation),
		// FpreAditionalInformasi09No:            CheckBoxChekingBool(false, user.HasPoliticallyExposedRelation),
		// FpreAditionalInformasi09Position:      TmClientAddInfoPoliticallyExposedRelationDetail(user.HasPoliticallyExposedRelation, user.PoliticallyExposedRelationDetails, 0),
		// FpreAditionalInformasi10Yes:           CheckBoxChekingBool(true, user.HasRelationToEmmiten),
		// FpreAditionalInformasi10No:            CheckBoxChekingBool(false, user.HasRelationToEmmiten),
		// FpreAditionalInformasi10CompanyName:   TmClientAddInfoRelationToEmittenDetail(user.HasRelationToEmmiten, user.RelationToEmittenDetails, 0),
		// FpreAditionalInformasi11Yes:           CheckBoxChekingBool(true, user.IsUsCitizen),
		// FpreAditionalInformasi11No:            CheckBoxChekingBool(false, user.IsUsCitizen),
		// FpreAditionalInformasi11Other:         "",
		// FpreAditionalInformasi12Yes:           CheckBoxChekingBool(true, user.IsUsGreencardHolder),
		// FpreAditionalInformasi12No:            CheckBoxChekingBool(false, user.IsUsGreencardHolder),
		// FpreAditionalInformasi13Yes:           CheckBoxChekingBool(true, user.IsUsPermanentResidence),
		// FpreAditionalInformasi13No:            CheckBoxChekingBool(false, user.IsUsPermanentResidence),
		// FpreAditionalInformasi14Yes:           CheckBoxChekingBool(true, user.IsUsPermanentResidence),
		// FpreAditionalInformasi15Yes:           CheckBoxChekingBool(false, user.IsUsPermanentResidence),
		// FpreAditionalInformasi16Yes:           "",
		// FpreAditionalInformasi10No:            "✓",
		// FpreAditionalInformasi17Yes:           CheckBoxChekingBool(true, user.HasRelationToHighRiskCountry),
		// FpreAditionalInformasi17No:            CheckBoxChekingBool(false, user.HasRelationToHighRiskCountry),
		// FpreAditionalInformasi18Yes:           HighriskCountryList(user.RelationToHighRiskCountryDetails),
		// FpreTujuanIvestasi01:                  CheckBoxChecking(user.InvestmentPurposeMappingName, "1"),
		// FpreTujuanIvestasi02:                  CheckBoxChecking(user.InvestmentPurposeMappingName, "4"),
		// FpreTujuanIvestasi03:                  CheckBoxChecking(user.InvestmentPurposeMappingName, "3"),
		// FpreTujuanIvestasi04:                  CheckBoxChecking(user.InvestmentPurposeMappingName, "2"),
		// FpreTujuanIvestasi05:                  CheckBoxChecking(user.InvestmentPurposeMappingName, "5"),
		// FpreTujuanIvestasi06:                  user.InvestmentPurposeFreeText,
	}

	err := AddTextToPdf(inputPath, outputPath, userData)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Complete, see output file: %s\n", outputPath)
}

// AddTextToPdf ...
func AddTextToPdf(inputPath, outputPath string, userData map[string]string) (err error) {
	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	pdfReader, err := pdf.NewPdfReader(f)
	if err != nil {
		return err
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}

	c := creator.New()

	// Load the pages.
	for i := 0; i < numPages; i++ {
		page, err := pdfReader.GetPage(i + 1)
		if err != nil {
			return err
		}

		err = c.AddPage(page)
		if err != nil {
			return err
		}

		for key, value := range userData {
			position := strings.Split(key, "-")
			if len(position) == 3 {
				page, _ := strconv.Atoi(position[2])
				if i == page || page == -1 {
					p := creator.NewParagraph(value)
					p.SetFont(fonts.NewFontTimesBold())
					xPos, _ := strconv.ParseFloat(position[0], 64)
					yPos, _ := strconv.ParseFloat(position[1], 64)
					p.SetPos(xPos, yPos)
					_ = c.Draw(p)
				}
			}
		}

	}

	err = c.WriteToFile(outputPath)
	return err
}
