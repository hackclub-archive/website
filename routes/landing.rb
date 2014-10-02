module HackEDU
  module Routes
    module Landing
      def self.registered(app)
        app.get '/' do
          @sponsors = [
            {
              name: 'Test',
              logo: '/images/open_source.svg'
            },
            {
              name: 'Test 2',
              logo: '/images/open_source.svg'
            },
            {
              name: 'Test 3',
              logo: '/images/open_source.svg'
            }
          ]
          erb :index
        end

        app.get '/contact' do
          erb :contact
        end

        app.get '/attributions' do
          @icons = [
            {
              name: 'Books',
              url: 'http://thenounproject.com/term/books/21509/',
              author: 'Piotrek Chuchla',
              author_url: 'http://www.piotrekchuchla.com'
            },
            {
              name: 'Community',
              url: 'http://thenounproject.com/term/community/5040/',
              author: 'Dmitry Baranovskiy',
              author_url: 'http://dmitry.baranovskiy.com'
            },
            {
              name: 'Open Source',
              url: 'http://thenounproject.com/term/open-source/8233/',
              author: 'Oriol Carbonell',
              author_url: 'http://www.hiperic.com'
            }
          ]
          erb :attributions
        end
      end
    end
  end
end
