import { defineStore } from "pinia";
import { data as UserPayload } from "@stores/users/data.ts";


export const useUserStore = defineStore('user', {
    state: () => ({}),
    actions: {
        fresh() {
            return UserPayload;
        },
    },
})
