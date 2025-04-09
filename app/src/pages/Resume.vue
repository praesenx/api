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
                  <h1 class="h1 font-aspekta mb-12">My resume</h1>
                  <!-- Page content -->
                  <div class="text-slate-500 dark:text-slate-400 space-y-12">

                    <Education />
                    <Experience />
                    <Awards />
                    <Recommendations />

                  </div>
                </section>

              </div>
            </div>

            <!-- Right sidebar -->
            <aside class="md:w-[240px] lg:w-[300px] shrink-0">
              <div class="space-y-6">

                <WidgetSkills />
                <WidgetLanguages />
                <WidgetReferences />

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
import Education from '../partials/Education.vue'
import Experience from '../partials/Experience.vue'
import Awards from '../partials/Awards.vue'
import Recommendations from '../partials/Recommendations.vue'
import WidgetSkills from '../partials/WidgetSkills.vue'
import WidgetLanguages from '../partials/WidgetLanguages.vue'
import WidgetReferences from '../partials/WidgetReferences.vue'
import Footer from '../partials/Footer.vue'

export default {
  name: 'Resume',
  components: {
    SideNavigation,
    Header,
    Education,
    Experience,
    Awards,
    Recommendations,
    WidgetSkills,
    WidgetLanguages,
    WidgetReferences,
    Footer,
  },
}
</script>