<template>
  <div class="max-w-7xl mx-auto">
    <div class="min-h-screen flex">

      <SideNavigation />

      <!-- Main content -->
      <main class="grow overflow-hidden px-6">
        <div class="w-full h-full max-w-[1072px] mx-auto flex flex-col">

          <Header />

          <!-- Content -->
          <div class="grow md:flex space-y-8 md:space-y-0 md:space-x-8 pt-12 md:pt-16 pb-16 md:pb-20">

            <!-- Middle area -->
            <div class="grow">
              <div class="max-w-[700px]">

                <section>
                  <!-- Page title -->
                  <h1 class="h1 font-aspekta mb-5">Hi. I'm Mark <span class="inline-flex relative text-sky-500 before:absolute before:inset-0 before:bg-sky-200 dark:before:bg-sky-500 before:opacity-30 before:-z-10 before:-rotate-2 before:translate-y-1/4">@mrk27</span> Ivings 🤟</h1>
                  <img class="w-full" src="../images/about.png" width="692" height="390" alt="About" />
                  <!-- Page content -->
                  <div class="text-slate-500 dark:text-slate-400 space-y-8">
                    <div class="space-y-8">
                      <h2 class="h3 font-aspekta text-slate-800 dark:text-slate-100">Short Bio</h2>
                      <p>
                        I'm a software engineer with more than 10 years of experience in a variety of domains. For the past few years, I've focused on highload server-side projects, distributed systems, and platform development - tinkering with infrastructure, all things containers and Cloud Native.
                      </p>
                      <p>
                        While there isn't a Wikipedia page about me (sorry folks!), a media bio is available below.
                      </p>
                    </div>
                    <div class="space-y-8">
                      <h2 class="h3 font-aspekta text-slate-800 dark:text-slate-100">Career</h2>
                      <p>
                        In my role as a Senior Software Engineer for Google Chrome, I am responsible for developing and maintaining the <a class="font-medium text-sky-500 hover:underline" href="#0">Chrome Web Browser</a>.
                      </p>
                      <p>
                        My work involves developing and testing new features, optimizing performance and security, and ensuring the browser works for users around the world. I also work closely with other Google teams ensure Chrome is well-integrated with other <a class="font-medium text-sky-500 hover:underline" href="#0">Google</a> products and services.
                      </p>
                      <p>
                        As CTO of AppForYou, <strong class="font-medium text-slate-800 dark:text-slate-100">I am responsible for leading</strong> the technical teamand developing the company's technology strategy. I work closely with the engineering team to ensure that the products and services we provide are secure.
                      </p>
                    </div>

                    <Experience />

                    <div class="space-y-8">
                      <h2 class="h3 font-aspekta text-slate-800 dark:text-slate-100">Let's Connect</h2>
                      <p>
                        I'm excited to connect with others via <a class="font-medium text-sky-500 hover:underline" href="#0">email</a> and <a class="font-medium text-sky-500 hover:underline" href="#0">Twitter</a> to chat about projects and ideas. Currently, I'm not taking on freelance projects, but I am open to hearing about potential opportunities, discussing them with you and then potentially collaborating if it's a good fit.
                      </p>
                    </div>
                  </div>
                </section>

              </div>
            </div>

            <!-- Right sidebar -->
            <aside class="md:w-[240px] lg:w-[300px] shrink-0">
              <div class="space-y-6">

                <WidgetNewsletter />
                <WidgetSponsor />

              </div>
            </aside>

          </div>

          <Footer />

        </div>
      </main>

    </div>
  </div>
</template>

<script>
import SideNavigation from '../partials/SideNavigation.vue'
import Header from '../partials/Header.vue'
import Experience from '../partials/Experience.vue'
import WidgetNewsletter from '../partials/WidgetNewsletter.vue'
import WidgetSponsor from '../partials/WidgetSponsor.vue'
import Footer from '../partials/Footer.vue'

export default {
  name: 'About',
  components: {
    SideNavigation,
    Header,
    Experience,
    WidgetNewsletter,
    WidgetSponsor,
    Footer,
  },
}
</script>