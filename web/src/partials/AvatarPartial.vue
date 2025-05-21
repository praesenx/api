<template>
	<img :class="className" :src="props.avatar" :alt="props.alt" />
</template>

<script setup lang="ts">
import { computed } from "vue";
import photo from '@images/profile/me.jpg';

interface Props {
    avatar?: string;
    alt?: string;
    width: string;
}

const props = withDefaults(defineProps<Props>(), {
    width: 'w-12',
    avatar: photo,
    alt: 'gocanto',
});

const className = computed<string>(() => {
    return `border-2 border-fuchsia-400 dark:border-slate-400 rounded-full ${props.width}`;
})
</script>
