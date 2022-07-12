package nextpay

import "errors"

type NextPayError error

// ToDo: Translate them to english? maybe?
var (
	ErrCode0  NextPayError = errors.New("پرداخت تکمیل و با موفقیت انجام شده است")
	ErrCode1  NextPayError = errors.New("منتظر ارسال تراکنش و ادامه پرداخت") //
	ErrCode2  NextPayError = errors.New("پرداخت رد شده توسط کاربر یا بانک")
	ErrCode3  NextPayError = errors.New("پرداخت در حال انتظار جواب بانک")
	ErrCode4  NextPayError = errors.New("پرداخت لغو شده است")
	ErrCode20 NextPayError = errors.New("کد api_key ارسال نشده است")
	ErrCode21 NextPayError = errors.New("کد trans_id ارسال نشده است")
	ErrCode22 NextPayError = errors.New("مبلغ ارسال نشده")
	ErrCode23 NextPayError = errors.New("لینک ارسال نشده")
	ErrCode24 NextPayError = errors.New("مبلغ صحیح نیست")
	ErrCode25 NextPayError = errors.New("تراکنش قبلا انجام و قابل ارسال نیست")
	ErrCode26 NextPayError = errors.New("مقدار توکن ارسال نشده است")
	ErrCode27 NextPayError = errors.New("شماره سفارش صحیح نیست")
	ErrCode28 NextPayError = errors.New("مقدار فیلد سفارشی [custom_json_fields] از نوع json نیست")
	ErrCode29 NextPayError = errors.New("کد بازگشت مبلغ صحیح نیست")
	ErrCode30 NextPayError = errors.New("مبلغ کمتر از حداقل پرداختی است")
	ErrCode31 NextPayError = errors.New("صندوق کاربری موجود نیست")
	ErrCode32 NextPayError = errors.New("مسیر بازگشت صحیح نیست")
	ErrCode33 NextPayError = errors.New("کلید مجوز دهی صحیح نیست")
	ErrCode34 NextPayError = errors.New("کد تراکنش صحیح نیست")
	ErrCode35 NextPayError = errors.New("ساختار کلید مجوز دهی صحیح نیست")
	ErrCode36 NextPayError = errors.New("شماره سفارش ارسال نشد است")
	ErrCode37 NextPayError = errors.New("شماره تراکنش یافت نشد")
	ErrCode38 NextPayError = errors.New("توکن ارسالی موجود نیست")
	ErrCode39 NextPayError = errors.New("کلید مجوز دهی موجود نیست")
	ErrCode40 NextPayError = errors.New("کلید مجوزدهی مسدود شده است")
	ErrCode41 NextPayError = errors.New("خطا در دریافت پارامتر، شماره شناسایی صحت اعتبار که از بانک ارسال شده موجود نیست")
	ErrCode42 NextPayError = errors.New("سیستم پرداخت دچار مشکل شده است")
	ErrCode43 NextPayError = errors.New("درگاه پرداختی برای انجام درخواست یافت نشد")
	ErrCode44 NextPayError = errors.New("پاسخ دریاف شده از بانک نامعتبر است")
	ErrCode45 NextPayError = errors.New("سیستم پرداخت غیر فعال است")
	ErrCode46 NextPayError = errors.New("درخواست نامعتبر")
	ErrCode47 NextPayError = errors.New("کلید مجوز دهی یافت نشد [حذف شده]")
	ErrCode48 NextPayError = errors.New("نرخ کمیسیون تعیین نشده است")
	ErrCode49 NextPayError = errors.New("تراکنش مورد نظر تکراریست")
	ErrCode50 NextPayError = errors.New("حساب کاربری برای صندوق مالی یافت نشد")
	ErrCode51 NextPayError = errors.New("شناسه کاربری یافت نشد")
	ErrCode52 NextPayError = errors.New("حساب کاربری تایید نشده است")
	ErrCode60 NextPayError = errors.New("ایمیل صحیح نیست")
	ErrCode61 NextPayError = errors.New("کد ملی صحیح نیست")
	ErrCode62 NextPayError = errors.New("کد پستی صحیح نیست")
	ErrCode63 NextPayError = errors.New("آدرس پستی صحیح نیست و یا بیش از ۱۵۰ کارکتر است")
	ErrCode64 NextPayError = errors.New("توضیحات صحیح نیست و یا بیش از ۱۵۰ کارکتر است")
	ErrCode65 NextPayError = errors.New("نام و نام خانوادگی صحیح نیست و یا بیش از ۳۵ کاکتر است")
	ErrCode66 NextPayError = errors.New("تلفن صحیح نیست")
	ErrCode67 NextPayError = errors.New("نام کاربری صحیح نیست یا بیش از ۳۰ کارکتر است")
	ErrCode68 NextPayError = errors.New("نام محصول صحیح نیست و یا بیش از ۳۰ کارکتر است")
	ErrCode69 NextPayError = errors.New("آدرس ارسالی برای بازگشت موفق صحیح نیست و یا بیش از ۱۰۰ کارکتر است")
	ErrCode70 NextPayError = errors.New("آدرس ارسالی برای بازگشت ناموفق صحیح نیست و یا بیش از ۱۰۰ کارکتر است")
	ErrCode71 NextPayError = errors.New("موبایل صحیح نیست")
	ErrCode72 NextPayError = errors.New("بانک پاسخگو نبوده است لطفا با نکست پی تماس بگیرید")
	ErrCode73 NextPayError = errors.New("مسیر بازگشت دارای خطا میباشد یا بسیار طولانیست")
	ErrCode90 NextPayError = errors.New("بازگشت مبلغ بدرستی انجام شد")
	ErrCode91 NextPayError = errors.New("عملیات ناموفق در بازگشت مبلغ")
	ErrCode92 NextPayError = errors.New("در عملیات بازگشت مبلغ خطا رخ داده است")
	ErrCode93 NextPayError = errors.New("موجودی صندوق کاربری برای بازگشت مبلغ کافی نیست")
	ErrCode94 NextPayError = errors.New("کلید بازگشت مبلغ یافت نشد")
)

