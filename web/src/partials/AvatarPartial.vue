<template>
	<img :class="className" :src="avatar" :alt="alt" />
</template>

<script>
import photo from '@images/profile/me.png';

export default {
	name: 'AvatarPartial',
	props: {
		width: {
			type: String,
			required: false,
			default: 'w-12',
		},
		avatar: {
			type: String,
			required: false,
			default: photo,
		},
		alt: {
			type: String,
			required: false,
			default: 'gocanto',
		},
	},
	computed: {
		className() {
			return `border-2 border-indigo-400 rounded-full ${this.width}`;
		},
	},
};
</script>
