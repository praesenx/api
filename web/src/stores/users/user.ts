import { defineStore } from 'pinia';
import { data as UserPayload } from '@stores/users/data.ts';
import { User } from '@stores/users/userType.ts';

export const useUserStore = defineStore('user', {
	actions: {
		getUser() {
			return getFor('user-bucket', UserPayload);
		},
	},
});

function getFor(key: string, seed: User): User {
	const storedValue = localStorage.getItem(key);

	if (storedValue === null) {
		localStorage.setItem(key, JSON.stringify(seed));

		return seed;
	} else {
		try {
			return JSON.parse(storedValue);
		} catch (error) {
			localStorage.setItem(key, JSON.stringify(seed));

			return seed;
		}
	}
}
