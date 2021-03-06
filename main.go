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
	FpreRelationPeriodeServiceYear = "142-276-2"
	// FpreRelationPeriodeServiceMonth ...
	FpreRelationPeriodeServiceMonth = "220-276-2"
	// FpreRelationLineOfBusiness01 ...
	FpreRelationLineOfBusiness01 = "13-302-2"
	// FpreRelationLineOfBusiness02 ...
	FpreRelationLineOfBusiness02 = "70-302-2"
	// FpreRelationLineOfBusiness03 ...
	FpreRelationLineOfBusiness03 = "123-302-2"
	// FpreRelationLineOfBusiness04 ...
	FpreRelationLineOfBusiness04 = "201-302-2"
	// FpreRelationLineOfBusiness05 ...
	FpreRelationLineOfBusiness05 = "13-317-2"
	// FpreRelationLineOfBusiness06 ...
	FpreRelationLineOfBusiness06 = "70-317-2"
	// FpreRelationLineOfBusiness07 ...
	FpreRelationLineOfBusiness07 = "123-317-2"
	// FpreRelationLineOfBusiness08 ...
	FpreRelationLineOfBusiness08 = "201-317-2"
	// FpreRelationLineOfBusiness09 ...
	FpreRelationLineOfBusiness09 = "13-332-2"
	// FpreRelationLineOfBusiness10 ...
	FpreRelationLineOfBusiness10 = "70-332-2"
	// FpreRelationLineOfBusiness11 ...
	FpreRelationLineOfBusiness11 = "123-332-2"
	// FpreRelationLineOfBusiness12 ...
	FpreRelationLineOfBusiness12 = "201-332-2"
	// FpreRelationLineOfBusiness13 ...
	FpreRelationLineOfBusiness13 = "13-349-2"
	// FpreRelationLineOfBusiness14 ...
	FpreRelationLineOfBusiness14 = "70-349-2"
	// FpreRelationLineOfBusiness15 ...
	FpreRelationLineOfBusiness15 = "123-349-2"
	// FpreRelationLineOfBusiness16 ...
	FpreRelationLineOfBusiness16 = "201-349-2"
	// FpreRelationLineOfBusiness17 ...
	FpreRelationLineOfBusiness17 = "13-364-2"
	// FpreRelationLineOfBusiness18 ...
	FpreRelationLineOfBusiness18 = "70-365-2"
	// FpreRelationLineOfBusiness19 ...
	FpreRelationLineOfBusiness19 = "123-365-2"
	// FpreRelationLineOfBusiness20 ...
	FpreRelationLineOfBusiness20 = "201-365-2"
	// FpreRelationJobs01 ...
	FpreRelationJobs01 = "14-392-2"
	// FpreRelationJobs02 ...
	FpreRelationJobs02 = "81-392-2"
	// FpreRelationJobs03 ...
	FpreRelationJobs03 = "142-392-2"
	// FpreRelationJobs04 ...
	FpreRelationJobs04 = "201-392-2"
	// FpreRelationJobs05 ...
	FpreRelationJobs05 = "14-408-2"
	// FpreRelationJobs06 ...
	FpreRelationJobs06 = "81-408-2"
	// FpreRelationJobs07 ...
	FpreRelationJobs07 = "142-408-2"
	// FpreRelationBusinessAdderess ...
	FpreRelationBusinessAdderess = "15-435-2"
	// FpreRelationBusinessCity ...
	FpreRelationBusinessCity = "76-454-2"
	// FpreRelationBusinessPostalCode ...
	FpreRelationBusinessPostalCode = "233-454-2"
	// FpreRelationBusinessProvince ...
	FpreRelationBusinessProvince = "78-471-2"
	// FpreRelationBusinessPhone ...
	FpreRelationBusinessPhone = "94-486-2"
	// FpreRelationAmanualGross ...
	FpreRelationAmanualGross = "24-518-2"
	// FpreRelationAditionalIncome ...
	FpreRelationAditionalIncome = "24-543-2"
	//FpreAditionalInformasi01Yes ...
	FpreAditionalInformasi01Yes = "234-573-2"
	//FpreAditionalInformasi01No ...
	FpreAditionalInformasi01No = "260-573-2"
	//FpreAditionalInformasi01FreeText ...
	FpreAditionalInformasi01FreeText = "102-620-2"
	//FpreAditionalInformasi02Yes ...
	FpreAditionalInformasi02Yes = "234-641-2"
	//FpreAditionalInformasi02No ...
	FpreAditionalInformasi02No = "260-641-2"
	// FpreAditionalInformasiName ...
	FpreAditionalInformasiName = "48-670-2"
	// FpreAditionalInformasiHubungan ...
	FpreAditionalInformasiHubungan = "62-680-2"
	//FpreAditionalInformasi03Yes ...
	FpreAditionalInformasi03Yes = "234-702-2"
	//FpreAditionalInformasi03No ...
	FpreAditionalInformasi03No = "260-702-2"
	// FpreAditionalInformasi03Name ...
	FpreAditionalInformasi03Name = "48-718-2"
	// FpreAditionalInformasi03accountNumber ...
	FpreAditionalInformasi03accountNumber = "94-728-2"
	//FpreAditionalInformasi04Yes ...
	FpreAditionalInformasi04Yes = "234-762-2"
	//FpreAditionalInformasi04No ...
	FpreAditionalInformasi04No = "260-762-2"
	// FpreAditionalInformasi04Name ...
	FpreAditionalInformasi04Name = "48-762-2"
	// FpreAditionalInformasi04Hubungan ...
	FpreAditionalInformasi04Hubungan = "64-770-2"
	// FpreAditionalInformasi04accountNumber ...
	FpreAditionalInformasi04accountNumber = "94-778-2"
	// FpreAditionalInformasi05Yes ...
	FpreAditionalInformasi05Yes = "531-52-2"
	//FpreAditionalInformasi05No ...
	FpreAditionalInformasi05No = "557-52-2"
	// FpreAditionalInformasi05CompanyName ...
	FpreAditionalInformasi05CompanyName = "408-80-2"
	// FpreAditionalInformasi06Yes ...
	FpreAditionalInformasi06Yes = "531-92-2"
	//FpreAditionalInformasi06No ...
	FpreAditionalInformasi06No = "557-92-2"
	// FpreAditionalInformasi06CompanyName ...
	FpreAditionalInformasi06CompanyName = "408-120-2"
	// FpreAditionalInformasi06Position ...
	FpreAditionalInformasi06Position = "380-130-2"
	// FpreAditionalInformasi07Yes ...
	FpreAditionalInformasi07Yes = "531-144-2"
	//FpreAditionalInformasi07No ...
	FpreAditionalInformasi07No = "557-144-2"
	// FpreAditionalInformasi07CompanyName ...
	FpreAditionalInformasi07CompanyName = "403-164-2"
	// FpreAditionalInformasi07Position ...
	FpreAditionalInformasi07Position = "360-176-2"
	// FpreAditionalInformasi07Saham ...
	FpreAditionalInformasi07Saham = "504-172-2"
	// FpreAditionalInformasi08Yes ...
	FpreAditionalInformasi08Yes = "531-190-2"
	//FpreAditionalInformasi08No ...
	FpreAditionalInformasi08No = "557-190-2"
	// FpreAditionalInformasi08Position ...
	FpreAditionalInformasi08Position = "360-214-2"
	// FpreAditionalInformasi09Yes ...
	FpreAditionalInformasi09Yes = "531-234-2"
	//FpreAditionalInformasi09No ...
	FpreAditionalInformasi09No = "557-234-2"
	// FpreAditionalInformasi09Position ...
	FpreAditionalInformasi09Position = "360-268-2"
	// FpreAditionalInformasi10Yes ...
	FpreAditionalInformasi10Yes = "531-288-2"
	//FpreAditionalInformasi10No ...
	FpreAditionalInformasi10No = "557-288-2"
	// FpreAditionalInformasi10CompanyName ...
	FpreAditionalInformasi10CompanyName = "408-311-2"
	// FpreAditionalInformasi11Yes ...
	FpreAditionalInformasi11Yes = "531-324-2"
	//FpreAditionalInformasi11No ...
	FpreAditionalInformasi11No = "557-324-2"
	//FpreAditionalInformasi11Other ...
	FpreAditionalInformasi11Other = "531-341-2"
	// FpreAditionalInformasi12Yes ...
	FpreAditionalInformasi12Yes = "531-406-2"
	//FpreAditionalInformasi12No ...
	FpreAditionalInformasi12No = "557-406-2"
	// FpreAditionalInformasi13Yes ...
	FpreAditionalInformasi13Yes = "531-440-2"
	//FpreAditionalInformasi13No ...
	FpreAditionalInformasi13No = "557-440-2"
	//FpreAditionalInformasi14Yes ...
	FpreAditionalInformasi14Yes = "360-516-2"
	//FpreAditionalInformasi15Yes ...
	FpreAditionalInformasi15Yes = "360-575-2"
	// FpreAditionalInformasi16Yes ...
	FpreAditionalInformasi16Yes = "531-599-2"
	//FpreAditionalInformasi16No ...
	FpreAditionalInformasi16No = "557-599-2"
	// FpreAditionalInformasi17Yes ...
	FpreAditionalInformasi17Yes = "531-683-2"
	//FpreAditionalInformasi17No ...
	FpreAditionalInformasi17No = "557-683-2"
	// FpreAditionalInformasi18Yes ...
	FpreAditionalInformasi18Yes = "20-76-3"
	// FpreTujuanIvestasi01 ...
	FpreTujuanIvestasi01 = "92-76-3"
	// FpreTujuanIvestasi02 ...
	FpreTujuanIvestasi02 = "155-77-3"
	// FpreTujuanIvestasi03 ...
	FpreTujuanIvestasi03 = "223-77-3"
	// FpreTujuanIvestasi04 ...
	FpreTujuanIvestasi04 = "20-107-3"
	// FpreTujuanIvestasi05 ...
	FpreTujuanIvestasi05 = "92-107-3"
	// FpreTujuanIvestasi06 ...
	FpreTujuanIvestasi06 = "453-650-3"
	// FprePernyataanName ...
	FprePernyataanName = "117-129-11"
	// FprePernyataanAddress ...
	FprePernyataanAddress = "117-147-11"
	// FprePernyataanNik ...
	FprePernyataanNik = "117-183-11"
)

