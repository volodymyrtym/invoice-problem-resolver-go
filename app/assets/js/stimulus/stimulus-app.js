import {Application} from "@hotwired/stimulus";

import LoginController from "./login_controller";
import DailyActivityCreateController from "./daily-activities-create-controller";
import DailyActivityListController from "./daily-activities-list-controller";

const application = Application.start();

application.register("login", LoginController);
application.register("daily-activity-create", DailyActivityCreateController);
application.register("daily-activity-list", DailyActivityListController);
