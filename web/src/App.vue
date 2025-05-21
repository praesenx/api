<template>
	<router-view />
</template>

<script setup lang="ts">
import { useUserStore } from "@stores/users/user.ts"

const userStore = useUserStore()

console.log(userStore.fresh().profile)

</script>
