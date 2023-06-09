﻿using UserService.Model;

namespace UserService.Service
{
    public interface IAppUserService
    {
        public List<AppUser> Get();
        public Task<AppUser> GetById(string id);
        Task RegisterUser(RegistrationUser dto);
        Task<bool> CheckIfUsernameExistsAsync(string username);
        Task<bool> UsernameMatchesPasswordAsync(Credentials dto);
        Task<string> LogInUserAsync(Credentials dto);
        Task<User> GetCurrentUser(string id);
        Task UpdateUser(User userDto);
        Task ChangePassword(ChangePasswordRequest request);
        Task<string> GetFullNameById(string id);

    }
}
