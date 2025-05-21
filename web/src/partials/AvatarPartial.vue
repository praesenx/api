<template>
<!--	<img :class="className" :src="props.avatar" :alt="props.alt" />-->

    <div class="flex items-center gap-4">
        <img class="w-10 h-10 rounded-full" :src="props.avatar" alt="">
        <div class="font-medium dark:text-white">
            <div>Jese Leos</div>
            <div class="text-sm text-gray-500 dark:text-gray-400">Joined in August 2014</div>
        </div>
    </div>

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
