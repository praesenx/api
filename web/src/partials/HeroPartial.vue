<template>
	<section>
		<div class="max-w-[700px]">
			<div class="pt-8 pb-10">
				<AvatarPartial width="w-14" />
				<h1 class="h1 font-aspekta mb-5 mt-5">
					I write about coding, engineering, and
					<span class="inline-flex relative text-sky-500 before:absolute before:inset-0 before:bg-sky-200 dark:before:bg-sky-500 before:opacity-30 before:-z-10 before:-rotate-2 before:translate-y-1/4"
						>leadership</span
					>
					as a service.
				</h1>
				<p class="text-lg text-slate-500 dark:text-slate-400 mb-5">Writer, Speaker, Developer, Founder, and Leadership.</p>
				<p class="mb-5">
					I'm a full-stack Software Engineer leader with over two decades of building complex web systems and products, specialising in areas like e-commerce, banking, cross-payment solutions, cyber
					security, and customer success.
				</p>
			</div>
		</div>
	</section>
</template>

<script>
import AvatarPartial from "@partials/AvatarPartial.vue";

export default {
	name: 'HeroPartial',
    components: {
        AvatarPartial,
    }
};
</script>
