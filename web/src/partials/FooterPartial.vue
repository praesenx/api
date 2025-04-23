<template>
	<footer class="border-t border-slate-200 dark:border-slate-800">
		<div class="py-8">
			<div class="text-center md:flex md:items-center md:justify-between">
				<!-- Social links -->
				<ul class="inline-flex mb-4 md:order-1 md:ml-4 md:mb-0 space-x-2">
					<li>
						<a class="flex justify-center items-center text-sky-500 hover:text-sky-600 transition duration-150 ease-in-out" href="https://x.com/gocanto" aria-label="X">
							<svg class="w-8 h-8 fill-current" viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg">
								<path d="m13.063 9 3.495 4.475L20.601 9h2.454l-5.359 5.931L24 23h-4.938l-3.866-4.893L10.771 23H8.316l5.735-6.342L8 9h5.063Zm-.74 1.347h-1.457l8.875 11.232h1.36l-8.778-11.232Z"></path>
							</svg>
						</a>
					</li>

					<li>
						<a class="flex justify-center items-center text-sky-500 hover:text-sky-600 transition duration-150 ease-in-out" href="https://www.youtube.com/@gocanto" aria-label="YouTube">
							<svg class="w-8 h-8 fill-current" viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg">
								<path
									d="M23.8 12.8c-.2-1.3-.8-2.2-2.2-2.4C19.4 10 16 10 16 10s-3.4 0-5.6.4c-1.4.2-2.1 1.1-2.2 2.4C8 14.1 8 16 8 16s0 1.9.2 3.2c.2 1.3.8 2.2 2.2 2.4 2.2.4 5.6.4 5.6.4s3.4 0 5.6-.4c1.4-.3 2-1.1 2.2-2.4.2-1.3.2-3.2.2-3.2s0-1.9-.2-3.2ZM14 19v-6l5 3-5 3Z"
								></path>
							</svg>
						</a>
					</li>

					<li>
						<a class="flex justify-center items-center text-sky-500 hover:text-sky-600 transition duration-150 ease-in-out" href="https://github.com/gocanto" aria-label="Github">
							<svg class="w-8 h-8 fill-current" viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg">
								<path
									d="M16 8.2c-4.4 0-8 3.6-8 8 0 3.5 2.3 6.5 5.5 7.6.4.1.5-.2.5-.4V22c-2.2.5-2.7-1-2.7-1-.4-.9-.9-1.2-.9-1.2-.7-.5.1-.5.1-.5.8.1 1.2.8 1.2.8.7 1.3 1.9.9 2.3.7.1-.5.3-.9.5-1.1-1.8-.2-3.6-.9-3.6-4 0-.9.3-1.6.8-2.1-.1-.2-.4-1 .1-2.1 0 0 .7-.2 2.2.8.6-.2 1.3-.3 2-.3s1.4.1 2 .3c1.5-1 2.2-.8 2.2-.8.4 1.1.2 1.9.1 2.1.5.6.8 1.3.8 2.1 0 3.1-1.9 3.7-3.7 3.9.3.4.6.9.6 1.6v2.2c0 .2.1.5.6.4 3.2-1.1 5.5-4.1 5.5-7.6-.1-4.4-3.7-8-8.1-8z"
								></path>
							</svg>
						</a>
					</li>

					<li>
						<a class="flex justify-center items-center text-sky-500 hover:text-sky-600 transition duration-150 ease-in-out" href="https://www.linkedin.com/in/gocanto/" aria-label="LinkedIn">
							<svg class="w-8 h-10 fill-current" viewBox="0 0 32 32" xmlns="http://www.w3.org/2000/svg">
								<path
									d="M12.51 8.796v1.697a3.738 3.738 0 0 1 3.288-1.684c3.455 0 4.202 2.16 4.202 4.97V19.5h-3.2v-5.072c0-1.21-.244-2.766-2.128-2.766-1.827 0-2.139 1.317-2.139 2.676V19.5h-3.19V8.796h3.168ZM7.2 6.106a1.61 1.61 0 0 1-.988 1.483 1.595 1.595 0 0 1-1.743-.348A1.607 1.607 0 0 1 5.6 4.5a1.601 1.601 0 0 1 1.6 1.606Z"
								/>
								<path d="M7.2 8.809H4V19.5h3.2V8.809Z" />
							</svg>
						</a>
					</li>
				</ul>

				<!-- Copyright -->
				<div class="text-sm text-slate-500 dark:text-slate-400">Copyright &copy; gocanto. All rights reserved.</div>
			</div>
		</div>
	</footer>
</template>

<script>
export default {
	name: 'FooterPartial',
};
</script>
