﻿using Grpc.Core;

namespace UserService.Exceptions
{
    public static class ExceptionStatusCode
    {
        private static Dictionary<Type, StatusCode> exceptionStatusCodes = new Dictionary<Type, StatusCode>
        {
            {typeof(Exception), StatusCode.Internal},
            {typeof(UserAlreadyExistsException), StatusCode.AlreadyExists},
            {typeof(IncorrectCredentialsException), StatusCode.NotFound},
            {typeof(UserNotFoundException), StatusCode.NotFound},

        };

        public static StatusCode GetExceptionStatusCode(Exception ex)
        {
            bool exceptionFound = exceptionStatusCodes.TryGetValue(ex.GetType(), out var statusCode);
            return exceptionFound ? statusCode : StatusCode.Internal;
        }
    }
}
