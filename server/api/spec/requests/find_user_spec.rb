require 'rails_helper'

RSpec.describe "FindUser", type: :request do
  describe "POST /find_user" do
    before do
      user = create :user, :with_stupid_password
      login_user(user)

      create :user, name: 'Shimakaze'
      create :user, name: 'Yamato'
    end

    context 'with a valid request' do
      before do
        post find_user_path, name: 'Shimakaze'
      end

      it "should work" do
        expect(response).to have_http_status(200)
      end
      it "should return a reasponse that includes a user" do
        expect(response_json[:user][:name]).to eq 'Shimakaze'
      end
    end

    context 'with invalid params' do
      it "should not work" do
        expect { post find_user_path }.to raise_error(Virtus::CoercionError)
      end
    end

    context 'with a no-existing user name' do
      it "should return empty" do
        post find_user_path, name: 'Kaga'
        expect(response_json[:user]).to be_nil
      end
    end
  end
end