func main() {
	// inputPath := "fpre.pdf"
	// outputPath := "output.pdf"

	inputPath := "fpre-syariah.pdf"
	outputPath := "output-syariah.pdf"

	userData := map[string]string{
		FpreCode:                              "L8MSSNU2",
		FpreName:                              "Thoriq Aziz Asuro",
		FpreService:                           "Online",
		FpreBranch:                            "Mandiri",
		FpreSales:                             "",
		FpreEventCode:                         "",
		FpreClientID:                          "",
		FpreCheckBoxDocument01:                "V",
		FpreCheckBoxDocument02:                "V",
		FpreCheckBoxDocument03:                "V",
		FpreCheckBoxDocument04:                "",
		FpreCheckBoxDocument05:                "",
		FpreCheckBoxDocument06:                "",
		FpreCheckBoxDocument07:                "",
		FpreCheckBoxDocument08:                "",
		FpreCheckBoxDocument09:                "",
		FpreDate01:                            "01",
		FpreDate02:                            "08",
		FpreDate03:                            "2020",
		FpreCheckBoxRRB01:                     "V",
		FpreCheckBoxRRB02:                     "",
		FpreCheckBoxAccountType01:             "",
		FpreCheckBoxAccountType02:             "V",
		FpreCheckBoxAccountType03:             "V",
		FpreCheckBoxBehalfOf01:                "V",
		FpreCheckBoxBehalfOf02:                "",
		FpreBehalfOfOther:                     "",
		FpreFullName:                          "Thoriq Aziz Asuro",
		FpreKtpName:                           "Thoriq Aziz Asuro",
		FpreCheckBoxGender01:                  "V",
		FpreCheckBoxGender02:                  "V",
		FpreCheckBoxIdentity01:                "V",
		FpreCheckBoxIdentity02:                "",
		FpreIdentityNumber:                    "7382738273",
		FpreCheckBoxValidity:                  "V",
		FprePlaceOfBirth:                      "Surabaya",
		FpreDateOfBirth01:                     "05",
		FpreDateOfBirth02:                     "08",
		FpreDateOfBirth03:                     "1998",
		FpreCheckBoxMaritalStatusSingel:       "V",
		FpreCheckBoxMaritalStatusWindow:       "",
		FpreCheckBoxMaritalStatusMarried:      "V",
		FpreCheckBoxMaritalStatusWindower:     "",
		FpreCheckBoxCitizenShip01:             "V",
		FpreCheckBoxCitizenShip02:             "",
		FpreMotherName:                        "Siti",
		FpreAddressID:                         "Jl Kayon",
		FpreRT:                                "08",
		FpreRW:                                "06",
		FprePostalCode:                        "0452",
		FpreVillage:                           "Ampel",
		FpreSubDistrict:                       "semampir",
		FpreCity:                              "Surabaya",
		FpreProvince:                          "Jawa Timur",
		FpreCheckBoxID:                        "V",
		FpreAddress:                           "Jl Kayon 2",
		FpreCurrentRt:                         "08",
		FpreCurrentRW:                         "05",
		FpreCurrentPostalCode:                 "0349",
		FpreCurrentVillage:                    "surabaya",
		FpreCurrentSubDistrict:                "Semampir",
		FpreCurrentCity:                       "Suarabaya",
		FpreCurrentProvince:                   "Jawa Timur",
		FpreResidentialStatus01:               "V",
		FpreResidentialStatus02:               "V",
		FpreResidentialStatus03:               "V",
		FpreResidentialStatus04:               "V",
		FpreResidentialStatus05:               "V",
		FpreResidentialStatus06:               "V",
		FpreHomePhone:                         "9239283",
		FprePhoneCellular:                     "938298382",
		FpreEmail:                             "Thoriq.azizasuro@gmail.com",
		FpreNpwpCountryID:                     "V",
		FpreNpwpCountryOther:                  "",
		FpreNpwpCountryOtherID:                "",
		FpreNpwpNumber:                        "02932093",
		FpreNpwpName:                          "thoriq",
		FpreEdu01:                             "V",
		FpreEdu02:                             "V",
		FpreEdu03:                             "V",
		FpreEdu04:                             "V",
		FpreEdu05:                             "V",
		FpreEdu06:                             "V",
		FpreEdu07:                             "V",
		FpreEdu08:                             "V",
		FpreEdu09:                             "",
		FpreReligion01:                        "V",
		FpreReligion02:                        "V",
		FpreReligion03:                        "V",
		FpreReligion04:                        "V",
		FpreReligion05:                        "V",
		FpreReligion06:                        "V",
		FpreReligion07:                        "V",
		FpreOce01:                             "V",
		FpreOce02:                             "V",
		FpreOce03:                             "V",
		FpreOce04:                             "V",
		FpreOce05:                             "V",
		FpreOce06:                             "V",
		FpreOce07:                             "V",
		FpreOce08:                             "V",
		FpreOce09:                             "V",
		FpreOce10:                             "V",
		FpreOce11:                             "V",
		FpreOce12:                             "V",
		FpreOce14:                             "27372",
		FpreBusinessName:                      "Twiscode",
		FprePeriodeService01:                  "2020",
		FprePeriodeService02:                  "01",
		FpreLineBusiness01:                    "V",
		FpreLineBusiness02:                    "V",
		FpreLineBusiness03:                    "V",
		FpreLineBusiness04:                    "V",
		FpreLineBusiness05:                    "V",
		FpreLineBusiness06:                    "V",
		FpreLineBusiness07:                    "V",
		FpreLineBusiness08:                    "V",
		FpreLineBusiness09:                    "V",
		FpreLineBusiness10:                    "V",
		FpreLineBusiness11:                    "V",
		FpreLineBusiness12:                    "V",
		FpreLineBusiness13:                    "V",
		FpreLineBusiness14:                    "V",
		FpreLineBusiness15:                    "V",
		FpreLineBusiness16:                    "V",
		FpreLineBusiness17:                    "V",
		FpreLineBusiness18:                    "V",
		FpreLineBusiness19:                    "V",
		FpreLineBusiness20:                    "V",
		FpreDesignation01:                     "V",
		FpreDesignation02:                     "V",
		FpreDesignation03:                     "V",
		FpreDesignation04:                     "V",
		FpreDesignation05:                     "V",
		FpreDesignation06:                     "V",
		FpreDesignation07:                     "V",
		FpreDesignation08:                     "",
		FpreBusinessAddress:                   "Jl Kayon 3",
		FpreBusinessCity:                      "surabaya",
		FpreBusinessPostalCode:                "0328",
		FpreBusinessProvice:                   "Jawa Timur",
		FpreBusinessPhone:                     "09323",
		FpreSourceOfFound01:                   "V",
		FpreSourceOfFound02:                   "V",
		FpreSourceOfFound03:                   "V",
		FpreSourceOfFound04:                   "V",
		FpreSourceOfFound05:                   "V",
		FpreSourceOfFound06:                   "V",
		FpreSourceOfFound07:                   "V",
		FpreSourceOfFound08:                   "V",
		FpreSourceOfFound09:                   "V",
		FpreSourceOfFound10:                   "Lainnyaa",
		FpreMainSourceIncome01:                "V",
		FpreMainSourceIncome02:                "V",
		FpreMainSourceIncome03:                "V",
		FpreMainSourceIncome04:                "V",
		FpreMainSourceIncome05:                "V",
		FpreMainSourceIncome06:                "V",
		FpreMainSourceIncome07:                "V",
		FpreMainSourceIncome08:                "V",
		FpreMainSourceIncome09:                "V",
		FpreSalary:                            "10000",
		FprAditionalSourceIncome01:            "V",
		FprAditionalSourceIncome02:            "V",
		FprAditionalSourceIncome03:            "V",
		FprAditionalSourceIncome04:            "V",
		FprAditionalSourceIncome05:            "V",
		FprAditionalSourceIncome06:            "V",
		FprAditionalSourceIncome07:            "V",
		FprAditionalSourceIncome08:            "V",
		FprAditionalSourceIncome09:            "V",
		FpreAditionalIncome:                   "100.000,00",
		FpreMallingAddress01:                  "V",
		FpreMallingAddress02:                  "V",
		FpreMallingAddress03:                  "V",
		FpreRelationAccountName:               "SROBUWU",
		FpreRelationBankName:                  "XYZ",
		FpreRelationBranch:                    "AWEF",
		FpreRelationAccountNumber:             "08080811",
		FpreRelationName:                      "MASON GREENWOOD",
		FpreRelationOccupation01:              "V",
		FpreRelationOccupation02:              "V",
		FpreRelationOccupation03:              "V",
		FpreRelationOccupation04:              "V",
		FpreRelationOccupation05:              "V",
		FpreRelationOccupation06:              "V",
		FpreRelationOccupation07:              "V",
		FpreRelationOccupation08:              "V",
		FpreRelationOccupation09:              "V",
		FpreRelationOccupation10:              "V",
		FpreRelationOccupation11:              "V",
		FpreRelationOccupation12:              "V",
		FpreRelationOccupation13:              "FREE TEXT",
		FpreRelationOccupation14:              "Employee",
		FpreRelationBusinessName:              "TEST",
		FpreRelationPeriodeServiceYear:        "2020",
		FpreRelationPeriodeServiceMonth:       "08",
		FpreRelationLineOfBusiness01:          "V",
		FpreRelationLineOfBusiness02:          "V",
		FpreRelationLineOfBusiness03:          "V",
		FpreRelationLineOfBusiness04:          "V",
		FpreRelationLineOfBusiness05:          "V",
		FpreRelationLineOfBusiness06:          "V",
		FpreRelationLineOfBusiness07:          "V",
		FpreRelationLineOfBusiness08:          "V",
		FpreRelationLineOfBusiness09:          "V",
		FpreRelationLineOfBusiness10:          "V",
		FpreRelationLineOfBusiness11:          "V",
		FpreRelationLineOfBusiness12:          "V",
		FpreRelationLineOfBusiness13:          "V",
		FpreRelationLineOfBusiness14:          "V",
		FpreRelationLineOfBusiness15:          "V",
		FpreRelationLineOfBusiness16:          "V",
		FpreRelationLineOfBusiness17:          "V",
		FpreRelationLineOfBusiness18:          "V",
		FpreRelationLineOfBusiness19:          "V",
		FpreRelationLineOfBusiness20:          "V",
		FpreRelationJobs01:                    "V",
		FpreRelationJobs02:                    "V",
		FpreRelationJobs03:                    "V",
		FpreRelationJobs04:                    "V",
		FpreRelationJobs05:                    "V",
		FpreRelationJobs06:                    "V",
		FpreRelationJobs07:                    "V",
		FpreRelationBusinessAdderess:          "JL Menuju Bahagia",
		FpreRelationBusinessCity:              "Subay",
		FpreRelationBusinessPostalCode:        "604",
		FpreRelationBusinessProvince:          "Jawa Tim",
		FpreRelationBusinessPhone:             "phone",
		FpreRelationAmanualGross:              "1.450.000,00",
		FpreRelationAditionalIncome:           "1.850.000,00",
		FpreAditionalInformasi01Yes:           "V",
		FpreAditionalInformasi01No:            "V",
		FpreAditionalInformasi01FreeText:      "RELASI RELASI",
		FpreAditionalInformasi02Yes:           "V",
		FpreAditionalInformasi02No:            "V",
		FpreAditionalInformasiName:            "RELASI NAME",
		FpreAditionalInformasiHubungan:        "RELASI HUB",
		FpreAditionalInformasi03Yes:           "V",
		FpreAditionalInformasi03No:            "V",
		FpreAditionalInformasi03Name:          "0",
		FpreAditionalInformasi03accountNumber: "0",
		FpreAditionalInformasi04Yes:           "V",
		FpreAditionalInformasi04No:            "V",
		FpreAditionalInformasi04Name:          "name",
		FpreAditionalInformasi04Hubungan:      "hub",
		FpreAditionalInformasi04accountNumber: "number",
		FpreAditionalInformasi05Yes:           "V",
		FpreAditionalInformasi05No:            "V",
		FpreAditionalInformasi05CompanyName:   "TWIS",
		FpreAditionalInformasi06Yes:           "V",
		FpreAditionalInformasi06No:            "V",
		FpreAditionalInformasi06CompanyName:   "CODE",
		FpreAditionalInformasi06Position:      "TWISC",
		FpreAditionalInformasi07Yes:           "V",
		FpreAditionalInformasi07No:            "V",
		FpreAditionalInformasi07CompanyName:   "CODE",
		FpreAditionalInformasi07Position:      "INFOR JABATAN",
		FpreAditionalInformasi07Saham:         "100s",
		FpreAditionalInformasi08Yes:           "V",
		FpreAditionalInformasi08No:            "V",
		FpreAditionalInformasi08Position:      "DAN LAIN LAIN",
		FpreAditionalInformasi09Yes:           "V",
		FpreAditionalInformasi09No:            "V",
		FpreAditionalInformasi09Position:      "WANNA COME HOME",
		FpreAditionalInformasi10Yes:           "V",
		FpreAditionalInformasi10No:            "V",
		FpreAditionalInformasi10CompanyName:   "SOMETIMES",
		FpreAditionalInformasi11Yes:           "V",
		FpreAditionalInformasi11No:            "V",
		FpreAditionalInformasi11Other:         "V",
		FpreAditionalInformasi12Yes:           "X",
		FpreAditionalInformasi12No:            "X",
		FpreAditionalInformasi13Yes:           "V",
		FpreAditionalInformasi13No:            "V",
		FpreAditionalInformasi14Yes:           "X",
		FpreAditionalInformasi15Yes:           "Y",
		FpreAditionalInformasi16Yes:           "A",
		FpreAditionalInformasi16No:            "Z",
		FpreAditionalInformasi17Yes:           "V",
		FpreAditionalInformasi17No:            "V",
		FpreAditionalInformasi18Yes:           "X",
		FpreTujuanIvestasi01:                  "X",
		FpreTujuanIvestasi02:                  "X",
		FpreTujuanIvestasi03:                  "D",
		FpreTujuanIvestasi04:                  "K",
		FpreTujuanIvestasi05:                  "J",
		FpreTujuanIvestasi06:                  "NAME",

		FprePernyataanName:    "TEST",
		FprePernyataanAddress: "SUBAY41, JL Uwuwuwuwu Awef Awef",
		FprePernyataanNik:     "***34345",
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

	//prepare-image
	img, err := creator.NewImageFromFile("assets/Kacab.png")
	if err != nil {
		return err
	}

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
					xPos, _ := strconv.ParseFloat(position[0], 64)
					yPos, _ := strconv.ParseFloat(position[1], 64)

					if value != "V" {
						p := creator.NewParagraph(value)
						p.SetFont(fonts.NewFontTimesBold())
						p.SetPos(xPos, yPos)
						_ = c.Draw(p)
					} else {
						img.ScaleToWidth(10)
						img.SetPos(xPos, yPos)
						_ = c.Draw(img)
					}
				}
			}
		}

	}

	err = c.WriteToFile(outputPath)
	return err
}
