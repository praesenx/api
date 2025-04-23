<template>
	<div class="max-w-7xl mx-auto">
		<div class="min-h-screen flex">
			<SideNavigation />

			<!-- Main content -->
			<main class="grow overflow-hidden px-6">
				<div class="rounded-lg w-full h-full max-w-[1072px] mx-auto flex flex-col">
					<HeaderPartial />

					<!-- Content -->
					<div class="grow md:flex space-y-8 md:space-y-0 md:space-x-8 pt-12 md:pt-16 pb-16 md:pb-20">
						<!-- Middle area -->
						<div class="grow">
							<div class="max-w-[700px]">
								<section>
									<!-- Page title -->
									<h2 class="h2 font-aspekta mb-5">
										Hi. I'm Gus
										<span
											class="inline-flex relative text-sky-500 before:absolute before:inset-0 before:bg-sky-200 dark:before:bg-sky-500 before:opacity-30 before:-z-10 before:-rotate-2 before:translate-y-1/4"
										>
											<a href="https://x.com/gocanto" title="follow me on twitter" target="_blank">@gocanto</a>
										</span>
									</h2>

									<img class="rounded-lg w-full mb-5" src="../images/profile/about.jpg" alt="About" />

									<!-- Page content -->
									<div class="text-slate-500 dark:text-slate-400 space-y-8">
										<div class="space-y-4">
											<h3 class="h3 font-aspekta text-slate-800 dark:text-slate-100">About</h3>
											<p class="block mb-3">
												I am a dedicated engineering leader passionate about building seamless, high-quality experiences for organizations and open source. With over fifteen years of experience in software
												development, workplace technology, and infrastructure management, I specialize in Golang, Node.js, and PHP.&nbsp; I am proficient in Laravel, Symfony, and modern web frameworks like
												Next.Js. I have led teams designing and implementing scalable, high-performance systems, ensuring reliability and efficiency across complex technical environments.
											</p>
											<p class="block mb-3">
												Beyond technical expertise, I have a strong leadership background in managing cross-functional teams, optimizing workflows, and implementing best practices that drive productivity and
												innovation. I thrive in fast-paced environments that demand strategic thinking, problem-solving, and a commitment to delivering high-quality results.
											</p>
											<p class="block">I constantly seek opportunities to enhance technology-driven processes, improve operational efficiency, and empower teams with the tools and strategies to succeed.</p>
										</div>
										<div class="space-y-4">
											<h3 class="h3 font-aspekta text-slate-800 dark:text-slate-100">Career</h3>
											<span class="block mb-3">
												I have worked closely with financial services companies, delivering secure and compliant solutions that align with industry regulations. I understand financial institutions' technical and
												operational demands and have implemented robust architectures that support high-availability systems, data security, and transactional integrity.
											</span>
											<span class="block mb-3">
												In addition to my software engineering and leadership experience, I have extensive workplace technology and IT infrastructure management expertise. I have overseen cloud infrastructure,
												on-premises network administration, and enterprise IT solutions, ensuring secure, scalable, and efficient operations. My knowledge extends to user management, security policies, and system
												integration, making me well-versed in enterprise collaboration tools.
											</span>
											<span class="block">
												I am well-versed in network protocols, including DHCP, TCP/IP, and VPN configurations, ensuring seamless connectivity and system security across distributed teams and cloud-based
												infrastructures. My expertise in managing networking challenges and experience in cloud services and workplace technology allows me to contribute effectively to IT operations and digital
												transformation initiatives.
											</span>
										</div>

										<Experience />

										<div class="space-y-8">
											<h2 class="h3 font-aspekta text-slate-800 dark:text-slate-100">Let's Connect</h2>
											<p>
												I'm excited to connect with others via <a class="font-medium text-sky-500 hover:underline" href="#0">email</a> and
												<a class="font-medium text-sky-500 hover:underline" href="#0">Twitter</a> to chat about projects and ideas. Currently, I'm not taking on freelance projects, but I am open to hearing about
												potential opportunities, discussing them with you and then potentially collaborating if it's a good fit.
											</p>
										</div>
									</div>
								</section>
								<!-- content --->
							</div>
						</div>

						<!-- Right sidebar -->
						<aside class="md:w-[240px] lg:w-[300px] shrink-0">
							<div class="space-y-6">
								<WidgetNewsletter />

								<WidgetSponsorPartial />
							</div>
						</aside>
					</div>

					<FooterPartial />
				</div>
			</main>
		</div>
	</div>
</template>

<script>
import SideNavPartial from '../partials/SideNavPartial.vue';
import HeaderPartial from '../partials/HeaderPartial.vue';
import ExperienceSection from '../partials/ExperienceSection.vue';
import WidgetNewsletterPartial from '../partials/WidgetNewsletterPartial.vue';
import WidgetSponsorPartial from '../partials/WidgetSponsorPartial.vue';
import FooterPartial from '../partials/FooterPartial.vue';

export default {
	name: 'AboutPage',
	components: {
		SideNavigation: SideNavPartial,
		HeaderPartial,
		Experience: ExperienceSection,
		WidgetNewsletter: WidgetNewsletterPartial,
		WidgetSponsorPartial,
		FooterPartial,
	},
};
</script>
