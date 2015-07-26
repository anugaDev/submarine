using TyphenApi.Type.Submarine;

namespace TyphenApi.Controller.WebApi
{
    public class Submarine : IWebApiController<Error>
    {
        public IWebApiRequestSender RequestSender { get; private set; }
        public ISerializer Serializer { get; private set; }
        public IDeserializer Deserializer { get; private set; }

        public Submarine()
        {
            RequestSender = new WebApiRequestSenderWWW();

            var jsonSerializer = new JSONSerializer();
            Serializer = jsonSerializer;
            Deserializer = jsonSerializer;
        }

        public void OnBeforeRequestSend(IWebApiRequest request)
        {
            request.Headers["Content-Type"] = "application/json";
        }

        public void OnRequestError(WebApiError<Error> error)
        {
        }

        public void OnRequestSuccess(IWebApiRequest request, IWebApiResponse response)
        {
        }
    }
}