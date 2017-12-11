// Copyright © 2017 EasyStack Inc. Shawn Wang <shawn.wang@easystack.cn>
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package ostree

import (
	"log"
	"os/exec"

	"gitlab.com/EasyStack/yakety/lib/error"
)

var execCommand = exec.Command

func IsOstreeHost() bool {
	_, err := execCommand("rpm-ostree", "status", "--json").Output()
	if err == nil {
		return true
	} else {
		log.Println(error.HOST_NOT_SUPPORT)
		return false
	}
}
