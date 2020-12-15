import { addMinutes, isBefore } from 'date-fns';
import { setValue, getValue } from '../utils/local-storage';

export function login() {
	const now = new Date();
	const date = addMinutes(now, 1);

	setValue('login', true);
	setValue('expire', Number(date));
}

export function checkLogin() {
	if (getValue('login')) {
		const expireDateNumber = getValue('expire');
		const date = new Date(expireDateNumber);
		const now = new Date();

		return isBefore(now, date);
	}
	return false;
}
