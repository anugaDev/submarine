﻿namespace TyphenApi
{
    public interface IWebApiController<ErrorT> where ErrorT : TypeBase
    {
        void OnBeforeRequestSend(IWebApiRequest request);
        void OnRequestError(WebApiError<ErrorT> error);
        void OnRequestSuccess(IWebApiRequest request, IWebApiResponse response);
    }
}