var Codes map[int]error = map[int]error{
	0:   ErrCode0,
	-1:  ErrCode1, //
	-2:  ErrCode2,
	-3:  ErrCode3,
	-4:  ErrCode4,
	-20: ErrCode20,
	-21: ErrCode21,
	-22: ErrCode22,
	-23: ErrCode23,
	-24: ErrCode24,
	-25: ErrCode25,
	-26: ErrCode26,
	-27: ErrCode27,
	-28: ErrCode28,
	-29: ErrCode29,
	-30: ErrCode30,
	-31: ErrCode31,
	-32: ErrCode32,
	-33: ErrCode33,
	-34: ErrCode34,
	-35: ErrCode35,
	-36: ErrCode36,
	-37: ErrCode37,
	-38: ErrCode38,
	-39: ErrCode39,
	-40: ErrCode40,
	-41: ErrCode41,
	-42: ErrCode42,
	-43: ErrCode43,
	-44: ErrCode44,
	-45: ErrCode45,
	-46: ErrCode46,
	-47: ErrCode47,
	-48: ErrCode48,
	-49: ErrCode49,
	-50: ErrCode50,
	-51: ErrCode51,
	-52: ErrCode52,
	-60: ErrCode60,
	-61: ErrCode61,
	-62: ErrCode62,
	-63: ErrCode63,
	-64: ErrCode64,
	-65: ErrCode65,
	-66: ErrCode66,
	-67: ErrCode67,
	-68: ErrCode68,
	-69: ErrCode69,
	-70: ErrCode70,
	-71: ErrCode71,
	-72: ErrCode72,
	-73: ErrCode73,
	-90: ErrCode90,
	-91: ErrCode91,
	-92: ErrCode92,
	-93: ErrCode93,
	-94: ErrCode94,
}

func findErrorByCode(currentCode, rightCode int) error {
	if currentCode != rightCode {
		if err, ok := Codes[currentCode]; ok {
			return err
		}
	}
	return nil
}
