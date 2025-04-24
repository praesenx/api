<template>
	<div class="space-y-8">
		<h2 class="h3 font-aspekta text-slate-800 dark:text-slate-100">Work Experience</h2>
		<ul class="space-y-8">
			<!-- Item -->
			<li v-for="item in sortedExperience" :key="item.uuid" class="relative group">
				<div
					class="flex items-start before:absolute before:left-0 before:h-full before:w-px before:bg-slate-200 dark:before:bg-slate-800 before:self-start before:ml-[28px] before:-translate-x-1/2 before:translate-y-8 group-last-of-type:before:hidden"
				>
					<div
						class="absolute left-0 h-14 w-14 flex items-center justify-center border border-slate-200 dark:border-slate-800 dark:bg-linear-to-t dark:from-slate-800 dark:to-slate-800/30 bg-white dark:bg-slate-900 rounded-full"
					>
						<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16">
							<path
								fill="#DE4D3B"
								d="M15.912 6.545h-7.75v3.273h4.393c-.702 2.182-2.437 2.91-4.428 2.91-2.12.002-3.992-1.368-4.601-3.37a4.693 4.693 0 0 1 1.959-5.307 4.854 4.854 0 0 1 5.722.33l2.414-2.265C10.947-.31 6.964-.69 3.867 1.186.771 3.06-.683 6.735.308 10.18c.99 3.445 4.185 5.822 7.819 5.82 4.477 0 8.525-2.91 7.785-9.456Z"
							/>
						</svg>
					</div>
					<div class="pl-20 space-y-1">
						<div class="text-xs text-slate-500 uppercase">{{ item.start_date }} <span class="text-slate-400 dark:text-slate-600">·</span> {{ item.end_date }}</div>
						<div class="font-aspekta font-[650] text-slate-800 dark:text-slate-100">{{ item.position }}</div>
						<div class="text-sm font-medium text-slate-800 dark:text-slate-100">{{ item.company }}</div>
						<div class="text-sm text-slate-500 dark:text-slate-400">{{ item.summary }}</div>
					</div>
				</div>
			</li>
		</ul>
	</div>
</template>

<script>
export default {
	name: 'ExperiencePartial',
    props: {
        experience: {
            type: Array,
            required: true,
        }
    },
    computed: {
        sortedExperience () {
            // eslint-disable-next-line vue/no-mutating-props
            return this.experience.sort((a, b) => a.start_date - b.start_date)
        }
    }
};
</script>
