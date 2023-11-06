package auth

const (
	forgetPasswordEmailTemplate = `
		<!DOCTYPE html>
			<html>
			<head>
				<style>
					body {
						font-family: Arial, sans-serif;
						background-color: #f7f7f7;
						margin: 0;
						padding: 20px;
					}
					.email-container {
						background-color: #ffffff;
						max-width: 600px;
						margin: 0 auto;
						padding: 20px;
						border-radius: 5px;
						box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
					}
					.details-header {
						font-weight: bold;
					}
				</style>
					</head>
					<body>
						<div class="email-container">
						<p>Dear {{.NamaLengkap}},</p>
						<p>You have requested to forget your password!</p>

						<p><span class="details-header">Click here to validate your request</span></p>
						<ul>
							<li><a href="https://www.w3schools.com?token={{.Token}}">Validate your request</a></li>
						</ul>

						<p>Please keep this email for your records. In case you have any questions or need to make any changes to your registration details, please do not hesitate to contact us at <a href="mailto:bistleague@std.stei.itb.ac.id">bistleague@std.stei.itb.ac.id</a> or +62 81290908333.</p>
						<p>Sincerely,</p>
						<p>Bist league 6</p>
						<p> do not reply this email </p>
						</div>
					</body>
					</html>
`
	forgetPasswordValidateTokenEmailTemplate = `
	<!DOCTYPE html>
			<html>
			<head>
				<style>
					body {
						font-family: Arial, sans-serif;
						background-color: #f7f7f7;
						margin: 0;
						padding: 20px;
					}
					.email-container {
						background-color: #ffffff;
						max-width: 600px;
						margin: 0 auto;
						padding: 20px;
						border-radius: 5px;
						box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
					}
					.details-header {
						font-weight: bold;
					}
				</style>
					</head>
					<body>
						<div class="email-container">
						<p>Dear {{.NamaLengkap}},</p>
						<p>Change your password now!</p>

						<p><span class="details-header">Here are the details of your account:</span></p>
						<ul>
							<li><span>Email:</span> {{.Email}}</li>
							<li><span>New Password:</span> {{.Password}}</li>
						</ul>

						<p>Please keep this email for your records. In case you have any questions or need to make any changes to your registration details, please do not hesitate to contact us at <a href="mailto:bistleague@std.stei.itb.ac.id">bistleague@std.stei.itb.ac.id</a> or +62 81290908333.</p>
						<p>Sincerely,</p>
						<p>Bist league 6</p>
						<p> do not reply this email </p>
						</div>
					</body>
					</html>
`
)
