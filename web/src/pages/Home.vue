<template>
  <div class="max-w-7xl mx-auto">
    <div class="min-h-screen flex">

      <SideNavigation />

      <!-- Main content -->
      <main class="grow overflow-hidden px-6">
        <div class="w-full h-full max-w-[1072px] mx-auto flex flex-col">

          <Header />
          <Hero />

          <!-- Content -->
          <div class="grow md:flex space-y-8 md:space-y-0 md:space-x-8 pb-16 md:pb-20">

            <!-- Middle area -->
            <div class="grow">
              <div class="max-w-[700px]">
                <div class="space-y-10">

                  <ArticlesList />
                  <Talks />
                  <FeaturedProjects />

                </div>
              </div>
            </div>

            <!-- Right sidebar -->
            <aside class="md:w-[240px] lg:w-[300px] shrink-0">
              <div class="space-y-6">

                <WidgetNewsletter />
                <WidgetSponsor />
                <WidgetBook />

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
import Hero from '../partials/Hero.vue'
import ArticlesList from '../partials/ArticlesList.vue'
import Talks from '../partials/Talks.vue'
import FeaturedProjects from '../partials/FeaturedProjects.vue'
import WidgetNewsletter from '../partials/WidgetNewsletter.vue'
import WidgetSponsor from '../partials/WidgetSponsor.vue'
import WidgetBook from '../partials/WidgetBook.vue'
import Footer from '../partials/Footer.vue'

export default {
  name: 'Home',
  components: {
    SideNavigation,
    Header,
    Hero,
    ArticlesList,
    Talks,
    FeaturedProjects,
    WidgetNewsletter,
    WidgetSponsor,
    WidgetBook,
    Footer,
  },
}
</script>